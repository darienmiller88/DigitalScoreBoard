import { createRouter, createWebHistory } from 'vue-router'
import Home from "../pages/Home/Home.vue"
import Register from '../pages/Register/Register.vue'
import GenerateJeopardyQuestions from '../pages/GenerateJeopardyQuestions/GenerateJeopardyQuestions.vue'
import Games from '../pages/Games/Games.vue'
import TournamentMode from '../pages/TournamentMode/TournamentMode.vue'
// import AddNewUser from '../pages/AddNewUser/AddNewUser.vue'

// Create the router instance
const router = createRouter({
    history: createWebHistory(), 
    routes: [
        {
            path: "/",
            component: Home
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