<script setup lang="ts">
    import { Icon } from "@iconify/vue"

    defineProps<{
        isPhoneMenuActive: boolean
        menuClick: () => void
    }>()
</script>

<template>
    <div :class="`menu ${isPhoneMenuActive ? 'slide-in': ''}`">
        <div class="icon-wrapper">
            <Icon icon="tabler:circle-letter-x" width="50" @click="menuClick"/>
        </div>
        <div class="links">
            <RouterLink to="/" class="link-item">Home</RouterLink>
            <RouterLink to="/register" class="link-item">Register</RouterLink>
            <RouterLink to="/generate-questions" class="link-item">Generate Questions</RouterLink>
        </div>
    </div>
</template>

<style scoped lang="scss">
    .menu{
        display: grid;
        grid-template-rows: 10% auto;

        background-color: rgba(128, 128, 128, 0.966);
        width: 100vw;
        height: 100vh;
        z-index: 1000;

        transition: 0.5s;
        position: fixed;
        top: 0;
        left: 100vw;

        .icon-wrapper{
            display: flex;
            justify-content: end;
            padding: 20px;
            // border: 2px solid white;
        }

        .links{
            display: flex;
            flex-direction: column;
            // border: 2px solid yellow;
           

            .link-item{
                // border: 2px solid red;
                text-align: center;
                color: var(--primary-color);
                font-size: 35px;
                padding: 30px;
                transition: 0.2s;

                &:active{

                }
            }
        }
    }

    .slide-in{
        left: 0;
    }
</style>