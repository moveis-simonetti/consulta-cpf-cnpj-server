<?php

require __DIR__ . '/vendor/autoload.php';

use Guzzle\Http\Client;

$host = 'http://localhost:3000';
$cpf = '02462208992';
$datnasc = '15111978';
$client = new Client;

//busca o captcha
$request = $client->get($host . '/captcha/cpf/' . $cpf);
$response = $request->send();
$captchaUrl = $host . "/" . $response->getBody();
echo $captchaUrl;

//busca os dados do cpf
//obs: se o CPF já foi pesquisado recentemente ele permanece no cache e pode ser acessado com qualquer capcha pois
//os dados vão ser recuperados do cache. O cache está configurado para salvar por 500 segundos.
$captcha = 'pZWp5c';
$request = $client->get($host . '/cpf/' . $cpf . '/' . $datnasc . '/' . $captcha);
$response = $request->send();
$cpf = json_decode($response->getBody());
var_dump($cpf);
//exemplo do retorno
// object(stdClass)#23 (6) {
//   ["Numero"]=>
//   string(14) "024.622.089-92"
//   ["Nome"]=>
//   string(18) "ELTON LUIS MINETTO"
//   ["DataNascimento"]=>
//   string(10) "15/11/1978"
//   ["Situacao"]=>
//   string(7) "REGULAR"
//   ["DataInscricao"]=>
//   string(10) "05/08/1996"
//   ["DigitoVerificador"]=>
//   string(2) "00"
// }
