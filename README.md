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


## 1. Установка библиотеки gqlgen

<table cellpadding="10">
  <tr>
    <td><img width="974" height="516" alt="image" src="https://github.com/user-attachments/assets/003621d9-2e97-4b45-aaf8-1b3985eda966" /></td>
  </tr>
</table>

## 2. Подготовка модели и временного хранилища

<table cellpadding="10">
  <tr>
    <td><img width="974" height="517" alt="image" src="https://github.com/user-attachments/assets/5d2df44a-0439-4f7b-b2ff-5224fcda0fa1" /></td>
  </tr>
</table>

## 3. Запуск GraphQL-сервера

<table cellpadding="10">
  <tr>
    <td><img width="974" height="521" alt="image" src="https://github.com/user-attachments/assets/347422c0-f217-49b0-8eec-8c2b09690025" /></td>
  </tr>
</table>

<table cellpadding="10">
  <tr>
    <td><img width="974" height="519" alt="image" src="https://github.com/user-attachments/assets/aa1b6bdd-53fd-4389-80e4-a175bb5e586c" /></td>
  </tr>
</table>

## 4. Проверка работы через Playground

### 4.1 Запрос списка задач

<table cellpadding="10">
  <tr>
    <td><img width="974" height="519" alt="image" src="https://github.com/user-attachments/assets/997e7724-faa5-4503-950b-0a9a5d2274b4" /></td>
  </tr>
</table>

### 4.2 Запрос одной задачи по идентификатору

<table cellpadding="10">
  <tr>
    <td><img width="974" height="518" alt="image" src="https://github.com/user-attachments/assets/c333d564-df32-4293-8571-c63ba85bfd6b" /></td>
  </tr>
</table>

### 4.3 Создание задачи

<table cellpadding="10">
  <tr>
    <td><img width="974" height="513" alt="image" src="https://github.com/user-attachments/assets/83f57797-80ce-47eb-adee-a1af0949a051" /></td>
  </tr>
</table>

### 4.4 Обновление задачи

<table cellpadding="10">
  <tr>
    <td><img width="974" height="518" alt="image" src="https://github.com/user-attachments/assets/11a12145-c126-499c-9f75-257dd50a1ea7" /></td>
  </tr>
</table>

### 4.5 Удаление Задачи

<table cellpadding="10">
  <tr>
    <td><img width="974" height="518" alt="image" src="https://github.com/user-attachments/assets/6e1e42d9-5712-46bc-aec4-a9f0b6c0dfd2" /></td>
  </tr>
</table>

## 5. Интеграция с существующим сервисом tasks

За основу сервиса таск я взял 9-ую практическую работу

Общий слой (ну тут просто база из 9-ой практики)

| Файлики:) | Назначение |
|-----|------|
| `internal/task/model.go` | Модель Task (`id`, `title`, `description`, `done`) | 
| `internal/task/repo.go` | In-memory репозиторий, CRUD | 
| `internal/service/task_service.go` | Бизнес-логика (без Redis) | 
| `internal/httpapi/handler.go` | REST из pr9, `id` как строка (`t_001`) | 

### GraphQL
- `graph/resolver.go` — в Resolver добавлен `TaskService`
- `graph/schema.resolvers.go` — резолверы вызывают `r.TaskService`
- `server.go` — создаёт repo + service, передаёт в резолвер

### REST services/tasks
- `services/tasks/cmd/tasks/main.go` — тот же repo + service + REST API
- Удалён отдельный `services/tasks/go.mod` — всё в корневом модуле
- Обновлены Dockerfile и `docker-compose.yml` (сборка из корня)

В двух терминалах запускаем шарманку

<table cellpadding="10">
  <tr>
    <td><img width="974" height="516" alt="image" src="https://github.com/user-attachments/assets/e63df0f6-365c-464e-bb33-db93514fff94" /></td>
  </tr>
</table>

<table cellpadding="10">
  <tr>
    <td><img width="1907" height="1021" alt="image" src="https://github.com/user-attachments/assets/92f76f52-2087-448c-b21a-5b0f245aa935" /></td>
  </tr>
</table>

Теперь делаем вот такие интересные запросики:
- `curl http://localhost:8082/v1/tasks`
- `curl http://localhost:8082/v1/tasks/t_001`

<table cellpadding="10">
  <tr>
    <td><img width="974" height="518" alt="image" src="https://github.com/user-attachments/assets/988f2816-cb66-40f9-92b4-b22b8b0418cf" /></td>
  </tr>
</table>

<table cellpadding="10">
  <tr>
    <td><img width="974" height="517" alt="image" src="https://github.com/user-attachments/assets/2f965000-d333-4c70-824a-ec97facc6aa7" /></td>
  </tr>
</table>

## 6. Реализация двух Query и трёх Mutation

Создаем три задачки

<table cellpadding="10">
  <tr>
    <td><img width="974" height="518" alt="image" src="https://github.com/user-attachments/assets/d8631e01-55f9-4338-8f79-2b84a25ff8c1" /></td>
  </tr>
</table>

<table cellpadding="10">
  <tr>
    <td><img width="974" height="515" alt="image" src="https://github.com/user-attachments/assets/633912c1-d9ef-4e98-8cb2-1b38c2ebd208" /></td>
  </tr>
