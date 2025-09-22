<script setup lang="ts">
    import { darkModeStore } from "../../stores/darkModeStore"
    import { storeToRefs } from 'pinia';
    import { onMounted } from "vue";

    const store = darkModeStore()
    const { toggleDarkMode } = store
    const { isDarkMode } = storeToRefs(store)

    const addOrRemoveDarkModeClass = () => {
        toggleDarkMode()

        if (isDarkMode.value) {
            document.documentElement.setAttribute('data-theme', "dark");
        } else {
            document.documentElement.setAttribute('data-theme', "light");
        }        
    }

     // Run on page load to set the initial class based on the dark mode state
     onMounted(() => {
        if (isDarkMode.value) {
            document.documentElement.setAttribute('data-theme', "dark");
        } else {
            document.documentElement.setAttribute('data-theme', "light");
        }   
    });
</script>

<template>
    <div>
        <button :class="`mode-toggle ${isDarkMode ? '' : 'toggle-right'}`" @click="addOrRemoveDarkModeClass">
            <div class="slider"></div>
        </button>
    </div>
</template>

<style scoped lang="scss">
    $toggle-width: 45px;
    $padding:      5px;
    $slider-width: 15px;

    $toggle-width-large: $toggle-width * 3.5;
    $padding-large:      $padding      * 4.5;
    $slider-width-large: $slider-width * 2.5;

    .mode-toggle{  
        display: flex;

        background-color: var(--toggle-background);
        border-radius: 20px;
        border: none;

        width: $toggle-width;
        height: fit-content;
        padding: $padding $padding;
        margin-right: 10px;

        @media (min-width: 3840px) {
            width: $toggle-width-large;
            padding: ($padding-large) ($padding-large);
            margin-right: 30px;
            border-radius: 40px;
        }

        .slider{
            background-color: var(--bg-color);
            width: $slider-width;
            height: $slider-width;
            border-radius: 50%;   
            
            transition: 0.5s;
            position: relative;

            @media (min-width: 3840px) {
                width: $slider-width-large;
                height: $slider-width-large;
            }
        }

        &:hover{
            cursor: pointer;
        }
    }

    .toggle-right{
        .slider{
            transform: translateX(calc($toggle-width - $slider-width - ($padding * 2)));

            @media (min-width: 3840px) {
                transform: translateX(calc($toggle-width-large - $slider-width-large - ($padding-large * 2)));
            }
        }
    }
</style>