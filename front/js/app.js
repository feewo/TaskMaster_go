// API
const rootUrl = '/api';
var idUser = 0;

async function postData(url, data) {
    const response = await fetch(url, {
        headers: {
            'X-API-KEY': 'FC52783F63184532B379EECD56DFC009E0131854354C4FA293EC5581CC6547F7',
            'Authorization': document.cookie.match(/login=(.+?)(;|$)/)[1]
        },
        method: 'POST',
        body: data,
    });

    if (!response.ok) {
        throw new Error('Ошибка');
    }

    return await response.json();
}

async function putData(url, data) {
    const response = await fetch(url, {
        headers: {
            'X-API-KEY': 'FC52783F63184532B379EECD56DFC009E0131854354C4FA293EC5581CC6547F7',
            'Authorization': document.cookie.match(/login=(.+?)(;|$)/)[1]
        },
        method: 'PUT',
        body: data,
    });

    if (!response.ok) {
        throw new Error('Ошибка');
    }

    return await response.json();
}

async function getData(url) {
    const response = await fetch(url, {
        headers: {
            'X-API-KEY': 'FC52783F63184532B379EECD56DFC009E0131854354C4FA293EC5581CC6547F7',
            'Authorization': document.cookie.match(/login=(.+?)(;|$)/)[1]
        },
        method: 'GET',
    });

    if (!response.ok) {
        throw new Error('Ошибка');
    }

    return await response.json();
}

document.addEventListener('DOMContentLoaded', () => {
    const url = document.URL;
    const urlArr = url.split('/');

    const taskItems = document.querySelectorAll('.task__item');

    if (urlArr[urlArr.length - 1] == 'task.html') {
        getTaskData();

        if (taskItems.length > 0) {
            interaction(taskItems);
        }
    }

    if (urlArr[urlArr.length - 1] == 'account.html') {
        getAccData();
        exit();
    }

    forms();
})

// Main functions

const addTask = (idUser) => {
    const taskCreate = document.querySelector('#taskCreate');

    taskCreate.addEventListener('click', () => {
        const taskTop = document.querySelector('.task__top');

        taskTop.insertAdjacentHTML('afterend', `
        <div class="task__item">
            <div class="task__item-head">
                <div class="task__item-title">Текст задачи</div>
                <input type="text" class="task__item-inputtitle hide" value="Текст задачи">
                <div class="task__item-icons">
                    <button class="task__item-btn task__item--edit">
                        <svg class="task__item-icon">
                            <use xlink:href="#change"></use>
                        </svg>
                    </button>

                    <button class="task__item-btn task__item-btn--edit task__item-btn--ready">
                        <svg class="task__item-icon">
                            <use xlink:href="#ready"></use>
                        </svg>
                    </button>

                    <button class="task__item-btn task__item-btn--edit task__item-btn--delete">
                        <svg class="task__item-icon">
                            <use xlink:href="#delete"></use>
                        </svg>
                    </button>
                </div>
            </div>

            <div class="task__item-container">
                <div class="task__item-point">
                    <div class="task__item-left">
                        <input type="checkbox" class="task__item-checkbox" id="point1">
                        <label for="point1" class="task__item-flag"></label>
                        <label for="point1" class="task__item-label">Текст подзадачи</label>
                        <input type="text" class="task__item-input hide" value="Текст подзадачи">
                    </div>

                    <button class="task__item-right">
                        <svg class="task__item-delete">
                            <use xlink:href="#delete"></use>
                        </svg>
                    </button>
                </div>

                <button class="task__add">
                    <svg class="task__add-icon">
                        <use xlink:href="#plus"></use>
                    </svg>

                    <div class="task__add-text">Добавить подзадачу</div>
                </button>
            </div>
        </div>
        `)

        addTaskRequest(idUser);

        const taskItems = document.querySelectorAll('.task__item');
        interaction(taskItems[0]);
    });
}

