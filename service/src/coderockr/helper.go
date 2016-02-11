package coderockr

import "strings"

func FormatCpf(cpf string) string {
	if len(cpf) < 11 || len(cpf) == 14 {
		return cpf
	}

	return cpf[0:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:11]
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
