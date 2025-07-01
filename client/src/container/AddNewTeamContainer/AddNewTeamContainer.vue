<script setup lang="ts">
    import { onMounted, ref } from 'vue';
    import { selectedTeamStore } from '../../stores/selectedLocationStore';
    import { optionsStore } from "../../stores/optionsStore"
    import { teamCardsStore } from "../../stores/teamCardsStore";
    import { storeToRefs } from 'pinia';
    import { Location } from "../../types/types"
    import { scoreBoardApi } from '../../api/api';
    
    //Components
    import Select from '../../components/Select/Select.vue';
    import Loading from '../../components/Loading/Loading.vue';

    //ref variables
    const { teamCards } = storeToRefs(teamCardsStore())
    const { remainingLocationOptions } = storeToRefs(optionsStore())
    const { selectedTeam } = storeToRefs(selectedTeamStore())

    //store methods
    const { setRemainingLocationOptions } = optionsStore()
    const { addTeamCard } = teamCardsStore()
    const { setSelectedTeam } = selectedTeamStore()
    
    const isLoading = ref<boolean>(true)

    //On click to load in the current selected team based on the name from the options.
    const optionClicked = async (event: Event) => {
        const selectedValue = (event.target as HTMLSelectElement).value;

        try {
            const locationResponse = await scoreBoardApi.get<Location>(`/get-location/${selectedValue}`)
            
            setSelectedTeam({
                players: locationResponse.data.users.map(user => user.username),
                score: 0,
                team_name: locationResponse.data.location_name
            })
        } catch (error) {
            console.log("err in clicking option:", error);
        }
    }

    const addTeam = () => {        
        
        //Create a new anonymous Team card with the following data:
        //The location for which team is playing, a score of 0, and no players (yet).
        addTeamCard({
            team_name: selectedTeam.value.team_name,
            score: 0,
            players: []
        })

        //After adding the card, remove it from options.
        setRemainingLocationOptions(teamCards.value)
        
        //Set the current visible option for the all of the 
        selectedTeam.value.team_name = remainingLocationOptions.value[0]    
    }

    onMounted(async () => {
        
        //On pade load, load the remaining teams left to the total remaining teams
        setRemainingLocationOptions(teamCards.value)
        
        isLoading.value = false
    })
</script>

<template>
    <div>
        <Loading :height="50" :usePrimary="true" v-if="isLoading"/>
        <div class="team-game-location" v-if="!isLoading && teamCards.length < 2">
            <Select 
                :options="remainingLocationOptions"
                :onChange="optionClicked"
                :selectModel="selectedTeam.team_name"
            />
            
            <button 
                @click="addTeam"
                class="add-team-button"
            >Add Team</button> 
        </div>
    </div>
</template>


<style scoped lang="scss">
    .team-game-location{
        text-align: center;
        
        @media screen and (min-width: 768px) {
            display: flex;
            align-items: center;
            justify-content: center;
        }

        font-size: 30px;
        transition: 0.5s;

        .add-team-button{
            padding: 10px 20px;
            margin: 0px 10px;

            border: none;
            border-radius: 5px;

            background-color: dodgerblue;
            color: aliceblue;
            font-size: 20px;
            transition: 0.3s;

            @media screen and (min-width: 768px) {
                &:hover{
                    cursor: pointer;
                    font-size: 26px;
                }
            }
        }
    }
</style>