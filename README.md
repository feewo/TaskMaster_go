# Golang REST API для управления задачами

## Описание

Простой REST API проект с авторизацией, написанный на языке Golang, использующий библиотеку GORM для взаимодействия с базой данных. API предоставляет набор ресурсов для управления данными, таких как создание, чтение, обновление и удаление.

![logo](https://github.com/feewo/TaskMaster_go/assets/57757873/8e8956c0-2c24-485c-8984-6ea65cc025ec)

## Технологический стек

1. Golang
2. GORM
3. Docker
4. Unit-тесты
5. REST API

## Запуск проекта

1. docker-compose build
2. docker-compose up

## API Endpoints
**Для получения данных из API необходимо авторизоваться, отправив заголовок Authorization с вашим токеном пользователя в каждом запросе.**

Пример:

```
Authorization: your_token
```

### Задачи

Создание новой задачи:

```
POST /task
{
  "title": "Программирование",
  "UserID": 1,
}
```

Получение списка задач:

```
GET /task
```

Получение определенной задачи:
login
```
GET /task/{id_task}
```

Обновление задачи:

```
PUT /task/{id_task}
{
  "title": "Кулинария",
  "UserID": 1,
}
```

Удаление задачи:

```
DELETE /task/{id_task}
```

### Подзадачи

Создание новой подзадачи:

```
POST /taskpoint
{
  "title": "Выучить golang",
  "TaskID": 1,
  "Ready": Ready,
}
```

Получение списка подзадач:

```
GET /taskpoint
```

Получение определенной подзадачи:

```
GET /taskpoint/{id_taskpoint}
```

Обновление подзадачи:

```
PUT /taskpoint/{id_taskpoint}
{
  "title": "Выучить django",
  "TaskID": 1,
  "Ready": Ready,
}
```

Удаление подзадачи:

```
DELETE /taskpoint/{id_taskpoint}
```

Получение списка подзадач по указанной задаче

```
GET /taskpoint_task/{id_task}
```

### Пользователи

Создание нового пользователя (не требует токена):

```
POST /user
{
  "login": "user123",
  "surname": "Иванов",
  "name": "Иван",
  "patronymic": "Иванович",
  "email": "test@mail.ru",
  "password": "12345678",
}
```

Получение списка пользователей:

```
GET /user
```

Получение определенного пользователя:

```
GET /user/{id_user}
```

Обновление пользователя:

```
PUT /user/{id_user}
{
  "login": "user123",
  "surname": "Петров",
  "name": "Петр",
  "patronymic": "Петрович",
  "email": "test@mail.ru",
}
```

Удаление пользователя:

```
DELETE /user/{id_user}
```

Авторизация пользователя:

```
POST /token
{
  "login": "user123",
  "password": "12345678",
}
```

Удаление токена пользователя (выход из лк):

```
DELETE /token
```

Получение информации о пользователе по токену
```
GET /user_token
```

## Разработчики
+ https://github.com/feewo - backend-разработчик
+ https://github.com/EkaterinaKopenkina - frontend-разработчик

## Контакты

+ Почта: [capi62@yandex.ru](mailto:capi62@yandex.ru)
+ Telegram: https://t.me/feewo
