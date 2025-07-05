<script setup lang="ts">
    import { useWindowSize } from "@vueuse/core"

    const { width } = useWindowSize();
    const availablePlayers = ["darien miller niller nuiller", "vicky", "sean", "kash", "darien", "vicky", "sean", "kash"]
</script>

<template>
    <div class="available-players">
        <div v-for="player in availablePlayers" class="available-player">
            <span>{{ player }}</span>
            <button v-if="width >= 768">Add To Game</button>
            <button v-else>Add</button>
        </div>
    </div>
</template>

<style scoped lang="scss">
    .available-players{
        display: grid;
        // flex-direction: column;
        // align-items: center;
        // border: 2px solid black;
        width: fit-content;
        margin: auto;
        margin-bottom: 20px;
        max-height: 20vh;
        overflow-y: scroll;

        @media (min-width: 768px) {
            max-height: 30vh;
        }

        .available-player{
            display: grid;
            grid-template-columns: 75% auto;
            // border: 2px salmon solid;
            width: 75vw;
            margin: 5px;

            //tablets
            @media (min-width: 768px) {
                grid-template-columns: 65% auto;
                width: 65vw;
            }

            // laptops
            @media (min-width: 1025px) {
                grid-template-columns: 75% auto;
                width: 55vw;
            }
            
            span{
                border: 2px solid var(--primary-color);
                border-right: none;
                border-top-left-radius: 8px;
                border-bottom-left-radius: 8px;

                color: var(--primary-color);
                // margin-right: 10px;
                padding: 10px;
                white-space: nowrap;
                font-size: 24px;
                overflow-x: scroll;

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
            }

            button{
                border: none;
                font-size: 24px; 
                font-weight: 600;  
                background-color: var(--primary-color);
                color: var(--bg-color);
                border-top-right-radius: 8px;
                border-bottom-right-radius: 8px;
                transition: 0.2s;

                &:hover{
                    cursor: pointer;
                    background-color: var(--primary-color-transparent);
                }

                &:active{
                    // box-shadow: 6px 6px 6px var(--primary-color-transparent);
                    transform: translateY(-2px);
                    // transform: scale(0.98);
                }
            }
        }
    }
</style>