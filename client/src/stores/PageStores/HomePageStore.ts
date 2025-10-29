import { defineStore } from 'pinia';
import { ref } from 'vue';
import { AvailablePlayer } from '../../types/types';

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

    //When a new player is added by the client, add it to the list of remaining players that can be added
    //to a game in progress.
    const addPlayerToRemainingListOfPlayers = (playerName: string) => {
        remainingPlayersInGame.value = [...remainingPlayersInGame.value, { player_name: playerName, isAddedToGame: false} ]
    }

    //When a player name is editted, reflect this change in either the list of players currently in a game,
    //or the list of players yet to be added to a game.
    // const editPlayerInActiveGame = () => {

    // }

    //Add the player to the game, and set their availibility to true.
    const addAvailablePlayerToGame = (playerIndex: number, playerName: string) => {
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
        
        remainingPlayersInGame.value.forEach(player => player.isAddedToGame = false)
    }
    
    return { 
        availablePlayersToAdd, 
        addPlayerToRemainingListOfPlayers,
        isGameCreated,
        currentPlayersInGame, 
        setAvailablePlayers, 
        setCurrentPlayers, 
        remainingPlayersInGame,
        currentLocation, 
        setCurrentLocation,
        toggleGameCreatedStatus,
        addAvailablePlayerToGame,
        removeAvailablePlayerFromGame
    }
}, {
    persist: true
})