const interaction = (taskItems) => {
    if (taskItems.length > 1) {
        taskItems.forEach(taskItem => {
            taskItem.addEventListener('click', (event) => {
                const clickEdit = event.target.closest('.task__item--edit');
                const clickDeletePoint = event.target.closest('.task__item-right');
                const clickDelete = event.target.closest('.task__item-btn--delete');
                const clickCanel = event.target.closest('.task__item-btn--canel');
                const clickReady = event.target.closest('.task__item-btn--ready');
                const clickAdd = event.target.closest('.task__add');
    
                if (clickEdit) {
                    editMode(clickEdit);
                }
                if (clickDeletePoint) {
                    deletePoint(clickDeletePoint);
                }
                if (clickDelete) {
                    deleteTask(clickDelete);
                }
                if (clickReady) {
                    const btns = clickReady.parentElement;
                    const itemPoints = btns.parentElement.nextElementSibling.querySelectorAll('.task__item-point');
                    const taskTitleInput = btns.previousElementSibling;
                    const taskTitle = taskTitleInput.previousElementSibling;
                    
                    defaultMode(clickReady);
                    pointDefaultValue(itemPoints, taskTitleInput, taskTitle);
                }
                if (clickAdd) {
                    addPoint(clickAdd);
                }
            })
        })
        return;
    }
    
    const taskItem = taskItems;
    taskItem.addEventListener('click', (event) => {
        const clickEdit = event.target.closest('.task__item--edit');
        const clickDeletePoint = event.target.closest('.task__item-right');
        const clickDelete = event.target.closest('.task__item-btn--delete');
        const clickCanel = event.target.closest('.task__item-btn--canel');
        const clickReady = event.target.closest('.task__item-btn--ready');
        const clickAdd = event.target.closest('.task__add');

        if (clickEdit) {
            editMode(clickEdit);
        }
        if (clickDeletePoint) {
            deletePoint(clickDeletePoint);
        }
        if (clickDelete) {
            deleteTask(clickDelete);
        }
        if (clickReady) {
            const btns = clickReady.parentElement;
            const itemPoints = btns.parentElement.nextElementSibling.querySelectorAll('.task__item-point');
            const taskTitleInput = btns.previousElementSibling;
            const taskTitle = taskTitleInput.previousElementSibling;
            
            defaultMode(clickReady);
            pointDefaultValue(itemPoints, taskTitleInput, taskTitle);
        }
        if (clickCanel) {
            const btns = clickCanel.parentElement;
            const itemPoints = btns.parentElement.nextElementSibling.querySelectorAll('.task__item-point');
            const taskTitleInput = btns.previousElementSibling;
            const taskTitle = taskTitleInput.previousElementSibling;

            defaultMode(clickCanel);
            pointDefault(itemPoints, taskTitleInput, taskTitle);
        }
        if (clickAdd) {
            addPoint(clickAdd);
        }
    })
}

const forms = () => {
    const formBtns = document.querySelectorAll('[data-form]');

    formBtns.forEach(formBtn => {
        formBtn.addEventListener('click', (event) => {
            event.preventDefault();
            const formId = formBtn.dataset.form;
            const form = document.querySelector(formId);
            
            switch(formId) {
                case "#regForm":
                    registration(form);
                    break;
                case "#logForm":
                    login(form);
                    break;
                case "#formCreateTask":
                    createTask(form);
                    break;
                case "#formCreatePoint":
                    createPoint(form);
                    break;
                case "#formEdit":
                    edit(form);
                    break;
                case "#accountSave":
                    account(form);
                    break;
            }
        })
    })
}

const exit = () => {
    const btn = document.querySelector('#exit');

    btn.addEventListener('click', (event) => {
        event.preventDefault();

        cleanCookie();
        window.location.replace("/");
    })
}

// API functoins

const submitFormReg = (formData) => {
    postData(`${rootUrl}/user`, formData)
    .then (result => {
        window.location.replace("login.html");
    })
    .catch (error => {
        const err = document.querySelector('#regError');
        notice(err);
    })
}   

const submitFormLog = (formData) => {
    postData(`${rootUrl}/token`, formData)
    .then (result => {
        document.cookie = `login=${result["Token"]}`;
        window.location.replace("account.html");
    })
    .catch (error => {
        const err = document.querySelector('#logError');
        notice(err);
    })
}   

const submitFormAcc = (formData) => {
    putData(`${rootUrl}/user/${idUser}`, formData)
    .then (result => {
        const err = document.querySelector('#accSuccess');
        notice(err);
    })
    .catch (error => {
        const err = document.querySelector('#accError');
        notice(err);
    })
}   

const getAccData = () => {
    getData(`${rootUrl}/user_token`)
    .then(result => {
        loadAccInfo(result);
        idUser = result["ID"];
    })
    .catch(error => {
        console.log(error);
    })
}

