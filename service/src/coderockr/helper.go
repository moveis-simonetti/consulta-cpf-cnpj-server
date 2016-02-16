package coderockr

import "strings"
// import "fmt"

type CpfData struct {
	Numero            string
	Nome              string
	DataNascimento    string
	Situacao          string
	DataInscricao     string
	DigitoVerificador string
}

type CnpjData struct {
	NumeroInscricao              string
	DataAbertura                 string
	NomeEmpresarial              string
	NomeFantasia                 string
	AtividadeEconomicaPrincipal  string
	AtividadeEconomicaSecundaria string
	NaturezaJuridica             string
	Logradouro                   string
	Numero                       string
	Complemento                  string
	Cep                          string
	Bairro                       string
	Municipio                    string
	Uf                           string
	EnderecoEletronico           string
	Telefone                     string
	EnteFederativoResponsavel    string
	Situacao                     string
	DataSituacao                 string
	MotivoSituacao               string
	SituacaoEspecial             string
	DataSituacaoEspecial         string
}

func FormatCpf(cpf string) string {
	if len(cpf) < 11 || len(cpf) == 14 {
		return cpf
	}

	return cpf[0:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:11]
}

func FormatCnpj(cnpj string) string {
	if len(cnpj) < 14 || len(cnpj) == 18 {
		return cnpj
	}
	return cnpj[0:2] + "." + cnpj[2:5] + "." + cnpj[5:8] + "/" + cnpj[8:12] + "-" + cnpj[12:14]
}

func FormatData(data string) string {
	if len(data) < 8 || len(data) == 10 {
		return data
	}

	return data[0:2] + "/" + data[2:4] + "/" + data[4:8]
}

func FormatCookie(cookie string) string {
	stringSlice := strings.Split(cookie, " ")
	if len(stringSlice) > 5 {
		return stringSlice[len(stringSlice)-5] + "=" + stringSlice[len(stringSlice)-1]
	}

	stringSlice = strings.Split(cookie, "\t")
	if len(stringSlice) > 1 {
		return stringSlice[len(stringSlice)-2] + "=" + stringSlice[len(stringSlice)-1]
	}

	return cookie
}

func FormatCpfData(unformated string) CpfData {
	stringSlice := strings.Split(unformated, "\n")
	cpf := CpfData{}
	for _, text := range stringSlice {
		data := strings.Split(text, ":")
		if len(data) == 2 {
			switch data[0] {
			case "No do CPF":
				cpf.Numero = strings.Trim(data[1], " ")
			case "Nome da Pessoa F\xedsica":
				cpf.Nome = strings.Trim(data[1], " ")
			case "Data de Nascimento":
				cpf.DataNascimento = strings.Trim(data[1], " ")
			case "Situa\xe7\xe3o Cadastral":
				cpf.Situacao = strings.Trim(data[1], " ")
			case "Data da Inscri\xe7\xe3o":
				cpf.DataInscricao = strings.Trim(data[1], " ")
			case "Digito Verificador":
				cpf.DigitoVerificador = strings.Trim(data[1], " ")
			default:
				println(data[0])
			}
		}

	}
	return cpf
}

