import { defineStore } from 'pinia';
import { ref } from 'vue';

//The navbar currently has four items: "Create game", "Add Player", "Team Mode", "View Games"
export enum NavBarItem {
    CREATE_NEW_GAME,
    ADD_PLAYER,
    TEAM_MODE,
    VIEW_GAMES
}

export const NavBarStore = defineStore("NavBarStore", () => {
    
    //Initialize the first navbar item to be highlighted.
    const activeNavBarItem = ref<NavBarItem>(NavBarItem.CREATE_NEW_GAME)

    const setActiveNavBarItem = (navbarItem: NavBarItem) => {
        activeNavBarItem.value = navbarItem
    }

    return {
        activeNavBarItem,
        setActiveNavBarItem
    }
}, {
    persist: true
})