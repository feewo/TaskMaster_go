// Edit
const taskItems = document.querySelectorAll('.task__item');

taskItems.forEach(taskItem => {
    taskItem.addEventListener('click', (event) => {
        const clickEdit = event.target.closest('.task__item--edit');
        const clickDeletePoint = event.target.closest('.task__item-right');
        const clickDelete = event.target.closest('.task__item-btn--delete');
        const clickReady = event.target.closest('.task__item-btn--ready');

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
            defaultMode(clickReady);
        }
    })
})

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
    const btnAdd = btns.querySelector('.task__item-plus');
    const btnsEdit = btns.querySelectorAll('.task__item-btn--edit');

    const taskItem = btns.parentElement.parentElement;
    const btnsDelete = taskItem.querySelectorAll('.task__item-right');
    const btnsChange = taskItem.querySelectorAll('.task__item-btnedit');

    // Hide
    btnEdit.style.opacity = '0';
    btnAdd.style.opacity = '0';

    setTimeout(() => {
        btnEdit.classList.add('hide');
        btnAdd.classList.add('hide');
    }, 200)

    // Show
    btnsEdit.forEach(btn => {
        btn.classList.add('show');

        setTimeout(() => {
            btn.style.opacity = '1';
        }, 200)
    })

    showBtns(btnsEdit);
    showBtns(btnsDelete);
    showBtns(btnsChange);
}


const defaultMode = (btnReady) => {
    const btns = btnReady.parentElement;
    const btnEdit = btns.querySelector('.task__item--edit');
    const btnAdd = btns.querySelector('.task__item-plus');
    const btnsEdit = btns.querySelectorAll('.task__item-btn--edit');

    const taskItem = btns.parentElement.parentElement;
    const btnsDelete = taskItem.querySelectorAll('.task__item-right');
    const btnsChange = taskItem.querySelectorAll('.task__item-btnedit');

    // Hide
    hideBtns(btnsEdit);
    hideBtns(btnsDelete);
    hideBtns(btnsChange);

    // Show
    btnAdd.classList.remove('hide');
    btnEdit.classList.remove('hide');

    setTimeout(() => {
        btnAdd.style.opacity = '1';
        btnEdit.style.opacity = '1';
    }, 200)
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

// Account
const changeBtns = document.querySelectorAll('.account__btn--change');

if (changeBtns) {
    changeBtns.forEach(changeBtn => {
        changeBtn.addEventListener ('click', (event) => {
            event.preventDefault();

            const accountInfo = changeBtn.parentElement;
            const accountForm = accountInfo.nextElementSibling;

            accountInfo.classList.add('hide');

            setTimeout(() => {
                accountInfo.style.display = "none";
            }, 200)

            setTimeout(() => {
                accountForm.style.display = "flex";
                setTimeout(() => {
                    accountForm.classList.add('show');
                }, 50)
            }, 200)
        })
    })
}

const accountReady = (readyBtn) => {
    const accountForm = readyBtn.parentElement;
    const accountInfo = accountForm.previousElementSibling;

    accountForm.classList.remove('show');
    setTimeout(() => {
        accountForm.style.display = "none";
    }, 200)

    setTimeout(() => {
        accountInfo.style.display = "flex";
        setTimeout(() => {
            accountInfo.classList.remove('hide');
        }, 50)
    }, 200)
}

// Modal
const btnsModal = document.querySelectorAll('[data-modal]');

const modalShow = (modal, modalContent) => {
    const body = document.querySelector('body');

    body.classList.add('no-scroll');
    modal.classList.add('active');

    setTimeout(() => {
        modalContent.style.opacity = '1';
        modalContent.style.transform = 'translateY(0)';
    }, 200)
}

const modalHide = (modal, modalContent) => {
    modal.addEventListener('click', () => {
        const body = document.querySelector('body');

        modalContent.style.opacity = '0';
        modalContent.style.transform = 'translateY(-250px)';

        setTimeout(() => {
            body.classList.remove('no-scroll');
            modal.classList.remove('active');
        }, 500)
    }, {"once": true})

    modalContent.addEventListener('click', (event) => {
        event.preventDefault();

        if (!event.target.closest('.modal__close')){
            event.stopPropagation()
        }
    })
}

const modalClose = (modal) => {
    const body = document.querySelector('body');
    const modalContent = modal.querySelector('.modal__content');

    modalContent.style.opacity = '0';
    modalContent.style.transform = 'translateY(-250px)';

    setTimeout(() => {
        body.classList.remove('no-scroll');
        modal.classList.remove('active');
    }, 500) 
}

btnsModal.forEach(btnModal => {
    btnModal.addEventListener('click', (event) => {
        event.preventDefault();

        const dataModal = btnModal.dataset.modal;
        const modal = document.querySelector(dataModal);
        const modalContent = modal.querySelector('.modal__content');

        // Show
        modalShow(modal, modalContent);

        // Hide
        modalHide(modal, modalContent);
    });
});

// Forms
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
            case "#accountForm":
                account(form, formBtn);
                break;
        }
    })
})

