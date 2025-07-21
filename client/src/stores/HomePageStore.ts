import { defineStore } from 'pinia';
import { ref } from 'vue';
import { AvailablePlayer } from '../types/types';

export const HomePageStore = defineStore("HomePageStore", () => {

    //EVERY player from the ADAPT site that wad chosen.
    const availablePlayersToAdd = ref<AvailablePlayer[]>([])

    //The players that has been added to a game by the client. It's basically a checkout line.
    const currentPlayersInGame = ref<string[]>([])

    //The players that have not been added to a game yet.
    const remainingPlayersInGame = ref<AvailablePlayer[]>([])
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

        availablePlayersToAdd.value[playerIndex].isAddedToGame = true
        setRemainingPlayers()
    }

    //Remove the player from the game, and set their availibility to true.
    const removeAvailablePlayerFromGame = (playerIndex: number, playerName: string) => {
        currentPlayersInGame.value = currentPlayersInGame.value.filter(_playerName => {
            return _playerName != playerName
        })

        availablePlayersToAdd.value[playerIndex].isAddedToGame = false
        setRemainingPlayers()
    }

    const setRemainingPlayers = () => {
        remainingPlayersInGame.value = availablePlayersToAdd.value.filter(playerName =>
            !currentPlayersInGame.value.includes(playerName.player_name)
        )
    }
    
    return { 
        availablePlayersToAdd, 
        isGameCreated,
        currentPlayersInGame, 
        setAvailablePlayers, 
        setCurrentPlayers, 
        remainingPlayersInGame,
        currentLocation, 
        setCurrentLocation,
        toggleGameCreatedStatus,
        addAvailalePlayerToGame,
        removeAvailablePlayerFromGame
    }
}, {
    persist: true
})