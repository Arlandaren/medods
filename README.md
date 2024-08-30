Установка

1. Скачать docker engine
2. из корневой директори прописать команды docker-compose build docker-compose up
3. Сервер запущен на порту 8081!!!


Endpoint 1:
Метод: GET
Путь: api/auth/token/:user_id
Вывод: {"access": string, "refresh": string}
Endpoint 2:
Метод: POST
Путь: api/auth/token/refresh
Параметры запроса: Refresh токен в формате JSON: {"access_token": string,"refresh_token": string}
Вывод: {"access": string, "refresh": string}
