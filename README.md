# Order App

## Requisitos

- Docker
- Docker Compose

## Como executar

1. Clone o repositório
2. Navegue até o diretório do projeto
3. Execute `docker compose up`

## Endpoints

- REST: Porta `8080`
- gRPC: Porta `50051`
- GraphQL: Porta `8080`

## REST
- Listar pedidos

```http
GET /order HTTP/1.1
Content-Type: application/json
```

- Criar pedido
```http
POST /order HTTP/1.1
Content-Type: application/json

{
"price": 100.5,
"tax": 0.5
}
```

# gRPC
- Listar pedidos

```sh
grpcurl -plaintext -d '{}' localhost:50051 pb.OrderService/ListOrders
```

- Criar pedido
```sh
grpcurl -plaintext -d '{"price": 100.5, "tax": 0.5}' localhost:50051 pb.OrderService/CreateOrder
```

# GraphQL

- Listar pedidos
```graphql
query {
    listOrders {
        id
        price
        tax
    }
}
```

- Criar pedido
```graphql
mutation {
    createOrder(price: 100.5, tax: 0.5) {
        id
        price
        tax
    }
}
```

## Estrutura da Order

```json
{
    "id": "a",
    "price": 100.5,
    "tax": 0.5
}
```

## Testando com o api.http

Abra o arquivo `api/http/api.http` em um editor que suporte e execute as requests.

## Testando o Serviço gRPC
Para testar o serviço gRPC, use o grpcurl:

1. Instale o grpcurl.
2. Execute o comando:

```bash
grpcurl -plaintext -d '{}' localhost:50051 pb.OrderService/ListOrders
```

## Notas

Cada vez que o docker-compose é executado, a tabela orders será recriada e populada com 10 registros iniciais.
