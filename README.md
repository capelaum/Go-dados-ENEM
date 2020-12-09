
# Projeto LP-2020-UnB

Projeto consiste em analisar dados sobre o ENEM dos anos de 2017, 2018 e 2019, através dos microdados fornecidos pelo [inep](http://inep.gov.br/web/guest/microdados), os dados coletados propiciam as seguintes análises:

### Análise Anual

- Número de participantes 
- Médias das áreas de conhecimento 
- Médias das áreas de conhecimento de cada Raça
- Médias das áreas de conhecimento de cada Tipo de Escola

### Análise Estadual

- Número de participantes 
- Médias das áreas de conhecimento 
- Médias das áreas de conhecimento de cada Raça
- Número de participantes de cada Raça 
- Número de participantes de Escola Pública e Privada
- Número de participantes de Escola Pública e Privada de cada Raça

Objetivo é demonstrar os diferenciais e pontos fortes da linguagem Golang, 
ao realizar a leitura dos microdados e armazenamento dos dados obtidos utilizando 
processos de execuções paralelas, buscando obter um menor tempo de execução e melhor performance.

## Pré-requisitos

Para executar o trabalho é preciso baixar as bases de dados referentes aos anos do ENEM de 2017,2018 e 2019 em http://inep.gov.br/web/guest/microdados e descompactar os arquivos zip na pasta raiz do diretório do projeto. Cada diretório referente a um ano deverá ser renomeado para o formato: **microdados_enem_2017**.

Baixar pacotes externos utilizados atravé dos comandos:

```bash
go get golang.org/x/text/encoding/charmap

go get github.com/montanaflynn/stats
```

Em seguida, todos os pacotes do diretório `./Packages` (opcional) do projeto podem ser movidos para o `$GOPATH` local com o seguinte comando:

```
env GIT_TERMINAL_PROMPT=1 go get github.com/munak98/LP-2020/Packages/extract
```

## Rodando o trabalho

Dentro da pasta Src, o trabalho pode ser executado com o comando:

```
go run main.go 
```
