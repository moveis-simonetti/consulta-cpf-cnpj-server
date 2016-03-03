package coderockr

import (
	"strings"
)

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
			case "Nome da Pessoa Física":
				cpf.Nome = strings.Trim(data[1], " ")
			case "Data de Nascimento":
				cpf.DataNascimento = strings.Trim(data[1], " ")
			case "Situação Cadastral":
				cpf.Situacao = strings.Trim(data[1], " ")
			case "Data da Inscrição":
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
		switch text {
		case "NÚMERO DE INSCRIÇÃO":
			value := strings.Trim(stringSlice[index+1], " ")
			cnpj.NumeroInscricao = value[0:18] + " " + value[18:]
		case "DATA DE ABERTURA":
			cnpj.DataAbertura = strings.Trim(stringSlice[index+1], " ")
		case "NOME EMPRESARIAL":
			cnpj.NomeEmpresarial = strings.Trim(stringSlice[index+1], " ")
		case "TÍTULO DO ESTABELECIMENTO (NOME DE FANTASIA)":
			cnpj.NomeFantasia = strings.Trim(stringSlice[index+1], " ")
		case "CÓDIGO E DESCRIÇÃO DA ATIVIDADE ECONÔMICA PRINCIPAL":
			cnpj.AtividadeEconomicaPrincipal = strings.Trim(stringSlice[index+1], " ")
		case "CÓDIGO E DESCRIÇÃO DAS ATIVIDADES ECONÔMICAS SECUNDÁRIAS":
			value := strings.Trim(stringSlice[index+1], " ")
			cnpj.AtividadeEconomicaSecundaria = value
		case "CÓDIGO E DESCRIÇÃO DA NATUREZA JURÍDICA":
			cnpj.NaturezaJuridica = strings.Trim(stringSlice[index+1], " ")
		case "LOGRADOURO":
			cnpj.Logradouro = strings.Trim(stringSlice[index+1], " ")
		case "NÚMERO":
			cnpj.Numero = strings.Trim(stringSlice[index+1], " ")
		case "COMPLEMENTO":
			cnpj.Complemento = strings.Trim(stringSlice[index+1], " ")
		case "CEP":
			cnpj.Cep = strings.Trim(stringSlice[index+1], " ")
		case "BAIRRO/DISTRITO":
			cnpj.Bairro = strings.Trim(stringSlice[index+1], " ")
		case "MUNICÍPIO":
			cnpj.Municipio = strings.Trim(stringSlice[index+1], " ")
		case "UF":
			cnpj.Uf = strings.Trim(stringSlice[index+1], " ")
		case "ENDEREÇO ELETRÔNICO":
			cnpj.EnderecoEletronico = strings.Trim(stringSlice[index+1], " ")
		case "TELEFONE":
			cnpj.Telefone = strings.Trim(stringSlice[index+1], " ")
		case "ENTE FEDERATIVO RESPONSÁVEL (EFR)":
			cnpj.EnteFederativoResponsavel = strings.Trim(stringSlice[index+1], " ")
		case "SITUAÇÃO CADASTRAL":
			cnpj.Situacao = strings.Trim(stringSlice[index+1], " ")
		case "DATA DA SITUAÇÃO CADASTRAL":
			cnpj.DataSituacao = strings.Trim(stringSlice[index+1], " ")
		case "MOTIVO DE SITUAÇÃO CADASTRAL":
			cnpj.MotivoSituacao = strings.Trim(stringSlice[index+1], " ")
		case "SITUAÇÃO ESPECIAL":
			cnpj.SituacaoEspecial = strings.Trim(stringSlice[index+1], " ")
		case "DATA DA SITUAÇÃO ESPECIAL":
			cnpj.DataSituacaoEspecial = strings.Trim(stringSlice[index+1], " ")
		}
	}

	return cnpj

}
