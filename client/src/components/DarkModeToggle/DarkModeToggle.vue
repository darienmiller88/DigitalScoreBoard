<script setup lang="ts">
    import { darkModeStore } from "../../stores/darkModeStore"
    import { storeToRefs } from 'pinia';

    const store = darkModeStore()
    const { toggleDarkMode } = store
    const { isDarkMode } = storeToRefs(store)
</script>

<template>
    <div>
        <button :class="`mode-toggle ${isDarkMode ? '' : 'toggle-right'}`" @click="toggleDarkMode">
            <div class="slider"></div>
        </button>
    </div>
</template>

<style scoped lang="scss">
    $toggle-width: 45px;
    $padding:      5px;
    $slider-width: 15px;

    .mode-toggle{  
        display: flex;

        background-color: var(--primary-color);
        border-radius: 20px;
        border: none;

        width: $toggle-width;
        height: fit-content;
        padding: $padding $padding;
        margin-right: 10px;

        .slider{
            background-color: var(--main-bg-color);
            width: $slider-width;
            height: $slider-width;
            border-radius: 50%;   
            
            transition: 0.5s;
            position: relative;
        }

        &:hover{
            cursor: pointer;
        }
    }

    .toggle-right{
        background-color: var(--main-bg-color);

        .slider{
            background-color: var(--primary-color);
            transform: translateX(calc($toggle-width - $slider-width - ($padding * 2)));
        }
    }
</style>