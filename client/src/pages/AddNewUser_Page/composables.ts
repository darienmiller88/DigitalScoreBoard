import { ref } from "vue"

export const currentLocation = ref<string>("")
export let players = ref<string[]>([])

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

export const addNewPlayerToArray = (playerName: string) => {
    players.value = [...players.value, playerName]
}

export const editPlayerName = (playerIndex: number, newName: string) => {
    players.value = players.value.map((player, i) =>
        i === playerIndex ? newName : player
    )
}