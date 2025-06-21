import { ref } from "vue"

export const currentLocation = ref<string>("")

export const setLocation = (location: string) => {
    currentLocation.value = location
}