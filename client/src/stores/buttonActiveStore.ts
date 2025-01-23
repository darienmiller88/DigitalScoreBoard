import { defineStore } from 'pinia';
import { ref } from 'vue';

export enum ButtonState {
    CREATE_NEW_GAME = 0,
    ADD_NEW_USER = 1,
    CREATE_NEW_TEAM_GAME = 2,
}

export const buttonActiveStore = defineStore("button-active", () => {
    const currentButtonGroupState = ref<ButtonState>(ButtonState.ADD_NEW_USER)

    const setButtonActive = (state: ButtonState) => {
        currentButtonGroupState.value = state
    }

    return { setButtonActive, currentButtonGroupState }
}, {
    persist: true
})