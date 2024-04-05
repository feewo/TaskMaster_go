// API

const rootUrl = "http://localhost:8080";
const exit = document.querySelector('#exit');
const tables = document.querySelectorAll('.table');
const table__errors = document.querySelectorAll('.table__error');
const table__adds = document.querySelectorAll('.table__add');
const auth = document.querySelectorAll('.auth');
let cookieLogin = document.cookie.match(/login=(.+?)(;|$)/);
let users = [];
let tasks = [];
let taskPoints = [];

async function getData(url) {
    const response = await fetch(url, {
        method: 'GET',
        headers: {
            'Authorization': document.cookie.match(/login=(.+?)(;|$)/)[1]
        }
    });

    if (!response.ok) {
        if (response.status === 401) {
            throw new Error(`Не авторизован`);
        }
        if (response.status === 403) {
            throw new Error(`Нет доступа`);
        }
        if (response.status >= 400) {
            throw new Error(`Ошибка`);
        }
    }

    return await response.json();
}

async function postData(url, data) {
    const response = await fetch(url, {
        method: 'POST',
        headers: {
            'Authorization': document.cookie.match(/login=(.+?)(;|$)/)[1],
            'Content-Type': 'application/json'
        },
        body: data,
    });

    if (!response.ok) {
        if (response.status === 401) {
            throw new Error(`Не авторизован`);
        }
        if (response.status === 403) {
            throw new Error(`Нет доступа`);
        }
        if (response.status >= 400) {
            throw new Error(`Ошибка`);
        }
    }

    return await response.json();
}

async function postDataAuth(url, data) {
    const response = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: data,
    });
    if (!response.ok) {
        if (response.status === 401) {
            throw new Error(`Не авторизован`);
        }
        if (response.status === 403) {
            throw new Error(`Нет доступа`);
        }
        if (response.status >= 400) {
            throw new Error(`Ошибка`);
        }
    }

    return await response.json();
}

async function deleteData(url) {
    const response = await fetch(url, {
        method: 'DELETE',
        headers: {
            'Authorization': document.cookie.match(/login=(.+?)(;|$)/)[1]
        }
    });

    if (!response.ok) {
        if (response.status === 401) {
            throw new Error(`Не авторизован`);
        }
        if (response.status === 403) {
            throw new Error(`Нет доступа`);
        }
        if (response.status >= 400) {
            throw new Error(`Ошибка`);
        }
    }

    return response;
}

async function putData(url, data) {
    const response = await fetch(url, {
        method: 'PUT',
        headers: {
            'Authorization': document.cookie.match(/login=(.+?)(;|$)/)[1],
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data),
    });

    if (!response.ok) {
        if (response.status === 401) {
            throw new Error(`Не авторизован`);
        }
        if (response.status === 403) {
            throw new Error(`Нет доступа`);
        }
        if (response.status >= 400) {
            throw new Error(`Ошибка`);
        }
    }

    return await response.json();
}

function getForm(form) {
    const formData = new FormData(form);
    const data = {}
    formData.forEach((value, key) => {
        console.log(value);
        if (key.includes('id') || key.includes('UserID') || key.includes('TaskID')) {
            value = +value;
        }

        // Если перебираемый элемент является чекбоксом, то пара ключ-значение будет выглядить как key: 'on/off'. Мы строку 'on/off' меняем на булево значение
        if (value === 'on') {
            value = true;
        } else if (value === 'off') {
            value = false;
        }

        data[key] = value;
    });
    
    return JSON.stringify(data)
}

const showErrorMessage = (errorSelector, errorActiveClass) => {
    const error = document.querySelector(errorSelector);

    error.classList.add(errorActiveClass);
};

// Users
const getTable = async ({url, table, errorSelector, errorActiveClass, errorForbiddenSelector, errorForbiddenActiveClass}) => {
    try {
        // к исходному массиву, который был передан в качестве аргумента (например users) добавляем новые данные, которые пришли с сервера
        // https://newwebmaster.ru/merge-arrays-in-js/
        table.push(...await getData(`${rootUrl}${url}`));
    } catch(error) {
        if (error.message === "Не авторизован") {
            logout();
        } else if (error.message === "Нет доступа") {
            showErrorMessage(errorForbiddenSelector, errorForbiddenActiveClass)
        } else {
            showErrorMessage(errorSelector, errorActiveClass)
        }
    }
};

const createTable = ({rows, tableSelector, tableHeaders, createRow}) => {
    const table = document.querySelector(tableSelector);

    table.insertAdjacentHTML('beforeend', tableHeaders);

    rows.forEach((row) => {
        table.insertAdjacentHTML('beforeend', createRow(row));
    })
};

