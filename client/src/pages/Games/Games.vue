<script setup lang="ts">
    import Game from '../../components/Game/Game.vue';
    import { SavedGame } from "../../types/types"

    const game1: SavedGame = {
        id: "1",
        created_at: new Date().toLocaleString(),
        average_points: 90,
        total_points: 5000,
        winner: {
            username: "darien",
            score: 100
        },
        location: {
            id: "1",
            users: [
                {
                    username: "Reina",
                    score: 300
                },
                {
                    username: "Marth",
                    score: 400
                },
            ],
            location_name: "Pelham"
        }
    }
</script>

<template>
    <div class="games">
       <Game 
          v-bind="game1"
       />
    </div>
</template>

<style scoped>
    .games{
        border: 3px dotted red;
    }
</style>