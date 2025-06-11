<script setup lang="ts">
    import { onMounted, ref } from 'vue';
    import { SavedGame } from '../../types/types';
    import { scoreBoardApi } from '../../api/api';
    import { PlayerCard, Team} from "../../types/types"
    import Loading from '../Loading/Loading.vue';
    import Game from '../../components/Game/Game.vue';
    import Modal from '../Modal/Modal.vue';
    import ViewSavedGamePlayers from '../ViewSavedGamePlayers/ViewSavedGamePlayers.vue';
    import ViewSavedGameTeams from '../ViewSavedGameTeams/ViewSavedGameTeams.vue';

    const isLoading = ref<boolean>(true)
    let showPeopleWhoPlayed = ref<boolean>(false)
    let showTeamsWhoPlayed = ref<boolean>(false)

    let games: SavedGame[] = [
        {
            id: "1",
            created_at: new Date().toLocaleString(),
            average_points: 90,
            total_points: 5000,
            winner: {
                username: "darien",
                score: 1000
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
                     {
                        username: "Reina",
                        score: 300
                    },
                    {
                        username: "Marth",
                        score: 400
                    }
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
        },
        {
            id: "4",
            created_at: new Date().toLocaleString(),
            average_points: 90,
            total_points: 5000,
            winner: {
                username: "Lawrence",
                score: 4000
            },
            teams: [
                {
                    team_name: "Lawrence",
                    score: 5000,
                    players: ["Michelle", "Nijmah", "Betty", "Denise"]
                },
                {
                    team_name: "Elmwood",
                    score: 4000,
                    players: ["Shaniqua", "Micki", "William", "Luis"]
                }
            ],
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
                    {
                        username: "Betty",
                        score: 700
                    },
                    {
                        username: "Denise",
                        score: 400
                    },
                    {
                        username: "Shaniqua",
                        score: 700
                    },
                    {
                        username: "Micki",
                        score: 400
                    },
                    {
                        username: "Luis",
                        score: 700
                    },
                    {
                        username: "William",
                        score: 400
                    },
                ],
                location_name: "Elmwood"
            }
        }
    ]

    let playersInSavedGame = ref<PlayerCard[]>([])
    let teamsInSavedGame = ref<Team[]>([])

    const openViewSavedGamesPlayers = (game: SavedGame) => {
        let location = game.location

        if (location) {
            playersInSavedGame.value = location.users
        }

        showPeopleWhoPlayed.value = true
    }

    const openViewSavedGamesTeams = (game: SavedGame) => {
        let teams = game.teams

        if (teams) {
            teamsInSavedGame.value = teams            
        }
        
        showTeamsWhoPlayed.value = true
    }

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
          :game="game"
          :key="index"
          :isSavedGameATeamGame="game.teams !== undefined"
          :viewTeamsInSavedGame="openViewSavedGamesTeams"
          :viewPlayersInSavedGame = "openViewSavedGamesPlayers"
        />
    </div>

    <!-- Modal to show the teams that played if the client submitted a team game -->
    <Modal 
        :modalHeader="'View People who played'"
        :show="showPeopleWhoPlayed"
        :modalContent="ViewSavedGamePlayers"
        :onHide="() => showPeopleWhoPlayed = false"
        :modalProps="{ playersInSavedGame: playersInSavedGame }"
    />


    <!-- Modal to show the teams that played if the client submitted a team game -->
    <Modal 
        :modalHeader="'View Teams who played'"
        :show="showTeamsWhoPlayed"
        :modalContent="ViewSavedGameTeams"
        :onHide="() => showTeamsWhoPlayed = false"
        :modalProps="{ teamsInSavedGame: teamsInSavedGame }"
    />
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