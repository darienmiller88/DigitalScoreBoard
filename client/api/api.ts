import axios from "axios"

const userApiURL = window.location.hostname === "localhost" 
? 
"http://localhost:8080/api/v1/users" 
:
"https://messengerv2.fly.dev/api/v1/users"

export const userApi = axios.create({
    baseURL: userApiURL,
    withCredentials: true,
})