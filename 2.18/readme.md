## Запуск
go run . -port=8081 

### В POST принимает JSON вида: 
{
    "user_id": 42,
    "date": "2025-07-24T10:00:00Z",
    "event": "logssin"
}
### В Get 
Добавляем qwery параметр, например:

?user_id=42&date=2025-07-24

## Структура
.
├── domain
│   ├── errors.go
│   ├── event.go
│   ├── handlers.go
│   ├── logger.go
│   ├── serv.go
│   └── storage.go
├── httpsh
│   ├── handlers.go
│   ├── middleware.go
│   ├── router.go
│   └── server.go
├── main.go
├── serv
│   └── servImpl.go
├── storage
│   └── storageImpl.go
├── utils
│   ├── flagos.go
│   ├── logger.go
│   └── validate.go
└── work.md

## Эндпойнты
   POST /create_event — создание нового события;
   POST /update_event — обновление существующего;
   POST /delete_event — удаление;

   GET /events_for_day — получить все события на день;
   GET /events_for_week — события на неделю;
   GET /events_for_month — события на месяц.

   