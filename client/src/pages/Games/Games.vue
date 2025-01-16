<script setup lang="ts">
    import Game from '../../components/Game/Game.vue';
    import { SavedGame } from "../../types/types"
    import { onMounted, ref } from 'vue';
    import { scoreBoardApi } from "../../api/api"
    import { Icon } from "@iconify/vue"

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
    <div class="title">Saved Games</div>
    <div class="loading-wrapper" v-if="isLoading">
        <Icon icon="line-md:loading-twotone-loop" color="#61dafb" height="60" width="60"/>
    </div>
    <div class="games" v-else>
       <Game 
          v-for="(game, index) in games"
          v-bind="game"
          :key="index"
       />
    </div>
</template>

<style scoped lang="scss">
    .title{
        font-size: 40px;
        text-align: center;
        color: var(--main-text-color);
        margin: 15px;

        @media screen and (min-width: 768px) {
            font-size: 50px;
        }
    }

    .loading-wrapper{
        text-align: center;
        margin-top: 60px;
    }

    .games{
        // border: 1px solid var(--main-text-color); 
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
            box-shadow: 0 4px 10px var(--main-text-color);
        }

    }
</style>