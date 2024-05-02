# r3ddd1t service


### Дерево проекта
```
reddit
├── Dockerfile              # Dockerfile для сборки Docker-образа
├── docker-compose.yml      # Docker Compose файл
├── go.mod                  # Go модуль
├── go.sum                  # Проверочная сумма модулей
├── gqlgen.yml              # Конфигурация GQLGen
├── internal                # Внутренние пакеты
│   ├── entities            # Сущности (модели данных)
│   │   ├── comment.go      # Сущность комментария
│   │   ├── models_gen.go   # Сгенерированные модели
│   │   └── post.go         # Сущность поста
│   ├── infrastructure      # Компоненты инфраструктуры
│   │   ├── database        # Работа с базами данных
│   │   │   └── postgres.go # Подключение к PostgreSQL
│   │   └── server          # HTTP-серверы и роутинг
│   │       └── http_server.go
│   ├── repositories        # Репозитории данных
│   │   ├── memory          # Реализации репозиториев в памяти
│   │   │   ├── comment.go  # Репозиторий комментариев в памяти
│   │   │   └── post.go     # Репозиторий постов в памяти
│   │   ├── postgres        # Реализации репозиториев на PostgreSQL
│   │   │   ├── comment_repository.go  # Репозиторий комментариев
│   │   │   └── post_repository.go     # Репозиторий постов
│   ├── resolvers           # Резолверы для GraphQL
│   │   ├── generated.go    # Сгенерированные резолверы
│   │   └── resolvers.go    # Основные резолверы
│   ├── usecases            # Сценарии использования (Use Cases)
│   │   ├── comment_usecase.go  # Сценарий использования для комментариев
│   │   └── post_usecase.go     # Сценарий использования для постов
├── pkg                     # Пакеты общего назначения
│   └── logger              # Пакет логирования
│       └── logger.go
├── readme.md               # Описание проекта
├── schema.graphql          # Схема GraphQL
└── scripts                 # Скрипты для инициализации базы данных
    └── init.sql

```

### Как запустить
```
docker compose up --build
```

### Ендпоинт для запросов
```http://localhost:8080/playground```

### Примеры запросов
```azure
mutation{
  
  createComment(postID: 1, content:"плохой пост", parentID:1){
    id
  }
}

{
  comments(postID: 7){
    content
  }
},

{
  comments(postID: 1, parentID: 1, limit: 5){
    id
    content
    parentID
  }
}

mutation {
  createPost(title: "I hate Graphql!", allowComments:false, content:"I really hate graphQL))))"){
    id
  }
}

mutation{
  
  createComment(postID: 6, content:"Хороший пост!!", parentID:3){
    id
  }
}

{
  comments(postID: 1){
    content
  }
}
```