const getTaskData = () => {
    getData(`${rootUrl}/user_token`)
    .then(result => {
        const idUserTask = result["ID"];
        loadName(result["Name"]);
        loadTaskInfo(idUserTask);
        addTask(idUserTask);
    })
    .catch(error => {
        console.log(error);
    })
}

const addTaskRequest = (idUser) => {
    formDataTask = `{
        "Title":"Текст задачи",
        "UserID":${idUser}
    }
    `
    postData(`${rootUrl}/task`, formDataTask)
    .then(resultTask => {
        formDataTaskpoint = `{
            "Title":"Текст подзадачи",
            "Ready":false,
            "TaskID":${resultTask["ID"]}
        }
        `
    
        postData(`${rootUrl}/taskpoint`, formDataTaskpoint)
        .then(resultTaskpoint => {
            console.log('success');
        })
        .catch(error => {
            const err = document.querySelector('#taskError');
            notice(err);
        })
    })
    .catch(error => {
        const err = document.querySelector('#taskError');
        notice(err);
    })
}

const changeCheckbox = () => {
    const inputs = document.querySelectorAll('.task__item-checkbox');

    inputs.forEach(input => {
        const idInput = input.id.slice(5);
        const titlePoint = input.nextElementSibling.nextElementSibling.textContent;
        const taskId = Number(input.parentElement.parentElement.parentElement.id.slice(13));

        input.addEventListener('input', () => {
            const formData = `
            {
                "Title":"${titlePoint}",
                "Ready":${input.checked},
                "TaskID":${taskId}
            }
            `

            putData(`${rootUrl}/taskpoint/${idInput}`, formData)
            .then(result => {
                console.log(result);
            })
            .catch(error => {
                console.log(error);
            })
        })  
    })
}

// Supportive functions

const loadName = (name) => {
    const headName = document.querySelector('.head__name');

    headName.insertAdjacentHTML('beforeend', `
        ${name}
    `)
}

const loadTaskInfo = (id) => {
    getData(`${rootUrl}/task_user/${id}`)
    .then (tasks => {
        const loader = document.querySelector('.loader__container');

        tasks.forEach(task => {
            loader.insertAdjacentHTML('afterend', `
            <div class="task__item">
                <div class="task__item-head">
                    <div class="task__item-title">${task["Title"]}</div>
                    <input type="text" class="task__item-inputtitle hide" value="${task["Title"]}">
                    <div class="task__item-icons">
                        <button class="task__item-btn task__item--edit">
                            <svg class="task__item-icon">
                                <use xlink:href="#change"></use>
                            </svg>
                        </button>

                        <button class="task__item-btn task__item-btn--edit task__item-btn--ready">
                            <svg class="task__item-icon">
                                <use xlink:href="#ready"></use>
                            </svg>
                        </button>

                        <button class="task__item-btn task__item-btn--edit task__item-btn--canel">
                            <svg class="task__item-icon">
                                <use xlink:href="#canel"></use>
                            </svg>
                        </button>

                        <button class="task__item-btn task__item-btn--edit task__item-btn--delete">
                            <svg class="task__item-icon">
                                <use xlink:href="#delete"></use>
                            </svg>
                        </button>
                    </div>
                </div>

                <div class="task__item-container" id="taskContainer${task["ID"]}">
                    <button class="task__add">
                        <svg class="task__add-icon">
                            <use xlink:href="#plus"></use>
                        </svg>

                        <div class="task__add-text">Добавить подзадачу</div>
                    </button>
                </div>
            </div>
            `)

            const taskId = task["ID"];
            getData(`${rootUrl}/taskpoint_task/${taskId}`)
            .then(taskpoints => {
                const taskContainer = document.querySelector(`#taskContainer${task["ID"]}`);

                taskpoints.forEach(taskpoint => {
                    let check;
                    taskpoint["Ready"] ? check = 'checked' : check = '';

                    taskContainer.insertAdjacentHTML('beforeend', `
                    <div class="task__item-point">
                        <div class="task__item-left">
                            <input type="checkbox" class="task__item-checkbox" id="point${taskpoint["ID"]}" ${check}>
                            <label for="point${taskpoint["ID"]}" class="task__item-flag" value="1"></label>
                            <label for="point${taskpoint["ID"]}" class="task__item-label">${taskpoint["Title"]}</label>
                            <input type="text" class="task__item-input hide" value="${taskpoint["Title"]}">
                        </div>

                        <button class="task__item-right">
                            <svg class="task__item-delete">
                                <use xlink:href="#delete"></use>
                            </svg>
                        </button>
                    </div>
                    `)
                })

                changeCheckbox();
            })
            .catch(error => {
                console.log(error);
            })
        })

        loader.style.display = "none";
    })
    .catch(error => {
        console.log(error);
    })
}

