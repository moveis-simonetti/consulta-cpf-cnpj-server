# Micro-serviço que forece consulta de CPF/CNPJ

Projeto em Golang para buscar dados de CPF/CNPJ e armazenar em cache 

# Compilar

O primeiro passo é ter o Go instalado. No site oficial é possível encontrar binários para os principais sistemas operacionais

    https://golang.org

Além de ter o Go instalado no sistema operacional é necessário executar:

    export GOPATH=/path/consulta-cpf-cnpj/service
    go get github.com/go-martini/martini
    go get github.com/andelf/go-curl
    go get github.com/ryanuber/go-filecache
    go get github.com/PuerkitoBio/goquery
    go build

###Observação

Este projeto foi testado com o Golang versão 1.5.2


# Executar

O binário chamado service será criado. Basta executá-lo e ele ficará ouvindo na porta 3000 por novas requisições

# Uso

Basta acessar a URL como nos exemplos abaixo

    http://localhost:3000/captcha/cpf/CPF

Exemplo:

    http://localhost:3000/captcha/cpf/02462208992
    

O retorno será um JSON com o conteúdo do Captcha a ser preenchido pelo usuário. O usuário deve preencher o captcha e a informação deve ser enviada para 

    http://localhost:3000/cpf/CPF/DATANASC/CAPTCHA

Exemplo:

    http://localhost:3000/cpf/02462208992/15111978/r5HMnx
