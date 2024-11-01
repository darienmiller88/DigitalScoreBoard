import axios from "axios"

const userApiURL = window.location.hostname === "localhost" 
? 
"http://localhost:8080/api/v1/users" 
:
"https://messengerv2.fly.dev/api/v1/users"

const scoreBoardApiUrl = window.location.hostname === "localhost" 
?
"http://localhost:9080/api/v1/scoreboard"
:
"https://adaptscoreboardapi.fly.dev/api/v1/scoreboard"

export const userApi = axios.create({
    baseURL: userApiURL,
    withCredentials: true,
})

const scoreBoardApi = axios.create({
    baseURL: scoreBoardApiUrl,
    withCredentials: true
})