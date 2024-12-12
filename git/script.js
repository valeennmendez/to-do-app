console.log("Conectado...")


const url = "https://api-to-do-mmut.onrender.com/"

//let autenticado = false;

function addTask(){


    const list = document.querySelector(".list ul")
    const btn = document.getElementById("button")
    const input = document.getElementById("inputTask")


    btn.addEventListener("click", function(e){
        const value = input.value;
        const body ={
            task: value
        }
        console.log(value)

        fetch(`${url}create-task`,{
            method: "POST",
            headers:{
                "Content-Type": "application/json"
            },
            body: JSON.stringify(body),
            credentials: "include"
        })
        .then(response =>{
            if(!response.ok){
                console.error(response)
            }
            return response.json()
        })
        .then(data =>{
            if(data.message === "task created"){
                const li = document.createElement("li")
                const taskID = data.task
                li.innerHTML = `${value} <i data-id="${taskID}" class="fa-solid fa-x deletebtn"></i>`            
                list.appendChild(li)
                const deleteBtn = li.querySelector(".deletebtn")
                deleteBtn.addEventListener("click", function(){
                    li.remove()
                })
                input.value = ""
                
            }
        })
    })


/*     btn.addEventListener("click", function(e){
        e.preventDefault()
        const inputValue = input.value
        if(inputValue != ""){
            const li = document.createElement("li")
            li.innerHTML = `${inputValue} <i class="fa-solid fa-x deletebtn"></i>`            
            list.appendChild(li)
            const deleteBtn = li.querySelector(".deletebtn")
            deleteBtn.addEventListener("click", function(){
                li.remove()
            })
            input.value = ""
        }
    })
 */
    input.addEventListener("keydown", function(e){
        if(e.keyCode === 13) {
            e.preventDefault()
            btn.click()
        }
    })
}


function Task(){
    fetch(url,{
        method: "GET",
        credentials: "include"
    })
    .then(response =>{
        if(!response.ok){
            console.error(response)
        }
        return response.json()
    })
    .then(data =>{
        data.forEach(task => {
            const list = document.querySelector(".list ul")

            const li = document.createElement("li")
            li.innerHTML = `${task.task} <i data-id="${task.ID}" class="fa-solid fa-x deletebtn"></i>`
            list.appendChild(li)
            const deleteBtn = li.querySelector(".deletebtn")
            deleteBtn.addEventListener("click", function(){
                li.remove()
            })
        });
    })
    .catch(error => console.error(error))
}

function getTask(){


        console.log("haciendo fetch...")
        fetch(`${url}task`,{
            method: "GET",
            credentials: "include"
        })
        .then(response => response.json())
        .then(data =>{
            const tasks = data.tasks; 
        
        // Itera sobre el array 'tasks'
            tasks.forEach(task => {
                console.log(task.task); // Muestra la tarea en la consola

                const list = document.querySelector(".list ul")

                const li = document.createElement("li")
                li.innerHTML = `${task.task} <i data-id="${task.ID}" class="fa-solid fa-x deletebtn"></i>`
                list.appendChild(li)
                const deleteBtn = li.querySelector(".deletebtn")
                deleteBtn.addEventListener("click", function(){
                    li.remove()
                })

            });
        })
        .catch(error => console.error('Error:', error));

}

function DeleteTask(){
    const lista = document.querySelector(".list")

    lista.addEventListener("click", function(e){
        const deleteButton = e.target.closest(".deletebtn")

        if(deleteButton){
            const id = deleteButton.getAttribute("data-id")
            
            fetch(`${url}delete-task/${id}`,{
                method: "DELETE",
                credentials: "include"
            })
            .then(response =>{
                if(!response.ok){
                    console.error(response)
                }
                return response.json()
            })
            .then(data =>{
                if(data.message ==="task deleted successfully"){
                    console.log("Task deleted")
                }
            })
        }

    })

}

function ActiveForms(){
    const registerBtn = document.getElementById("registerbtn")
    const loginBtn = document.getElementById("loginbtn");
    const formRegister = document.getElementById("formRegister")
    const formLogin = document.getElementById("formLogin")
     
    if(registerBtn){
        registerBtn.addEventListener("click", function(){
            formRegister.classList.add("active")
            formLogin.classList.remove("active")
        })
    }


    if(loginBtn){
        loginBtn.addEventListener("click", function(){
            formLogin.classList.add("active")
            formRegister.classList.remove("active")
        })
    }

}


function RegisterUser(){

    const registerForm = document.getElementById("registerFrom")


    registerForm.addEventListener("submit", function(){

        const username = document.getElementById("username-input").value
        const password = document.getElementById("password-input").value

        const data = {
            username: username,
            password: password,
        }

        fetch(`${url}register`,{
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data),
            credentials: "include"
        })
        .then(response =>{
            if(!response.ok){
                console.error(response)
            }
            return response.json()
        })
        .then(data =>{
            const error = document.getElementById("errorRegister")

                if(data.message === "user created successfully"){
                    alert("User created successfully")
                }


/*             if(data.error === "user exist"){
                error.innerHTML = "User exist"
                return
            } */

        })
        .catch(error => console.error(error))
    })

}


function Login(){
    
    const loginForm = document.getElementById("loginForm")
    const formLogin = document.getElementById("formLogin")

    loginForm.addEventListener("submit", function(e){
        e.preventDefault();
        const username = document.getElementById("username").value
        const password = document.getElementById("password").value


        const data = {
            username: username,
            password: password
        }

        fetch(`${url}login`,{
            method: "POST",
            headers:{
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data),
            credentials: "include"
        })
        .then(response =>{
            if(!response.ok){
                console.error(response)
            }
            return response.json()
        })
        .then(data =>{
            const error = document.getElementById("errorLogin")

            if(data.error === "invalid user or password"){
                error.innerHTML = "Invalid user or password"
                return
            }

            if(data.message ==="login successfull"){
                e.preventDefault()
                autenticado = true;
                console.log(autenticado)
                formLogin.classList.remove("active")
                getTask();
            }
        })
        .catch(error => console.error(error))

    })
}




document.addEventListener("DOMContentLoaded", function(){
    addTask()

    DeleteTask()

    RegisterUser()

    Login()

    ActiveForms()

    //getTask()
})
