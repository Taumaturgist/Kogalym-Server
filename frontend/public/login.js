window.loginApi = function () {
    fetch('/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'X-CSRF-Token': this.csrf,
        },
        body: JSON.stringify({
            login: this.login,
            password: this.password,
        }),
    })
        .then(response => {
            if (response.ok) {
                // Если авторизация успешна, перенаправьте пользователя на главную страницу
                window.location.href = '/';
            } else {
                return response.json(); // Парсим ответ как JSON
            }
        })
        .then(data => {
            if (data && data.data) {
                this.error = data.data;
            }
        });
}