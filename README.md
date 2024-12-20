> [!IMPORTANT]
> 
> 


```shell
go test ./pkg/calc
```

```shell
curl -X POST http://localhost:8080/api/v1/calculate \
     -d '{"expression":"2+2"}'
```

```shell
{"result":"4.000000"}
```

```shell
Calc
├── cmd
│   └── main.go
├── config
│   ├── config.go
│   └── config.json
├── internal
│   ├── application
│   │   └── application.go
│   ├── Errors
│   │   └── error.go
│   └── http
│       ├── handler
│       │   ├── handler.go
│       │   └── handler_test.go
│       └── server
│           └── http.go
├── pkg
│   └── calc
│       ├── calc.go
│       └── calc_test.go
├── logger
│   └── logger.go
├── go.mod
└── README.md

```