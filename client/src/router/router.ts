import { createRouter, createWebHistory } from 'vue-router'
import Home from "../pages/Home/Home.vue"
import Register from '../pages/Register/Register.vue'
import GenerateJeopardyQuestions from '../pages/GenerateJeopardyQuestions/GenerateJeopardyQuestions.vue'

// Create the router instance
const router = createRouter({
    history: createWebHistory(), // Uses HTML5 history mode
    routes: [
        {
            path: "/",
            component: Home
        },
        {
            path: "/register",
            component: Register
        },
        {
            path: "/generate-questions",
            component: GenerateJeopardyQuestions
        }
    ]
})
  

export default router