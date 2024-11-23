const express = require('express');
const bodyParser = require('body-parser');
const app = express();
const port = 3000;

// Middleware для обработки JSON тела запроса
app.use(bodyParser.json());

app.post('/register', (req, res) => {
    const userData = req.body;

    // Выведем данные пользователя в консоль (их можно отправить в базу данных)
    console.log('Received user data:', userData);

    // Ответим пользователю, что регистрация прошла успешно
    res.json({ message: 'Registration successful', user: userData });
});

app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});
