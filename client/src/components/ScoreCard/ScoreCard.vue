<script setup lang="ts">
    import { darkModeStore } from '../../stores/darkModeStore';
    import { storeToRefs } from 'pinia';
    import { scoreCardsStore } from "../../stores/scoreCardsStore"
    import { selectedLocationStore } from '../../stores/selectedLocationStore';
    import { scoreBoardApi } from '../../api/api';

    const { isDarkMode } = storeToRefs(darkModeStore())
    const { removeCard, addPoints, minusPoints, resetPoints } = scoreCardsStore()
    const { selectedLocation } = storeToRefs(selectedLocationStore())
    const props = defineProps<{ 
        cardIndex: number
        pointValue: number
        score: number
        username: string 
    }>()

    const removeCardFromLocation = async () => {
        removeCard(props.cardIndex)

        try {
            await scoreBoardApi.delete(`/remove-user-from-location/${selectedLocation.value}`, {data: {username: props.username}})
        } catch (error) {
            console.log("err:", error);
        }
    }
</script>

<template>
     <div class="user-card">
        <div :class="`username ${isDarkMode ? 'dark-mode' : 'light-mode'}`">{{ username }}</div>
        <div class="divider"></div>
        <button 
            :class="`remove ${isDarkMode ? 'dark-mode-button' : 'light-mode-button'}`" 
            @click="removeCardFromLocation()"
        >Remove User</button>
        <br>
        <button class="reset" @click="() => resetPoints(cardIndex)" >
            Reset points
        </button>
        <div class="score-wrapper">
            <button class="minus-points" @click="() => minusPoints(cardIndex, pointValue)">-</button>
            <div :class="`score ${isDarkMode ? 'dark-mode' : 'light-mode'}`">{{ score }}</div>
            <button class="add-points" @click="() => addPoints(cardIndex, pointValue)">+</button>
        </div>
    </div>
</template>

<style scoped lang="scss">
    .dark-mode{
        transition: 0.5s;
        color: var(--main-text-color);
    }

    .light-mode{
        transition: 0.5s;
        color: var(--main-text-color);
    }

    .dark-mode-button{
        color: var(--main-text-color);
        border: 2px solid var(--primary-color);
    }

    .light-mode-button{
        color: var(--main-text-color);
        border: 2px solid var(--main-text-color);
    }

    .user-card{
        border: 2px solid var(--main-text-color);
        border-radius: 10px;

        width: 90vw;
        box-shadow: 8px 8px 5px rgba(173, 216, 230, 0.548);
        margin: 20px 0px; 

        transition: 0.3s;
        text-align: center;

        @media screen and (min-width: 768px) {
            // margin-bottom/: 0px;
            width: 40vw;
            margin: 20px 0px; 
        }

        @media screen and (min-width: 992px) {
            width: 25vw;
            margin: 20px 0px; 
        }

        @media screen and (min-width: 1400px) {
            width: 15vw;
            // height: 30vh;
            margin: 20px 0px; 
        }

        &:hover{
            box-shadow: 10px 10px 15px var(--main-text-color);
            transform: translateY(-5px);
        }
        
        .divider{
            width: 99%;
            border: 2px solid var(--main-text-color);
            margin-top: 5px;
            margin-bottom: 20px;
        }

        .remove{
            border-radius: 10px;

            padding: 10px 20px;
            color: var(--main-text-color);
            background-color: transparent;

            transition: 0.5s;
            font-size: 15px;

            &:hover{
                cursor: pointer;
                background-color: var(--main-text-color-transparent);
            }
        }

        .reset{
            margin: 10px 0px;
            padding: 5px 10px;

            background-color: var(--main-text-color);
            color: var(--primary-color);
            font-weight: bold;

            border-radius: 50px;
            border: none;
            transition: 0.3s; 

            &:hover{
                cursor: pointer;
                background-color: var(--main-text-color-transparent);
            }
        }

        .score-wrapper{
            display: grid;
            grid-template-columns: auto 160px auto;
            // margin-top: 20px;

            @media screen and (min-width: 768px) {
                grid-template-columns: auto 150px auto;
            }

            @media screen and (min-width: 992px) {
                grid-template-columns: auto 170px auto;
            }

            @media screen and (min-width: 1400px) {
                grid-template-columns: auto 200px auto;
            }

            .score{
                display: flex;
                align-items: center;
                justify-content: center;

                font-size: 25px;
                padding: 0px 20px;
            }

            button{
                border-radius: 15px;
                border: 2px solid var(--main-bg-color);
                
                font-size: 30px;
                padding: 8px 25px;

                transition: 0.5s;

                @media screen and (min-width: 768px) {
                    padding: 10px 30px;
                    font-size: 20px;
                }

                &:hover{
                    cursor: pointer;
                    color: var(--primary-color);
                    border: 2px solid var(--primary-color);
                }
            }

            .minus-points{
                background-color: red;
            }

            .add-points{
                background-color: var(--main-text-color);
            }
        }

        .username{
            font-size: 35px;            
            margin: 5px 15px;
            // border: 3px dotted aliceblue;
            overflow-x: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
        } 
    }
</style>