package main

import (
	"bufio"
	"coderockr"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	curl "github.com/andelf/go-curl"
	iconv "github.com/andelf/iconv-go"
	"github.com/go-martini/martini"
	"github.com/ryanuber/go-filecache"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	m := martini.Classic()
	m.Use(martini.Static("cache"))

	m.Get("/", func() string {
		return "Micro-serviço que forece consulta de CPF/CNPJ"
	})

	m.Get("/captcha/cpf/:id", func(params martini.Params, writer http.ResponseWriter) string {
		writer.Header().Set("Content-Type", "application/json")
		return getCaptcha("cpf", params["id"])
	})

	m.Get("/captcha/cnpj/:id", func(params martini.Params, writer http.ResponseWriter) string {
		writer.Header().Set("Content-Type", "application/json")
		return getCaptcha("cnpj", params["id"])
	})

	m.Get("/cpf/:id/:datnasc/:captcha", func(params martini.Params, writer http.ResponseWriter) string {
		writer.Header().Set("Content-Type", "application/json")
		cpf := getCpf(params["id"], params["datnasc"], params["captcha"])
		if cpf == "" {
			writer.WriteHeader(http.StatusNotFound)
		}
		return cpf
	})

	m.Get("/cnpj/:id/:captcha", func(params martini.Params, writer http.ResponseWriter) string {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		cnpj := getCnpj(params["id"], params["captcha"])
		if cnpj == "" {
			writer.WriteHeader(http.StatusNotFound)
		}
		return cnpj
	})

	m.Run()
}

func getCaptcha(captchaType string, id string) string {
	os.MkdirAll("cache/"+captchaType+"/"+id, 0777)

	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_COOKIEJAR, "cache/"+captchaType+"/"+id+"/cookie.jar")
	easy.Setopt(curl.OPT_VERBOSE, true)
	if captchaType == "cpf" {
		easy.Setopt(curl.OPT_URL, "http://www.receita.fazenda.gov.br/Aplicacoes/ATCTA/CPF/captcha/gerarCaptcha.asp")
	}
	if captchaType == "cnpj" {
		easy.Setopt(curl.OPT_URL, "http://www.receita.fazenda.gov.br/pessoajuridica/cnpj/cnpjreva/captcha/gerarCaptcha.asp")
	}

	easy.Setopt(curl.OPT_WRITEFUNCTION, func(ptr []byte, userdata interface{}) bool {
		file := userdata.(*os.File)
		if _, err := file.Write(ptr); err != nil {
			return false
		}
		return true
	})

	fp, _ := os.Create("cache/" + captchaType + "/" + id + "/captcha.png")
	defer fp.Close() // defer close

	easy.Setopt(curl.OPT_WRITEDATA, fp)

	easy.Setopt(curl.OPT_VERBOSE, true)

	if err := easy.Perform(); err != nil {
		println("ERROR", err.Error())
	}

	return captchaType + "/" + id + "/captcha.png"
}

func getCookieContent(path string) string {
	f, err := os.Open(path)
	defer f.Close() // defer close
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	bf := bufio.NewReader(f)
	for {
		switch line, err := bf.ReadString('\n'); err {
		case nil:
			// valid line, echo it.  note that line contains trailing \n.
			if len(line) >= 25 && line[0:26] == "www.receita.fazenda.gov.br" {
				return line[0 : len(line)-1] //remove \n
			}
		default:
			fmt.Printf("ERROR: %v\n", err)
			return ""
		}
	}
	return ""
}

func getCpf(id string, datnasc string, captcha string) string {
	cached := getFromCache("cpf", id)
	if cached != "" {
		return cached
	}
	cookie := coderockr.FormatCookie(getCookieContent("cache/cpf/" + id + "/cookie.jar"))
	easy := curl.EasyInit()
	defer easy.Cleanup()
	unformatedId := id
	id = coderockr.FormatCpf(id)
	datnasc = coderockr.FormatData(datnasc)

	easy.Setopt(curl.OPT_HTTPHEADER, []string{"Accept:text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8", "Content-Type:application/x-www-form-urlencoded", "refer:http://www.receita.fazenda.gov.br/aplicacoes/atcta/cpf/ConsultaPublica.asp", "Cookie:" + cookie})
	easy.Setopt(curl.OPT_VERBOSE, true)
	easy.Setopt(curl.OPT_URL, "http://www.receita.fazenda.gov.br/aplicacoes/atcta/cpf/ConsultaPublicaExibir.asp")
	postdata := "txtTexto_captcha_serpro_gov_br=" + captcha + "&tempTxtCPF=" + id + "&tempTxtNascimento=" + datnasc + "&temptxtToken_captcha_serpro_gov_br=\"\"&temptxtTexto_captcha_serpro_gov_br=" + captcha + "&Enviar=Consultar"
	easy.Setopt(curl.OPT_POST, true)
	easy.Setopt(curl.OPT_POSTFIELDS, postdata)
	easy.Setopt(curl.OPT_POSTFIELDSIZE, len(postdata))

	result := " "

	// make a callback function
	executionCallback := func(buf []byte, userdata interface{}) bool {
		result = result + string(buf)
		return true
	}

	easy.Setopt(curl.OPT_WRITEFUNCTION, executionCallback)

	if err := easy.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return ""
	}

	output, err := iconv.ConvertString(result, "iso-8859-1", "utf-8")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	cpfData := ""
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader((output)))
	doc.Find("span").Each(func(j int, s *goquery.Selection) {
		if s.HasClass("clConteudoDados") {
			cpfData = cpfData + s.Text() + "\n"
		}
	})
	if cpfData == "" {
		return cpfData
	}

	cpf, err := json.Marshal(coderockr.FormatCpfData(cpfData))
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return saveOnCache("cpf", unformatedId, string(cpf))
}

