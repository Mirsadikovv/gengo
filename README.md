# gengo

CLI-генератор Go-микросервисов под архитектуру `go`.
Автоматически создаёт полный boilerplate: модели, DTO, сервисы, хендлеры, роутинг.

---

## Установка

```bash
# Локально (из директории gengo/)
make install-local

# Из remote (после публикации)
make install
```

Убедись, что `$(go env GOPATH)/bin` добавлен в `PATH`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

---

## Команды

### `gengo --new` — новый проект с нуля

```bash
gengo --new
```

```
Enter project name:  user_service
Enter module name:   user
Enter entity name:   user
Enter mod path:      github.com/myorg/user_service
```

Создаёт полную структуру проекта:

```
user_service/
├── main.go                          # точка входа
├── dev.go                           # загрузка .env.dev (build tag: dev)
├── prod.go                          # загрузка env из OS (build tag: !dev)
├── .env.example                     # шаблон переменных окружения
├── .gitignore
├── Makefile
└── src/
    ├── main.go                      # Env, Exec(), migration()
    └── module/
        ├── auth_service/            # bootstrap-модуль аутентификации
        │   ├── dto/auth_dto.go
        │   └── middleware/auth_middleware.go
        └── user/                    # первый модуль
            ├── cmd.go               # регистрация хендлеров (/api/v1)
            ├── dto/user_dto.go
            ├── model/user_model.go
            ├── handler/user_handler.go
            └── service/user_service.go
```

---

### `gengo --add` — добавить модуль в существующий проект

Запускать из **корня проекта** (где лежит `go.mod`):

```bash
cd user_service
gengo --add
```

```
Enter module name:  payment
Enter entity name:  invoice
```

Создаёт `src/module/payment/` и **автоматически обновляет** `src/main.go`:
- добавляет импорты `payment_cmd` и `invoice_model`
- вставляет `payment_cmd.Cmd(router, db, log, authMiddleware)` в блок вызовов
- добавляет `&invoice_model.Invoice{}` в `db.AutoMigrate(...)`

> `src/main.go` должен содержать маркеры `// @gengo:modules` и `// @gengo:migration`.
> При создании проекта через `--new` они добавляются автоматически.

---

## Сгенерированные эндпоинты

Для каждой сущности генерируются 6 REST-эндпоинтов:

| Метод    | Путь                       | Действие          |
|----------|----------------------------|-------------------|
| `POST`   | `/api/v1/{entity}/create`  | Создать запись    |
| `PATCH`  | `/api/v1/{entity}/:id`     | Обновить запись   |
| `DELETE` | `/api/v1/{entity}/:id`     | Удалить (soft)    |
| `GET`    | `/api/v1/{entity}/page`    | Список с пагинацией |
| `GET`    | `/api/v1/{entity}/search`  | Поиск (ILIKE)     |
| `GET`    | `/api/v1/{entity}/:id`     | Получить по ID    |

---

## Структура сгенерированного кода

### Model (`{entity}_model.go`)

```go
type User struct {
    Id        int64      `gorm:"primaryKey"`
    Name      string     `gorm:"not null"`
    CreatedAt *time.Time `gorm:"autoCreateTime"`
    UpdatedAt *time.Time `gorm:"autoUpdateTime"`
    DeletedAt *time.Time `gorm:"softDelete"`
}

func (User) TableName() string { return "users" }
```

### DTO (`{entity}_dto.go`)

```go
type UserPage = response.PageData[User]

type User struct { ... }
type CreateUser struct { Name string `validate:"required"` }
type UpdateUser struct { Name *string }
type UserQueryParams struct { Name string `query:"name"` }
```

### Service (`{entity}_service.go`)

```go
type UserService interface {
    Create(dto *user_dto.CreateUser) (int64, error)
    Update(dto *user_dto.UpdateUser, filter pg.Filter) error
    Delete(filter pg.Filter) error
    Page(ctx, paginate, filter) (*user_dto.UserPage, error)
    Find(ctx, filter) ([]user_dto.User, error)
    FindOne(ctx, filter) (*user_dto.User, error)
}
```

Реализация использует хелперы из `shared_service/pg`:
`pg.Create`, `pg.Update`, `pg.Delete`, `pg.PageWithScan`, `pg.FindWithScan`, `pg.FindOneWithScan`.

### Handler (`{entity}_handler.go`)

- Биндинг через `req.BindBody()` / `req.BindQuery()`
- Ответы через `req.OK()`, `req.NoContent()`, `req.BadRequest()`
- Swagger-аннотации для каждого метода

---

## Переменные окружения

