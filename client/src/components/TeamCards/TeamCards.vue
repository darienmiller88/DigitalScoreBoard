<script setup lang="ts">
    import TeamCard from '../TeamCard/TeamCard.vue';
    import { storeToRefs } from 'pinia';
    import { buttonActiveStore, ButtonState } from '../../stores/buttonActiveStore';
    import { Team } from "../../types/types"
    
    const { currentButtonGroupState } = storeToRefs(buttonActiveStore())
    const teams: Team[] = [
        {
            team_name: "L.L.R.D",
            score: 0,
            players: []
        },
        {
            team_name: "Vought",
            score: 0,
            players: []
        },
        {
            team_name: "Landshark",
            score: 0,
            players: []
        },
        {
            team_name: "Trainheads",
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
            :team_name="team.team_name"
            :players="team.players"
            :score="team.score"
        />
    </div>
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