const updateCellOrDeleteRow = ({tableSelector, url, idName, errorSelector, errorActiveClass, errorForbiddenSelector, errorForbiddenActiveClass}) => {
    const table = document.querySelector(tableSelector);

    table.addEventListener('click', (e) => {
        // Если кликнули не по кнопке "Изменить, Закрыть, Сохранить, Удалить", то ничего не делаем
        if (e.target && !e.target.classList.contains('table__btn')) return;
        e.preventDefault();

        const btnActiveClass = 'table__btn_active';
            
        if (e.target.textContent === 'Изменить') {
            const btnUpdate = e.target;
            const btnSave = btnUpdate.nextElementSibling;
            const btnClose = btnUpdate.nextElementSibling.nextElementSibling;
            const input = btnUpdate.previousElementSibling;

            btnUpdate.classList.remove(btnActiveClass);
            btnSave.classList.add(btnActiveClass);
            btnClose.classList.add(btnActiveClass);
            input.disabled = '';
            input.focus();
        }

        if (e.target.textContent === 'Закрыть') {
            const btnClose = e.target;
            const btnUpdate = btnClose.previousElementSibling.previousElementSibling;
            const btnSave = btnClose.previousElementSibling;
            const input = btnClose.previousElementSibling.previousElementSibling.previousElementSibling;
        
            btnUpdate.classList.add(btnActiveClass);
            btnSave.classList.remove(btnActiveClass);
            btnClose.classList.remove(btnActiveClass);
            input.disabled = 'disabled';
        }

        if (e.target.textContent === 'Сохранить') {
            const btnSave = e.target;
            const id = btnSave.parentElement.parentElement.parentElement.querySelector('.table__cell').querySelector('.table__input').getAttribute('data-id');
            const input = btnSave.previousElementSibling.previousElementSibling;
            console.log(input.type)
            if (input.name === idName || input.type == "number") {
                value = +input.value;
            }

            if (input.type === 'checkbox') {
                value = input.checked;
            }

            putData(`${rootUrl}${url}${id}`, {
                [input.name]: value
            })
                .then(() => {
                    const btnSave = e.target;
                    const btnUpdate = btnSave.previousElementSibling;
                    const btnClose = btnSave.nextElementSibling;
                    const input = btnSave.previousElementSibling.previousElementSibling;
                
                    btnClose.classList.remove(btnActiveClass);
                    btnSave.classList.remove(btnActiveClass);
                    btnUpdate.classList.add(btnActiveClass);
                    input.disabled = 'disabled';
                })
                .catch((error) => {
                    if (error.message === "Не авторизован") {
                        logout();
                    } else if (error.message === "Нет доступа") {
                        showErrorMessage(errorForbiddenSelector, errorForbiddenActiveClass)
                    } else {
                        showErrorMessage(errorSelector, errorActiveClass)
                    }
                });
        }

        if (e.target.textContent === 'Удалить') {
            const btnDelete = e.target;
            const row = btnDelete.parentElement.parentElement;
            const id = btnDelete.parentElement.parentElement.querySelector('[data-id]').getAttribute('data-id');

            deleteData(`${rootUrl}${url}${id}`)
                .then(() => {
                    row.remove();
                })
                .catch((error) => {
                    if (error.message === "Не авторизован") {
                        logout();
                    } else if (error.message === "Нет доступа") {
                        showErrorMessage(errorForbiddenSelector, errorForbiddenActiveClass)
                    } else {
                        showErrorMessage(errorSelector, errorActiveClass)
                    }
                });
        }
    })
};

const addRow = ({formSelector, url, tableSelector, createRow, errorSelector, errorActiveClass, errorForbiddenSelector, errorForbiddenActiveClass}) => {
    const form = document.querySelector(formSelector);

    form.addEventListener('submit', async (e) => {
        e.preventDefault();

        try {
            const row = await postData(`${rootUrl}${url}`, getForm(form));
            console.log(row)
            const table = document.querySelector(tableSelector);

            table.insertAdjacentHTML('beforeend', createRow(row));
        } catch(error) {
            if (error.message === "Не авторизован") {
                logout();
            } else if (error.message === "Нет доступа") {
                showErrorMessage(errorForbiddenSelector, errorForbiddenActiveClass)
            } else {
                showErrorMessage(errorSelector, errorActiveClass)
            }
        } finally {
            e.target.reset();
        }
    });
};