| Переменная          | Описание                  | По умолчанию |
|---------------------|---------------------------|--------------|
| `HTTP_HOST`         | Хост сервера              | `localhost`  |
| `HTTP_PORT`         | Порт сервера              | `8080`       |
| `DB_HOST`           | Хост PostgreSQL           | —            |
| `DB_USER`           | Пользователь БД           | —            |
| `DB_PASSWORD`       | Пароль БД                 | —            |
| `DB_NAME`           | Имя БД                    | —            |
| `DB_PORT`           | Порт БД                   | —            |
| `DB_SSL_MODE`       | SSL режим                 | `disable`    |
| `DB_TIME_ZONE`      | Часовой пояс              | `UTC`        |
| `JWT_SECRET`        | Секрет JWT                | —            |
| `JWT_EXPIRED`       | Время жизни токена (сек)  | —            |
| `JWT_REFRESH_EXPIRED` | Время жизни refresh (сек) | —          |
| `REDIS_ADDR`        | Адрес Redis               | —            |
| `SWAGGER_USER`      | Basic Auth для Swagger    | —            |
| `SWAGGER_PASSWORD`  | Basic Auth для Swagger    | —            |

---

## shared_service

`shared_service` — встроенная библиотека gengo, публикуется вместе с ним на GitHub.
Сгенерированные проекты импортируют её напрямую как публичную зависимость:

```go
import "github.com/Mirsadikovv/gengo/shared_service/pg"
import "github.com/Mirsadikovv/gengo/shared_service/request"
import "github.com/Mirsadikovv/gengo/shared_service/response"
import "github.com/Mirsadikovv/gengo/shared_service/logger"
import "github.com/Mirsadikovv/gengo/shared_service/middleware"
import "github.com/Mirsadikovv/gengo/shared_service/jwt"
import "github.com/Mirsadikovv/gengo/shared_service/redis"
```

Никакой дополнительной настройки не требуется — `go mod tidy` подтянет зависимость
из `github.com/Mirsadikovv/gengo` автоматически.

### Пакеты shared_service

| Пакет         | Описание                                          |
|---------------|---------------------------------------------------|
| `pg`          | GORM-хелперы: Create, Update, Delete, Page, Find  |
| `request`     | HTTP-хелперы: BindBody, BindQuery, OK, NoContent  |
| `response`    | Типы ответов: PageData[T], ID, Error              |
| `logger`      | Интерфейс логгера                                 |
| `middleware`  | AuthEchoMiddleware (JWT + GORM)                   |
| `jwt`         | JWT сервис (access + refresh токены)              |
| `redis`       | Redis клиент                                      |
| `sharedutil`  | Env, Validator, утилиты                           |
| `swagger`     | Basic Auth middleware для Swagger UI              |
| `chromium`    | Генерация PDF через chromedp                      |
| `parser`      | Парсер go.mod (version)                           |

---

## Структура gengo

```
gengo/
├── main.go                    # CLI: --new / --add
├── Makefile                   # install / install-local
├── go.mod                     # module github.com/Mirsadikovv/gengo
├── internal/
│   ├── new_project.go         # логика --new
│   └── add_module.go          # логика --add
├── service/
│   ├── util.go                # конвертация регистров (toCamel, toSnake, ...)
│   ├── folder.go              # создание директорий
│   ├── file.go                # рендеринг шаблонов
│   └── injector.go            # string-инъекция в src/main.go по маркерам
├── stuble/
│   ├── stuble.go              # go:embed декларации
│   ├── module/                # шаблоны модуля/сущности
│   │   ├── cmd.tpl
│   │   ├── dto.tpl
│   │   ├── model.tpl
│   │   ├── service.tpl
│   │   └── handler.tpl
│   ├── project/               # шаблоны корня проекта
│   │   ├── main_entry.tpl
│   │   ├── dev.tpl
│   │   ├── prod.tpl
│   │   ├── src_main.tpl
│   │   ├── env.tpl
│   │   ├── gitignore.tpl
│   │   └── makefile.tpl
│   └── auth/                  # bootstrap auth_service
│       ├── auth_dto.tpl
│       └── auth_middleware.tpl
└── shared_service/            # встроенная библиотека (часть модуля gengo)
    ├── pg/                    # GORM-хелперы
    ├── request/               # HTTP request/response хелперы
    ├── response/              # типы ответов
    ├── logger/                # интерфейс логгера
    ├── middleware/            # auth middleware
    ├── jwt/                   # JWT сервис
    ├── redis/                 # Redis клиент
    ├── sharedutil/            # env, validator, утилиты
    ├── swagger/               # Swagger Basic Auth
    ├── chromium/              # PDF генерация
    └── parser/                # парсер версий
```
