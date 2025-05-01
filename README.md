## Criar os Endpoints
- Endpoint REST (GET /orders)
- Service ListOrders com GRPC
- Query ListOrders GraphQL

## Criar Migrações
* ????
* migrate create -ext=sql -dir-migrations -seq init
* migrate -path=migrate -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up
* migrate -path=migrate -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down

## arquivo api.http
* request para criar
* request listar as orders.

## Porta da Aplicação
* Web REST 8000
* GraphQL 8080
* gRPC 50051

### Acesso ao serviço gRPC
# Exemplo com grpcurl para listar serviços
```bash
grpcurl -plaintext localhost:50051 list
```
# Exemplo para criar uma ordem
```bash
grpcurl -plaintext -d '{"id":"1", "price":10.5, "tax":1.5}' localhost:50051 pb.OrderService/CreateOrder

```
# Exemplo para listar ordens

```bash
grpcurl -plaintext -d '{}' localhost:50051 pb.OrderService/ListOrders

```

## Acessando Usabndo Evans

```bash
evans --verbose -r repl

package pb


show service

service OrderService

call ListOrders


```
