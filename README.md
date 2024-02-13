# go-todo

## Env

```
$ cp -pr .env{.example,}
```

## Load Modules

```
$ go mod tidy -v
```

## Migrate

```
$ go run cmd/migrate.go
```

## Run PostgreSQL

```
$ docker-compose up -d --build
$ docker-compose exec db bash
e07b4f515b65:/# psql -h db -p 5432 -U todo-user -d todos
Password for user todo-user: 
psql (15.1)
todos=# \dt
Did not find any relations.
（マイグレーション後）
todos=# \dt
         List of relations
 Schema | Name  | Type  |   Owner   
--------+-------+-------+-----------
 public | todos | table | todo-user
(1 row)
（インサート後）
todos=# select * from todos;
 id |  title  | status |          created_at           |          updated_at           
----+---------+--------+-------------------------------+-------------------------------
  1 | テスト1 | 処理中 | 2024-02-12 16:13:17.196959+00 | 2024-02-12 16:13:17.196959+00
```

## Run Echo Server (Not docker yet)

```
$ go run ./cmd/main.go       
postgres://todo-user:todo-password@localhost:5432/todos
データベースとの接続に成功しました

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.11.4
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8080
```

## Call API

```
$ curl -X POST -H "Content-Type: application/json" -d '{"title":"テスト1", "status":"処理中"}' http://localhost:8080/todos
{"id":0,"title":"テスト1","status":"処理中","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}

$ curl http://localhost:8080/todos
[{"id":1,"title":"テスト1","status":"処理中","created_at":"2024-02-13T01:13:17.196959+09:00","updated_at":"2024-02-13T01:13:17.196959+09:00"}]

$ curl http://localhost:8080/todos/1
{"id":1,"title":"テスト1","status":"処理中","created_at":"2024-02-13T01:13:17.196959+09:00","updated_at":"2024-02-13T01:13:17.196959+09:00"}

$ curl -X PUT -H "Content-Type: application/json" -d '{"title":"テスト1", "status":"未処理"}' http://localhost:8080/todos/1
"object does not exist"（バグ）

$ curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/todos/1

"object does not exist"（バグ）
```