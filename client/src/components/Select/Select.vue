<script setup lang="ts">
    import { selectedLocationStore, selectedTeamLocationStore } from '../../stores/selectedLocationStore';
    import { storeToRefs } from "pinia";
    import { darkModeStore } from "../../stores/darkModeStore"
    import { computed, onMounted } from 'vue';
    import { ButtonState, buttonActiveStore } from "../../stores/buttonActiveStore"

    const { isDarkMode } = storeToRefs(darkModeStore())
    const { selectedLocationName } = storeToRefs(selectedLocationStore())
    const { selectedTeamLocation } = storeToRefs(selectedTeamLocationStore())
    const { currentButtonGroupState } = storeToRefs(buttonActiveStore())
    
    const props = defineProps<{
        optionClicked: (event: Event) => void
        options: string[]
    }>()

    const selectModel = computed(() => {
        return currentButtonGroupState.value == ButtonState.CREATE_NEW_TEAM_GAME ? selectedTeamLocation : selectedLocationName
    })

    onMounted(() => {
        console.log("selectModel: ", selectModel)
    })
</script>

<template>
     <select 
        v-model="selectedLocationName"
        name="locations" 
        :class="`${isDarkMode ? 'dark-mode-select' : 'light-mode-select'}`" 
        @change="props.optionClicked"
    >
        <option v-for="(option, index) in props.options" :value="option" :key="index+1">
            {{ option }}
        </option>
    </select>   
</template>

<style scoped lang="scss">
     select {
        font-size: 18px;
        padding: 2px 5px;
        transition: 0.5s;

        @media screen and (min-width: 768px) {
            font-size: 28px;
        }
    }

    .dark-mode-select{
        background-color: var(--main-bg-color);
        color: var(--primary-color);
        border: 2px solid var(--primary-color);
    }

    .light-mode-select{
        background-color: var(--primary-color);
        color: var(--main-bg-color);
        border: 2px solid var(--main-bg-color);
    }
</style>