<script setup lang="ts">
    import { onMounted, ref } from 'vue';
    import Game from '../../components/Game/Game.vue';
    import { SavedGame } from '../../types/types';
    import { scoreBoardApi } from '../../api/api';
    import Loading from '../Loading/Loading.vue';

    const isLoading = ref<boolean>(true)

    let games: SavedGame[] = [
        {
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
        },
        {
            id: "3",
            created_at: new Date().toLocaleString(),
            average_points: 90,
            total_points: 5000,
            winner: {
                username: "Michelle",
                score: 400
            },
            location: {
                id: "3",
                users: [
                    {
                        username: "Nijmah",
                        score: 300
                    },
                    {
                        username: "Michelle",
                        score: 400
                    },
                ],
                location_name: "Lawrence"
            }
        }
    ]

    onMounted(async () => {
        try {
            const savedGameResult = await scoreBoardApi.get<SavedGame[]>("/get-saved-games")
            games = [...games, ...savedGameResult.data]
            
            console.log("games:", games);
        } catch (error) {
            console.log("error:", error);
        }

        isLoading.value = false
    })
</script>

<template>
    <Loading :height="60" :usePrimary="true" v-if="isLoading"/>
    <div class="games" v-else>
       <Game 
          v-for="(game, index) in games"
          v-bind="game"
          :key="index"
       />
    </div>
</template>

<style scoped lang="scss">
    .games{
        overflow-y: scroll;

        height: 70vh;
        width: 95vw;
        margin: auto;
        transition: 0.3s;

        @media screen and (min-width: 768px) {
            width: fit-content;
            padding: 0px 30px;
        }

        &:hover{
            box-shadow: 0 4px 10px var(--primary-color);
        }
    }
</style>