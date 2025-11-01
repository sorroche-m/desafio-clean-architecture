# Desafio Clean Architecture

Este projeto é uma implementação do desafio de Clean Architecture. O objetivo principal é criar o caso de uso para listagem de pedidos (*Orders*), expondo esta funcionalidade através de múltiplas interfaces: REST, gRPC e GraphQL.

### Requisitos do Desafio
* Endpoint REST: `GET /order`
* Serviço gRPC: `ListOrders`
* Query GraphQL: `ListOrders`
* Banco de dados provisionado via Docker (com migrações).
* Arquivo `api.http` para testes manuais de criação e listagem.

---

## 🚀 Começando

Siga estas instruções para configurar e rodar o ambiente de desenvolvimento local.

### 1. Pré-requisitos

Certifique-se de que você tem as seguintes ferramentas instaladas:
* [Docker](https://www.docker.com/get-started/) e [Docker Compose](https://docs.docker.com/compose/install/)
* [Make](https://www.gnu.org/software/make/)
* [grpcurl](https://github.com/fullstorydev/grpcurl) (para testar o gRPC)

### 2. Endpoints da Aplicação

A aplicação expõe os seguintes serviços nas respectivas portas:

* **Serviços HTTP (REST e GraphQL):** Porta **`8080`**
* **Serviço gRPC:** Porta **`50051`**

### 3. Banco de Dados

O banco de dados é provisionado via Docker e as migrações são aplicadas na subida. Para iniciar, execute:

```bash
docker compose build && docker compose up -d
```

Isso irá subir o banco e preparar o ambiente para a aplicação.

### 4. Protobuf / gRPC

Os arquivos `.proto` estão localizados na pasta `pkg/proto`.

Para gerar os arquivos Go dos protos, execute:

```bash
make install-tools
make generate
```

Isso instalará os plugins necessários e gerará os arquivos `.pb.go`.

---

## 🧪 Testes da API (Manuais)

Os arquivos para testes manuais da API estão localizados na pasta `test/`.

Utilizamos o formato `.http` (ex: `test/api.http`), que permite executar requisições HTTP diretamente do seu editor de código.

###  Executando com VSCode (Recomendado)

A forma mais fácil de rodar esses testes é usando a extensão **[REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)** no VSCode.

1.  **Instale a extensão:** Procure por `REST Client` no painel de Extensões do VSCode e instale-a.
2.  **Abra o arquivo:** Navegue e abra o arquivo `test/api.http`.
3.  **Envie a Requisição:** Acima de cada definição de rota (ex: `GET http://...`), você verá um link de texto `Send Request`. Clique nele para executar a chamada.

## 📡 Exemplo de Chamada gRPC (Usando grpcurl)

Para testar os *endpoints* gRPC diretamente da linha de comando, é necessário ter a ferramenta **`grpcurl`** instalada.

Abaixo, um exemplo de como criar um pedido (`Order`) chamando o método `CreateOrder` do serviço `OrderService`:

```bash
grpcurl -plaintext -proto pkg/proto/order.proto -d '{
    "customer_id": "customer-123",
    "amount": 299.99
}' localhost:50051 order.OrderService/CreateOrder
```

