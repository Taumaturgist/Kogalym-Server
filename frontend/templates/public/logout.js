window.logoutApi = function () {
    fetch('/logout', {
        method: 'POST', headers: {
            'Content-Type': 'application/json',
        },
    })
        .then(response => {
            if (response.ok) {
                // Если авторизация успешна, перенаправьте пользователя на главную страницу
                window.location.href = '/login';
            } else {
                return response.json(); // Парсим ответ как JSON
            }
        })
        .then(data => {
            if (data && data.error) {
                this.error = data.error;
            }
        });
}