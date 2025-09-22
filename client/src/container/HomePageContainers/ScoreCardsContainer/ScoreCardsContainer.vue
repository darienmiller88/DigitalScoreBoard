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
        transition: 0.3s;

        &:hover{
            box-shadow: 0 4px 10px var(--primary-color);
        }

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

        //2k resolution
        @media (min-width: 2560px){
            display: grid;
            gap: 25px;
            max-height: 75vh;
            max-width: 95vw;
        }

        //4k resolution
        @media (min-width: 3840px){
            display: grid;
            gap: 50px;
            max-height: 75vh;
            max-width: 90vw;
            padding: 0px 30px;
            border: 2px solid black;
        }
    }
</style>