func getCnpj(id string, captcha string) string {
	cached := getFromCache("cnpj", id)
	if cached != "" {
		return cached
	}
	unformatedId := id
	id = coderockr.FormatCnpj(id)

	cookie := coderockr.FormatCookie(getCookieContent("cache/cnpj/" + unformatedId + "/cookie.jar"))

	start := curl.EasyInit()
	defer start.Cleanup()
	start.Setopt(curl.OPT_HTTPHEADER, []string{"Accept:text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8", "Content-Type:application/x-www-form-urlencoded", "refer:http://www.receita.fazenda.gov.br/pessoajuridica/cnpj/cnpjreva/cnpjreva_solicitacao.asp", "Cookie:" + cookie})
	start.Setopt(curl.OPT_VERBOSE, false)
	start.Setopt(curl.OPT_URL, "http://www.receita.fazenda.gov.br/pessoajuridica/cnpj/cnpjreva/cnpjreva_solicitacao2.asp")
	if err := start.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	cookie = "flag=1;" + cookie

	firstUrl := curl.EasyInit()
	defer firstUrl.Cleanup()
	refer := "http://www.receita.fazenda.gov.br/pessoajuridica/cnpj/cnpjreva/cnpjreva_solicitacao2.asp"
	firstUrl.Setopt(curl.OPT_HTTPHEADER, []string{"Accept:text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8", "Content-Type:application/x-www-form-urlencoded", "refer:" + refer, "Cookie:" + cookie})
	firstUrl.Setopt(curl.OPT_VERBOSE, false)
	firstUrl.Setopt(curl.OPT_URL, "http://www.receita.fazenda.gov.br/pessoajuridica/cnpj/cnpjreva/valida.asp")
	postdata := "origem=comprovante&cnpj=" + unformatedId + "&txtTexto_captcha_serpro_gov_br=" + captcha + "&submit1=Consultar&search_type=cnpj"
	firstUrl.Setopt(curl.OPT_POST, true)
	firstUrl.Setopt(curl.OPT_POSTFIELDS, postdata)
	firstUrl.Setopt(curl.OPT_POSTFIELDSIZE, len(postdata))
	if err := firstUrl.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	secondUrl := curl.EasyInit()
	defer secondUrl.Cleanup()
	secondUrl.Setopt(curl.OPT_HTTPHEADER, []string{"Accept:text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8", "Content-Type:application/x-www-form-urlencoded", "refer:" + refer, "Cookie:" + cookie})
	secondUrl.Setopt(curl.OPT_VERBOSE, false)
	secondUrl.Setopt(curl.OPT_URL, "http://www.receita.fazenda.gov.br/pessoajuridica/cnpj/cnpjreva/Cnpjreva_Vstatus.asp?origem=comprovante&cnpj="+unformatedId)
	if err := secondUrl.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	thirdUrl := curl.EasyInit()
	defer thirdUrl.Cleanup()
	thirdUrl.Setopt(curl.OPT_HTTPHEADER, []string{"Accept:text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8", "Content-Type:application/x-www-form-urlencoded", "refer:" + refer, "Cookie:" + cookie})
	thirdUrl.Setopt(curl.OPT_VERBOSE, false)
	thirdUrl.Setopt(curl.OPT_URL, "http://www.receita.fazenda.gov.br/pessoajuridica/cnpj/cnpjreva/Cnpjreva_Campos.asp")
	if err := thirdUrl.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	lastUrl := curl.EasyInit()
	defer lastUrl.Cleanup()
	lastUrl.Setopt(curl.OPT_HTTPHEADER, []string{"Accept:text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8", "Content-Type:application/x-www-form-urlencoded", "refer:" + refer, "Cookie:" + cookie})
	lastUrl.Setopt(curl.OPT_VERBOSE, false)
	lastUrl.Setopt(curl.OPT_URL, "http://www.receita.fazenda.gov.br/pessoajuridica/cnpj/cnpjreva/Cnpjreva_Comprovante.asp")
	result := " "

	// make a callback function
	executionCallback := func(buf []byte, userdata interface{}) bool {
		result = result + string(buf)
		return true
	}

	lastUrl.Setopt(curl.OPT_WRITEFUNCTION, executionCallback)

	if err := lastUrl.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return ""
	}

	output, err := iconv.ConvertString(result, "iso-8859-1", "utf-8")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	cnpjData := ""
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader((output)))
	doc.Find("font").Each(func(j int, s *goquery.Selection) {
		cnpjData = cnpjData + s.Text() + "<br>"
	})
	cnpjData = strings.Replace(cnpjData, " ", "|", -1)
	cnpjData = strings.Replace(cnpjData, "\t", "|", -1)
	cnpjData = strings.Replace(cnpjData, "\n", "|", -1)
	if cnpjData == "" {
		return ""
	}

	cnpj := coderockr.FormatCnpjData(cnpjData)
	json, err := json.Marshal(cnpj)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	return saveOnCache("cnpj", unformatedId, string(json))
}

func getFromCache(cacheType string, id string) string {
	fc := filecache.New("cache/"+cacheType+"/"+id+"/result.json", 500*time.Second, nil)

	fh, err := fc.Get()
	if err != nil {
		return ""
	}

	content, err := ioutil.ReadAll(fh)
	if err != nil {
		return ""
	}

	return string(content)
}

func saveOnCache(cacheType string, id string, content string) string {
	updater := func(path string) error {
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.Write([]byte(content))
		return err
	}

	fc := filecache.New("cache/"+cacheType+"/"+id+"/result.json", 500*time.Second, updater)

	_, err := fc.Get()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return ""
	}

	return content
}
