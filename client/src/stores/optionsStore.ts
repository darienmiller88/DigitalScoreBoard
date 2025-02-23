import { defineStore } from 'pinia';
import { ref } from 'vue';

export const optionsStore = defineStore("optionsStore", () => {
    const options = ref<string[]>([])

    //Overides the Location options with a set of new one's.
    const setOptions = (newOptions: string[]) => {
        options.value = newOptions
    }

    //add an option to the list of Location options.
    const addOption = (option: string) => {
        options.value.push(option)
    }
   
    return { options, setOptions, addOption }
})