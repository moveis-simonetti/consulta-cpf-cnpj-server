# Cliente PHP para consumir o serviço de consulta de CPF/CNPJ

## Instalar Composer

    php -r "readfile('https://getcomposer.org/installer');" > composer-setup.php
    php -r "if (hash('SHA384', file_get_contents('composer-setup.php')) === 'fd26ce67e3b237fffd5e5544b45b0d92c41a4afe3e3f778e942e43ce6be197b9cdc7c251dcde6e2a52297ea269370680') { echo 'Installer verified'; } else { echo 'Installer corrupt'; unlink('composer-setup.php'); }"
    php composer-setup.php
    php -r "unlink('composer-setup.php');"


## Instalar dependências

    php composer.phar install

## Exemplo de busca de CPF

### Primeiro passo - buscar o Captcha

    <?php

    require __DIR__ . '/vendor/autoload.php';

    use Guzzle\Http\Client;

    $host = 'http://localhost:3000';
    $cpf = '02462208992';
    $client = new Client;

    $request = $client->get($host . '/captcha/cpf/' . $cpf);
    $response = $request->send();
    $captchaUrl = $host . "/" . $response->getBody();
    echo $captchaUrl;

O exemplo acima receberá a url onde está armazenada a imagem gerada para o Captcha que deve ser respondido pelo usuário.


## Segundo passo - receber os dados do CPF

    <?php

    require __DIR__ . '/vendor/autoload.php';

    use Guzzle\Http\Client;

    $host = 'http://localhost:3000';
    $cpf = '02462208992';
    $datnasc = '15111978';
    $client = new Client;

    //neste exemplo vamos considerara que o usuário preencheu o captcha com o conteúdo abaixo
    $captcha = 'pZWp5c';
    $request = $client->get($host . '/cpf/' . $cpf . '/' . $datnasc . '/' . $captcha);
    $response = $request->send();
    $cpf = json_decode($response->getBody());


**Observações**:

    - se o CPF já foi pesquisado recentemente ele permanece no cache e pode ser acessado com qualquer captcha pois os dados vão ser recuperados do cache. O cache está configurado para salvar por 500 segundos e esta configuração pode ser alterada no serviço, no arquivo _main.go_
    - se não for encontrado ou o captcha estiver errado a request vai retornar com o status 404

##  Busca de CNPJ

### Primeiro passo - buscar o Captcha

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

O exemplo acima receberá a url onde está armazenada a imagem gerada para o Captcha que deve ser respondido pelo usuário.

## Segundo passo - receber os dados do CNPJ

    <?php

    require __DIR__ . '/vendor/autoload.php';

    use Guzzle\Http\Client;

    $host = 'http://localhost:3000';
    $cnpj = '10349094000162';
    $client = new Client;

    //neste exemplo vamos considerara que o usuário preencheu o captcha com o conteúdo abaixo
    $captcha = 'pZWp5c';
    $request = $client->get($host . '/cpf/' . $cpf . '/' . $datnasc . '/' . $captcha);
    $response = $request->send();
    $cpf = json_decode($response->getBody());

**Observações**:

    - se o CNPJ já foi pesquisado recentemente ele permanece no cache e pode ser acessado com qualquer captcha pois os dados vão ser recuperados do cache. O cache está configurado para salvar por 500 segundos e esta configuração pode ser alterada no serviço, no arquivo _main.go_
    - se não for encontrado ou o captcha estiver errado a request vai retornar com o status 404