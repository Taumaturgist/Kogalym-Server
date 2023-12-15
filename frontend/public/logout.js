window.logout = () => {
    return {
        isActiveTab: 'groups',
        changeActiveTab(tab) {
            this.isActiveTab = tab;
        },

        logoutApi: function () {
            fetch('/logout', {
                method: 'POST', headers: {
                    'Content-Type': 'application/json',
                    'X-CSRF-Token': this.csrf,
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
                    if (data && data.data) {
                        this.error = data.data;
                    }
                });
        },
    };
};