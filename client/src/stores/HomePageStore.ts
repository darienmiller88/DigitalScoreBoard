import { defineStore } from 'pinia';
import { ref } from 'vue';
import { PlayerCard } from '../types/types';

export const HomePageStore = defineStore("HomePageStore", () => {
    const availablePlayersToAdd = ref<string[]>([])
    const currentPlayersInGame = ref<PlayerCard[]>([])
    const currentLocation = ref<string>("")

    const setAvailablePlayers = (newPlayers: string[]) => {
        availablePlayersToAdd.value = newPlayers
    }

    const setCurrentPlayers = (newPlayers: PlayerCard[]) => {
        currentPlayersInGame.value = newPlayers
    }

    const setCurrentLocation = (newLocation: string) => {
        currentLocation.value = newLocation
    }
    
    return { 
        availablePlayersToAdd, 
        currentPlayersInGame, 
        setAvailablePlayers, 
        setCurrentPlayers, 
        currentLocation, 
        setCurrentLocation
    }
}, {
    persist: true
})