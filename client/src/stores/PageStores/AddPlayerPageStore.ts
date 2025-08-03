import { defineStore } from 'pinia';
import { ref } from 'vue';

export const AddPlayerPageStore = defineStore("AddPlayerPageStore", () => {
    const currentLocation = ref<string>("")
    const players = ref<string[]>([])

    const setLocation = (location: string) => {
        currentLocation.value = location
    }
    
    const setPlayers = (newPlayers: string[]) => {
        players.value = newPlayers
    }
    
    const removePlayerFromArray = (playerIndex: number) => { 
        players.value = players.value.filter((_, index) => {
            return playerIndex != index
        })
    }
    
    const addNewPlayerToArray = (playerName: string) => {
        players.value = [...players.value, playerName]
    }

    const editPlayerName = (playerIndex: number, newName: string) => {
        players.value = players.value.map((player, i) =>
            i === playerIndex ? newName : player
        )
    }

    return{
        currentLocation,
        players,
        setLocation, 
        setPlayers,
        removePlayerFromArray,
        addNewPlayerToArray,
        editPlayerName
    }
}, {
    persist: true
})