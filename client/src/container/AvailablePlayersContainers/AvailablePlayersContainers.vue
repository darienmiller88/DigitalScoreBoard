<script setup lang="ts">
    import { storeToRefs } from "pinia";
    import { HomePageStore } from "../../stores/HomePageStore";
    import { useWindowSize } from "@vueuse/core"

    const { availablePlayersToAdd, currentPlayersInGame } = storeToRefs(HomePageStore())
    const { width } = useWindowSize();
</script>

<template>
    <div v-if="availablePlayersToAdd.length">
        <div class="players-indicator" >Number of Players added: {{ currentPlayersInGame.length }} </div>
        <div class="available-players" >
            <div v-for="player in availablePlayersToAdd" class="available-player">
                <span>{{ player }}</span>
                <button v-if="width >= 768">Add To Game</button>
                <button v-else>Add</button>
            </div>
        </div>
    </div>
    <div class="no-players" v-else>
        No players available for selected location.
    </div>
</template>

<style scoped lang="scss">
    .players-indicator{
        text-align: center;
        color: var(--primary-color);
        margin: 10px;
        font-size: 20px;

        @media (min-width: 768px) {
            font-size: 28px;
        }
    }

    .no-players{
        text-align: center;
        margin-bottom: 20px;
        color: var(--primary-color);

        @media (min-width: 768px) {
            font-size: 30px;
        }
    }

    .available-players{
        display: grid;
        // flex-direction: column;
        // align-items: center;
        border: 2px solid var(--primary-color);
        border-radius: 5px;

        width: fit-content;
        margin: auto;
        margin-bottom: 20px;
        padding: 8px;
        
        max-height: 20vh;
        overflow-y: scroll;
        transition: 0.3s;

        @media (min-width: 768px) {
            border: none;
            max-height: 30vh;
        }

        @media (min-width: 1025px) {
            border: none;
            box-shadow: var(--primary-color-transparent) 1.95px 1.95px 2.6px;            

            &:hover{
                box-shadow: var(--primary-color-transparent) 0px 30px 60px -10px, rgba(0, 0, 0, 0.3) 0px 30px 60px -30px, rgba(10, 37, 64, 0.35) 0px -2px 6px 0px inset;
            }
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
                transition: 0.3s;
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

                @media (min-width: 1025px) {
                    font-size: 30px;

                    &:hover{
                        color: var(--toggle-background);
                    }
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