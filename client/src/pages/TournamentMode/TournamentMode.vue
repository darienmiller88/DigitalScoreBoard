<script setup lang="ts">
    import TeamCards from '../../components/TeamCards/TeamCards.vue';
    import SaveGame from '../../components/SaveGame/SaveGame.vue';
    import LocationSelection from '../../components/LocationSelection/LocationSelection.vue';
    import { scoreCardsStore } from "../../stores/scoreCardsStore"
    import { ButtonState } from "../../stores/buttonActiveStore"
    import { ref } from 'vue';

     //Stateful methods
    const { resetAllPoints } = scoreCardsStore()
    const isLoading = ref<boolean>(true)
</script>

<template>
    <!-- Contains the ADAPT locations in the select tag -->
    <LocationSelection :buttonState="ButtonState.CREATE_NEW_TEAM_GAME"/>

    <div class="reset-all-points-wrapper">
        <button type="button" @click="resetAllPoints" >
            Reset All Points
        </button>
    </div>

    <!-- Shows Create new Team game form when "Create new team game" button is clicked  --> 
    <TeamCards />

    <!-- Saves a game to the server -->
    <SaveGame />
</template>

<style scoped lang="scss">
    .reset-all-points-wrapper{
        text-align: center;
        margin: 25px;

        button{
            border-radius: 10px;
            border: none;

            padding: 10px 18px;
            font-size: 16px;
            font-weight: 700;
            transition: 0.3s;
            
            background-color: red;
            color: aliceblue;

            &:hover{
                cursor: pointer;
                background-color: rgba($color: #e30909, $alpha: .8);
            }
            
            &:active{
                transform: translateY(-5px);
            }
        }
    }
</style>