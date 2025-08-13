<script setup lang="ts">
    import { ref } from "vue";
    import adaptLogo from "../../assets/adapt_logo-nobg.png"

    const isSignUpForm = ref<boolean>(true)
    // const isActive = ref<boolean>(true)

    const handleSubmit = () => {

    }
</script>

<template>
    <div class="signup">
        <form @submit.prevent="handleSubmit">
            <img :src="adaptLogo" alt="">
            <div class="header">
                <span v-if="isSignUpForm">Sign up Form</span>
                <span v-else>Login Form</span>
            </div>

            <div class="button-wrapper">
                <div :class="isSignUpForm ? 'highlight move-left' : 'highlight move-right'"></div>
                <button @click="() => isSignUpForm = true">Sign up</button>
                <button @click="() => isSignUpForm = false">Log in</button>
            </div>

            <button @click="() => $router.push('/')"> continue to home</button>
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
    
        form{
            background-color: #344966;
            padding-left: 30px;
            padding-right: 100px;
    
            img{
                width: 200px;
                height: auto;
            }

            .header{
                font-size: 25px;
            }

            .button-wrapper{
                display: flex;
                border: 2px solid black;
                width: fit-content;
                position: relative;
                z-index: 0;

                .highlight{
                    position: absolute;
                    top: 0;
                    left: 0;
                    background-color: blue;
                    width: 50%;
                    height: 100%;
                    z-index: 1;
                    transition: transform 0.3s ease; // optional smooth slide
                }

                .move-left{
                    transform: translateX(0);
                }

                .move-right{
                    transform: translateX(100%);
                }

                button{
                    border: none;
                    background-color: transparent;
                    padding: 10px 20px;
                    z-index: 2;

                    &:hover{
                        cursor: pointer;
                    }
                }
            }
        }
    }
</style>