import { defineStore } from 'pinia';
import { ref } from 'vue';
import { PlayerCard, AvailablePlayer } from '../types/types';

export const HomePageStore = defineStore("HomePageStore", () => {
    const availablePlayersToAdd = ref<AvailablePlayer[]>([])
    const currentPlayersInGame = ref<PlayerCard[]>([])
    const currentLocation = ref<string>("")

    const setAvailablePlayers = (newPlayers: AvailablePlayer[]) => {
        availablePlayersToAdd.value = newPlayers
    }

    const setCurrentPlayers = (newPlayers: PlayerCard[]) => {
        currentPlayersInGame.value = newPlayers
    }

    const setCurrentLocation = (newLocation: string) => {
        currentLocation.value = newLocation
    }

    //Add the player to the game, and set their availibility to false.
    const addAvailalePlayerToGame = (playerIndex: number, playerName: string) => {
        currentPlayersInGame.value = [...currentPlayersInGame.value, {
            username: playerName,
            score: 0
        }]

        availablePlayersToAdd.value[playerIndex].isAddedToGame = !availablePlayersToAdd.value[playerIndex].isAddedToGame
    }

    //Remove the player from the game, and set their availibility to true.
    const removeAvailablePlayerFromGame = (playerIndex: number, playerName: string) => {
        currentPlayersInGame.value = currentPlayersInGame.value.filter(player => {
            return player.username != playerName
        })

        availablePlayersToAdd.value[playerIndex].isAddedToGame = !availablePlayersToAdd.value[playerIndex].isAddedToGame
    }
    
    return { 
        availablePlayersToAdd, 
        currentPlayersInGame, 
        setAvailablePlayers, 
        setCurrentPlayers, 
        currentLocation, 
        setCurrentLocation,
        addAvailalePlayerToGame,
        removeAvailablePlayerFromGame
    }
}, {
    persist: true
})