package main

import (
    "golang.org/x/net/publicsuffix"
    "io/ioutil"
    "log"
    "net/http"
    "net/http/cookiejar"
)

func main() {
    options := cookiejar.Options{
        PublicSuffixList: publicsuffix.List,
    }
    jar, err := cookiejar.New(&options)
    if err != nil {
        log.Fatal(err)
    }
    client := http.Client{Jar: jar}
    resp, err := client.Get("http://www.receita.fazenda.gov.br/aplicacoes/atcta/cpf/consultapublica.asp")
    if err != nil {
        log.Fatal(err)
    }
    data, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    log.Println(string(data))
}