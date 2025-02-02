<script setup lang="ts">
    import { storeToRefs } from "pinia";
    import { darkModeStore } from "../../stores/darkModeStore"
    import { buttonActiveStore, ButtonState } from "../../stores/buttonActiveStore"
    import { useWindowSize } from "@vueuse/core";

    const { width } = useWindowSize();

    //Stateful variables
    const { isDarkMode } = storeToRefs(darkModeStore())
    const{ currentButtonGroupState } = storeToRefs(buttonActiveStore())

    //Stateful methods
    const { setButtonActive } = buttonActiveStore()
</script>

<template>
    <div class="button-group">
        <button 
            :class="`
                ${isDarkMode ? 'dark-mode-button-group' : 'light-mode-button-group'}
                ${currentButtonGroupState == ButtonState.CREATE_NEW_GAME ? 'active' : ''}
            `" 
            @click="setButtonActive(ButtonState.CREATE_NEW_GAME)"
        >Create new game</button>
        <span v-if="width >= 768" :class="`${isDarkMode ? 'dark-mode-span' : 'light-mode-span'}`"></span>
        <button 
            :class="`
                ${isDarkMode ? 'dark-mode-button-group' : 'light-mode-button-group'}
                ${currentButtonGroupState == ButtonState.ADD_NEW_USER ? 'active' : ''}
            `"
            @click="setButtonActive(ButtonState.ADD_NEW_USER)"
        >Add new user</button>
        <span v-if="width >= 768" :class="`${isDarkMode ? 'dark-mode-span' : 'light-mode-span'}`"></span>
        <button 
            :class="`
                ${isDarkMode ? 'dark-mode-button-group' : 'light-mode-button-group'}
                ${currentButtonGroupState == ButtonState.CREATE_NEW_TEAM_GAME ? 'active' : ''}
            `"
            @click="setButtonActive(ButtonState.CREATE_NEW_TEAM_GAME)"
        >Create new team game</button>
    </div>
</template>

<style scoped lang="scss"> 
    .button-group{
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: stretch;
        
        padding: 0px 15px;
        margin-top: 30px;

        @media screen and (min-width: 768px) {
            flex-direction: row;
            padding: 0px;
        }

        span{
            margin: 0px 15px;
            transition: 0.3s;
        }

        .dark-mode-span{
            border: 1px solid aliceblue;
        }

        .light-mode-span{
            border: 1px solid black;
        }

        .active{
            background-color: dodgerblue;
        }

        button{
            background-color: black;
            color: aliceblue;
            border: none;
            border-radius: 8px;

            padding: 15px 28px;
            margin-bottom: 18px;

            font-size: 20px;
            transition: 0.3s;

            box-shadow: 0px 4px 8px rgba(100, 100, 100, 0.25), /* Bottom shadow */
                0px -4px 8px rgba(100, 100, 100, 0.25), /* Top shadow */
                4px 0px 8px rgba(100, 100, 100, 0.25), /* Right shadow */
                -4px 0px 8px rgba(100, 100, 100, 0.25);

            @media screen and (min-width: 768px) {
                margin-bottom: 0px;

                &:hover{
                    cursor: pointer;
                    padding: 15px 38px;
                    font-size: 28px;
                }
            }

            &:active{
                transform: translateY(-5px);
            }
        }
    }
</style>