<script setup lang="ts">
    import { useWindowSize } from "@vueuse/core"
    import { Icon } from "@iconify/vue"
    import { ref } from "vue";

    import PhoneNavMenu from "../PhoneNavMenu/PhoneNavMenu.vue";
    import DarkModeToggle from "../DarkModeToggle/DarkModeToggle.vue";

    const { width } = useWindowSize();
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
            <RouterLink to="/" activeClass="active link-item" class="link-item">Create Game</RouterLink>

            <!-- Route to "/add-new-players" -->
            <RouterLink to="/add-new-players" activeClass="active link-item" class="link-item">Add Player</RouterLink>  
            
            <!-- Route to "/team-mode" -->
            <RouterLink to="/team-mode" activeClass="active link-item" class="link-item">Team Mode</RouterLink>

            <!-- Route to /view games -->
            <RouterLink to="/view-games" activeClass="active link-item" class="link-item">View Games</RouterLink>
            
            <RouterLink to="/signup" class="link-item">Sign out</RouterLink>
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

        @media (min-width: 3840px) {
            font-size: 75px;
        }

        img{
            height: 30px;
            width: auto;
            margin-right: 8px;

            @media (min-width: 3840px) {
                height: 70px;
                margin: 20px;
            }
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

            @media (min-width: 3840px) {
                font-size: 70px;
                padding: 40px 50px;
                height: unset;
            }

            &:hover{
                background-color: rgb(50, 50, 219);
            }
        }

        .active{
            background-color: blue;
        }
    }
</style>