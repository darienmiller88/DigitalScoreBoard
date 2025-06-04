<script setup lang="ts">
    import { Team } from "../../types/types"

    const props = defineProps<{
        teamsInSavedGame: Team[]
    }>()
</script>

<template>
    <div class="teams"> 
        <div class="team" v-for="team in props.teamsInSavedGame">
            {{ team.team_name }}

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
        border: 2px solid black;

        .team{
            border: 2px solid red;

            .players{
                border: 2px solid purple;

                .player{

                }
            }
        }
    }
</style>