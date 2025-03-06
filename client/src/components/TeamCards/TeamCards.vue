<script setup lang="ts">
    import TeamCard from '../TeamCard/TeamCard.vue'
    import { Card } from "../../types/types"
    import { scoreBoardApi } from "../../api/api"
    import { Team } from "../../types/types"
    import { ref } from "vue"; 
    import { storeToRefs } from 'pinia'
    import { buttonActiveStore, ButtonState } from '../../stores/buttonActiveStore'
    import { teamCardsStore } from "../../stores/teamCardsStore"
    import Modal from "../Modal/Modal.vue"
    import AddNewPlayer from "../AddNewPlayer/AddNewPlayer.vue"
    import ViewPlayers from '../ViewPlayers/ViewPlayers.vue';
    
    const { currentButtonGroupState } = storeToRefs(buttonActiveStore())
    const { teamCards } = storeToRefs(teamCardsStore())
    
    let showAddTeamPlayerModal = ref(false)
    let showTeamPlayersModal = ref(false)

    //Data to be ent to the "AddNewPlayer" component via the modal.
    let addTeamNewPlayerData = ref<{
        players: string[],
        locationName: string
    }>({
        players: [],
        locationName: ""
    })

    //Data to be ent to the "ViewPlayers" component via the modal.
    let viewPlayersData = ref<{
        teamPlayers: string[]
    }>({
        teamPlayers: []
    })

    const openAddTeamPlayerModal = (team: Team) => {
        addTeamNewPlayerData.value.players = team.players
        addTeamNewPlayerData.value.locationName = team.team_name
        showAddTeamPlayerModal.value = true
        
        console.log("addTeamNewPlayerData", addTeamNewPlayerData.value)
    }

    const openViewTeamPlayersModal = async (team: Team) => {
        try {
            const teamPlayersResponse = await scoreBoardApi.get<Card[]>(`/get-all-users/${team.team_name}`)
            
            viewPlayersData.value.teamPlayers = teamPlayersResponse.data.map(team => team.username)
            showTeamPlayersModal.value = true
        } catch (error) {
            console.log("error:", error)
        }
    }
</script>

<template>
    <div class="team-cards">
        <TeamCard
            v-if="currentButtonGroupState == ButtonState.CREATE_NEW_TEAM_GAME"
            v-for="(team, index) in teamCards" 
            :key="index"
            :card-index="index"
            :score="team.score"
            :point-value="100"
            :team="team"
            :openAddTeamPlayerModal="() => openAddTeamPlayerModal(team)"
            :openViewTeamPlayers="() => openViewTeamPlayersModal(team)"
        />
    </div>

    <Modal 
      :modalHeader="'Add team member'"
      :show="showAddTeamPlayerModal"
      :modalContent="AddNewPlayer"
      :onHide="() => showAddTeamPlayerModal = false"
      :modalProps="addTeamNewPlayerData"
    />

    <Modal 
      :modalHeader="'View team players'"
      :show="showTeamPlayersModal"
      :modalContent="ViewPlayers"
      :onHide="() => showTeamPlayersModal = false"
      :modalProps="viewPlayersData"
    />
</template>

<style scoped lang="scss">
    .team-cards{
        display: grid;
        grid-template-columns: 1fr;    
        margin: auto;
        padding: 10px 10px;

        // border: 2px solid white; 
        width: fit-content; 
        max-width: 95vw;
        max-height: 40vh;
        overflow-y: scroll;

        @media only screen and (min-width: 768px){
            display: grid;
            grid-template-columns: repeat(2, 1fr);
            gap: 50px;
            max-height: 75vh;
        }

        @media only screen and (min-width: 992px){
            display: grid;
            grid-template-columns: repeat(3, 1fr);
            gap: 40px;
        }

        @media only screen and (min-width: 1400px){
            display: grid;
            grid-template-columns: repeat(4, 1fr);
            gap: 30px;
        }
    }
</style>