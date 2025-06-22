import { ref } from "vue"

export const currentLocation = ref<string>("")
export const players = ref<string[]>([])

export const setLocation = (location: string) => {
    currentLocation.value = location
}

export const setPlayers = (newPlayers: string[]) => {
    players.value = newPlayers
}

export const removePlayerFromArray = (playerIndex: number) => { 
    players.value = players.value.filter((_, index) => {
        return playerIndex != index
    })
}