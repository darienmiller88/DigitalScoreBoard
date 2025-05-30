import { createRouter, createWebHistory } from 'vue-router'
import Home from "../pages/Home_Page/Home_Page.vue"
import Register from '../pages/Register/Register.vue'
import GenerateJeopardyQuestions from '../pages/GenerateJeopardyQuestions/GenerateJeopardyQuestions.vue'
import Games from '../pages/Games_Page/Games_Page.vue'
import TournamentMode from '../pages/TeamMode_Page/TeamMode_Page.vue'
import AddNewUser from '../pages/AddNewUser_Page/AddNewUser_Page.vue'

// Create the router instance
const router = createRouter({
    history: createWebHistory(), 
    routes: [
        {
            path: "/",
            component: Home
        },
        {
            path: "/add-new-user",
            component: AddNewUser
        },
        {
            path: "/view-games",
            component: Games
        },
        {
            path: "/register",
            component: Register
        },
        {
            path: "/generate-questions",
            component: GenerateJeopardyQuestions
        },
        {
            path: "/tournament-mode",
            component: TournamentMode
        }
    ]
})
  
export default router