func FormatCnpjData(unformated string) CnpjData {
	cnpj := CnpjData{}
	unformated = strings.Replace(unformated, "||", "", -1)
	unformated = strings.Replace(unformated, "|", " ", -1)
	stringSlice := strings.Split(unformated, "<br>")
	for index, text := range stringSlice {
		text = strings.Trim(text, " ")
		// fmt.Printf("TEXTO:%+q\n", text)
		switch text {
			case "N?MERO DE INSCRI??O":
			case "N\xdaMERO DE INSCRI\xc7\xc3O":
				value := strings.Trim(stringSlice[index+1], " ")
				cnpj.NumeroInscricao = value[0:18] + " " + value[18:]
			case "DATA DE ABERTURA":
				cnpj.DataAbertura = strings.Trim(stringSlice[index+1], " ")
			case "NOME EMPRESARIAL":
				cnpj.NomeEmpresarial = strings.Trim(stringSlice[index+1], " ")
			case "T?TULO DO ESTABELECIMENTO (NOME DE FANTASIA)":
			case "T\xcdTULO DO ESTABELECIMENTO (NOME DE FANTASIA)":
				cnpj.NomeFantasia = strings.Trim(stringSlice[index+1], " ")
			case "C?DIGO E DESCRI??O DA ATIVIDADE ECON?MICA PRINCIPAL":
			case "C\xd3DIGO E DESCRI\xc7\xc3O DA ATIVIDADE ECON\xd4MICA PRINCIPAL":
				cnpj.AtividadeEconomicaPrincipal = strings.Trim(stringSlice[index+1], " ")
			case "C?DIGO E DESCRI??O DAS ATIVIDADES ECON?MICAS SECUND?RIAS":
			case "C\xd3DIGO E DESCRI\xc7\xc3O DAS ATIVIDADES ECON\xd4MICAS SECUND\xc1RIAS":
				value := strings.Trim(stringSlice[index+1], " ")
				if (value == "N?o informada" || value == "N\xe3o informada") {
					value = "Não informada"
				}
				cnpj.AtividadeEconomicaSecundaria = value
			case "C?DIGO E DESCRI??O DA NATUREZA JUR?DICA":
			case "C\xd3DIGO E DESCRI\xc7\xc3O DA NATUREZA JUR\xcdDICA":
				cnpj.NaturezaJuridica = strings.Trim(stringSlice[index+1], " ")
			case "LOGRADOURO":
				cnpj.Logradouro = strings.Trim(stringSlice[index+1], " ")
			case "N?MERO":
			case "N\xdaMERO":
				cnpj.Numero = strings.Trim(stringSlice[index+1], " ")
			case "COMPLEMENTO":
				cnpj.Complemento = strings.Trim(stringSlice[index+1], " ")
			case "CEP":
				cnpj.Cep = strings.Trim(stringSlice[index+1], " ")
			case "BAIRRO/DISTRITO":
				cnpj.Bairro = strings.Trim(stringSlice[index+1], " ")
			case "MUNIC?PIO":
			case "MUNIC\xcdPIO":
				cnpj.Municipio = strings.Trim(stringSlice[index+1], " ")
			case "UF":
				cnpj.Uf = strings.Trim(stringSlice[index+1], " ")
			case "ENDERE?O ELETR?NICO":
			case "ENDERE\xc7O ELETR\xd4NICO":
				cnpj.EnderecoEletronico = strings.Trim(stringSlice[index+1], " ")
			case "TELEFONE":
				cnpj.Telefone = strings.Trim(stringSlice[index+1], " ")
			case "ENTE FEDERATIVO RESPONS?VEL (EFR)":
			case "ENTE FEDERATIVO RESPONS\xc1VEL (EFR)":
				cnpj.EnteFederativoResponsavel = strings.Trim(stringSlice[index+1], " ")
			case "SITUA??O CADASTRAL":
			case "SITUA\xc7\xc3O CADASTRAL":
				cnpj.Situacao = strings.Trim(stringSlice[index+1], " ")
			case "DATA DA SITUA??O CADASTRAL":
			case "DATA DA SITUA\xc7\xc3O CADASTRAL":
				cnpj.DataSituacao = strings.Trim(stringSlice[index+1], " ")
			case "MOTIVO DE SITUA??O CADASTRAL":
			case "MOTIVO DE SITUA\xc7\xc3O CADASTRAL":
				cnpj.MotivoSituacao = strings.Trim(stringSlice[index+1], " ")
			case "SITUA??O ESPECIAL":
			case "SITUA\xc7\xc3O ESPECIAL":
				cnpj.SituacaoEspecial = strings.Trim(stringSlice[index+1], " ")
			case "DATA DA SITUA??O ESPECIAL":
			case "DATA DA SITUA\xc7\xc3O ESPECIAL":
				cnpj.DataSituacaoEspecial = strings.Trim(stringSlice[index+1], " ")
		}
	}

	return cnpj

}
