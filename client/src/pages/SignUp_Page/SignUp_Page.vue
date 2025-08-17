<script setup lang="ts">
    import { ref } from "vue";
    import adaptLogo from "../../assets/adapt_logo-nobg.png"

    const isSignUpForm = ref<boolean>(true)
    const emailAddress = ref<string>("")
    const password = ref<string>("")
    const confirmPassword = ref<string>("")
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

            <div class="button-wrapper form-item">
                <div :class="isSignUpForm ? 'highlight move-left' : 'highlight move-right'"></div>
                <button :class="isSignUpForm ? 'active' : ''" @click="() => isSignUpForm = true">Sign up</button>
                <button :class="!isSignUpForm ? 'active' : ''" @click="() => isSignUpForm = false">Log in</button>
            </div>

            <div class="input-wrapper">
                <input class="form-item" type="text" placeholder="Email Address" v-model="emailAddress">
                <div class="password-wrapper">
                    <input class="form-item" type="password" placeholder="Password" v-model="password">
                    <div v-if="!isSignUpForm" class="forgot-password">Forgot password?</div>
                </div>
                <input v-if="isSignUpForm" class="form-item" type="password" placeholder="Confirm Password" v-model="confirmPassword">
            </div>


            <!-- <button @click="() => $router.push('/')"> continue to home</button> -->
            <button class="submit form-item" type="submit">
                <span v-if="isSignUpForm">Sign Up</span>
                <span v-else>Log in</span>
            </button>

            <div v-if="!isSignUpForm" class="not-a-member">Not a member? <span @click="() => isSignUpForm = true">Sign up now</span></div>
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
            display: flex;
            align-items: center;
            justify-content: center;
            flex-direction: column;
            background-color: #344966;
            padding: 20px 0px;
            

            @media (min-width: 768px) {
                padding: 40px;
            }

            @media (min-width: 1025px) {
                height: 60vh;
            }

            .form-item{
                width: 80vw;

                @media (min-width: 768px) {
                    width: 60vw;
                }

                @media (min-width: 1025px) {
                    width: 20vw;
                }
            }
    
            img{
                width: 200px;
                height: auto;

                @media (min-width: 768px) {
                    width: 400px;
                }

                @media (min-width: 1025px) {
                    width: 300px;
                }
            }

            .header{
                font-size: 25px;
                margin: 20px;
            }

            .input-wrapper{
                display: flex;
                flex-direction: column;
                margin: 10px;

                input{
                    margin: 8px;
                    padding: 10px;
                    font-size: 18px;
                    border-radius: 10px;
                    border: 1px solid rgb(200, 200, 200);
                }

                .forgot-password{
                    font-size: 17px;
                    margin: 5px 8px;
                    color: aliceblue;
                    transition: 0.3s;

                    &:hover{
                        color: rgb(140, 140, 140);
                        cursor: pointer;
                    }
                }
            }


            .button-wrapper{
                display: flex;
                border: 1px solid rgb(200, 200, 200);
                border-radius: 10px;
                width: fit-content;
                position: relative;
                z-index: 0;

                .highlight{
                    position: absolute;
                    top: 0;
                    left: 0;
                    width: 50%;
                    height: 100%;
                    z-index: 1;
                    transition: transform 0.3s ease; 
                    border-radius: 10px;
                     
                    background: linear-gradient(to right,#003366, #0073e6);
                }

                .move-left{
                    transform: translateX(0);
                }

                .move-right{
                    transform: translateX(100%);
                }

                .active{
                    transition: 0.3s;
                    color: white;
                }

                button{
                    border: none;
                    background-color: transparent;
                    padding: 10px 40px;
                    z-index: 2;
                    font-size: 18px;
                    font-weight: 600;
                    transition: 0.3s;

                    &:hover{
                        cursor: pointer;
                    }
                }
            }

            .submit{
                border: none;
                border-radius: 10px;
                background: linear-gradient(to right,#003366, #0073e6);
                color: white;
                font-weight: 600;
                padding: 10px;
                font-size: 18px;
                transition: 0.5s;

                &:hover{
                    cursor: pointer;
                    background: linear-gradient(to right,#0073e6, #003366);
                }
            }

            .not-a-member{
                font-size: 16px;
                margin: 12px;
                transition: 0.3s;

                span{
                    color: #0073e6;
                }

                &:hover{
                    cursor: pointer;
                    color: aliceblue;
                }
            }
            
        }
    }
</style>