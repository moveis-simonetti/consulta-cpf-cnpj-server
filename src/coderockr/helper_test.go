package coderockr

import (
	"testing"
)

type testpair struct {
	unformated string
	formated   string
}

type testCnpfPair struct {
	unformated string
	formated   CnpjData
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

func TestFormatCpfData(t *testing.T) {
	unformated := "No do CPF: 024.622.089-92\nNome da Pessoa F\u00edsica: ELTON LUIS MINETTO                                          \nData de Nascimento: 15/11/1978\nSitua\u00e7\u00e3o Cadastral: REGULAR\nData da Inscri\u00e7\u00e3o: 05/08/1996\nDigito Verificador: 00\n"
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

func TestFormatCnpjData(t *testing.T) {
	unformated := "||||||||||||||Comprovante|de|Inscri\u00e7\u00e3o|e|de|Situa\u00e7\u00e3o|Cadastral<br>|||Contribuinte,|||<br>Confira|os|dados|de|Identifica\u00e7\u00e3o|da|Pessoa|Jur\u00eddica|e,|se|houver|qualquer|diverg\u00eancia,||||providencie|junto|\u00e0|RFB|a|sua|atualiza\u00e7\u00e3o|cadastral.<br>||||||||||REP\u00daBLICA|FEDERATIVA|DO|BRASIL||||||||||<br>||||||||||CADASTRO|NACIONAL|DA|PESSOA|JUR\u00cdDICA||||||||||<br>|||N\u00daMERO|DE|INSCRI\u00c7\u00c3O|||<br>|||10.349.094/0001-62||||||MATRIZ||||||<br>||||||COMPROVANTE|DE|INSCRI\u00c7\u00c3O|E|DE|SITUA\u00c7\u00c3O|CADASTRAL|||||||<br>|||DATA|DE|ABERTURA|||<br>|||19/09/2008|||<br>|||NOME|EMPRESARIAL|||<br>|||CODEROCKR|DESENVOLVIMENTO|DE|PROGRAMAS|LTDA|-|ME|||<br>|||T\u00cdTULO|DO|ESTABELECIMENTO|(NOME|DE|FANTASIA)|||<br>||||CODEROCKR|&|CIA|LTDA|ME|||<br>|||C\u00d3DIGO|E|DESCRI\u00c7\u00c3O|DA|ATIVIDADE|ECON\u00d4MICA|PRINCIPAL|||<br>||||62.01-5-01|-|Desenvolvimento|de|programas|de|computador|sob|encomenda||||<br>|||C\u00d3DIGO|E|DESCRI\u00c7\u00c3O|DAS|ATIVIDADES|ECON\u00d4MICAS|SECUND\u00c1RIAS|||<br>|||||N\u00e3o|informada|||||<br>|||C\u00d3DIGO|E|DESCRI\u00c7\u00c3O|DA|NATUREZA|JUR\u00cdDICA|||<br>|||206-2|-|SOCIEDADE|EMPRESARIA|LIMITADA|||<br>|||LOGRADOURO|||<br>||||R|HENRIQUE|MEYER|||<br>|||N\u00daMERO|||<br>|||40|||<br>|||COMPLEMENTO|||<br>|||LOJA||1|||||||||||||||||||BOX|||1||||<br>|||CEP|||<br>||||89.201-405||||<br>|||BAIRRO/DISTRITO|||<br>|||CENTRO|||<br>|||MUNIC\u00cdPIO|||<br>|||JOINVILLE|||<br>|||UF|||<br>|||SC|||<br>|||ENDERE\u00c7O|ELETR\u00d4NICO|||<br>||||adeassisfrutuoso@yahoo.com.br||||<br>|||TELEFONE|||<br>|||||||(49)|3323-8205|||||<br>|||ENTE|FEDERATIVO|RESPONS\u00c1VEL|(EFR)|||<br>||||||*****||||||<br>|||SITUA\u00c7\u00c3O|CADASTRAL|||<br>||||ATIVA||||<br>|||DATA|DA|SITUA\u00c7\u00c3O|CADASTRAL|||<br>|||19/09/2008|||<br>|||MOTIVO|DE|SITUA\u00c7\u00c3O|CADASTRAL|||<br>||||||||<br>|||SITUA\u00c7\u00c3O|ESPECIAL|||<br>||||********||||<br>|||DATA|DA|SITUA\u00c7\u00c3O|ESPECIAL|||<br>||||********||||<br>||Aprovado|pela|Instru\u00e7\u00e3o|Normativa|RFB|n\u00ba|1.470,|de|30|de|maio|de|2014.|<br>||||||||||||Emitido|no|dia|19/02/2016|\u00e0s||16:36:16||(data|e|hora|de|Bras\u00edlia).|||||||||||||P\u00e1gina:|1/1||||||||<br>Emitido|no|dia|19/02/2016|\u00e0s||16:36:16||(data|e|hora|de|Bras\u00edlia).|<br>P\u00e1gina:|1/1<br>"
	expected := CnpjData{"10.349.094/0001-62 MATRIZ", "19/09/2008", "CODEROCKR DESENVOLVIMENTO DE PROGRAMAS LTDA - ME", "CODEROCKR & CIA LTDA ME", "62.01-5-01 - Desenvolvimento de programas de computador sob encomenda", "Não informada", "206-2 - SOCIEDADE EMPRESARIA LIMITADA", "R HENRIQUE MEYER", "40", "LOJA1 BOX 1", "89.201-405", "CENTRO", "JOINVILLE", "SC", "adeassisfrutuoso@yahoo.com.br", "(49) 3323-8205", "*****", "ATIVA", "19/09/2008", "", "********", "********"}
	v := FormatCnpjData(unformated)
	if v != expected {
		t.Error(
			"expected", expected,
			"got", v,
		)
	}

	unformated = "||||||||||||||Comprovante|de|Inscri\u00e7\u00e3o|e|de|Situa\u00e7\u00e3o|Cadastral<br>|||Contribuinte,|||<br>Confira|os|dados|de|Identifica\u00e7\u00e3o|da|Pessoa|Jur\u00eddica|e,|se|houver|qualquer|diverg\u00eancia,||||providencie|junto|\u00e0|RFB|a|sua|atualiza\u00e7\u00e3o|cadastral.<br>||||||||||REP\u00daBLICA|FEDERATIVA|DO|BRASIL||||||||||<br>||||||||||CADASTRO|NACIONAL|DA|PESSOA|JUR\u00cdDICA||||||||||<br>|||N\u00daMERO|DE|INSCRI\u00c7\u00c3O|||<br>|||21.101.726/0001-90||||||MATRIZ||||||<br>||||||COMPROVANTE|DE|INSCRI\u00c7\u00c3O|E|DE|SITUA\u00c7\u00c3O|CADASTRAL|||||||<br>|||DATA|DE|ABERTURA|||<br>|||23/09/2014|||<br>|||NOME|EMPRESARIAL|||<br>|||TAMARCADO|SERVICOS|DE|AGENDAMENTO|OLINE|LTDA|||<br>|||T\u00cdTULO|DO|ESTABELECIMENTO|(NOME|DE|FANTASIA)|||<br>||||********|||<br>|||C\u00d3DIGO|E|DESCRI\u00c7\u00c3O|DA|ATIVIDADE|ECON\u00d4MICA|PRINCIPAL|||<br>||||62.02-3-00|-|Desenvolvimento|e|licenciamento|de|programas|de|computador|customiz\u00e1veis||||<br>|||C\u00d3DIGO|E|DESCRI\u00c7\u00c3O|DAS|ATIVIDADES|ECON\u00d4MICAS|SECUND\u00c1RIAS|||<br>||||||74.90-1-04|-|Atividades|de|intermedia\u00e7\u00e3o|e|agenciamento|de|servi\u00e7os|e|neg\u00f3cios|em|geral,|exceto|imobili\u00e1rios||||||<br>|||C\u00d3DIGO|E|DESCRI\u00c7\u00c3O|DA|NATUREZA|JUR\u00cdDICA|||<br>|||206-2|-|SOCIEDADE|EMPRESARIA|LIMITADA|||<br>|||LOGRADOURO|||<br>||||R|BARAO|DO|RIO|BRANCO|||<br>|||N\u00daMERO|||<br>|||149|||<br>|||COMPLEMENTO|||<br>|||SALA||3||||<br>|||CEP|||<br>||||88.350-201||||<br>|||BAIRRO/DISTRITO|||<br>|||CENTRO|||<br>|||MUNIC\u00cdPIO|||<br>|||BRUSQUE|||<br>|||UF|||<br>|||SC|||<br>|||ENDERE\u00c7O|ELETR\u00d4NICO|||<br>||||||||<br>|||TELEFONE|||<br>|||||||(47)|3351-2107|||||<br>|||ENTE|FEDERATIVO|RESPONS\u00c1VEL|(EFR)|||<br>||||||*****||||||<br>|||SITUA\u00c7\u00c3O|CADASTRAL|||<br>||||ATIVA||||<br>|||DATA|DA|SITUA\u00c7\u00c3O|CADASTRAL|||<br>|||23/09/2014|||<br>|||MOTIVO|DE|SITUA\u00c7\u00c3O|CADASTRAL|||<br>||||||||<br>|||SITUA\u00c7\u00c3O|ESPECIAL|||<br>||||********||||<br>|||DATA|DA|SITUA\u00c7\u00c3O|ESPECIAL|||<br>||||********||||<br>||Aprovado|pela|Instru\u00e7\u00e3o|Normativa|RFB|n\u00ba|1.470,|de|30|de|maio|de|2014.|<br>||||||||||||Emitido|no|dia|19/02/2016|\u00e0s||16:10:34||(data|e|hora|de|Bras\u00edlia).|||||||||||||P\u00e1gina:|1/1||||||||<br>Emitido|no|dia|19/02/2016|\u00e0s||16:10:34||(data|e|hora|de|Bras\u00edlia).|<br>P\u00e1gina:|1/1<br>"
	v = FormatCnpjData(unformated)
	expected = CnpjData{"21.101.726/0001-90 MATRIZ", "23/09/2014", "TAMARCADO SERVICOS DE AGENDAMENTO OLINE LTDA", "********", "62.02-3-00 - Desenvolvimento e licenciamento de programas de computador customizáveis", "74.90-1-04 - Atividades de intermediação e agenciamento de serviços e negócios em geral, exceto imobiliários", "206-2 - SOCIEDADE EMPRESARIA LIMITADA", "R BARAO DO RIO BRANCO", "149", "SALA3", "88.350-201", "CENTRO", "BRUSQUE", "SC", "", "(47) 3351-2107", "*****", "ATIVA", "23/09/2014", "", "********", "********"}
	if v != expected {
		t.Error(
			"expected", expected,
			"got", v,
		)
	}
}
