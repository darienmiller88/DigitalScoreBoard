import { defineStore } from 'pinia';
import { ref } from 'vue';

export const buttonActiveStore = defineStore("button-active", () => {
    const currentButtonGroupState = ref<boolean[]>([false, true, false])

    

    return {  }
}, {
    persist: true
})