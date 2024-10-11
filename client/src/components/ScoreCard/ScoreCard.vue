<script setup lang="ts">

defineProps<{ 
    cardIndex: number
    pointValue: number
    score: number
    username: string 
    addPoints: (index: number, pointsToAdd: number) => void
    minusPoints: (index: number, pointsToSubtract: number) => void
}>()

</script>

<template>
     <div class="user-card">
        <div class="username">{{ username }}</div>
        <div class="divider"></div>
        <div class="score-wrapper">
            <button class="minus-points" @click="() => minusPoints(cardIndex, pointValue)">-</button>
            <div class="score">{{ score }}</div>
            <button class="add-points" @click="() => addPoints(cardIndex, pointValue)">+</button>
        </div>
    </div>
</template>


<style scoped lang="scss">
    .user-card{
        border: 2px solid var(--main-text-color);
        border-radius: 10px;
        width: fit-content;

        text-align: center;
        
        .divider{
            width: 99%;
            border: 2px solid var(--main-text-color);
            margin-top: 5px;
            margin-bottom: 20px;
        }

        .score-wrapper{
            display: flex;
            align-items: center;
        
            margin-top: 20px;

            .score{
                color: var(--primary-color);
                font-size: 20px;
                margin: 0px 10px;
            }

            button{
                padding: 10px 30px;
                border-radius: 15px;
                border: 2px solid var(--main-bg-color);

                transition: 0.5s;

                &:hover{
                    cursor: pointer;
                    color: var(--primary-color);
                    border: 2px solid var(--primary-color);
                }
            }

            .minus-points{
                margin-right: 15px;
                background-color: red;
            }

            .add-points{
                margin-left: 15px;
                background-color: var(--main-text-color);
            }
        }


        .username{
            font-size: 35px;
            color: var(--primary-color);
            
            margin: 0px 10px;
            /* border: 2px solid wheat; */
        } 
    }
</style>