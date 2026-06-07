<h1 align="center"> Привет! Я <a target="_blank"> Кармеев Артур из группы ЭФМО-01-25 </a> 
<img src="https://github.com/blackcater/blackcater/raw/main/images/Hi.gif" height="32"/></h1>
<h3 align="center"> Данная практика была простой :wink: </h3>

<h3 align="center"> Практическая работа №11: Создание GraphQL API с использованием gqlgen. Запросы и мутации </h3>


    
Структура работы:    
    
    
    └── pz11-graphql/
        ├── .gitignore
        ├── go.mod
        ├── go.sum
        ├── gqlgen.yml
        ├── README.md
        ├── server.go
        ├── graph/
        │   ├── generated.go
        │   ├── resolver.go
        │   ├── schema.graphqls
        │   ├── schema.resolvers.go
        │   └── model/
        │       └── models_gen.go
        ├── deploy/
        │   └── docker-compose.yml
        ├── .git/
        ├── services/
        │   └── tasks/
        │       ├── .dockerignore
        │       ├── Dockerfile
        │       └── cmd/
        │           └── tasks/
        │               └── main.go
        ├── internal/
        │   ├── task/
        │   │   ├── model.go
        │   │   └── repo.go
        │   ├── service/
        │   │   └── task_service.go
        │   └── httpapi/
        │       └── handler.go
        └── .github/
            └── workflows/
                └── ci.yml
