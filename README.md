# Micro-serviço que forece consulta de CPF/CNPJ

Projeto em Golang para buscar dados de CPF/CNPJ e armazenar em cache 

# Compilar

O primeiro passo é ter o Go instalado. No site oficial é possível encontrar binários para os principais sistemas operacionais

    https://golang.org

Além de ter o Go instalado no sistema operacional é necessário executar:

    export GOPATH=/path/consulta-cpf-cnpj-server
    go get github.com/PuerkitoBio/goquery
    go get github.com/andelf/go-curl
    go get github.com/go-martini/martini
    go get github.com/ryanuber/go-filecache
    go get github.com/andelf/iconv-go
    go build

###Observação

Este projeto foi testado com as versões 1.5.2 e 1.5.3 do Go


# Executar

O binário chamado service será criado. Basta executá-lo e ele ficará ouvindo na porta 3000 por novas requisições

    /path/consulta-cpf-cnpj-server/consulta-cpf-cnpj-server &

**Observação:**  o & no final do comando acima faz com que o programa fique executando em "background" em sistemas Linux e MacOSX. Desta forma o aplicativo continua rodando mesmo quando o terminal onde foi executado for fechado. Para parar o aplicativo basta fazer o kill do mesmo com o comando:

    kill -9 `ps aux | grep consulta-cpf-cnpj-server/consulta-cpf-cnpj-server`

# Usando o Docker

Uma alternativa é usar o Docker para executar o servidor. 
O primeiro passo é instalar o _docker_ e o _docker compose_ usando os pacotes do sistema operacional que irá rodar o serviço (Linux, Windows ou Mac OSX). 

Devido ao fato da imagem ser privada é necessário realizar primeiro o login no Docker Hub com algum usuário que tenha permissão de acesso a imagem. Os comandos para realizar o login e inicializar a imagem são:

  docker login -e email@domain.com -p senha
  docker-compose up -d

O primeiro comando é necessário apenas na primeira execução. 

Para finalizar o serviço basta executar

  docker-compose stop

# Uso

Basta acessar a URL como nos exemplos abaixo

    http://localhost:3000/captcha/cpf/CPF

Exemplo:

    http://localhost:3000/captcha/cpf/02462208992
    

O retorno será um JSON com a url da imagem com o conteúdo do Captcha a ser preenchido pelo usuário. O usuário deve preencher o captcha e a informação deve ser enviada para 

    http://localhost:3000/cpf/CPF/DATANASC/CAPTCHA

Exemplo:

    http://localhost:3000/cpf/02462208992/15111978/r5HMnx

O resultado será um arquivo JSON como o exemplo abaixo:

    
    {
      "Numero": "024.622.089-92",
      "Nome": "ELTON LUIS MINETTO",
      "DataNascimento": "15/11/1978",
      "Situacao": "REGULAR",
      "DataInscricao": "05/08/1996",
      "DigitoVerificador": "00"
    }

O mesmo para o CNPJ:

    http://localhost:3000/captcha/cnpj/10349094000162
    http://localhost:3000/cnpj/10349094000162/r5HMnx

**Observação:** para a consulta do CNPJ basta o CNPJ e o valor do captcha

O resultado será um arquivo JSON como o exemplo abaixo:


    {
      "NumeroInscricao": "10.349.094/0001-62 MATRIZ",
      "DataAbertura": "19/09/2008",
      "NomeEmpresarial": "CODEROCKR DESENVOLVIMENTO DE PROGRAMAS LTDA - ME",
      "NomeFantasia": "CODEROCKR & CIA LTDA ME",
      "AtividadeEconomicaPrincipal": "62.01-5-01 - Desenvolvimento de programas de computador sob encomenda",
      "AtividadeEconomicaSecundaria": "Não informada",
      "NaturezaJuridica": "206-2 - SOCIEDADE EMPRESARIA LIMITADA",
      "Logradouro": "R HENRIQUE MEYER",
      "Numero": "40",
      "Complemento": "LOJA1 BOX 1",
      "Cep": "89.201-405",
      "Bairro": "CENTRO",
      "Municipio": "JOINVILLE",
      "Uf": "SC",
      "EnderecoEletronico": "adeassisfrutuoso@yahoo.com.br",
      "Telefone": "(49) 3323-8205",
      "EnteFederativoResponsavel": "*****",
      "Situacao": "ATIVA",
      "DataSituacao": "19/09/2008",
      "MotivoSituacao": "",
      "SituacaoEspecial": "********",
      "DataSituacaoEspecial": "********"
    }

**Observação:** Caso o CPF/CNPJ não for encontrado, ou o Captcha for digitado errado o retorno será uma resposta HTTP com o código 404.
