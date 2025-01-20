import { defineStore } from 'pinia';
import { ref } from 'vue';

export const buttonActiveStore = defineStore("button-active", () => {
    const currentButtonGroupState = ref<boolean[]>([false, true, false])

    const setButtonActive = (index: number) => {
        for (let i = 0; i < currentButtonGroupState.value.length; i++) {
            currentButtonGroupState.value[i] = false
        }

        currentButtonGroupState.value[index] = true
    }

    return { currentButtonGroupState, setButtonActive }
}, {
    persist: true
})