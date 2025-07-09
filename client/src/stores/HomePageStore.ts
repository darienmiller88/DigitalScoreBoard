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