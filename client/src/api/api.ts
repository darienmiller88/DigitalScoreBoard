import axios from "axios"

const userApiURL =   window.location.hostname === "localhost" || window.location.hostname === "127.0.0.1"
? 
"http://localhost:8000/api/v1/users" 
:
"https://digitalscoreboard.up.railway.app/api/v1/users"

const scoreBoardApiUrl =  window.location.hostname === "localhost" || window.location.hostname === "127.0.0.1"
?
"http://localhost:8000/api/v1/scoreboard"
:
"https://digitalscoreboard.up.railway.app/api/v1/scoreboard"

export const userApi = axios.create({
    baseURL: userApiURL,
    withCredentials: true,
})

export const scoreBoardApi = axios.create({
    baseURL: scoreBoardApiUrl,
    withCredentials: true
})