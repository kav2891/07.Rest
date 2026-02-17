📘 API Документация

Все ответы имеют заголовок:

`Content-Type: application/json`

Формат ошибки:

`{
"error": "сообщение"
}`

---

🔹 GET `/health`

Health-check сервиса.

✅ Ответ 200
`{
"status": "ok"
}`

---

🔹 GET `/tasks`

Получить список задач.

✅ Ответ 200

`[
{
"id": 1,
"title": "Buy milk",
"done": false,
"created_at": "2026-02-17T12:00:00Z"
}
]`

---

🔹 POST `/tasks`

Создать новую задачу.

📥 Тело запроса

`{
"title": "Buy milk",
"done": false
}`

✅ Ответ 201

`{
"id": 1,
"title": "Buy milk",
"done": false,
"created_at": "2026-02-17T12:00:00Z"
}`

❌ Ошибка 400

`{
"error": "title is required"
}`

---

🔹 GET `/tasks/{id}`

Получить задачу по ID.

✅ Ответ 200

`{
"id": 1,
"title": "Buy milk",
"done": false,
"created_at": "2026-02-17T12:00:00Z"
}`

❌ Ответ 404

`{
"error": "task not found"
}`

---

🔹 PUT `/tasks/{id}`

Обновить задачу полностью.

📥 Тело запроса

`{
"title": "Buy bread",
"done": true
}`

✅ Ответ 200

`{
"id": 1,
"title": "Buy bread",
"done": true,
"created_at": "2026-02-17T12:00:00Z"
}`

---

🔹 DELETE `/tasks/{id}`

Удалить задачу.

✅ Ответ 204

Без тела ответа.

❌ Ответ 404

`{
"error": "task not found"
}`

---

## 🧪 Примеры тестирования (curl)

### Создание
`curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{"title":"Test task","done":false}'`

### Получить список
`curl http://localhost:8080/tasks`

### Получить по ID
`curl http://localhost:8080/tasks/1`

### Обновить
`curl -X PUT http://localhost:8080/tasks/1 \
-H "Content-Type: application/json" \
-d '{"title":"Updated","done":true}'`

### Удалить

`curl -X DELETE http://localhost:8080/tasks/1`