# Micro-serviço que forece consulta de CPF/CNPJ

Projeto em Golang para buscar dados de CPF/CNPJ e armazenar em cache 

# Compilar

Além de ter o Go instalado no sistema operacional é necessário executar:

    export GOPATH=/path/consulta-cpf-cnpj/service
    go get github.com/go-martini/martini
    go get github.com/andelf/go-curl
    go get github.com/ryanuber/go-filecache
    go build

###Observação

Este projeto foi testado com o Golang versão 1.5.2


# Executar

O binário chamado service será criado. Basta executá-lo e ele ficará ouvindo na porta 3000 por novas requisições

# Uso

Basta acessar a URL como nos exemplos abaixo

    http://localhost:3000/cpf/02462208992
    http://localhost:3000/cnpj/10349094000162

O retorno será um JSON com o conteúdo do Captcha a ser preenchido pelo usuário. O usuário deve preencher o captcha e a informação deve ser enviada para 

    http://localhost:3000/cpf-detail/02462208992/captcha
    http://localhost:3000/cnpj-detail/10349094000162/captcha