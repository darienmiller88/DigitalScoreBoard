<script setup lang="ts">
    //Libraries
    import { onMounted, ref } from 'vue';
    import { SavedGame } from '../../../types/types';
    import { scoreBoardApi } from '../../../api/api';
    import { PlayerCard, Team } from "../../../types/types"

    //Components
    import Loading from '../../../components/Loading/Loading.vue';
    import Game from '../../../components/Game/Game.vue';
    import Modal from '../../../components/Modal/Modal.vue';
    import ViewSavedGamePlayers from '../../../components/ModalComponents/SavedGamesPageModalComponents/ViewSavedGamePlayers/ViewSavedGamePlayers.vue';
    import ViewSavedGameTeams from '../../../components/ModalComponents/SavedGamesPageModalComponents/ViewSavedGameTeams/ViewSavedGameTeams.vue';
    import DeleteSavedGame from '../../../components/ModalComponents/SavedGamesPageModalComponents/DeleteSavedGame/DeleteSavedGame.vue';

    const isLoading = ref<boolean>(true)
    let showPeopleWhoPlayed = ref<boolean>(false)
    let showTeamsWhoPlayed = ref<boolean>(false)
    let showDeleteGameModal= ref<boolean>(false)

    let games: SavedGame[] = []
    let playersInSavedGame = ref<PlayerCard[]>([])
    let teamsInSavedGame = ref<Team[]>([])
    let gameIndexToDelete = ref<number>(0)
    let gameToDelete = ref<SavedGame>({
        id: '',
        winner: {
            username: '',
            score: 0
        },
        location_name: '',
        created_at: '',
        total_points: 0,
        average_points: 0
    })

    const openViewSavedGamesPlayers = (game: SavedGame) => {
        let players: PlayerCard[] | undefined = game.players

        if (players) {
            playersInSavedGame.value = players
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

    const openDeleteSavedGameModal = (gameIndex: number) => {
        showDeleteGameModal.value = true

        gameIndexToDelete.value = gameIndex
        // removeGameFromArray(gameIndex)
    }

    const removeGameFromArray = async (gameIndex: number) => {
        games = games.filter((_, index) => {
            return index != gameIndex
        })
    }

    onMounted(async () => {
        try {
            const savedGameResult = await scoreBoardApi.get<SavedGame[]>("/get-saved-games")

            //Convert each date to a more readable format
            savedGameResult.data.forEach(game => {                
                game.created_at = new Date(game.created_at).toLocaleString()
                 
                if (game.players) {
                    game.players.sort((a, b) => b.score - a.score)
                }
            })

            games = [...savedGameResult.data].reverse()
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
          :gameIndex="index"
          :key="index"
          :isSavedGameATeamGame="game.teams !== undefined"
          :viewTeamsInSavedGame="openViewSavedGamesTeams"
          :viewPlayersInSavedGame = "openViewSavedGamesPlayers"
          :openDeletePlayerModal="openDeleteSavedGameModal"
        />
    </div>

    <!-- Modal to show the people that played if the client submitted a single player game -->
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

    <!-- Modal to give the client a chance to review if they want to delete a saced game -->
    <Modal 
        :modalHeader="'Delete Saved Game'"
        :show="showDeleteGameModal"
        :modalContent="DeleteSavedGame"
        :onHide="() => showDeleteGameModal = false"
        :modalProps="{ 
            gameIndex: gameIndexToDelete,
            hideModal: () => showDeleteGameModal = false,
            removeGameFromArray: removeGameFromArray,
            game: gameToDelete
        }"
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