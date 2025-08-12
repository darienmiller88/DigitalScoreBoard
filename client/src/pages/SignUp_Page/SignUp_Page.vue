<script setup lang="ts">
    import { ref } from "vue";
    import adaptLogo from "../../../public/adapt_logo-nobg.png"

    const isSignUpForm = ref<boolean>(true)

    const handleSubmit = () => {

    }
</script>

<template>
    <div class="signup">
        <form v-if="isSignUpForm" class="signup-form" @submit.prevent="handleSubmit">
            <img :src="adaptLogo" alt="">
            <div>Sign up form</div>

            <button @click="() => isSignUpForm = false">log in</button>
            <button @click="() => $router.push('/')"> continue to home</button>
        </form>
        <form v-else class="login-form" @submit.prevent="handleSubmit">
            <img :src="adaptLogo" alt="">
            <div>Login form</div>

            <button @click="() => isSignUpForm = true">log in</button>
            <button @click="() => $router.push('/')">continue to home</button>
        </form>
    </div>
    
</template>

<style scoped lang="scss">
    .signup{
        display: flex;
        justify-content: center;
        align-items: center;

        font-size: 40px;
        width: 100vw;
        height: 100vh;
        background-color: #003366;
    }

    form{
        background-color: #344966;
        padding-left: 30px;
        padding-right: 100px;

        img{
            width: 200px;
            height: auto;
        }
    }
</style>