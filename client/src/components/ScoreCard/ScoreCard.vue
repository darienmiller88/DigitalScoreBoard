<script setup lang="ts">

defineProps<{ 
    cardIndex: number
    pointValue: number
    score: number
    username: string 
    addPoints: (index: number, pointsToAdd: number) => void
    minusPoints: (index: number, pointsToSubtract: number) => void
    removeUser: (index: number) => void
}>()

</script>

<template>
     <div class="user-card">
        <div class="username">{{ username }}</div>
        <div class="divider"></div>
        <div class="remove">
            <button @click="() => removeUser(cardIndex)">Remove</button>
        </div>
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

        .remove button{
            border: 2px var(--primary-color) solid;
            border-radius: 10px;

            padding: 10px 20px;
            color: var(--main-text-color);
            background-color: var(--main-bg-color);

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

                color: var(--primary-color);
                font-size: 20px;
                padding: 0px 20px;
            }

            button{
                border-radius: 15px;
                border: 2px solid var(--main-bg-color);
                
                font-size: 20px;
                padding: 10px 30px;

                transition: 0.5s;

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
            color: var(--primary-color);
            
            margin: 0px 10px;
        } 
    }
</style>