window.loginApi = function () {
    fetch('/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
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
            if (data && data.error) {
                this.error = data.error;
            } else {
                this.error = 'Authentication failed. Please try again.';
            }
        })
        .catch(error => {
            this.error = 'An error occurred. Please try again later.';
        });
}