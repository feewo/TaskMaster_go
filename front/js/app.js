// API

const rootUrl = "http://localhost:8080";
const exit = document.querySelector('#exit');
const table = document.querySelector('.table');
const auth = document.querySelectorAll('.auth');

async function getData(url) {
    const response = await fetch('localhost/user', {
        method: 'GET',
        headers: {
            'Authorization': document.cookie.match(/login=(.+?)(;|$)/)[1]
        }
    });

    if (!response.ok) {
        if (response.status >= 400) {
            throw new Error(`Ошибка`);
        }
    }

    return await response.json();
}

async function postData(url, data) {
    const response = await fetch(url, {
        method: 'POST',
        body: data,
    });
    window.result = await response.json();
    if (!response.ok) {
        if (response.status >= 400) {
            throw new Error(`Ошибка`);
        }
    }
}

async function deleteData(url) {
    const response = await fetch(url, {
        method: 'DELETE',
        headers: {
            'Authorization': document.cookie.match(/login=(.+?)(;|$)/)[1]
        }
    });

    if (!response.ok) {
        if (response.status >= 400) {
            throw new Error(`Ошибка`);
        }
    }

    return await response.json();
}

async function putData(url, data) {
    const response = await fetch(url, {
        method: 'GET',
        headers: {
            'Authorization': document.cookie.match(/login=(.+?)(;|$)/)[1]
        },
        body: data,
    });

    if (!response.ok) {
        if (response.status >= 400) {
            throw new Error(`Ошибка`);
        }
    }

    return await response.json();
}

document.addEventListener('DOMContentLoaded', () => {
    var cookieLogin = document.cookie.match(/login=(.+?)(;|$)/);
    

    if (cookieLogin) {
        exit.classList.add('active');
        table.classList.add('active');
        auth.forEach(item => {
            item.classList.remove('active');
        })
    } else {
        exit.classList.remove('active');
        table.classList.remove('active');
        auth.forEach(item => {
            item.classList.add('active');
        })
    }
})

// Form
// const fromBtns = document.querySelectorAll('[data-form]');
const forms = document.querySelectorAll('.auth__form');

forms.forEach(form => {
    form.addEventListener('submit', (event) => {
        event.preventDefault();
        console.log(form.id)

        const formId  = form.id;
        
        switch (formId) {
            case "formReg":
                registration(form);
                break;
            case "formLog":
                login(form);
                break;
        }
    })
})

function getForm(form) {
    const fd = new FormData(form)
    const data = {}
    fd.forEach((v, k) => data[k] = v)
    return JSON.stringify(data)
}

// Registration
const registration = (form) => {
    postData(`${rootUrl}/user`, getForm(form))
    .then((result) => {
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
}

// Login
const login = (form) => {
    postData(`${rootUrl}/token`, getForm(form))
    .then((result) => {
        document.cookie = `login=${window.result["Token"]}`;
        cookieLogin = document.cookie.match(/login=(.+?)(;|$)/);
        exit.classList.add('active');
        table.classList.add('active');
        auth.forEach(item => {
            item.classList.remove('active');
        })
    })
    .catch((error) => {
        formLog = document.querySelector('#formLog')
        formLog.innerHTML +="<div class='error'>Неправильные данные</div>"
        if (error.message) {
            console.log(error);
        }
    });
}

// Exit
exit.addEventListener('click', () => {
    deleteData(`${rootUrl}/token`)
    .then((result) => {
        exit.classList.remove('active');
        table.classList.remove('active');
        auth.forEach(item => {
            item.classList.add('active');
        })
    })
    cleanCookie();
})

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
}