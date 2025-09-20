<script setup lang="ts">
    import { scoreCardsStore } from "../../stores/scoreCardsStore"
    import { HomePageStore } from "../../stores/PageStores/HomePageStore";
    import { ref } from "vue";

    const { removeCard, addPoints, minusPoints, resetPoints } = scoreCardsStore()
    const { removeAvailablePlayerFromGame } = HomePageStore()
    const maxPoints: number = 99999
    let pointsToAdd = ref<number>(0)
    let isExceedMaxPointsError = ref<boolean>(false)
    
    const props = defineProps<{ 
        cardIndex: number
        pointValue: number
        score: number
        username: string 
    }>()

    const removeCardFromPlayerList = () => {
        removeAvailablePlayerFromGame(props.cardIndex, props.username)
        removeCard(props.cardIndex)
    }

    const addPointsToTotal = () => {
        if (pointsToAdd.value + props.score > maxPoints) {
            isExceedMaxPointsError.value = true

            setTimeout(() => {
                isExceedMaxPointsError.value = false
            }, 2000);
        } else {
            addPoints(props.cardIndex, pointsToAdd.value)
            pointsToAdd.value = 0
        }
    }

    const subtractPointsFromTotal = () => {
        if (pointsToAdd.value + props.score > maxPoints) {
            isExceedMaxPointsError.value = true

            setTimeout(() => {
                isExceedMaxPointsError.value = false
            }, 2000);
        } else {
            minusPoints(props.cardIndex, pointsToAdd.value)
            pointsToAdd.value = 0
        }
    }
</script>

<template>
     <div class="score-card">
        <div class="username">{{ username }}</div>
        <div class="divider"></div>
        <button class="remove" @click="() => removeCardFromPlayerList()" >
            Remove Player
        </button>
        <div class="input-wrapper">
            <span @click="subtractPointsFromTotal">-</span>
            <input type="number" v-model.number="pointsToAdd" placeholder="points" min="1" >
            <span @click="addPointsToTotal">+</span>
        </div>
        <div class="error" v-if="isExceedMaxPointsError">Total must be {{ maxPoints }} or lower</div> 
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

            @media (min-width: 3840px) {
                font-size: 50px;
                padding: 15px 25px;
                margin: 10px;
            }
        }

        .input-wrapper{
            display: flex;
            justify-content: center;
            margin: 10px 0px;
            
            span{
                font-size: 26px;
                margin: 0px 10px;
                color: var(--toggle-background);

                &:hover{
                    cursor: pointer;
                }

                @media (min-width: 3840px){
                    font-size: 65px;
                    margin: 0px 30px;
                }
            }

            input{
                background-color: var(--input-background-color);
                border: 2px solid var(--input-border);
                color: var(--input-color);

                transition: 0.3s;
                width: 40%;
                padding: 10px;
                border-radius: 10px;
                font-size: 16px;

                @media (min-width: 3840px){
                    font-size: 55px;
                }
            }
        }

        .error{
            color: red;
            margin: 10px
        }

        .reset{
            // margin: 10px 0px;
            padding: 7px 12px;

            background-color: var(--primary-color);
            color: var(--bg-color);
            font-weight: bold;

            border-radius: 50px;
            border: none;
            transition: 0.3s; 

            @media (min-width: 768px) {
                padding: 10px 30px;
                font-size: 20px;
            }
            
            @media (min-width: 3840px) {
                margin: 20px 0px;
                padding: 20px 50px;
                font-size: 55px;
            }

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

            @media (min-width: 3840px) {
                grid-template-columns: auto 480px auto;
            }

            .score{
                display: flex;
                align-items: center;
                justify-content: center;
                color: var(--primary-color);

                font-size: 25px;
                padding: 0px 20px;

                @media (min-width: 3840px){
                    font-size: 75px;
                }
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

                @media (min-width: 3840px) {
                    padding: 10px 30px;
                    font-size: 80px;
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

           @media (min-width: 3840px){
                font-size: 80px;
            }
        } 
    }
</style>