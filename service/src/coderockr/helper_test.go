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

func TestFormatCpfData(t *testing.T) {
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

func TestFormatCnpjData(t *testing.T) {
	unformated := "||||||||||||||Comprovante|de|Inscri\xe7\xe3o|e|de|Situa\xe7\xe3o|Cadastral<br>|||Contribuinte,|||<br>Confira|os|dados|de|Identifica\xe7\xe3o|da|Pessoa|Jur\xeddica|e,|se|houver|qualquer|diverg\xeancia,||||providencie|junto|\xe0|RFB|a|sua|atualiza\xe7\xe3o|cadastral.<br>||||||||||REP\xdaBLICA|FEDERATIVA|DO|BRASIL||||||||||<br>||||||||||CADASTRO|NACIONAL|DA|PESSOA|JUR\xcdDICA||||||||||<br>|||N\xdaMERO|DE|INSCRI\xc7\xc3O|||<br>|||10.349.094/0001-62||||||MATRIZ||||||<br>||||||COMPROVANTE|DE|INSCRI\xc7\xc3O|E|DE|SITUA\xc7\xc3O|CADASTRAL|||||||<br>|||DATA|DE|ABERTURA|||<br>|||19/09/2008|||<br>|||NOME|EMPRESARIAL|||<br>|||CODEROCKR|DESENVOLVIMENTO|DE|PROGRAMAS|LTDA|-|ME|||<br>|||T\xcdTULO|DO|ESTABELECIMENTO|(NOME|DE|FANTASIA)|||<br>||||CODEROCKR|&|CIA|LTDA|ME|||<br>|||C\xd3DIGO|E|DESCRI\xc7\xc3O|DA|ATIVIDADE|ECON\xd4MICA|PRINCIPAL|||<br>||||62.01-5-01|-|Desenvolvimento|de|programas|de|computador|sob|encomenda||||<br>|||C\xd3DIGO|E|DESCRI\xc7\xc3O|DAS|ATIVIDADES|ECON\xd4MICAS|SECUND\xc1RIAS|||<br>|||||N\xe3o|informada|||||<br>|||C\xd3DIGO|E|DESCRI\xc7\xc3O|DA|NATUREZA|JUR\xcdDICA|||<br>|||206-2|-|SOCIEDADE|EMPRESARIA|LIMITADA|||<br>|||LOGRADOURO|||<br>||||R|HENRIQUE|MEYER|||<br>|||N\xdaMERO|||<br>|||40|||<br>|||COMPLEMENTO|||<br>|||LOJA||1|||||||||||||||||||BOX|||1||||<br>|||CEP|||<br>||||89.201-405||||<br>|||BAIRRO/DISTRITO|||<br>|||CENTRO|||<br>|||MUNIC\xcdPIO|||<br>|||JOINVILLE|||<br>|||UF|||<br>|||SC|||<br>|||ENDERE\xc7O|ELETR\xd4NICO|||<br>||||adeassisfrutuoso@yahoo.com.br||||<br>|||TELEFONE|||<br>|||||||(49)|3323-8205|||||<br>|||ENTE|FEDERATIVO|RESPONS\xc1VEL|(EFR)|||<br>||||||*****||||||<br>|||SITUA\xc7\xc3O|CADASTRAL|||<br>||||ATIVA||||<br>|||DATA|DA|SITUA\xc7\xc3O|CADASTRAL|||<br>|||19/09/2008|||<br>|||MOTIVO|DE|SITUA\xc7\xc3O|CADASTRAL|||<br>||||||||<br>|||SITUA\xc7\xc3O|ESPECIAL|||<br>||||********||||<br>|||DATA|DA|SITUA\xc7\xc3O|ESPECIAL|||<br>||||********||||<br>||Aprovado|pela|Instru\xe7\xe3o|Normativa|RFB|n\xba|1.470,|de|30|de|maio|de|2014.|<br>||||||||||||Emitido|no|dia|16/02/2016|\xe0s||15:54:21||(data|e|hora|de|Bras\xedlia).|||||||||||||P\xe1gina:|1/1||||||||<br>Emitido|no|dia|16/02/2016|\xe0s||15:54:21||(data|e|hora|de|Bras\xedlia).|<br>P\xe1gina:|1/1<br>"
	v := FormatCnpjData(unformated)
	expected := CnpjData{"10.349.094/0001-62 MATRIZ", "19/09/2008", "CODEROCKR DESENVOLVIMENTO DE PROGRAMAS LTDA - ME", "CODEROCKR & CIA LTDA ME", "62.01-5-01 - Desenvolvimento de programas de computador sob encomenda", "Não informada", "206-2 - SOCIEDADE EMPRESARIA LIMITADA", "R HENRIQUE MEYER", "40", "LOJA1 BOX 1", "89.201-405", "CENTRO", "JOINVILLE", "SC", "adeassisfrutuoso@yahoo.com.br", "(49) 3323-8205", "*****", "ATIVA", "19/09/2008", "", "********", "********"}
	if v.NumeroInscricao != expected.NumeroInscricao {
		t.Error(
			"expected", expected.NumeroInscricao,
			"got", v.NumeroInscricao,
		)
	}
	if v.DataAbertura != expected.DataAbertura {
		t.Error(
			"expected", expected.DataAbertura,
			"got", v.DataAbertura,
		)
	}
	if v.NomeEmpresarial != expected.NomeEmpresarial {
		t.Error(
			"expected", expected.NomeEmpresarial,
			"got", v.NomeEmpresarial,
		)
	}
	if v.NomeFantasia != expected.NomeFantasia {
		t.Error(
			"expected", expected.NomeFantasia,
			"got", v.NomeFantasia,
		)
	}
	if v.AtividadeEconomicaPrincipal != expected.AtividadeEconomicaPrincipal {
		t.Error(
			"expected", expected.AtividadeEconomicaPrincipal,
			"got", v.AtividadeEconomicaPrincipal,
		)
	}
	if v.AtividadeEconomicaSecundaria != expected.AtividadeEconomicaSecundaria {
		t.Error(
			"expected", expected.AtividadeEconomicaSecundaria,
			"got", v.AtividadeEconomicaSecundaria,
		)
	}
	if v.NaturezaJuridica != expected.NaturezaJuridica {
		t.Error(
			"expected", expected.NaturezaJuridica,
			"got", v.NaturezaJuridica,
		)
	}
	if v.Logradouro != expected.Logradouro {
		t.Error(
			"expected", expected.Logradouro,
			"got", v.Logradouro,
		)
	}
	if v.Numero != expected.Numero {
		t.Error(
			"expected", expected.Numero,
			"got", v.Numero,
		)
	}
	if v.Complemento != expected.Complemento {
		t.Error(
			"expected", expected.Complemento,
			"got", v.Complemento,
		)
	}
	if v.Cep != expected.Cep {
		t.Error(
			"expected", expected.Cep,
			"got", v.Cep,
		)
	}
	if v.Bairro != expected.Bairro {
		t.Error(
			"expected", expected.Bairro,
			"got", v.Bairro,
		)
	}
	if v.Municipio != expected.Municipio {
		t.Error(
			"expected", expected.Municipio,
			"got", v.Municipio,
		)
	}
	if v.Uf != expected.Uf {
		t.Error(
			"expected", expected.Uf,
			"got", v.Uf,
		)
	}
	if v.EnderecoEletronico != expected.EnderecoEletronico {
		t.Error(
			"expected", expected.EnderecoEletronico,
			"got", v.EnderecoEletronico,
		)
	}
	if v.Telefone != expected.Telefone {
		t.Error(
			"expected", expected.Telefone,
			"got", v.Telefone,
		)
	}
	if v.EnteFederativoResponsavel != expected.EnteFederativoResponsavel {
		t.Error(
			"expected", expected.EnteFederativoResponsavel,
			"got", v.EnteFederativoResponsavel,
		)
	}
	if v.Situacao != expected.Situacao {
		t.Error(
			"expected", expected.Situacao,
			"got", v.Situacao,
		)
	}
	if v.DataSituacao != expected.DataSituacao {
		t.Error(
			"expected", expected.DataSituacao,
			"got", v.DataSituacao,
		)
	}
	if v.MotivoSituacao != expected.MotivoSituacao {
		t.Error(
			"expected", expected.MotivoSituacao,
			"got", v.MotivoSituacao,
		)
	}
	if v.SituacaoEspecial != expected.SituacaoEspecial {
		t.Error(
			"expected", expected.SituacaoEspecial,
			"got", v.SituacaoEspecial,
		)
	}
	if v.DataSituacaoEspecial != expected.DataSituacaoEspecial {
		t.Error(
			"expected", expected.DataSituacaoEspecial,
			"got", v.DataSituacaoEspecial,
		)
	}
}
