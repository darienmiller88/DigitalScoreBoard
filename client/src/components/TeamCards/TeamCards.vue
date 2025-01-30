<script setup lang="ts">
    import TeamCard from '../TeamCard/TeamCard.vue'
    import { ref } from "vue"; 
    import { storeToRefs } from 'pinia'
    import { buttonActiveStore, ButtonState } from '../../stores/buttonActiveStore'
    import { Team } from "../../types/types"
    import Modal from "../Modal/Modal.vue"
    import AddNewPlayer from "../AddNewPlayer/AddNewPlayer.vue"
    
    let show = ref(false)
    const { currentButtonGroupState } = storeToRefs(buttonActiveStore())
    const teams: Team[] = [
        {
            team_name: "Lawrence",
            score: 0,
            players: []
        },
        {
            team_name: "Pelham Bay",
            score: 0,
            players: []
        },
        {
            team_name: "Elmwood",
            score: 0,
            players: []
        },
        {
            team_name: "Flushing",
            score: 0,
            players: []
        }
    ]
</script>

<template>
    <div class="team-cards">
        <TeamCard
            v-if="currentButtonGroupState == ButtonState.CREATE_NEW_TEAM_GAME"
            v-for="(team, index) in teams" 
            :key="index"
            :team="team"
            :open-modal="() => show = true"
        />
    </div>

    <Modal 
      :modal-header="'Add team member'"
      :show="show"
      :modal-content="AddNewPlayer"
      :onHide="() => show = false"
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
        max-width: 90vw;
        max-height: 50vh;
        overflow-y: scroll;

        @media only screen and (min-width: 768px){
            display: grid;
            grid-template-columns: repeat(4, 1fr);
            gap: 30px;
            max-height: 75vh;
        }
    }
</style>