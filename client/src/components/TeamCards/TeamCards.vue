<script setup lang="ts">
    import TeamCard from '../TeamCard/TeamCard.vue'
    import { ref } from "vue"; 
    import { storeToRefs } from 'pinia'
    import { buttonActiveStore, ButtonState } from '../../stores/buttonActiveStore'
    import { teamCardsStore } from "../../stores/teamCardsStore"
    import Modal from "../Modal/Modal.vue"
    import AddNewPlayer from "../AddNewPlayer/AddNewPlayer.vue"
    
    let showAddTeamPlayerModal = ref(false)
    let showTeamPlayersModal = ref(false)
    const { currentButtonGroupState } = storeToRefs(buttonActiveStore())
    const { teamCards } = storeToRefs(teamCardsStore())

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
            :openAddTeamPlayerModal="() => showAddTeamPlayerModal = true"
        />
    </div>

    <Modal 
      :modal-header="'Add team member'"
      :show="showAddTeamPlayerModal"
      :modal-content="AddNewPlayer"
      :onHide="() => showAddTeamPlayerModal = false"
    />

    <Modal 
      :modal-header="'View team players'"
      :show="showTeamPlayersModal"
      :modal-content="AddNewPlayer"
      :onHide="() => showTeamPlayersModal = false"
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
        max-height: 50vh;
        overflow-y: scroll;

        @media only screen and (min-width: 768px){
            display: grid;
            grid-template-columns: repeat(2, 1fr);
            gap: 30px;
            max-height: 75vh;
        }

        @media only screen and (min-width: 992px){
            display: grid;
            grid-template-columns: repeat(3, 1fr);
            gap: 50px;
            max-height: 75vh;
        }

        @media only screen and (min-width: 1400px){
            display: grid;
            grid-template-columns: repeat(4, 1fr);
            gap: 30px;
            max-height: 75vh;
        }
    }
</style>