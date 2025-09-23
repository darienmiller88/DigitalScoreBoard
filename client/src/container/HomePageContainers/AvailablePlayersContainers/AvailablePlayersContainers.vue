<script setup lang="ts">
    import { storeToRefs } from "pinia";
    import { HomePageStore } from "../../../stores/PageStores/HomePageStore";
    import { scoreCardsStore } from "../../../stores/scoreCardsStore";
    import { useWindowSize } from "@vueuse/core"

    const { scoreCards } = storeToRefs(scoreCardsStore())
    const { addScoreCard } = scoreCardsStore()
    const { availablePlayersToAdd, currentPlayersInGame, remainingPlayersInGame } = storeToRefs(HomePageStore())
    const { addAvailalePlayerToGame, removeAvailablePlayerFromGame, toggleGameCreatedStatus } = HomePageStore()
    const { width } = useWindowSize();

    //IF there are scorecards added, that means a game has started, show the players who have not been added yet. 
    //OTherwise, show all available players to can be added.
    let players = scoreCards.value.length ? remainingPlayersInGame : availablePlayersToAdd

    const addPlayerToScoreCardList = (playerIndex: number, playerName: string) => {
        
        //if the player has a game in session, add them to the list of score cards
        if (scoreCards.value.length) {
            addScoreCard({
                username: playerName,
                score: 0
            })
        }

        addAvailalePlayerToGame(playerIndex, playerName)

        //When there are no remaining players left, return to the game screen.
        if (!remainingPlayersInGame.value.length) {
            toggleGameCreatedStatus(true)
        }
    }
</script>

<template>
    <div v-if="availablePlayersToAdd.length">
        <div class="players-indicator">Number of Players added: {{ currentPlayersInGame.length }} </div>
        <div class="available-players">
            <div v-for="(player, i) in players" class="available-player">
                <div class="player-name">{{ player.player_name }}</div>
                <div class="buttons-wrapper">
                    <div class="remove-player-wrapper" v-if="player.isAddedToGame">
                        <button class="base-btn" @click="() => removeAvailablePlayerFromGame(i, player.player_name)">Remove</button>
                    </div>
                    <div class="add-player-wrapper" v-else>
                        <button class="base-btn" @click="() => addPlayerToScoreCardList(i, player.player_name)">
                            <span v-if="width >= 768">Add To Game</span> 
                            <span v-else>Add</span>
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="no-players" v-else-if="!availablePlayersToAdd.length || !remainingPlayersInGame.length">
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

        @media (min-width: 3840px) {
            font-size: 55px;
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
        
        max-height: 35vh;
        overflow-y: scroll;
        transition: 0.3s;

        @media (min-width: 768px) {
            border: none;
            max-height: 30vh;
        }

        @media (min-width: 1025px) {
            border: none;
            box-shadow: var(--primary-color-transparent) 1.95px 1.95px 2.6px; 
            max-height: 40vh;           

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
            
            .player-name{
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

                @media (min-width: 3840px) {
                    font-size: 85px;
                    padding: 30px;
                }
            }

            .buttons-wrapper{
                .remove-player-wrapper{
                    height: 100%;

                    button{
                        background-color: rgb(241, 50, 10);

                        &:hover{
                            cursor: pointer;
                            background-color: rgba(241, 50, 10, 0.7);
                        }
                    }
                }
                
                .add-player-wrapper{
                    height: 100%;

                    button{
                        &:hover{
                            cursor: pointer;
                            background-color: var(--primary-color-transparent);
                        }
                    }
                }

                .base-btn{
                    border: none;
                    font-size: 24px; 
                    font-weight: 600;  
                    background-color: var(--toggle-background);
                    color: var(--bg-color);
                    border-top-right-radius: 8px;
                    border-bottom-right-radius: 8px;
                    height: 100%;
                    width: 100%;
                    transition: 0.2s;
        
                    @media (min-width: 3840px) {
                        font-size: 65px;
                    }

                    &:active{
                        transform: translateY(-2px);
                    }
                }
                
            }
        }
    }
</style>