const loadAccInfo = (data) => {
    const loader = document.querySelector('.loader__container');

    loader.insertAdjacentHTML('afterend', `
    <form action="" method="post" class="account__body" id="accountSave">
        <div class="account__block account__block--left">
            <h4 class="account__title">Общие данные</h4>

            <div class="auth__group">
                <input minlength="1" maxlength="225" type="text" class="auth__input" id="accountSurname" name="surname" value="${data["Surname"]}" placeholder=" " required>
                <label for="accountSurname" class="auth__label">Фамилия</label>   
            </div>

            <div class="auth__group">
                <input minlength="1" maxlength="225" type="text" class="auth__input" id="accountName" name="name" value="${data["Name"]}" placeholder=" " required>
                <label for="accountName" class="auth__label">Имя</label>
            </div>

            <div class="auth__group">
                <input minlength="1" maxlength="225" type="text" class="auth__input" id="accountPatronymic" name="patronymic" value="${data["Patronymic"]}" placeholder=" " required>
                <label for="accountPatronymic" class="auth__label">Отчество</label>
            </div>
        </div>

        <div class="account__block account__block--right">
            <h4 class="account__title">Безопасность</h4>

            <div class="auth__group">
                <input minlength="1" maxlength="225" type="text" class="auth__input" id="accountEmail" name="email" value="${data["Email"]}" placeholder=" " required>
                <label for="accountEmail" class="auth__label">Почта</label>
            </div>

            <div class="auth__group">
                <input minlength="1" maxlength="225" type="text" class="auth__input" id="accountLogin" name="login" value="${data["Login"]}" placeholder=" " required>
                <label for="accountLogin" class="auth__label">Логин</label>
            </div>
        </div>
    </form>
    `)

    loader.style.display = "none";
}

const addPoint = (btn) => {
    const container = btn.parentElement;

    btn.insertAdjacentHTML('beforebegin' , `
    <div class="task__item-point task__item-add">
        <div class="task__item-left">
            <input type="checkbox" class="task__item-checkbox" id="point6">
            <label for="point6" class="task__item-flag"></label>
            <label for="point2" class="task__item-label hide"></label>
            <input type="text" class="task__item-addinput">
        </div>

        <button class="task__item-right">
            <svg class="task__item-delete">
                <use xlink:href="#delete"></use>
            </svg>
        </button>
    </div>
    `)

    container.scrollTop = container.scrollHeight;
}

const deleteTask = (btn) => {
    const taskItem = btn.parentElement.parentElement.parentElement;
    
    taskItem.style.opacity = '0';
    setTimeout(() => {
        taskItem.classList.add('hide');
    }, 200)
}

const deletePoint = (btn) => {
    const point = btn.parentElement;

    point.style.opacity = '0';
    setTimeout(() => {
        point.classList.add('hide');
    }, 200)
}

const editMode = (btnEdit) => {
    const btns = btnEdit.parentElement;
    const itemPoints = btns.parentElement.nextElementSibling.querySelectorAll('.task__item-point');
    const btnsEdit = btns.querySelectorAll('.task__item-btn--edit');

    const taskItem = btns.parentElement.parentElement;
    const btnsDelete = taskItem.querySelectorAll('.task__item-right');

    const taskAdd = btns.parentElement.nextElementSibling.querySelector('.task__add');
    const taskTitleInput = btns.previousElementSibling;
    const taskTitle = taskTitleInput.previousElementSibling;

    // Hide
    btnEdit.style.opacity = '0';

    setTimeout(() => {
        btnEdit.classList.add('hide');
    }, 200)

    // Show
    taskAdd.classList.add('show');
    showBtns(btnsEdit);
    showBtns(btnsDelete);

    itemPoints.forEach(itemPoint => {
        const pointLabel = itemPoint.querySelector('.task__item-label');
        const pointInput = itemPoint.querySelector('.task__item-input');
        const pointLabelValue = pointLabel.textContent;
        const taskTitleValue = taskTitle.textContent;

        pointLabel.classList.add('hide');
        taskTitle.classList.add('hide');

        pointInput.value = pointLabelValue
        taskTitleInput.value = taskTitleValue;

        pointInput.classList.remove('hide');
        taskTitleInput.classList.remove('hide');
    })
}

