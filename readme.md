# Desafio Metaplane

## Clonando o projeto

Faça um git clone da seguinte url:

```shell
   git clone git@github.com:MatThHeuss/desafio-metaplane.git
```

Após isso, entre no diretório do projeto:

```shell
cd desafio-metaplane
```

## Como rodar o projeto

### Localmente

Para rodar o projeto localmente, basta executar o container docker.
Para isso, execute o comando abaixo:

```shell
  docker compose up -d
```

Após isso, é necessário iniciar a aplicação de dentro do container, para isso
execute o comando abaixo para entrar no container.

```shell
make enter-container
```

Dentro do container, basta executar o seguinte comando:

```shell
make start
```

### Localmente via Kubernetes

Para fazer um deploy local do kubernetes, basta seguir os passos a seguir:

```shell
 kubectl apply -f k8s/deployment.yaml
```

```shell
kubectl apply -f k8s/service.yaml
```

Após rodar os dois comandos, é necessário expor alguma porta para acessarmos a aplicação. Para isso,
execute o comando abaixo:

```shell
kubectl port-forward svc/metaplanesvc 8080:8080
```

### Acessando via GKE

Para acessar a aplicação via o Google Kubernetes Engine (GKE)
basta fazer requisições para o Ip abaixo:

```
http://34.151.225.187
```

### Executando os testes

Para executar os testes basta rodar o comando abaixo:

```shell
make test
```

### Endpoints

#### /saveLists

Endpoint utilizado para salvar as listas que serão mergeadas.
Exemplo de requisição:

```curl
curl --location 'http://localhost:8080/savelists' \
--header 'Content-Type: application/json' \
--data '{
    "numbers": [1, 2, 4]
}'
```
O Código aceita no máximo duas listas

```curl
curl --location 'http://34.151.225.187/savelists' \
--header 'Content-Type: application/json' \
--data '{
    "numbers": [1, 2, 4]
}'
```

#### /merge

Endpoint utilizado para realizar o merge das listas que foram salvas
no endpoint anterior.
Exemplo de requisição:

```curl
    curl --location --request POST 'http://localhost:8080/merge' \
--data ''
```

```curl
curl --location --request POST 'http://34.151.225.187/merge' \
--data ''
```
Após feito o merge, as listas são apagadas da memória para que possa ser
realizado vários testes.