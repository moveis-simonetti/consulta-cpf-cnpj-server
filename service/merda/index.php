<?php
require_once __DIR__ . '/vendor/autoload.php';

use GuzzleHttp\Cookie\CookieJar;
use GuzzleHttp\Client;
$client = new Client(
    [
        'base_uri' => 'http://www.receita.fazenda.gov.br/aplicacoes/atcta/cpf/',
        'connect_timeout' => 2,
    ]
);

// $res = $client->get('captcha/gerarCaptcha.asp');
// var_dump($res->getHeader('Set-Cookie')[0]);
// var_dump(file_put_contents('captcha2.png', $res->getBody()));

$res = $client->post('ConsultaPublicaExibir.asp',[
    //'cookies' => $jar,
    'headers' => [
        'Cookie' => 'ASPSESSIONIDSSQTBATT=DJANDMLBJPEGOIGKHNIFKMGO',
        'refer' => 'http://www.receita.fazenda.gov.br/aplicacoes/atcta/cpf/consultapublica.asp',
    ],
    'form_params' => [
        'tempTxtCPF' => '144.840.167-45',
        'tempTxtNascimento' => '10/05/1993',
        'temptxtTexto_captcha_serpro_gov_br' => 'BQYmCP',
        'txtTexto_captcha_serpro_gov_br' => 'BQYmCP',
        'Enviar' => 'Consultar'
    ]
]);
echo (string)$res->getBody();
file_put_contents('retorno.html', (string)$res->getBody());
