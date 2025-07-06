import { defineStore } from 'pinia';
import { ref } from 'vue';
import { PlayerCard } from '../types/types';

export const HomePageStore = defineStore("HomePageStore", () => {
    const availablePlayersToAdd = ref<string[]>([])
    const currentPlayers = ref<PlayerCard[]>([])

    const setAvailablePlayers = (newPlayers: string[]) => {
        availablePlayersToAdd.value = newPlayers
    }

    const setCurrentPlayers = (newPlayers: PlayerCard[]) => {
        currentPlayers.value = newPlayers
    }
    
    return { availablePlayersToAdd, currentPlayers, setAvailablePlayers, setCurrentPlayers }
}, {
    persist: true
})