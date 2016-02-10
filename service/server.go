package main

import (
    "fmt"
    curl "github.com/andelf/go-curl"
    "github.com/go-martini/martini"
    "github.com/ryanuber/go-filecache"
    "io/ioutil"
    "net/http"
    "os"
    "time"
    "strings"
    "golang.org/x/net/html"
    "github.com/PuerkitoBio/goquery"
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
        // writer.Header().Set("Content-Type", "application/json")
        writer.Header().Set("Content-Type", "text/html")
        return getCpf(params["id"], params["datnasc"], params["captcha"])
    })

    m.Get("/cnpj/:id", func(params martini.Params, writer http.ResponseWriter) string {
        writer.Header().Set("Content-Type", "application/json")
        return getCnpj(params["id"])
    })


    m.Run()
}

func getCaptcha(captchaType string, id string) string {
    easy := curl.EasyInit()
    defer easy.Cleanup()

    easy.Setopt(curl.OPT_COOKIEJAR, "cache/"+captchaType+"/"+id+"_cookie.jar")
    easy.Setopt(curl.OPT_VERBOSE, true)
    if (captchaType == "cpf") {
        easy.Setopt(curl.OPT_URL, "http://www.receita.fazenda.gov.br/Aplicacoes/ATCTA/CPF/captcha/gerarCaptcha.asp")
    }
    if (captchaType == "cnpj") {
        easy.Setopt(curl.OPT_URL, "http://www.receita.fazenda.gov.br/pessoajuridica/cnpj/cnpjreva/captcha/gerarCaptcha.asp")
    }

    easy.Setopt(curl.OPT_WRITEFUNCTION, func(ptr []byte, userdata interface{}) bool {
        file := userdata.(*os.File)
        if _, err := file.Write(ptr); err != nil {
            return false
        }
        return true
    })

    fp, _ := os.Create("cache/"+captchaType+"/"+id+"_captcha.png")
    defer fp.Close() // defer close

    easy.Setopt(curl.OPT_WRITEDATA, fp)

    easy.Setopt(curl.OPT_VERBOSE, true)

    if err := easy.Perform(); err != nil {
        println("ERROR", err.Error())
    }

    return captchaType+"/"+id+"_captcha.png"
}

// Helper function to pull the href attribute from a Token
func getClass(t html.Token) (ok bool, href string) {
    // Iterate over all of the Token's attributes until we find an "href"
    for _, a := range t.Attr {
        if a.Key == "href" {
            href = a.Val
            ok = true
        }
    }

    // "bare" return will return the variables (ok, href) as defined in
    // the function definition
    return
}

func getCpf(id string, datnasc string, captcha string) string {
    // cached := getFromCache("cnpj", id)
    // if cached != "" {
    //     return cached
    // }
    easy := curl.EasyInit()
    defer easy.Cleanup()
    id = "024.622.089-92"
    // captcha="T4PQ35"
    datnasc="15/11/1978"

    // easy.Setopt(curl.OPT_COOKIE, "ASPSESSIONIDSSQTBATT=LMOAEMLBMDEJIDDBJCBDOMKN")
    easy.Setopt(curl.OPT_HTTPHEADER, []string{"Accept:text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8", "Content-Type:application/x-www-form-urlencoded", "refer:http://www.receita.fazenda.gov.br/aplicacoes/atcta/cpf/ConsultaPublica.asp", "Cookie:ASPSESSIONIDSSQTBATT=PEDFFMLBHHFBNBBOHOJMEJLI"})
    // easy.Setopt(curl.OPT_HTTPHEADER, []string{"refer:http://www.receita.fazenda.gov.br/aplicacoes/atcta/cpf/ConsultaPublica.asp"})
    easy.Setopt(curl.OPT_VERBOSE, true)
    // easy.Setopt(curl.OPT_COOKIEJAR, "cache/cpf/"+id+"_cookie.jar")
    // easy.Setopt(curl.OPT_COOKIEFILE, "cache/cpf/"+id+"_cookie.jar")
    easy.Setopt(curl.OPT_URL, "http://www.receita.fazenda.gov.br/aplicacoes/atcta/cpf/ConsultaPublicaExibir.asp")
    postdata := "txtTexto_captcha_serpro_gov_br=" + captcha + "&tempTxtCPF=" + id + "&tempTxtNascimento=" + datnasc + "&temptxtToken_captcha_serpro_gov_br=\"\"&temptxtTexto_captcha_serpro_gov_br=" + captcha + "&Enviar=Consultar"
    fmt.Printf("Post data: %v\n", postdata)
    fmt.Printf("Post data: %v\n", len(postdata))
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
    }

    // f, _ := os.Create("cache/cpf/"+id+".html")
    // defer f.Close() // defer close
    // f.WriteString(result)

    // doc, err := html.Parse(strings.NewReader(result))
    // if err != nil {
    //     fmt.Printf("ERROR: %v\n", err)
    // }
    // var f func(*html.Node, bool)
    // f = func(n *html.Node, printText bool) {
    //     if printText && n.Type == html.TextNode {
    //         fmt.Printf("%q\n", n.Data)
    //         fmt.Printf("%q\n", n.Attr)
    //     }
    //     printText = printText || (n.Type == html.ElementNode && n.Data == "span")
    //     for c := n.FirstChild; c != nil; c = c.NextSibling {
    //         f(c, printText)
    //     }
    // }
    // f(doc, false)

    doc, _ := goquery.NewDocumentFromReader(strings.NewReader((result)))
    doc.Find("span").Each(func(j int, s *goquery.Selection) {
        if s.HasClass("clConteudoDados") {
            fmt.Printf("%q\n", s.Text())
        }
    })
    // htmlResult, _ := doc.Html()
    // fmt.Println(htmlResult)

    return result
    // return saveOnCache("cnpj", id, result)
}

func getCnpj(id string) string {
    cached := getFromCache("cnpj", id)
    if cached != "" {
        return cached
    }
    easy := curl.EasyInit()
    defer easy.Cleanup()

    easy.Setopt(curl.OPT_URL, "http://viacep.com.br/ws/"+id+"/json/")

    result := " "

    // make a callback function
    executionCallback := func(buf []byte, userdata interface{}) bool {
        result = string(buf)

        return true
    }

    easy.Setopt(curl.OPT_WRITEFUNCTION, executionCallback)

    if err := easy.Perform(); err != nil {
        fmt.Printf("ERROR: %v\n", err)
    }

    return saveOnCache("cnpj", id, result)
}

func getFromCache(cacheType string, id string) string {
    fc := filecache.New("cache/"+cacheType+"/"+id, 500*time.Second, nil)

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

    fc := filecache.New("cache/"+cacheType + "/" + id, 500*time.Second, updater)

    _, err := fc.Get()
    if err != nil {
        fmt.Printf("ERROR: %v\n", err)
        return ""
    }

    return content
}
