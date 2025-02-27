import { defineStore } from 'pinia';
import { ref } from 'vue';

export const optionsStore = defineStore("optionsStore", () => {
    const allLocationOptions = ref<string[]>([])
    const remainingLocationOptions = ref<string[]>([])

    //Overides the Location options with a set of new ones.
    const setAllLocationOptions = (newOptions: string[]) => {
        allLocationOptions.value = newOptions
    }

    //Overides the Location options with a set of new ones.
    const setRemainingLocationOptions = (newOptions: string[]) => {
        remainingLocationOptions.value = newOptions
    }

    //add an option to the list of Location options.
    const addOptionToRemainingLocationOptions = (option: string) => {
        remainingLocationOptions.value.push(option)
    }
   
    return { 
        allLocationOptions, 
        remainingLocationOptions, 
        setAllLocationOptions, 
        setRemainingLocationOptions,
        addOptionToRemainingLocationOptions
    }
})