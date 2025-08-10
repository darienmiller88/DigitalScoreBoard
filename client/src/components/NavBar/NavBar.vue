<script setup lang="ts">
    import { useWindowSize } from "@vueuse/core"
    import { Icon } from "@iconify/vue"
    import { ref } from "vue";
    import { storeToRefs } from "pinia";
    import { NavBarStore, NavBarItem } from "../../stores/NavBarStore";

    import PhoneNavMenu from "../PhoneNavMenu/PhoneNavMenu.vue";
    import DarkModeToggle from "../DarkModeToggle/DarkModeToggle.vue";

    const { width } = useWindowSize();
    const { activeNavBarItem } = storeToRefs(NavBarStore())
    const { setActiveNavBarItem } = NavBarStore()
    const isPhoneMenuActive = ref<boolean>(false)
</script>

<template>
    <nav>
        <RouterLink to="/" class="logo">
            <img src="../../assets/sb.png" alt="logo">
            <div class="logo-item"> Scoreboard </div>
        </RouterLink>
        
        <div v-if="width <= 1024" class="icon-wrapper">
            <DarkModeToggle />
            <Icon icon="mdi-light:menu" color="#f8f9fa" width="50" @click="() => isPhoneMenuActive = !isPhoneMenuActive"/>
        </div>
        <div class="links" v-else>
            <DarkModeToggle />

            <!-- Route to "/" -->
            <RouterLink 
                to="/" 
                :class="activeNavBarItem === NavBarItem.CREATE_NEW_GAME ? 'active link-item' : 'link-item'"
                @click="() => setActiveNavBarItem(NavBarItem.CREATE_NEW_GAME)"
            >Create Game</RouterLink>

            <!-- Route to "/add-new-players" -->
            <RouterLink 
                to="/add-new-players" 
                :class="activeNavBarItem === NavBarItem.ADD_PLAYER ? 'active link-item' : 'link-item'"
                @click="() => setActiveNavBarItem(NavBarItem.ADD_PLAYER)"    
            >Add Player</RouterLink>  
            
            <!-- Route to "/team-mode" -->
            <RouterLink 
                to="/team-mode" 
                :class="activeNavBarItem === NavBarItem.TEAM_MODE ? 'active link-item' : 'link-item'"
                @click="() => setActiveNavBarItem(NavBarItem.TEAM_MODE)"
            >Team Mode</RouterLink>

            <!-- Route to /view games -->
            <RouterLink 
                to="/view-games" 
                :class="activeNavBarItem === NavBarItem.VIEW_GAMES ? 'active link-item' : 'link-item'"
                @click="() => setActiveNavBarItem(NavBarItem.VIEW_GAMES)"
            >View Games</RouterLink>
            <!-- <RouterLink class="link-item">Sign out</RouterLink> -->
        </div>
    </nav>
    <PhoneNavMenu :isPhoneMenuActive="isPhoneMenuActive" :menuClick="() => isPhoneMenuActive = !isPhoneMenuActive"/>
</template>

<style scoped lang="scss">
    nav{
        display: flex;
        justify-content: space-between;
        background-color: var(--primary-color-transparent);
        border: 3px solid var(--toggle-background);

        position: sticky;
        top: 0;
        z-index: 1000; 
    }

    .icon-wrapper{
        display: flex;
        align-items: center;
        justify-content: center;

        padding: 10px;
    }

    .logo{
        display: flex;
        justify-content: center;
        align-items: center;
        width: fit-content;

        margin: 15px;
        font-size: 25px;
        text-decoration: none;

        img{
            height: 30px;
            width: auto;
            margin-right: 8px;
        }

        .logo-item{
            color: aliceblue;
            transition: 0.5s;

            &:hover{
                color: black;
            }
        }

        &:hover{
            cursor: pointer;
        }
    }

    .links{
        display: flex;
        justify-content: center;
        align-items: center;
        width: fit-content;

        .link-item{
            display: flex;
            align-items: center;

            font-size: 18px;
            padding: 0px 20px;

            color: aliceblue;

            transition: 0.5s;
            height: 100%;

            &:hover{
                background-color: rgb(50, 50, 219);
            }
        }

        .active{
            background-color: blue;
        }
    }
</style>