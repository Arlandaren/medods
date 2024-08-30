<h2>Установка</h2>
        <ol>
            <li><strong>Скачать Docker Engine</strong>
                <p>Инструкции по установке Docker можно найти <a href="https://docs.docker.com/get-docker/" target="_blank">здесь</a>.</p>
            </li>
            <li><strong>Собрать и запустить контейнеры</strong>
                <p>Из корневой директории проекта выполните следующие команды:</p>
                <pre><code>docker-compose build
docker-compose up</code></pre>
            </li>
            <li><strong>Сервер запущен на порту 8081</strong></li>
        </ol>

<h2>API</h2>

        <h3>Endpoint 1</h3>
        <ul>
            <li><strong>Метод:</strong> GET</li>
            <li><strong>Путь:</strong> /api/auth/token/:user_id</li>
            <li><strong>Вывод:</strong></li>
        </ul>
        <pre><code>{
    "access": "string",
    "refresh": "string"
}</code></pre>

        <h3>Endpoint 2</h3>
        <ul>
            <li><strong>Метод:</strong> POST</li>
            <li><strong>Путь:</strong> /api/auth/token/refresh</li>
            <li><strong>Параметры запроса:</strong> Refresh токен в формате JSON:</li>
        </ul>
        <pre><code>{
    "access_token": "string",
    "refresh_token": "string"
}</code></pre>
        <ul>
            <li><strong>Вывод:</strong></li>
        </ul>
        <pre><code>{
    "access": "string",
    "refresh": "string"
}</code></pre>