</table>

<table cellpadding="10">
  <tr>
    <td><img width="974" height="517" alt="image" src="https://github.com/user-attachments/assets/ea7f3cf4-e074-44f8-bd31-6948ebbf4d5d" /></td>
  </tr>
</table>

Теперь проверяем

<table cellpadding="10">
  <tr>
    <td><img width="974" height="515" alt="image" src="https://github.com/user-attachments/assets/0738a2f8-ebda-4eed-a690-93153b268b9f" /></td>
  </tr>
</table>

<table cellpadding="10">
  <tr>
    <td><img width="974" height="518" alt="image" src="https://github.com/user-attachments/assets/3e8f410b-7a1b-4a00-9989-e567d87f4f6a" /></td>
  </tr>
</table>

Обновлять и удалять не стану и так норм :upside_down_face:

## 7. Контрольные вопросы :unamused:

1. Что такое GraphQL и в чём его основное отличие от REST API?

GraphQL — язык запросов и способ организации API, при котором клиент обращается к одному endpoint (у нас /query) и сам указывает, какие поля нужны в ответе.

В REST обычно много URL (`/tasks`, `/tasks/{id}`), а формат ответа задаёт сервер. В GraphQL один URL, структура ответа определяется запросом клиента. GraphQL не заменяет REST полностью, а даёт другой способ получения данных — гибче по составу ответа, но сложнее в настройке.

2. Для чего используется GraphQL-схема?

Схема (`schema.graphqls`) — контракт API: какие типы есть (`Task`), какие операции доступны (`Query`, `Mutation`), какие аргументы и поля обязательны. По схеме gqlgen генерирует код, Playground показывает документацию, клиент и сервер договариваются об одном формате данных. Схема — «источник правды» для API.

3. Чем Query отличается от Mutation?

Query — операции чтения данных без изменения состояния: `tasks`, `task(id)`. Mutation — операции изменения: `createTask`, `updateTask`, `deleteTask`. В GraphQL их разделяют явно: запросы для чтения, мутации для записи. В вашей практике Query берёт данные из `TaskService`, Mutation создаёт, обновляет и удаляет задачи.

4. Что такое резолвер и какую роль он выполняет?

Резолвер — функция на сервере, которая выполняет конкретное поле или операцию из схемы. Например, `Tasks`, `Task`, `CreateTask` в `schema.resolvers.go`. Резолвер связывает GraphQL-запрос с реальной логикой: вызывает `TaskService`, читает репозиторий, возвращает данные в формате GraphQL. Схема описывает «что можно запросить», резолвер — «как это получить».

5. Почему GraphQL позволяет уменьшить over-fetching данных?

Over-fetching — когда сервер отдаёт больше полей, чем нужно клиенту. В REST ответ часто фиксирован. В GraphQL клиент пишет, например, только `id` и `title` — сервер вернёт только их, без `description` и `done`. Меньше лишних данных по сети и проще формировать ответ под экран (мобильное приложение, виджет и т.д.).

6. Для чего используется библиотека gqlgen в Go-проекте?

gqlgen — генератор GraphQL-сервера для Go. Сначала пишется схема, затем `gqlgen generate` создаёт типы, интерфейсы резолверов и серверный код. Разработчик дописывает только резолверы (бизнес-логику). Это schema-first подход: схема в центре, boilerplate генерируется автоматически, меньше ручной рутины и ошибок в типах.

7. Почему желательно использовать общий сервисный или репозиторный слой для REST и GraphQL?

Чтобы не дублировать бизнес-логику и данные. Если GraphQL и REST хранят задачи отдельно, легко получить рассинхрон: в Playground одно, в REST другое. Общий `TaskService` и `Repo` — одни правила создания, обновления и удаления для обоих API. Проще поддерживать и тестировать, архитектура чище.

8. Что произойдёт, если GraphQL-запрос попытается получить слишком большой объём данных?

Возможны медленный ответ, высокая нагрузка на сервер и БД, таймауты. Запрос вроде `tasks` со всеми вложенными полями по большой базе может «утянуть» много данных за один round-trip. На production обычно ставят лимиты: глубина запроса, сложность, пагинация, rate limiting. В учебной работе с in-memory и двумя задачами это не проявляется, но риск в GraphQL реален при сложных запросах.

9. Какие преимущества даёт Playground при тестировании API?

Playground встроен в сервер (`http://localhost:8080/`): не нужен отдельный Postman для базовых проверок. Есть подсветка синтаксиса, автодополнение по схеме, панель Variables, просмотр JSON-ответа и ошибок. Удобно гонять Query и Mutation из методички и делать скриншоты для отчёта. Для разработки и демонстрации API это быстрее, чем собирать curl вручную.

10. В каких случаях GraphQL особенно удобен на практике?

GraphQL удобен, когда много клиентов с разными потребностями (веб, мобильное приложение, админка) и каждому нужен свой набор полей. Хорошо подходит для агрегирующих API (один запрос вместо нескольких REST-вызовов), для продуктов с богатой связанной моделью и когда важно дать фронтенду гибкость без размножения endpoint’ов. Менее выгоден для простых CRUD-сервисов и файловых/стриминговых сценариев, где REST или gRPC проще.
















