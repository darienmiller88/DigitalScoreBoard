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
                Players on: <i>Team {{ team.team_name }}</i>
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
        height: 50vh;
        overflow-y: scroll;

        .team{
            transition: 0.3s;
            box-shadow: 0 1px 3px rgba(5, 5, 5, 0.6);
            background-color: rgba(245, 245, 245, 0.6);
            margin: 15px;
            border-radius: 10px;

            &:hover{
                box-shadow: 0 3px 6px rgba(15, 15, 15, 0.9);
            }

            .team-name{ 
                font-size: 25px;
                font-weight: 700;
                padding: 20px;
            }

            .players{
                height: 18vh;
                overflow-y: scroll;
                
                .player{
                    font-size: 22px;
                    padding: 8px;
                    margin: 10px;
                    transition: 0.3s;
                    border-radius: 10px;
                    
                    box-shadow: 0 1px 3px rgba(5, 5, 5, 0.6);

                    &:hover{
                        box-shadow: 0 3px 6px rgba(15, 15, 15, 0.9);
                    }
                }
            }
        }
    }
</style>