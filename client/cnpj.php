<?php

require __DIR__ . '/vendor/autoload.php';

use Guzzle\Http\Client;

$host = 'http://localhost:3000';
$cnpj = '10349094000162';
$client = new Client;

//busca o captcha
$request = $client->get($host . '/captcha/cnpj/' . $cnpj);
$response = $request->send();
$captchaUrl = $host . "/" . $response->getBody();
echo $captchaUrl;

//busca os dados do cnpj
//obs: se o CNPJ já foi pesquisado recentemente ele permanece no cache e pode ser acessado com qualquer capcha pois
//os dados vão ser recuperados do cache. O cache está configurado para salvar por 500 segundos.
$captcha = 'JSctRq';
$request = $client->get($host . '/cnpj/' . $cnpj . '/' . $captcha);
$response = $request->send();
$cnpj = json_decode($response->getBody());
//obs: se não for encontrado ou o captcha estiver errado a request vai retornar com o status 404
var_dump($cnpj);
//exemplo do retorno
// object(stdClass)#23 (22) {
//   ["NumeroInscricao"]=>
//   string(25) "10.349.094/0001-62 MATRIZ"
//   ["DataAbertura"]=>
//   string(10) "19/09/2008"
//   ["NomeEmpresarial"]=>
//   string(48) "CODEROCKR DESENVOLVIMENTO DE PROGRAMAS LTDA - ME"
//   ["NomeFantasia"]=>
//   string(23) "CODEROCKR & CIA LTDA ME"
//   ["AtividadeEconomicaPrincipal"]=>
//   string(69) "62.01-5-01 - Desenvolvimento de programas de computador sob encomenda"
//   ["AtividadeEconomicaSecundaria"]=>
//   string(14) "Não informada"
//   ["NaturezaJuridica"]=>
//   string(37) "206-2 - SOCIEDADE EMPRESARIA LIMITADA"
//   ["Logradouro"]=>
//   string(16) "R HENRIQUE MEYER"
//   ["Numero"]=>
//   string(2) "40"
//   ["Complemento"]=>
//   string(11) "LOJA1 BOX 1"
//   ["Cep"]=>
//   string(10) "89.201-405"
//   ["Bairro"]=>
//   string(6) "CENTRO"
//   ["Municipio"]=>
//   string(9) "JOINVILLE"
//   ["Uf"]=>
//   string(2) "SC"
//   ["EnderecoEletronico"]=>
//   string(29) "adeassisfrutuoso@yahoo.com.br"
//   ["Telefone"]=>
//   string(14) "(49) 3323-8205"
//   ["EnteFederativoResponsavel"]=>
//   string(5) "*****"
//   ["Situacao"]=>
//   string(5) "ATIVA"
//   ["DataSituacao"]=>
//   string(10) "19/09/2008"
//   ["MotivoSituacao"]=>
//   string(0) ""
//   ["SituacaoEspecial"]=>
//   string(8) "********"
//   ["DataSituacaoEspecial"]=>
//   string(8) "********"
// }
