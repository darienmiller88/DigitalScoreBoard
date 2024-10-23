<script setup lang="ts">
    import { darkModeStore } from '../../stores/darkModeStore';
    import { storeToRefs } from 'pinia';
    import { scoreCardsStore } from "../../stores/scoreCardsStore"

    const { isDarkMode } = storeToRefs(darkModeStore())
    const { removeCard, addPoints, minusPoints } = scoreCardsStore()

    defineProps<{ 
        cardIndex: number
        pointValue: number
        score: number
        username: string 
    }>()

</script>

<template>
     <div class="user-card">
        <div :class="`username ${isDarkMode ? 'dark-mode' : 'light-mode'}`">{{ username }}</div>
        <div class="divider"></div>
        <button :class="`remove ${isDarkMode ? 'dark-mode-button' : 'light-mode-button'}`"  @click="() => removeCard(cardIndex)">Remove</button>
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
        width: fit-content;
        box-shadow: 8px 8px 5px rgba(173, 216, 230, 0.548);
        transition: 0.3s;
        text-align: center;
        margin: 20px 0px; 

        @media screen and (min-width: 768px) {
            // margin-bottom/: 0px;
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

        .score-wrapper{
            display: grid;
            grid-template-columns: auto 100px auto;
        
            margin-top: 20px;

            .score{
                display: flex;
                align-items: center;
                justify-content: center;

                font-size: 20px;
                padding: 0px 20px;
            }

            button{
                border-radius: 15px;
                border: 2px solid var(--main-bg-color);
                
                font-size: 30px;
                padding: 12px 45px;

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
            margin: 0px 10px;
        } 
    }
</style>