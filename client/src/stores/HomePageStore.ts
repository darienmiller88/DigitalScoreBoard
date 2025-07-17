import { defineStore } from 'pinia';
import { ref } from 'vue';
import { AvailablePlayer } from '../types/types';

export const HomePageStore = defineStore("HomePageStore", () => {
    const availablePlayersToAdd = ref<AvailablePlayer[]>([])
    const currentPlayersInGame = ref<string[]>([])
    const currentLocation = ref<string>("")
    const isGameCreated = ref<boolean>(false)  

    const toggleGameCreatedStatus = (gameCreatedStatus: boolean) =>{
        isGameCreated.value = gameCreatedStatus
    }

    const setAvailablePlayers = (newPlayers: AvailablePlayer[]) => {
        availablePlayersToAdd.value = newPlayers
    }

    const setCurrentPlayers = (newPlayers: string[]) => {
        currentPlayersInGame.value = newPlayers
    }

    const setCurrentLocation = (newLocation: string) => {
        currentLocation.value = newLocation
    }

    //Add the player to the game, and set their availibility to false.
    const addAvailalePlayerToGame = (playerIndex: number, playerName: string) => {
        currentPlayersInGame.value = [...currentPlayersInGame.value, playerName]

        availablePlayersToAdd.value[playerIndex].isAddedToGame = !availablePlayersToAdd.value[playerIndex].isAddedToGame
    }

    //Remove the player from the game, and set their availibility to true.
    const removeAvailablePlayerFromGame = (playerIndex: number, playerName: string) => {
        currentPlayersInGame.value = currentPlayersInGame.value.filter(_playerName => {
            return _playerName != playerName
        })

        availablePlayersToAdd.value[playerIndex].isAddedToGame = !availablePlayersToAdd.value[playerIndex].isAddedToGame
    }
    
    return { 
        availablePlayersToAdd, 
        isGameCreated,
        currentPlayersInGame, 
        setAvailablePlayers, 
        setCurrentPlayers, 
        currentLocation, 
        setCurrentLocation,
        toggleGameCreatedStatus,
        addAvailalePlayerToGame,
        removeAvailablePlayerFromGame
    }
}, {
    persist: true
})