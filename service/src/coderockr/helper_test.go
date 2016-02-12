package coderockr

import "testing"

type testpair struct {
	unformated string
	formated   string
}

var testsCpf = []testpair{
	{"teste", "teste"},
	{"02462208992", "024.622.089-92"},
	{"71558137459", "715.581.374-59"},
	{"634.651.328-15", "634.651.328-15"},
}

var testsCnpj = []testpair{
	{"teste", "teste"},
	{"10349094000162", "10.349.094/0001-62"},
	{"10.349.094/0001-62", "10.349.094/0001-62"},
	{"11731841000195", "11.731.841/0001-95"},
}

var testsData = []testpair{
	{"teste", "teste"},
	{"151178", "151178"},
	{"15111978", "15/11/1978"},
	{"15/11/1978", "15/11/1978"},
}

var testsCookie = []testpair{
	{"teste", "teste"},
	{"www.receita.fazenda.gov.br    FALSE   /   FALSE   0   ASPSESSIONIDAACRRRSC    OEFFCIBCEHMAEKAAMEKEAKBJ", "ASPSESSIONIDAACRRRSC=OEFFCIBCEHMAEKAAMEKEAKBJ"},
	{"www.receita.fazenda.gov.br	FALSE	/	FALSE	0	ASPSESSIONIDCSRSBDQQ	JMNGBIBCNDBMBMOAMHCMCGCC", "ASPSESSIONIDCSRSBDQQ=JMNGBIBCNDBMBMOAMHCMCGCC"},
}

func TestFormatCpf(t *testing.T) {
	for _, pair := range testsCpf {
		v := FormatCpf(pair.unformated)
		if v != pair.formated {
			t.Error(
				"For", pair.unformated,
				"expected", pair.formated,
				"got", v,
			)
		}
	}
}

func TestFormatCnpj(t *testing.T) {
	for _, pair := range testsCnpj {
		v := FormatCnpj(pair.unformated)
		if v != pair.formated {
			t.Error(
				"For", pair.unformated,
				"expected", pair.formated,
				"got", v,
			)
		}
	}
}

func TestFormatData(t *testing.T) {
	for _, pair := range testsData {
		v := FormatData(pair.unformated)
		if v != pair.formated {
			t.Error(
				"For", pair.unformated,
				"expected", pair.formated,
				"got", v,
			)
		}
	}
}

func TestFormatCookie(t *testing.T) {
	for _, pair := range testsCookie {
		v := FormatCookie(pair.unformated)
		if v != pair.formated {
			t.Error(
				"For", pair.unformated,
				"expected", pair.formated,
				"got", v,
			)
		}
	}
}

func TestFormatCpfData(t *testing.T ) {
	unformated := "No do CPF: 024.622.089-92\nNome da Pessoa F\xedsica: ELTON LUIS MINETTO                                          \nData de Nascimento: 15/11/1978\nSitua\xe7\xe3o Cadastral: REGULAR\nData da Inscri\xe7\xe3o: 05/08/1996\nDigito Verificador: 00\n"
	v := FormatCpfData(unformated)
	expected := CpfData{"024.622.089-92", "ELTON LUIS MINETTO", "15/11/1978", "REGULAR", "05/08/1996", "00"}
	if v != expected {
		t.Error(
			"For", unformated,
			"expected", expected,
			"got", v,
		)
	}
}
