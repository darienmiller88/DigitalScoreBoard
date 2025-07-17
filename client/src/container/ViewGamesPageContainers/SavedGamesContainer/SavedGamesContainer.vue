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
    import ViewSavedGamePlayers from '../../../components/ViewSavedGamePlayers/ViewSavedGamePlayers.vue';
    import ViewSavedGameTeams from '../../../components/ViewSavedGameTeams/ViewSavedGameTeams.vue';

    const isLoading = ref<boolean>(true)
    let showPeopleWhoPlayed = ref<boolean>(false)
    let showTeamsWhoPlayed = ref<boolean>(false)

    let games: SavedGame[] = []
    let playersInSavedGame = ref<PlayerCard[]>([])
    let teamsInSavedGame = ref<Team[]>([])

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

    onMounted(async () => {
        try {
            const savedGameResult = await scoreBoardApi.get<SavedGame[]>("/get-saved-games")

            //Convert each date to a more readable format
            savedGameResult.data.forEach(game => {                
                game.created_at = new Date(game.created_at).toLocaleString()
            })

            games = [...games, ...savedGameResult.data].reverse()
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