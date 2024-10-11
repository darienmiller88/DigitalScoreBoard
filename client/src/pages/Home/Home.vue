<script setup lang="ts">
    import ScoreCard from '../../components/ScoreCard/ScoreCard.vue';
    import { Card } from "../../types/types"
    import { ref } from 'vue';

    const cards = ref<Card[]>([
        {
            username: "Darien",
            score: 100
        }, 
        {
            username: "Vicky",
            score: 500
        },
        
      
    ])

    const username = ref<string>("")

    const addPoints = (index: number, amountToAdd: number) => {
        cards.value[index].score += amountToAdd
    }

    const minusPoints = (index: number, amountToAdd: number) => {
        cards.value[index].score -= amountToAdd
    }

    const addUser = () => {
        const newCard: Card = {
            username: username.value,
            score: 0
        }

        cards.value = [...cards.value, newCard]
    }
</script>

<template>
    <div class="title">Digital Score Board</div>
    <form @submit.prevent="addUser">
        <div class="add-user-wrapper">
            <label for="add-user">Name</label><br>
            <input class="form-element" id="add-user" v-model="username" minlength="1" maxlength="12" type="text" name="addUser" placeholder="Add user to game" required>
        </div>
        <button class="form-element" type="submit">
            Add User To List
        </button>
    </form>

    <div id="user-cards-id" class="user-cards">
        <ScoreCard 
            v-for="(card, index) in cards" 
            :key="card.username" 
            :username="card.username" 
            :card-index="index"
            :point-value="100"
            :score="card.score" 
            :add-points="addPoints"
            :minus-points="minusPoints"
        />
    </div>

    <div class="save-wrapper">
        <button type="button" data-toggle="modal" data-target="#exampleModalCenter">
            Save Game
        </button>
    </div>
</template>

<style scoped lang="scss">
    .title{
        color: var(--main-text-color);
        text-align: center;
        font-size: 60px;
        margin: 10px 0px;
    }

    form{
        text-align: center;
        margin-top: 20px;

        .form-element{
            width: 60vw;
        }

        button{
            margin-top: 10px;
        }
    }

    .user-cards{
        display: grid;
        grid-template-columns: 1fr;    

        // border: 2px solid white; 
        width: fit-content; 
        max-width: 85vw;
        margin: auto;
        margin-top: 20px;

        @media only screen and (min-width: 768px){
            display: grid;
            grid-template-columns: repeat(4, 1fr);
            gap: 30px;
        }
    }


    .add-user-wrapper{
        color: var(--main-text-color);
        text-align: left;
        width: fit-content;

        margin: auto;

        label {
            margin: 0;
            font-size: 30px;
        }

        input{
            font-size: 25px;
            padding: 10px;

            border: 2px solid var(--main-text-color);
            border-radius: 5px;

            background-color: var(--main-bg-color);
            color: var(--primary-color);
        }
    }

    .save-wrapper{
        text-align: center;
        margin-top: 20px;
    }

    .save-wrapper button, form button{
        border: 2px var(--primary-color) solid;
        border-radius: 10px;

        padding: 15px 35px;
        color: var(--main-text-color);
        background-color: var(--main-bg-color);

        transition: 0.5s;
        font-size: 25px;
    }

    .save-wrapper button:hover, form button:hover{
        cursor: pointer;
        background-color: var(--main-text-color-transparent);
    }

    .save-wrapper button:active, form button:active{
        transform: translateY(-5px);
    }
</style>