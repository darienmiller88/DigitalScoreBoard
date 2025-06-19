<script setup lang="ts">
    import ScoreCards from '../../components/ScoreCards/ScoreCards.vue';
    import PageTitle from '../../components/PageTitle/PageTitle.vue';
    import SaveGame from '../../components/SaveGame/SaveGame.vue';
    import ResetPoints from '../../components/ResetPoints/ResetPoints.vue';
    import { onMounted } from 'vue';
    import { scoreCardsStore } from "../../stores/scoreCardsStore"

    //Stateful methods
    const { resetAllPoints, totalPoints } = scoreCardsStore()
  
    onMounted(async () => {
        // if (locationsFromLocalStorage.value.length) {
        //     locations = locationsFromLocalStorage.value
        //     options.value = locations.map(location => {
        //         return location.location_name
        //     })   
        // } else {
        // }

        try {
            //Find the target location in the list of locations from the database.
            // const targetLocation = locations.find(location => {
            //     return location.location_name === selectedLocation.value
            // })

            

            // //If the scoreCards array is empty, populate it with the users from the location.
            // if (targetLocation && scoreCards.value.length === 0) {
            //     setCards(targetLocation.users)
            // }

            // //If the list of users from the target location is greater than the list of users currently,
            // //append the difference to the current list of users
            // if (targetLocation && targetLocation.users.length > scoreCards.value.length) {
            //     const index = targetLocation.users.length - scoreCards.value.length

            //     scoreCards.value = [...scoreCards.value, ...targetLocation.users.slice(index)]
            // }

            // //If the list of users from the target location is greater than the list of users currently,
            // //append the difference to the current list of users
            // if (targetLocation && targetLocation.users.length <= scoreCards.value.length) {
            //     setCards(targetLocation.users)                
            // }

            // console.log("target location:", targetLocation, "scoreCards:", scoreCards.value);
            
            //If there are no current score cards set in local storage, retrieve them from the database
            // locationsResponse.data.forEach((location, i) => {
            //     if (location.location_name === selectedLocation.value) {
            //         setCards(locationsResponse.data[i].users)
            //     }
            // })
            // if (!scoreCards) {
            // } else{
            //     console.log("card:", scoreCards.value);
                
            // }         
        } catch (error) {
            console.log("err:", error);
        }
        
        // console.log("locations:", locations, "option:", options, "host:", window.location.hostname);
        // console.log("api:", scoreBoardApi.getUri());
        
    })
</script>

<template>
    <PageTitle :titleName="'Create New Game'"/>

    <ResetPoints :resetAllPoints="resetAllPoints"/>

    <div class="total-points">
        Total Points: {{ totalPoints() }} 
    </div>

    <!-- Shows all users when "Add new users" is clicked -->
    <ScoreCards />

    <!-- Saves a game to the server -->
    <SaveGame />
</template>

<style scoped lang="scss">
    .error{
        text-align: center;
        color: red;
        font-size: 25px;
    }

    .total-points{
        text-align: center;
        font-size: 30px;
        color: var(--primary-color);
    }  
</style>