const defaultMode = (btnReady) => {
    const btns = btnReady.parentElement;
    const btnEdit = btns.querySelector('.task__item--edit');
    const btnsEdit = btns.querySelectorAll('.task__item-btn--edit');

    const taskItem = btns.parentElement.parentElement;
    const btnsDelete = taskItem.querySelectorAll('.task__item-right');

    const taskAdd = btns.parentElement.nextElementSibling.querySelector('.task__add');

    // Hide
    taskAdd.classList.remove('show');
    hideBtns(btnsEdit);
    hideBtns(btnsDelete);

    // Show
    btnEdit.classList.remove('hide');

    setTimeout(() => {
        btnEdit.style.opacity = '1';
    }, 200)
}

const pointDefault = (itemPoints, taskTitleInput, taskTitle) => {
    itemPoints.forEach(itemPoint => {
        const pointLabel = itemPoint.querySelector('.task__item-label');
        const pointInput = itemPoint.querySelector('.task__item-input');

        pointInput.classList.add('hide');
        taskTitleInput.classList.add('hide');

        pointLabel.classList.remove('hide');
        taskTitle.classList.remove('hide');
    })
}

const pointDefaultValue = (itemPoints, taskTitleInput, taskTitle) => {
    itemPoints.forEach(itemPoint => {
        const pointLabel = itemPoint.querySelector('.task__item-label');
        var pointInput = itemPoint.querySelector('.task__item-input');

        if (!pointInput) {
            pointInput = itemPoint.querySelector('.task__item-addinput');
            pointInput.classList.remove('task__item-addinput');
            pointInput.classList.add('task__item-input');
        }

        const pointInputValue = pointInput.value;
        const taskTitleInputValue = taskTitleInput.value;

        pointInput.classList.add('hide');
        taskTitleInput.classList.add('hide');

        pointLabel.innerHTML = pointInputValue
        taskTitle.innerHTML = taskTitleInputValue;

        pointLabel.classList.remove('hide');
        taskTitle.classList.remove('hide');
    })
}

const showBtns = (btns) => {
    btns.forEach(btn => {
        btn.classList.add('show');

        setTimeout(() => {
            btn.style.opacity = '1';
        }, 200)
    })
}

const hideBtns = (btns) => {
    btns.forEach(btn => {
        btn.style.opacity = '0';

        setTimeout(() => {
            btn.classList.remove('show');
        }, 200)
    })
}

const getForm = (form) => {
    const fd = new FormData(form);
    const data = {}
    fd.forEach((v, k) => data[k] = v);
    return JSON.stringify(data);
}

const notice = (notice) => {
    notice.classList.add('active');
}

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

// Forms

const registration = (regForm) => {
    const formInputs = regForm.querySelectorAll('.auth__input');
    const regSurname = document.querySelector('#regSurname');
    const regName = document.querySelector('#regName');
    const regEmail = document.querySelector('#regEmail');
    const regPass = document.querySelector('#regPass');
    const regPassValue = regPass.value;
    const regLogin = document.querySelector('#regLogin');
    const rEmail = /^[A-Z0-9._%+-]+@[A-Z0-9-]+.+.[A-Z]{2,4}$/i;

    formClear(formInputs);

    if (regSurname.value == '') {
        formError(regSurname);
        formBlur(regSurname);
        return;
    }
    if (regName.value == '') {
        formError(regName);
        formBlur(regName);
        return;
    }
    if (!rEmail.test(regEmail.value)) {
        formError(regEmail);
        formEmailBlur(regEmail);
        return;
    }
    if (regLogin.value == '') {
        formError(regLogin);
        formBlur(regLogin);
        return;
    }
    if (regPassValue.length < 8) {
        formError(regPass);
        formBlur(regPass);
        return;
    }
    
    const formData = getForm(regForm);
    submitFormReg(formData);
}

