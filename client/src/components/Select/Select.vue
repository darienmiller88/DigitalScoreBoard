<script setup lang="ts">
    import { storeToRefs } from "pinia";
    import { darkModeStore } from "../../stores/darkModeStore"
    import { onMounted } from 'vue';

    const { isDarkMode } = storeToRefs(darkModeStore())
    
    const props = defineProps<{
        optionClicked: (event: Event) => void
        selectModel: string
        options: string[]
    }>()

    onMounted(() => {
        console.log("options:", props.options);
    })
</script>

<template>    
    <select 
        :value="selectModel"
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
        padding: 5px 8px;
        transition: 0.5s;
        border-radius: 10px;

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