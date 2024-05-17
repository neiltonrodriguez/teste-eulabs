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


se der conflito com a porta 3306, pare o seu serviço mysql, ou mude a porta no Dockerfile
```
# sudo service mysqld stop
ou
# sudo /etc/init.d/mysqld stop
```


##  Rotas
```
GET: http://localhost:8080/product/
GET: http://localhost:8080/product/:id
POST: http://localhost:8080/product/
PUT: http://localhost:8080/product/:id
DELETE: http://localhost:8080/product/:id

```

cURL create:
```
curl --request POST \
  --url http://localhost:8080/product/ \
  --data '{
  "name": "play station 4",
  "description": "teste de produto",
  "value": 5000.99
}'
```

cURL GetAll:
```
curl --request GET \
  --url http://localhost:8080/product/
```

cURL getById:
```
curl --request GET \
  --url http://localhost:8080/product/:id
```

cURL Update:
```
curl --request PUT \
  --url http://localhost:8080/product/9d89ea02-d84a-46c5-ba50-de89fd237e52 \
  --data '{
  "name": "Notebook i7",
  "description": "teste de produto lenovo",
  "value": 1999.33
}'
```

cURL Delete:
```
curl --request DELETE \
  --url http://localhost:8080/product/709a470b-75f3-4038-a752-477bcbdc4607
```


#### endereços para acesso:
```
PhpMyAdmin: http://localhost:8888/
Endereço da api: http://localhost:8080/
```

Developed by Neilton Rodrigues