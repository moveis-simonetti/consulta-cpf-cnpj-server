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
