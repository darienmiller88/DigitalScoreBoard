<script setup lang="ts">
    import { Icon } from "@iconify/vue"
    import { useWindowSize } from "@vueuse/core"

    const { width } = useWindowSize();

    defineProps<{
        isPhoneMenuActive: boolean
        menuClick: () => void
    }>()

</script>

<template>
    <div :class="`menu ${isPhoneMenuActive ? 'slide-in': ''}`">
        <div class="icon-wrapper">
            <Icon icon="tabler:circle-letter-x" :width="width < 820 ? 50 : 75" @click="menuClick"/>
        </div>
        <div class="links" @click="menuClick">
            <RouterLink to="/" class="link-item">Create New Game</RouterLink>
            <div class="divider"></div>
            <RouterLink to="/add-new-players" class="link-item">Add New Players</RouterLink> 
            <div class="divider"></div>    
            <RouterLink to="/team-mode" class="link-item">Team Mode</RouterLink>
            <div class="divider"></div>
            <RouterLink to="/view-games" class="link-item">View Games</RouterLink>
        </div>
    </div>
</template>

<style scoped lang="scss">
    @import url('https://fonts.googleapis.com/css2?family=Montserrat:ital,wght@0,100..900;1,100..900&family=Share+Tech&display=swap');
    
    .menu{
        display: grid;
        grid-template-rows: 10% auto;

        background-color: var(--bg-color);
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
            color: var(--primary-color);
        }

        .links{
            display: flex;
            flex-direction: column;
            align-items: center;

            padding: 60px;

            .divider{
                width: 40vw;
                border: 1px solid rgb(255, 215, 0);
                text-align: center;
            }
            
            .link-item{
                font-family: "Share Tech", sans-serif;
                font-weight: 600;
                font-style: normal;

                text-align: center;
                color: var(--primary-color);
                font-size: 25px;
                padding: 25px;
                transition: 0.2s;
                text-decoration: none;

                @media (min-width: 768px) {
                    font-size: 45px;
                }
            }
        }
    }

    .slide-in{
        left: 0;
    }
</style>