# Order App

## Requisitos

- Docker
- Docker Compose

## Como executar

1. Clone o repositório
2. Navegue até o diretório do projeto
3. Execute `docker compose up`

## Endpoints

- REST: `GET /order` na porta `8080`
- gRPC: Porta `50051`
- GraphQL: `POST /graphql` na porta `8080`

## Estrutura da Order

```json
{
    "id": "a",
    "price": 100.5,
    "tax": 0.5
}
```

## Testando com o api.http

Abra o arquivo `/app/http/api.http` em um editor que suporte e execute as requests.

## Testando o Serviço gRPC
Para testar o serviço gRPC, use o grpcurl:

1. Instale o grpcurl.
2. Execute o comando:

```bash
grpcurl -plaintext -d '{}' localhost:50051 order-app.OrderService/ListOrders
```

## Notas

Cada vez que o docker-compose é executado, a tabela orders será recriada e populada com 10 registros iniciais.
