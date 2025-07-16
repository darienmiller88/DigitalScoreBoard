<script setup lang="ts">
    import ScoreCard from '../../../components/ScoreCard/ScoreCard.vue'
    import { scoreCardsStore } from "../../../stores/scoreCardsStore"
    import { storeToRefs } from 'pinia';

    const { scoreCards } = storeToRefs(scoreCardsStore())
</script>

<template>
    <div
        id="user-cards-id" 
        class="user-cards" 
    >
        <ScoreCard 
            v-for="(card, index) in scoreCards" 
            :key="card.username" 
            :username="card.username" 
            :card-index="index"
            :point-value="100"
            :score="card.score"
        />
    </div>
</template>

<style scoped lang="scss">
    .user-cards{
        display: grid;
        grid-template-columns: 1fr;    

        width: fit-content; 
        max-width: 90vw;
        max-height: 50vh;
        overflow-y: scroll;

        margin: auto;
        padding: 0px 10px;

        @media (min-width: 768px){
            display: grid;
            grid-template-columns: repeat(2, 1fr);
            gap: 30px;
            max-height: 75vh;
        }

        @media (min-width: 1025px){
            display: grid;
            grid-template-columns: repeat(3, 1fr);
            gap: 30px;
            max-height: 75vh;
        }

        @media (min-width: 2560px){
            display: grid;
            grid-template-columns: repeat(5, 1fr);
            gap: 25px;
            max-height: 75vh;
            max-width: 95vw;
        }
    }
</style>