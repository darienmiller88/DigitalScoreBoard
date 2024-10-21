<script setup lang="ts">
    import { useWindowSize } from "@vueuse/core"
    import { darkModeStore } from "../../stores/darkModeStore";
    import { storeToRefs } from "pinia";
    import { Icon } from "@iconify/vue"
    import { ref } from "vue";
    import PhoneNavMenu from "../PhoneNavMenu/PhoneNavMenu.vue";
    import DarkModeToggle from "../DarkModeToggle/DarkModeToggle.vue";

    const { width } = useWindowSize();
    const { isDarkMode } = storeToRefs(darkModeStore())
    const isPhoneMenuActive = ref<boolean>(false)

    const menuClick = () => {
        isPhoneMenuActive.value = !isPhoneMenuActive.value
    }
</script>

<template>
    <nav :class="`${isDarkMode ? 'dark-mode' : 'light-mode'}`">
        <RouterLink to="/" class="logo">
            <img src="../../assets/sb.png" alt="logo">
            <div class="logo-item">
                Scoreboard
            </div>
        </RouterLink>
        
        <div v-if="width < 768" class="icon-wrapper">
            <DarkModeToggle />
            <Icon icon="mdi-light:menu" color="#f8f9fa" width="50" @click="menuClick"/>
        </div>
        <div class="links" v-else>
            <DarkModeToggle />
            <RouterLink to="/" class="link-item">Home</RouterLink>
            <RouterLink to="/register" class="link-item">Register</RouterLink>
            <RouterLink to="/generate-questions" class="link-item">Generate Questions</RouterLink>
        </div>
    </nav>
    <PhoneNavMenu :isPhoneMenuActive="isPhoneMenuActive" :menuClick="menuClick"/>
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
        // z-index: -1000; 
    }

    .icon-wrapper{
        display: flex;
        align-items: center;
        justify-content: center;

        padding: 10px;
        // border: 2px solid red;

        &:active{
            background-color: black;
        }
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