# TESTE EULABS
### CRUD de produtos

### Descrição
uma simples api para CRUD de produtos, usando golang e mysql

### Contexto de negócio
```
Docker
MySql
Golang 1.21.3
Echo Framework
```


## Executar localmente
na primeira vez, use:
```
# docker-compose up -d --build
# renomei o arquivo .env-example para .env
# acesse o phpMyAdmin em(http://localhost:8888/) e execute o script sql que está dentro de ./script-sql/model.sql para criar as tabelas
```


se faltar alguma dependência use o compose install dentro do container
```
# docker exec setup-php composer install
```


##  Rotas
```
GET: http://localhost:8080/product/
GET: http://localhost:8080/product/:id
POST: http://localhost:8080/product/
PUT: http://localhost:8080/product/:id
DELETE: http://localhost:8080/product/:id

```

payload para create e update:
```
{
  "name": "nome do produto",
  "description": "descrição produto",
  "value": 59.99
}
```


#### endereços para acesso:
```
PhpMyAdmin: http://localhost:8888/
Endereço da api: http://localhost:8080/
```

Developed by Neilton Rodrigues