// Login
const initIsAuthorized = async () => {
    exit.classList.add('active');
    tables.forEach(table => {
        table.classList.add('active');
    });
    table__adds.forEach(table__add => {
        table__add.classList.add('table__add_active');
    });
    auth.forEach(item => {
        item.classList.remove('active');
    });

    await getTable({
        url: '/user', 
        table: users, 
        errorSelector: '.table__error-user', 
        errorActiveClass: 'table__error-user_active', 
        errorForbiddenSelector: '.table__error-auth-user', 
        errorForbiddenActiveClass: 'table__error-auth-user_active',
    });
    createTable({
        rows: users, 
        tableSelector: '.user__table',
        tableHeaders: userHeaders,
        createRow: createRowUser,
    });
    updateCellOrDeleteRow({
        tableSelector: '.user__table', 
        url: '/user/', 
        idName: 'ID',
        errorSelector: '.table__error-user', 
        errorActiveClass: 'table__error-user_active', 
        errorForbiddenSelector: '.table__error-auth-user', 
        errorForbiddenActiveClass: 'table__error-auth-user_active',
    });
    addRow({
        formSelector: '.add__user', 
        url: '/user', 
        tableSelector: '.user__table', 
        createRow: createRowUser,
        errorSelector: '.table__error-user', 
        errorActiveClass: 'table__error-user_active', 
        errorForbiddenSelector: '.table__error-auth-user', 
        errorForbiddenActiveClass: 'table__error-auth-user_active',
    });

    await getTable({
        url: '/task', 
        table: tasks, 
        errorSelector: '.table__error-task', 
        errorActiveClass: 'table__error-task_active', 
        errorForbiddenSelector: '.table__error-auth-task', 
        errorForbiddenActiveClass: 'table__error-auth-task_active',
    });
    createTable({
        rows: tasks, 
        tableSelector: '.task__table',
        tableHeaders: taskHeaders,
        createRow: createRowTask,
    });
    updateCellOrDeleteRow({
        tableSelector: '.task__table', 
        url: '/task/', 
        idName: 'ID',
        errorSelector: '.table__error-task', 
        errorActiveClass: 'table__error-task_active', 
        errorForbiddenSelector: '.table__error-auth-task', 
        errorForbiddenActiveClass: 'table__error-auth-task_active',
    });
    addRow({
        formSelector: '.add__task', 
        url: '/task', 
        tableSelector: '.task__table', 
        createRow: createRowTask,
        errorSelector: '.table__error-task', 
        errorActiveClass: 'table__error-task_active', 
        errorForbiddenSelector: '.table__error-auth-task', 
        errorForbiddenActiveClass: 'table__error-auth-task_active',
    });

    await getTable({
        url: '/taskpoint', 
        table: taskPoints, 
        errorSelector: '.table__error-task-point', 
        errorActiveClass: 'table__error-task-point_active', 
        errorForbiddenSelector: '.table__error-auth-task-point', 
        errorForbiddenActiveClass: 'table__error-auth-task-point_active',
    });
    createTable({
        rows: taskPoints, 
        tableSelector: '.task-point__table',
        tableHeaders: taskPointHeaders,
        createRow: createRowTaskPoint,
    });
    updateCellOrDeleteRow({
        tableSelector: '.task-point__table', 
        url: '/taskpoint/', 
        idName: 'ID',
        errorSelector: '.table__error-task-point', 
        errorActiveClass: 'table__error-task-point_active', 
        errorForbiddenSelector: '.table__error-auth-task-point', 
        errorForbiddenActiveClass: 'table__error-auth-task-point_active',
    });
    addRow({
        formSelector: '.add__task-point', 
        url: '/taskpoint', 
        tableSelector: '.task-point__table', 
        createRow: createRowTaskPoint,
        errorSelector: '.table__error-task-point', 
        errorActiveClass: 'table__error-task-point_active', 
        errorForbiddenSelector: '.table__error-auth-task-point', 
        errorForbiddenActiveClass: 'table__error-auth-task-point_active',
    });
};

const login = (form) => {
    postDataAuth(`${rootUrl}/token`, getForm(form))
        .then((result) => {
            document.cookie = `login=${result["Token"]}`;
            cookieLogin = document.cookie.match(/login=(.+?)(;|$)/);

            initIsAuthorized();
        })
        .catch((error) => {
            formLog = document.querySelector('#formLog')
            formLog.innerHTML +="<div class='error'>Неправильные данные</div>"
            if (error.message) {
                console.log(error);
            }
        });
};

