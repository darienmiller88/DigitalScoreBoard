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
        font-weight: 600;
        padding: 5px 8px;
        transition: 0.5s;
        border-radius: 10px;

        background-color: var(--primary-color);
        color: var(--bg-color);

        @media screen and (min-width: 768px) {
            font-size: 28px;
        }
    }
</style>