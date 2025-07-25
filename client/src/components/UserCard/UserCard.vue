<script setup lang="ts">
    import { storeToRefs } from 'pinia';
    import { HomePageStore } from '../../stores/HomePageStore'
    import { useToast } from 'vue-toastification';
    import { ref } from 'vue';

    const { isGameCreated } = storeToRefs(HomePageStore())
    const toast = useToast()
    let isEditDisabled = ref<boolean>(false)

    const props = defineProps<{
        playerName: string
        playerIndex: number
        showModalAndSetCurrentPlayer: (username: string, playerIndex: number) => void
        removePlayer: (playerIndex: number) => void
    }>()

    // Prevents users from editing a players name when a game has already started.
    const preventEditWhenGameStarted = () => {
        if (isGameCreated.value) {
            isEditDisabled.value = true
            toast.error("Cannot edit name while game in session!", { timeout: 2000 })

            setTimeout(() => {
               isEditDisabled.value = false 
            }, 2000)
        } else {
            props.showModalAndSetCurrentPlayer(props.playerName, props.playerIndex)
        }
    }
</script>

<template>
    <div class="user-card">
        <div class="player">{{ playerName }}</div>
        <div class="divider"></div>
        <div class="remove-wrapper">
            <button @click="() => removePlayer(playerIndex)">Remove Player</button>
        </div>
        <div class="edit-wrapper">
            <button @click="preventEditWhenGameStarted" :disabled="isEditDisabled">Edit Name</button>
        </div>
    </div>
</template>

<style scoped lang="scss">
    .user-card{
        border: 2px solid var(--primary-color);
        border-radius: 10px;

        width: 85vw;
        box-shadow: 8px 8px 5px var(--primary-color-transparent);
        margin: 20px 0px; 
        
        transition: 0.3s;
        text-align: center;

        &:hover{
            transform: translateY(-5px);
            box-shadow: 10px 10px 15px var(--primary-color);
        }

        @media (min-width: 768px) {
            width: 40vw;
            margin: 20px 0px; 
        }

        @media (min-width: 992px) {
            // width: 25vw;
            margin: 20px 0px; 
        }

        @media (min-width: 1025px) {
            width: 25vw;
            margin: 20px 0px; 
        }

        .player{
            color: var(--primary-color);
            font-size: 25px;
            margin: 10px 0px;
            overflow-x: scroll;
            white-space: nowrap;
            padding: 0px 20px;

            &::-webkit-scrollbar {
                height: 4px;
            }

            &::-webkit-scrollbar-thumb {
                background-color: var(--primary-color); 
                border-radius: 3px;
            }

            &::-webkit-scrollbar-track {
                background: transparent;
            }

            // For wider tablets
            @media (min-width: 992px) {
                font-size: 35px;
            }

            //For laptops with smaller displays
            @media (min-width: 1025px) {
                font-size: 30px;
            }

            @media (min-width: 2560px) {
                font-size: 45px;
            }
        }

        

        .divider{
            width: 99%;
            border: 2px solid var(--primary-color);
            margin-top: 5px;
            margin-bottom: 20px;
        }

        .remove-wrapper{
            button{
                padding: 10px 20px;
                border-radius: 10px;
                
                border: 2px solid var(--bg-color);
                color: var(--bg-color);
                background-color: red;

                font-weight: 600;
                font-size: 16px;
                transition: 0.3s;

                &:hover{
                    cursor: pointer;
                    border: 2px solid var(--primary-color);
                    color: var(--primary-color);
                }
            }

            margin-bottom: 10px;
        }

        .edit-wrapper{
            button{
                padding: 10px 20px;
                border-radius: 10px;

                background-color: rgb(20, 174, 56);
                color: var(--bg-color);
                border: 2px solid var(--bg-color);

                font-weight: 600;
                font-size: 16px;
                transition:  0.3s;
                
                &:hover{
                    cursor: pointer;
                    border: 2px solid var(--primary-color);
                    color: var(--primary-color);
                }
            }

            margin-bottom: 10px;
        }
    }
</style>