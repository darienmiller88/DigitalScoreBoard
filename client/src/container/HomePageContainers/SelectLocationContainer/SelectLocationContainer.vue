<script setup lang="ts">
    import Select from '../../../components/Select/Select.vue';
    
    import { optionsStore } from '../../../stores/optionsStore';
    import { HomePageStore } from '../../../stores/PageStores/HomePageStore';
    import { storeToRefs } from 'pinia';
    import { scoreBoardApi } from '../../../api/api';
    import { onMounted } from 'vue';
    import { PlayerCard } from '../../../types/types';

    const { currentLocation } = storeToRefs(HomePageStore())
    const { setCurrentLocation, setAvailablePlayers, setCurrentPlayers } = HomePageStore()
    const { allLocationOptions } = optionsStore()


    const onChangeSelect = async (event: Event) => {
        const locationName = (event.target as HTMLSelectElement).value
        
        setCurrentPlayers([])
        setCurrentLocation(locationName)
        loadPlayersFromLocation(locationName)
    }

    const loadPlayersFromLocation = async (locationName: string) => {
        try {
            const playersResult = await scoreBoardApi.get<PlayerCard[]>(`/get-all-users/${locationName}`)

            //Set the current players to an empty array whenever new players are loaded from a location.

            //And then set the available list of players to the ones we just retrieved
            setAvailablePlayers(playersResult.data.map(player => ({
                player_name: player.username,
                isAddedToGame: false
            })))            
        } catch (error) {
            console.log(error);
        }
    }
    
    onMounted(() => {
        //Load the first location in IF there is no current location saved in local storage
        if (!currentLocation.value) {
            currentLocation.value = allLocationOptions[0]
        }

        loadPlayersFromLocation(currentLocation.value)        
    })
</script>

<template>
    <div class="choose-site-wrapper">
        <span>Choose site:</span>
        <Select :options="allLocationOptions" :selectModel="currentLocation" :onChange="onChangeSelect" />
    </div>
</template>

<style scoped lang="scss">
    .choose-site-wrapper{
        text-align: center;
        margin: 20px;

        span{
            font-size: 24px;
            margin-right: 10px;
            color: var(--primary-color);

            @media (min-width: 768px) {
                font-size: 30px;
            }
        }
    }
</style>