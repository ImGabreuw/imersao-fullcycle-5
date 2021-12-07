# transaction-processing

### Rodar os testes

```bash
$ go test ./...
```

### Criar banco de dados

```bash
$ sqlite3 test.db
```

* Copiar e colar no terminal do SQlite

  ```sql
  CREATE TABLE transactions
  (
    id            TEXT NOT NULL,
    account_id    TEXT NOT NULL,
    amount        REAL NOT NULL,
    status        TEXT NOT NULL,
    error_message TEXT NOT NULL,
    created_at    TEXT NOT NULL,
    updated_at    TEXT NOT NULL
  );
  ```
  
### Entrar no _Topic_ "transactions" como _Producer_

```bash
$ kafka-console-producer --bootstrap-server=localhost:9092 --topic=transactions

> {"id":"1","account_id":"1","credit_card_number":"4193523830170205","credit_card_name":"Wesley Silva","credit_card_expiration_month":12,"credit_card_expiration_year":2021,"credit_card_cvv":123,"amount":900}
```

### Entrar no _Topic_ "transactions_result" como _Consumer_

```bash
$ kafka-console-consumer --bootstrap-server=localhost:9092 --topic=transactions_result
```

* Ao publicar a mensagem no tópico "transactions", deve aparecer a seguinte mensagem no terminal do _Consumer_

  ```json
  {"id":"1","status":"approved","error_message":""}
  ```