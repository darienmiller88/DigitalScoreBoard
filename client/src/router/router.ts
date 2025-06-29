import { createRouter, createWebHistory } from 'vue-router'
import HomePage from "../pages/Home_Page/Home_Page.vue"
import GamesPage from '../pages/Games_Page/Games_Page.vue'
import TeamModePage from '../pages/TeamMode_Page/TeamMode_Page.vue'
import AddNewUserPage from '../pages/AddNewUser_Page/AddNewUser_Page.vue'

// Create the router instance
const router = createRouter({
    history: createWebHistory(), 
    routes: [
        {
            path: "/",
            component: HomePage
        },
        {
            path: "/add-new-players",
            component: AddNewUserPage
        },
        {
            path: "/view-games",
            component: GamesPage
        },
        {
            path: "/team-mode",
            component: TeamModePage
        } 
    ]
})
  
export default router