// Registration
const register = (form) => {
    postDataAuth(`${rootUrl}/user`, getForm(form))
        .then(() => {
            formReg = document.querySelector('#formReg')
            formReg.innerHTML +="<div class='success'>Успешно! Бегите авторизовываться :)</div>"
        })
        .catch((error) => {
            formReg = document.querySelector('#formReg')
            formReg.innerHTML +="<div class='error'>Неправильные данные</div>"
            if (error.message) {
                console.log(error);
            }
        });
};

const handleLoginAndRegister = () => {
    const forms = document.querySelectorAll('.auth__form');

    forms.forEach(form => {
        form.addEventListener('submit', (event) => {
            event.preventDefault();

            switch (form.id) {
                case "formReg":
                    register(form);
                    break;
                case "formLog":
                    login(form);
                    break;
            }
        })
    });
};

// Logout
const cleanCookie = () => {
    var cookies = document.cookie.split("; ");
    for (var c = 0; c < cookies.length; c++) {
        var d = window.location.hostname.split(".");
        while (d.length > 0) {
            var cookieBase = encodeURIComponent(cookies[c].split(";")[0].split("=")[0]) + '=; expires=Thu, 01-Jan-1970 00:00:01 GMT; domain=' + d.join('.') + ' ;path=';
            var p = location.pathname.split('/');
            document.cookie = cookieBase + '/';
            while (p.length > 0) {
                document.cookie = cookieBase + p.join('/');
                p.pop();
            };
            d.shift();
        }
    }
};

const initIsNotAuthorized = () => {
    exit.classList.remove('active');

    tables.forEach(table => {
        table.classList.remove('active');
    });
    table__errors.forEach(table__error => {
        table__error.classList.remove('table__error_active', 'table__error-auth-user_active', 'table__error-auth-task_active', 'table__error-auth-task-point_active');
    })
    table__adds.forEach(table__add => {
        table__add.classList.remove('table__add_active');
    });
    auth.forEach(item => {
        item.classList.add('active');
    });
};

const clearTables = () => {
    const rows = document.querySelectorAll('.table tbody');

    rows.forEach((row) => row.remove());
};

const logout = () => {
    deleteData(`${rootUrl}/token`)
        .then(() => {
            users = [];
            tasks = [];
            taskPoints = [];

            clearTables();
            cleanCookie();
            initIsNotAuthorized();
        });
};

const handleLogout = () => {
    exit.addEventListener('click', () => {
        logout();
    });
};

// Структура таблиц
const userHeaders = `
    <tr class="table__row">
        <th class="table__head">ID</th>
        <th class="table__head">Логин</th>
        <th class="table__head">Почта</th>
        <th class="table__head">Роль</th>
        <th class="table__head">Пароль</th>
        <th class="table__head">Удаление</th>
    </tr>`;

const createRowUser = (row) => {
    return `<tr class="table__row">
                <td class="table__cell">
                    <form class="table__form">
                        <input type="number" class="auth__input table__input" name="ID" value=${row.ID} data-id=${row.ID} disabled>
                        <button class="table__btn auth__input table__btn_active">Изменить</button>
                        <button class="table__btn auth__input" type="submit">Сохранить</button>
                        <button class="table__btn auth__input">Закрыть</button>
                    </form>
                </td>
                <td class="table__cell">
                    <form class="table__form">
                        <input type="text" class="auth__input table__input" name="Login" value=${row.Login} disabled>
                        <button class="table__btn auth__input table__btn_active">Изменить</button>
                        <button class="table__btn auth__input" type="submit">Сохранить</button>
                        <button class="table__btn auth__input">Закрыть</button>
                    </form>
                </td>
                <td class="table__cell">
                    <form class="table__form">
                        <input type="text" class="auth__input table__input" name="Email" value=${row.Email} disabled>
                        <button class="table__btn auth__input table__btn_active">Изменить</button>
                        <button class="table__btn auth__input" type="submit">Сохранить</button>
                        <button class="table__btn auth__input">Закрыть</button>
                    </form>
                </td>
                <td class="table__cell">
                    <form class="table__form">
                        <input type="text" class="auth__input table__input" name="Role" value=${row.Role} disabled>
                        <button class="table__btn auth__input table__btn_active">Изменить</button>
                        <button class="table__btn auth__input" type="submit">Сохранить</button>
                        <button class="table__btn auth__input">Закрыть</button>
                    </form>
                </td>
                <td class="table__cell">
                    <form class="table__form">
                        <input type="password" class="auth__input table__input" name="Password" value=${row.Password} disabled>
                        <button class="table__btn auth__input table__btn_active">Изменить</button>
                        <button class="table__btn auth__input" type="submit">Сохранить</button>
                        <button class="table__btn auth__input">Закрыть</button>
                    </form>
                </td>
                <td class="table__cell">
                    <button class="table__btn auth__input table__btn_active">Удалить</button>
                </td>
            </tr>`;
};

