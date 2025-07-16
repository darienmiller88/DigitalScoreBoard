<script setup lang="ts">
    // import { storeToRefs } from 'pinia';
    import { scoreCardsStore } from "../../stores/scoreCardsStore"
    // import { selectedLocationStore } from '../../stores/selectedLocationStore';
    // import { scoreBoardApi } from '../../api/api';

    const { removeCard, addPoints, minusPoints, resetPoints } = scoreCardsStore()
    // const { selectedLocation } = storeToRefs(selectedLocationStore())
    defineProps<{ 
        cardIndex: number
        pointValue: number
        score: number
        username: string 
    }>()


    //testing
    // const removeCardFromLocation = async () => {
    //     removeCard(props.cardIndex)

    //     //Remove
    //     try {
    //         await scoreBoardApi.delete(`/remove-user-from-location/${selectedLocation.value}`, {data: {username: props.username}})
    //     } catch (error) {
    //         console.log("err:", error);
    //     }
    // }
</script>

<template>
     <div class="score-card">
        <div class="username">{{ username }}</div>
        <div class="divider"></div>
        <button class="remove" @click="() => removeCard(cardIndex)" >
            Remove User
        </button>
        <br>
        <button class="reset" @click="() => resetPoints(cardIndex)" >
            Reset points
        </button>
        <div class="score-wrapper">
            <button class="minus-points" @click="() => minusPoints(cardIndex, pointValue)">-</button>
            <div class="score">{{ score }}</div>
            <button class="add-points" @click="() => addPoints(cardIndex, pointValue)">+</button>
        </div>
    </div>
</template>

<style scoped lang="scss">
    .score-card{
        border: 2px solid var(--primary-color);
        border-radius: 10px;

        width: 90vw;
        box-shadow: 8px 8px 5px var(--primary-color-transparent);
        margin: 20px 0px; 

        transition: 0.3s;
        text-align: center;

        @media (min-width: 768px) {
            width: 40vw;
            margin: 20px 0px; 
        }

        @media (min-width: 1000px) {
            width: 40vw;
        }

        @media (min-width: 1025px) {
            width: 25vw;
        }

        &:hover{
            box-shadow: 8px 8px 5px var(--primary-color);
            transform: translateY(-5px);
        }

        .username{
            color: var(--primary-color);
        }
        
        .divider{
            width: 99%;
            border: 2px solid var(--primary-color);
            margin-top: 5px;
            margin-bottom: 20px;
        }

        .remove{
            border-radius: 10px;
            border: 2px solid var(--primary-color);

            padding: 10px 20px;
            color: var(--primary-color);
            background-color: transparent;

            transition: 0.5s;
            font-size: 15px;

            &:hover{
                cursor: pointer;
                background-color: var(--primary-color-transparent);
            }
        }

        .reset{
            margin: 10px 0px;
            padding: 7px 12px;

            background-color: var(--primary-color);
            color: var(--bg-color);
            font-weight: bold;

            border-radius: 50px;
            border: none;
            transition: 0.3s; 

            &:hover{
                cursor: pointer;
                background-color: var(--primary-color-transparent);
            }
        }

        .score-wrapper{
            display: grid;
            grid-template-columns: auto 160px auto;

            @media (min-width: 768px) {
                grid-template-columns: auto 50% auto;
            }

            // @media (min-width: 992px) {
            //     grid-template-columns: auto 170px auto;
            // }

            @media (min-width: 1025px) {
                grid-template-columns: auto 160px auto;
            }

            .score{
                display: flex;
                align-items: center;
                justify-content: center;
                color: var(--primary-color);

                font-size: 25px;
                padding: 0px 20px;
            }

            button{
                border-radius: 15px;
                border: 2px solid var(--bg-color);
                
                font-size: 30px;
                padding: 8px 25px;

                transition: 0.5s;

                @media (min-width: 768px) {
                    padding: 10px 30px;
                    font-size: 20px;
                }

                &:hover{
                    cursor: pointer;
                    color: aliceblue;
                    border: 2px solid aliceblue;
                }
            }

            .minus-points{
                background-color: red;
            }

            .add-points{
                background-color: var(--primary-color);
            }
        }

        .username{
            font-size: 35px;            
            margin: 5px 15px;
            overflow-x: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
        } 
    }
</style>