const registration = (regForm) => {
    const formInputs = regForm.querySelectorAll('.auth__input');
    const regSurname = document.querySelector('#regSurname');
    const regName = document.querySelector('#regName');
    const regPatronymic = document.querySelector('#regPatronymic');
    const regEmail = document.querySelector('#regEmail');
    const regPass = document.querySelector('#regPass');
    const regPassValue = regPass.value;
    const regConfirm = document.querySelector('#regConfirm');
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
    if (regPatronymic.value == '') {
        formError(regPatronymic);
        formBlur(regPatronymic);
        return;
    }
    if (!rEmail.test(regEmail.value)) {
        formError(regEmail);
        formEmailBlur(regEmail);
        return;
    }
    if (regPassValue.length < 8) {
        formError(regPass);
        formBlur(regPass);
        return;
    }
    if (regConfirm.value != regPassValue) {
        formError(regConfirm);
        formBlur(regConfirm);
        return;
    }
    console.log('ОТПРАВКА ФОРМЫ');
}

const login = (logForm) => {
    const formInputs = logForm.querySelectorAll('.auth__input');
    const logEmail = document.querySelector('#loginEmail');
    const logPass = document.querySelector('#loginPass');
    const logPassValue = logPass.value;
    const rEmail = /^[A-Z0-9._%+-]+@[A-Z0-9-]+.+.[A-Z]{2,4}$/i;

    formClear(formInputs);

    if (!rEmail.test(logEmail.value)) {
        formError(logEmail);
        formEmailBlur(logEmail);
        return;
    }
    if (logPassValue.length < 8) {
        formError(logPass);
        formBlur(logPass);
        return;
    }
    console.log('ОТПРАВКА ФОРМЫ');
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

const account = (form, formBtn) => {
    const formInputs = form.querySelectorAll('.auth__input');
    const accountSurname = document.querySelector('#accountSurname');
    const accountName = document.querySelector('#accountName');
    const accountPatronymic = document.querySelector('#accountPatronymic');
    const accountLogin = document.querySelector('#accountLogin');
    const accountPass = document.querySelector('#accountPassword');
    const accountPassValue = accountPass.value;
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
    if (accountPatronymic.value == '') {
        formAccountError(accountPatronymic);
        formAccountBlur(accountPatronymic);
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
    if (accountPassValue.length < 8) {
        formAccountError(accountPass);
        formAccountBlur(accountPass);
        return;
    }
    accountReady(formBtn);
    console.log('ОТПРАВКА ФОРМЫ');
}

const formClear = (formInputs) => {
    formInputs.forEach(input => {
        let inputError = input.nextElementSibling;

        input.classList.remove('error');
        inputError.classList.remove('active');
    })
}

const formAccountClear = (formInputs) => {
    formInputs.forEach(input => {
        input.classList.remove('error');
    })
}

const formError = (element) => {
    element.classList.add('error');
    element.nextElementSibling.classList.add('active');
}

const formAccountError = (element) => {
    element.classList.add('error');
}

const formBlur = (input) => {
    input.addEventListener('blur', function() {
        if (input.value != '') {
            const inputError = input.nextElementSibling;

            input.classList.remove('error');
            inputError.classList.remove('active');
        }     
    })
}

const formAccountBlur = (input) => {
    input.addEventListener('blur', function() {
        if (input.value != '') {
            input.classList.remove('error');
        }     
    })
}

const formEmailBlur = (input) => {
    const rEmail = /^[A-Z0-9._%+-]+@[A-Z0-9-]+.+.[A-Z]{2,4}$/i;
    input.addEventListener('blur', function() {
        if (rEmail.test(input.value)) {
            const inputError = input.nextElementSibling;

            input.classList.remove('error');
            inputError.classList.remove('active');
        }     
    })
}

const formAccountEmailBlur = (input) => {
    const rEmail = /^[A-Z0-9._%+-]+@[A-Z0-9-]+.+.[A-Z]{2,4}$/i;
    input.addEventListener('blur', function() {
        if (rEmail.test(input.value)) {
            input.classList.remove('error');
        }     
    })
}

