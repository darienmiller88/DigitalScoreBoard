<script setup lang="ts">
    import { useWindowSize } from "@vueuse/core"
    import { darkModeStore } from "../../stores/darkModeStore";
    import DarkModeToggle from "../DarkModeToggle/DarkModeToggle.vue";
    import { storeToRefs } from "pinia";

    const { width } = useWindowSize();
    const { isDarkMode } = storeToRefs(darkModeStore())
</script>

<template>
    <nav :class="`${isDarkMode ? 'dark-mode' : 'light-mode'}`">
        <RouterLink to="/" class="logo">
            <img src="../../assets/sb.png" alt="logo">
            <div class="logo-item">
                Scoreboard
            </div>
        </RouterLink>
        <div class="links">
            <DarkModeToggle />
            <RouterLink to="/" class="link-item">Home</RouterLink>
            <RouterLink to="/register" class="link-item">Register</RouterLink>
            <RouterLink to="/generate-questions" class="link-item">Generate Questions</RouterLink>
        </div>
    </nav>
</template>

<style scoped lang="scss">
    .dark-mode{
        border: 3px solid var(--primary-color);
    }

    .light-mode{
        border: 3px solid var(--main-bg-color);
    }

    nav{
        display: flex;
        justify-content: space-between;
        background-color: var(--main-text-color-dark-transparent);

        position: sticky;
        top: 0;
        z-index: 1000; 
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
            color: var(--main-bg-color);
            transition: 0.5s;

            &:hover{
                color: var(--primary-color);
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

            color: var(--primary-color);

            transition: 0.5s;
            height: 100%;

            &:hover{
                background-color: blue;
            }
        }
    }
</style>