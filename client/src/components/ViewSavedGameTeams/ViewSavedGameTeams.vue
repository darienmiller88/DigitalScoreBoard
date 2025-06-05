<script setup lang="ts">
    import { Team } from "../../types/types"

    const props = defineProps<{
        teamsInSavedGame: Team[]
    }>()
</script>

<template>
    <div class="teams"> 
        <div class="team" v-for="team in props.teamsInSavedGame">
            <div class="team-name">
                {{ team.team_name }}
            </div> 

            <div class="players">
                <div class="player" v-for="player in team.players">
                    {{ player }}
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped lang="scss">
    .teams{
        text-align: center;
        border: 2px solid black;
        height: 30vh;
        overflow-y: scroll;

        .team{
            border: 2px solid red;

            .team-name{ 
                font-size: 25px;
                font-weight: 700;
                padding: 20px;
            }

            .players{
                border: 2px solid purple;
                height: 13vh;
                overflow-y: scroll;
                
                .player{
                    font-size: 18px;
                    padding: 10px;
                }
            }
        }
    }
</style>