const taskHeaders = `
    <tr class="table__row">
        <th class="table__head">ID</th>
        <th class="table__head">Название</th>
        <th class="table__head">Статус</th>
        <th class="table__head">UserID</th>
        <th class="table__head">Удаление</th>
    </tr>`;

const createRowTask = (row) => {
    return `<tr class="table__row">
                <td class="table__cell">
                    <form class="table__form">
                        <input type="number" class="auth__input table__input" name="ID" value=${row.ID} data-id=${row.ID} disabled>
                        <button class="table__btn auth__input table__btn_active">Изменить</button>
                        <button class="table__btn auth__input" type="submit">Сохранить</button>
                        <button class="table__btn auth__input">Закрыть</button>
                    </form>
                </td>
                <td class="table__cell">
                    <form class="table__form">
                        <input type="text" class="auth__input table__input" name="Title" value=${row.Title} disabled>
                        <button class="table__btn auth__input table__btn_active">Изменить</button>
                        <button class="table__btn auth__input" type="submit">Сохранить</button>
                        <button class="table__btn auth__input">Закрыть</button>
                    </form>
                </td>
                <td class="table__cell">
                    <form class="table__form">
                        <input type="checkbox" class="auth__input table__input" name="Ready" ${row.Ready && 'checked'}  disabled>
                        <button class="table__btn auth__input table__btn_active">Изменить</button>
                        <button class="table__btn auth__input" type="submit">Сохранить</button>
                        <button class="table__btn auth__input">Закрыть</button>
                    </form>
                </td>
                <td class="table__cell">
                    <form class="table__form">
                        <input type="number" class="auth__input table__input" name="UserID" value=${row.UserID} disabled>
                        <button class="table__btn auth__input table__btn_active">Изменить</button>
                        <button class="table__btn auth__input" type="submit">Сохранить</button>
                        <button class="table__btn auth__input">Закрыть</button>
                    </form>
                </td>
                <td class="table__cell">
                    <button class="table__btn auth__input table__btn_active">Удалить</button>
                </td>
            </tr>`;
};

const taskPointHeaders = `
    <tr class="table__row">
        <th class="table__head">ID</th>
        <th class="table__head">Название</th>
        <th class="table__head">Статус</th>
        <th class="table__head">TaskID</th>
        <th class="table__head">Удалить</th>
    </tr>`;

const createRowTaskPoint = (row) => {
    return `<tr class="table__row">
                <td class="table__cell">
                    <form class="table__form">
                        <input type="number" class="auth__input table__input" name="ID" value=${row.ID} data-id=${row.ID} disabled>
                        <button class="table__btn auth__input table__btn_active">Изменить</button>
                        <button class="table__btn auth__input" type="submit">Сохранить</button>
                        <button class="table__btn auth__input">Закрыть</button>
                    </form>
                </td>
                <td class="table__cell">
                    <form class="table__form">
                        <input type="text" class="auth__input table__input" name="Title" value=${row.Title} disabled>
                        <button class="table__btn auth__input table__btn_active">Изменить</button>
                        <button class="table__btn auth__input" type="submit">Сохранить</button>
                        <button class="table__btn auth__input">Закрыть</button>
                    </form>
                </td>
                <td class="table__cell">
                    <form class="table__form">
                        <input type="checkbox" class="auth__input table__input" name="Ready" ${row.Ready && 'checked'}  disabled>
                        <button class="table__btn auth__input table__btn_active">Изменить</button>
                        <button class="table__btn auth__input" type="submit">Сохранить</button>
                        <button class="table__btn auth__input">Закрыть</button>
                    </form>
                </td>
                <td class="table__cell">
                    <form class="table__form">
                        <input type="text" class="auth__input table__input" name="TaskID" value=${row.TaskID} disabled>
                        <button class="table__btn auth__input table__btn_active">Изменить</button>
                        <button class="table__btn auth__input" type="submit">Сохранить</button>
                        <button class="table__btn auth__input">Закрыть</button>
                    </form>
                </td>
                <td class="table__cell">
                    <button class="table__btn auth__input table__btn_active">Удалить</button>
                </td>
            </tr>`;
};

document.addEventListener('DOMContentLoaded', () => {
    handleLoginAndRegister();
    handleLogout();

    if (cookieLogin) {
        initIsAuthorized();
    } else {
        initIsNotAuthorized();
    }
})

