# 🛠 Manager-API
> Introdução do projeto

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=gameficame_nfe-vault-api&metric=alert_status&token=2142ceae3ff67ff7d5774306f75fb381165aa62f)](https://sonarcloud.io/summary/new_code?id=gameficame_nfe-vault-api)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=gameficame_nfe-vault-api&metric=bugs&token=2142ceae3ff67ff7d5774306f75fb381165aa62f)](https://sonarcloud.io/summary/new_code?id=gameficame_nfe-vault-api)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=gameficame_nfe-vault-api&metric=duplicated_lines_density&token=2142ceae3ff67ff7d5774306f75fb381165aa62f)](https://sonarcloud.io/summary/new_code?id=gameficame_nfe-vault-api)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=gameficame_nfe-vault-api&metric=sqale_index&token=2142ceae3ff67ff7d5774306f75fb381165aa62f)](https://sonarcloud.io/summary/new_code?id=gameficame_nfe-vault-api)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=gameficame_nfe-vault-api&metric=coverage&token=2142ceae3ff67ff7d5774306f75fb381165aa62f)](https://sonarcloud.io/summary/new_code?id=gameficame_nfe-vault-api)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=gameficame_nfe-vault-api&metric=ncloc&token=2142ceae3ff67ff7d5774306f75fb381165aa62f)](https://sonarcloud.io/summary/new_code?id=gameficame_nfe-vault-api)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=gameficame_nfe-vault-api&metric=code_smells&token=2142ceae3ff67ff7d5774306f75fb381165aa62f)](https://sonarcloud.io/summary/new_code?id=gameficame_nfe-vault-api)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=gameficame_nfe-vault-api&metric=security_rating&token=2142ceae3ff67ff7d5774306f75fb381165aa62f)](https://sonarcloud.io/summary/new_code?id=gameficame_nfe-vault-api)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=gameficame_nfe-vault-api&metric=vulnerabilities&token=2142ceae3ff67ff7d5774306f75fb381165aa62f)](https://sonarcloud.io/summary/new_code?id=gameficame_nfe-vault-api)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=gameficame_nfe-vault-api&metric=reliability_rating&token=2142ceae3ff67ff7d5774306f75fb381165aa62f)](https://sonarcloud.io/summary/new_code?id=gameficame_nfe-vault-api)

## Sumário da documentação
- [Funcionalidade](#funcionalidade)
- [Pré-requisitos](#pré-requisitos)
- [Configuração de ambiente](#configuração-de-ambiente)
- [Estrutura do projeto](#estrutura-do-projeto)
- [Como Testar](#como-testar)
- [Makefile](#makefile)

## Funcionalidade
Manager-API o projeto backend do Gestor na plataforma Gamefica

## Pré-requisitos
Antes de iniciar, certifique-se de ter os seguintes requisitos instalados em seu sistema:

#### `1. Configurando variável de ambiente:`
- Em seu sistema operacional, crie uma variável de ambiente chamada `GOPRIVATE` com o valor `github.com/gameficame`. Essa variável irá conter as dependências e pacotes que vamos necessitar ao trabalhar com o projeto manager-API. 

#### `2. Go (Linguagem de programação):`
 - O projeto está desenvolvido em Go. Para instalar a linguagem em sua máquina, acesse https://go.dev/doc/install e instale a versão mais atualizada da linguagem. Para verificar a versão e se a instalação ocorreram de forma correta, abra um terminal de comando e rode o comando `go version`.

#### `3. Clone o repositório do projeto em sua máquina local:`

```bash
git clone https://github.com/gameficame/manager-api.git
```


#### `3. Dependências do projeto:`
 - Também será necessário clonar o repositório com as dependências do projeto manager-api: 

```bash
git clone https://github.com/gameficame/go.git
git clone https://github.com/gameficame/manager-sdk.git
git clone https://github.com/gameficame/dyna-orm.git
```
 - ⚠️ *É importante que os diretórios com todos os projetos estejam unificados em um único diretório.*

- Pra que servem esses pacotes: 
- `Go`: É o repositório com os pacotes e bibliotecas utilizados para o desenvolvimento do manager-api.
- `Manager-sdk`: Pacote com as dependências utilizadas pelo 
server gRPC do Organization.
- `Dyna-orm`: Pacote integrado com a API ORM do DynamoDB.
- Após a clonagem de todos os repositórios, rode o comando `go mod tidy` em cada um deles para que sejam feitas as instalações dos pacotes necessários.

#### `4. Insomnium:`
 - Para testar os endpoints da nossa API em ambiente local e de produção, utilizamos o Insomnium. Para instalar, acesse https://github.com/ArchGPT/insomnium/releases. 
 - ⚠️ *Importante ser pelo menos a versão 0.2.3-a, a qual da suporte ao git sync.*

## Configuração de ambiente
Siga as etapas abaixo para configurar o ambiente de desenvolvimento:

#### `1. Configuração do Insomnium:`
 - Será necessário fazer a configuração do Insomnia para que as nossas branches estejam sincronizadas no Github. O link a seguir contém um tutorial passo a passo de como configurar o ambiente de branches e os plugins necessários no Insomnium: https://coda.io/d/Configurar-o-Insomnium_d9Ki6AdbKGg/Configurar-Insomnium_suI8j#_luTee

#### `2. Criação de profile AWS:`
 - Após configurar o Insomnium, será necessária a criação de um profile AWS, para que o projeto possa ser rodado na máquina local. O link a seguir contém um tutorial passo a passo de como configurar um profile AWS: https://coda.io/d/Creating-AWS-Profile_d4FmJOS1MG3/Creating-AWS-Profile_suHWD#_lu0dN

#### `3. Criação de arquivo ENV:`
 - Após a criação do profile AWS, será necessária a criação do arquivo env.yml no projeto, com a credencial da aws para que possam ser realizadas operações no Dynamo. Para isso, entre no arquivo `env-default.yml`, e o copie por inteiro. Após isso, crie um novo arquivo chamado `env.yml` dentro do diretório `configs`, e cole o conteúdo copiado do `env-default.yml`. Após isso, substitua o profile aws pelo que foi criado no momento em que foi criado o profile aws. 

 ```bash
aws:
  profile:
    dev: adicione-o-perfil-criado
```
- Após isso, adicione o arquivo `env.yml` ao arquivo `.gitignore`, para que ele fique apenas em sua máquina local.

## Estrutura do projeto
A estrutura do projeto deve seguir o layout recomendado pelo [github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout/tree/master).
```
|── build
│   └── Dockerfile
├── cmd
│   ├── manager-api
│       └── manager-api.go
│   └── statics-gen
│       └── statics-gen.go
├── configs
|   └── env-default.yml
├── internal
│   ├── consumer
│       ├── account.go
│       └── account_test.go
│   ├── consumer
│       ├── account_subdomain.go
│       ├── account.go
│       ├── database.go
│       ├── org_cnpj.go
│       ├── organization.go
│       ├── store_ein.go
│       ├── store.go
│       └── user.go
|   ├── grpc
│       ├── organization
│       |   └── organization.go
│       └── store
│           └── logo.go
│   ├── producer
│       ├── account.go
│       └── producer.go
│   ├── routes
│       ├── middleware
│       |   └── user.go
│       ├── initial_data.go
│       ├── routes.go
│       ├── stores.go
│       └── upload.go
|   └── service
│       ├── account
│       |   ├── account_test.go
│       |   └── account.go
│       ├── appclient
│       |   ├── appclient_test.go
│       |   └── appclient.go
│       ├── distribution
│       |   ├── distribution_test.go
│       |   └── distribution.go
│       ├── organization
│       |   ├── organization_test.go
│       |   └── organization.go
│       ├── provisioning
│       |   ├── provisioning_test.go
│       |   └── provisioning.go
│       ├── store
│       |   ├── store_test.go
│       |   └── store.go
│       ├── subdomain
│       |   ├── subdomain_test.go
│       |   └── subdomain.go
│       ├── upload
│       |   └── authorize.go
│       ├── user
│       |   └── user.go
│       ├── userpool
│       |   ├── userpool_test.go
│       |   └── userpool.go
│       └── service.go
├── pkg
│   ├── db
│       ├── str
│       |    └── model.go
|       └── model.go
│   ├── sqs
│       └── type.go
│   └── utils
│       ├── dynamo_serializer_test.go
│       └── dynamo_serializer.go
├── go.mod
├── go.sum
├── Makefile
├── README.md
```

- `Makefile`: Este arquivo contém os comandos e regras para automatizar tarefas de compilação, execução e gerenciamento do projeto.
- `build`: Este diretório contém os arquivos de configuração para a criação de imagens Docker.
- `cmd`: Esta pasta contém o código-fonte principal do aplicativo, incluindo o arquivo `manager-api.go` que contém a função `main`.
- `configs`: Esta pasta contém arquivos de configuração, como o arquivo `env.yml`, que pode conter variáveis de ambiente, configurações de banco de dados, chaves de API, etc.
- `go.mod` e `go.sum`: Esses arquivos são usados para gerenciar as dependências do projeto usando o sistema de gerenciamento de pacotes do Go.
- `internal`: Esta pasta contém código interno do projeto, que não é destinado a ser importado por outros pacotes externos. Aqui são aplicadas as camadas `routes`, `service`, `database`, `grpc`, `producer` e `consumer`, que são de extrema importância para a aplicação.
- `pkg`: Este pacote contém os modelos das entidades da aplicação. 

## Como testar
Para os testes unitários, utilizamos o pacote `testing` do Go. Este pacote nos permite criar testes unitários e testes de cobertura no nosso projeto. 

### Como criar um arquivo de teste
Por convenção, os arquivos de teste sempre estão na mesma pasta ou pacote que o seu arquivo a ser testado, e sempre devem terminar com `_test.go`. Esses arquivos não são compilados quando você roda `go.build`, então não há preocupação de eles acabarem parando em algum deployment. 

### O pacote
A linguagem Go tem um pacote bem completo chamado `testing`, ele será o utilizado para realizar nossos testes. Então é preciso importa-lo para seu arquivo `_test.go`.

```bash
package example

import "testing"
```

### A função de teste
Uma função de teste em Go possui esta sintaxe: `func TestXxxx(t *testing.T)`. Isso significa que todas as funções de teste devem começar com a palavra `Test`, seguida do nome do teste, que deve começar com a letra maiúscula. Testes em Go recebem apenas um parâmetro, que é um ponteiro do tipo `testing.T`. Este tipo irá conter os métodos que precisaremos para imprimir resultados, mostrar logs de erros na tela e apresentar mensagens de erro, como por exemplo o método `t.Errorf()`.

```bash
package example

import "testing"

func TestExample(t *testing.T){

    got := Add(4, 6)
    want := 10

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
```

### Rodando o teste 
Após escrever e salvar o arquivo, precisamos testá-lo. Para isso, precisamos rodar o comando `go test` no terminal com o projeto aberto. 

```bash
$ go test
```

- `PASS`: Significa que seu código está funcionando como o esperado.

```bash
Output
PASS
ok      ./example 0.988s
```

- `FAIL`: Significa que seu teste falhou e deve ser revisado.

```bash
Output
FAIL
FAIL      ./example 0.988s
```

### Go test -v
Também podemos rodar o comando `go test -v`, para adicionar verbosidade ao retorno do teste, assim podemos ter algumas informações a mais da função que foi testada. 

```bash
Output
=== RUN   TestExample
--- PASS: TestExample (0.00s)
PASS
ok      ./example 1.410s
```

### Testes de cobertura 
Para rodar testes de cobertura no projeto, utilizamos o Makefile para a automatização desses testes.

## Makefile
O Makefile é um utilitário de script utilizado para automatizar tarefas de compilação, execução e gerenciamento de projetos.

### Como instalar o pacote Makefile

- Chocolatey: Primeiramente deverá ser feito o download do gerenciador de pacotes Chocolatey. Para a instalação, 
digite a seguinte linha de código no Windows PowerShell, tendo sido executado como administrador: 
```bash
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
```

- MakeFile: Após a instalação do Chocolatey, deverá ser feita a instalação do pacote Makefile, com a seguinte linha de código:
```bash
choco install make
```

- A instalação do pacote está concluída. É importante ressaltar que para a execução de qualquer comando utilizando o pacote Makefile, deverá ser usado um terminal baseado em uma distribuição Linux, como o Git Bash, por que o próprio Makefile é um arquivo de configuração usado principalmente em sistemas baseados em UNIX, como o Linux.

## Comandos para o pacote Makefile

`make:` Executa a regra principal all, que limpa, compila e realiza outras ações conforme configurado.

`make clean:` Limpa os arquivos gerados, excluindo a pasta bin e todos os seus conteúdos.

`make build:` Compila o projeto e gera o binário na pasta bin.

`make test:` Executa os testes do projeto.

`make cover:` Executa os testes apontando a porcentagem que indica a média de testes ao longo de todo o projeto.

`make cover-html:` Executa os testes e abre uma página html no navegador, que mostra a porcentagem de testes no projeto e também indica as linhas de código testadas e não testadas.

`make deps:` Instala as dependências do projeto (usando o Go Modules).

`make run:` Executa o projeto.
Certifique-se de atualizar BINARY_NAME e PKG no Makefile de acordo com o seu projeto. Isso permitirá que você ajuste o nome do binário e o caminho do pacote conforme necessário.

`make docker:` Cria uma imagem Docker para o projeto e roda um container em seguida.