const login = (logForm) => {
    const formInputs = logForm.querySelectorAll('.auth__input');
    const logLogin = document.querySelector('#loginLogin');
    const logPass = document.querySelector('#loginPass');
    const logPassValue = logPass.value;

    formClear(formInputs);

    if (logLogin.value == "") {
        formError(logLogin);
        formBlur(logLogin);
        return;
    }
    if (logPassValue.length < 8) {
        formError(logPass);
        formBlur(logPass);
        return;
    }

    const formData = getForm(logForm);
    submitFormLog(formData);
}

const createTask = (form) => {
    const formInputs = form.querySelectorAll('.auth__input');
    const taskName = document.querySelector('#taskName');
    const modal = form.parentElement.parentElement;
    formClear(formInputs);

    if (taskName.value == '') {
        formError(taskName);
        formBlur(taskName);
        return;
    }
    modalClose(modal);
    console.log('ОТПРАВКА ФОРМЫ');
}

const createPoint = (form) => {
    const formInputs = form.querySelectorAll('.auth__input');
    const pointName = document.querySelector('#taskPointName');
    const modal = form.parentElement.parentElement;
    formClear(formInputs);

    if (pointName.value == '') {
        formError(pointName);
        formBlur(pointName);
        return;
    }
    modalClose(modal);
    console.log('ОТПРАВКА ФОРМЫ');
}

const edit = (form) => {
    const formInputs = form.querySelectorAll('.auth__input');
    const taskEdit = document.querySelector('#taskEdit');
    const modal = form.parentElement.parentElement;
    formClear(formInputs);

    if (taskEdit.value == '') {
        formError(taskEdit);
        formBlur(taskEdit);
        return;
    }
    modalClose(modal);
    console.log('ОТПРАВКА ФОРМЫ');
}

const account = (form) => {
    const formInputs = form.querySelectorAll('.auth__input');
    const accountSurname = document.querySelector('#accountSurname');
    const accountName = document.querySelector('#accountName');
    const accountLogin = document.querySelector('#accountLogin');
    const accountEmail = document.querySelector('#accountEmail');
    const rEmail = /^[A-Z0-9._%+-]+@[A-Z0-9-]+.+.[A-Z]{2,4}$/i;

    formAccountClear(formInputs);

    if (accountSurname.value == '') {
        formAccountError(accountSurname);
        formAccountBlur(accountSurname);
        return;
    }
    if (accountName.value == '') {
        formAccountError(accountName);
        formAccountBlur(accountName);
        return;
    }
    if (accountLogin.value == '') {
        formAccountError(accountLogin);
        formAccountBlur(accountLogin);
        return;
    }
    if (!rEmail.test(accountEmail.value)) {
        formAccountError(accountEmail);
        formAccountEmailBlur(accountEmail);
        return;
    }
  
    const formData = getForm(form);
    submitFormAcc(formData);
}

const formClear = (formInputs) => {
    formInputs.forEach(input => {
        let inputError = input.nextElementSibling;

        input.classList.remove('err');
        inputError.classList.remove('active');
    })
}

const formAccountClear = (formInputs) => {
    formInputs.forEach(input => {
        input.classList.remove('err');
    })
}

const formError = (element) => {
    element.classList.add('err');
    element.nextElementSibling.nextElementSibling.classList.add('active');
}

const formAccountError = (element) => {
    element.classList.add('err');
}

const formBlur = (input) => {
    input.addEventListener('blur', () => {
        if (input.value != '') {
            const inputError = input.nextElementSibling.nextElementSibling;

            input.classList.remove('err');
            inputError.classList.remove('active');
        }     
    })
}

const formAccountBlur = (input) => {
    input.addEventListener('blur', () =>  {
        if (input.value != '') {
            input.classList.remove('err');
        }     
    })
}

const formEmailBlur = (input) => {
    const rEmail = /^[A-Z0-9._%+-]+@[A-Z0-9-]+.+.[A-Z]{2,4}$/i;
    input.addEventListener('blur', () =>  {
        if (rEmail.test(input.value)) {
            const inputError = input.nextElementSibling.nextElementSibling;

            input.classList.remove('err');
            inputError.classList.remove('active');
        }     
    })
}

const formAccountEmailBlur = (input) => {
    const rEmail = /^[A-Z0-9._%+-]+@[A-Z0-9-]+.+.[A-Z]{2,4}$/i;
    input.addEventListener('blur', () =>  {
        if (rEmail.test(input.value)) {
            const inputError = input.nextElementSibling.nextElementSibling;

            input.classList.remove('err');
            inputError.classList.remove('active');
        }     
    })
}

