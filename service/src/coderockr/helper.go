package coderockr

import "strings"

type CpfData struct {
    Numero string
    Nome string
    DataNascimento string
    Situacao string
    DataInscricao string
    DigitoVerificador string
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
	if (len(stringSlice) > 5) {
		return stringSlice[len(stringSlice)-5] + "=" + stringSlice[len(stringSlice)-1]
	}

	stringSlice = strings.Split(cookie, "\t")
	if (len(stringSlice) > 1) {
		return stringSlice[len(stringSlice)-2] + "=" + stringSlice[len(stringSlice)-1]
	}

	return cookie
}

func FormatCpfData(unformated string) CpfData {
	stringSlice := strings.Split(unformated, "\n")
	cpf := CpfData{};
	for _, text := range stringSlice {
		data := strings.Split(text, ":")
		if (len(data) == 2) {
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