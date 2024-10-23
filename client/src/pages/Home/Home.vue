<script setup lang="ts">
    import ScoreCard from '../../components/ScoreCard/ScoreCard.vue';
    import { Card } from "../../types/types"
    import { onMounted, ref } from 'vue';
    import { darkModeStore } from "../../stores/darkModeStore"
    import { scoreCardsStore } from "../../stores/scoreCardsStore"
    import { selectedLocationStore } from '../../stores/selectedLocationStore';
    import { storeToRefs } from 'pinia';

    const username = ref<string>("")
    const options = ref<string[]>([
        "Pelham", 
        "Lawrence", 
        "Elmwood", 
        "Flushing", 
        "Grand Concourse", 
        "Port Richmond",
        "W 154th St",
        "West End",
    ])

    //Stateful methods
    const { addScoreCard } = scoreCardsStore()
    const { setSelectedLocation } = selectedLocationStore()

    //Stateful variables
    const { isDarkMode } = storeToRefs(darkModeStore())
    const { scoreCards } = storeToRefs(scoreCardsStore())
    const { selectedLocation } = storeToRefs(selectedLocationStore())

    const duplicateErrorMessage = ref<string>("")

    const addUser = () => {
        const newCard: Card = {
            username: username.value,
            score: 0
        }

        if (scoreCards.value.some(card => card.username == newCard.username)) {
            duplicateErrorMessage.value = `${newCard.username} already exists! Please select another username.`
            console.log("duplicate user:", duplicateErrorMessage);
            
            setTimeout(() => {
                duplicateErrorMessage.value = ""
            }, 3000);
        }else{
            addScoreCard(newCard)
            username.value = ""
        }
    }

    const optionClicked = (event: Event) => {
        const selectedValue = (event.target as HTMLSelectElement).value;

        setSelectedLocation(selectedValue)
    }

    onMounted(() => {
        selectedLocation.value = selectedLocation.value == "" ? options.value[0] : selectedLocation.value
    })
</script>

<template>
    <div class="title">Digital Score Board</div>
    <div :class="`location ${isDarkMode ? 'dark-mode-text' : 'light-mode-text'}`">Current Location: 
        <select 
            v-model="selectedLocation"
            name="locations" 
            id="locations" 
            :class="`${isDarkMode ? 'dark-mode-select' : 'light-mode-select'}`" 
            @change="optionClicked"
        >
            <option v-for="(option, index) in options" :value="option" :key="index">
                {{ option }}
            </option>
        </select>    
    </div>

    <form @submit.prevent="addUser">
        <div class="add-user-wrapper">
            <label for="add-user">Name</label><br>
            <input 
                :class="`form-element ${isDarkMode ? 'dark-mode-text' : 'light-mode-text'}`"
                id="add-user" 
                v-model="username" 
                minlength="1"
                maxlength="16" 
                type="text" 
                name="addUser" 
                placeholder="Add user to game" 
                required
            >
            <div class="error">{{ duplicateErrorMessage }}</div>
        </div>
        <button :class="`form-element ${isDarkMode ? 'dark-mode' : 'light-mode'}`" type="submit">
            Add User To List
        </button>
    </form>

    <div id="user-cards-id" class="user-cards">
        <ScoreCard 
            v-for="(card, index) in scoreCards" 
            :key="card.username" 
            :username="card.username" 
            :card-index="index"
            :point-value="100"
            :score="card.score"
        />
    </div>
    <!-- <div class="save-wrapper">
        <button type="button" :class="`${isDarkMode ? 'dark-mode' : 'light-mode'}`">
            Save Game
        </button>
    </div> -->
</template>

<style scoped lang="scss">
    .dark-mode{
        border: 2px var(--primary-color) solid;
        color: var(--main-text-color);
        background-color: transparent;
    }

    .light-mode{
        // border: 2px solid var(--main-text-color);
        color: var(--primary-color);
        background-color: var(--main-text-color)
    }

    .dark-mode-text{
        color: var(--primary-color);
    }

    .light-mode-text{
        color: var(--main-bg-color);
    }

    .title{
        color: var(--main-text-color);
        text-align: center;
        font-size: 40px;
        margin: 40px 0px;

        @media screen and (min-width: 768px) {
            font-size: 60px;
            margin: 10px 0px;
        }
    }

    .location{
        text-align: center;
        font-size: 30px;
        transition: 0.5s;

        .underline{
            text-decoration: underline;
        }
    }

    select {
        font-size: 18px;
        padding: 2px 5px;
        transition: 0.5s;

        @media screen and (min-width: 768px) {
            font-size: 28px;
        }
    }

    .dark-mode-select{
        background-color: var(--main-bg-color);
        color: var(--primary-color);
        border: 2px solid var(--primary-color);
    }

    .light-mode-select{
        background-color: var(--primary-color);
        color: var(--main-bg-color);
        border: 2px solid var(--main-bg-color);
    }


    form{
        text-align: center;
        margin-top: 20px;

        .form-element{
            width: 85vw;

            @media screen and (min-width: 768px) {
                width: 60vw;
            }
        }

        button{
            margin-top: 10px;
        }
    }

    .error{
        text-align: center;
        color: red;
        font-size: 25px;
        // margin: 15px 0px;
    }

    .user-cards{
        display: grid;
        grid-template-columns: 1fr;    

        // border: 2px solid white; 
        width: fit-content; 
        max-width: 90vw;
        max-height: 50vh;
        overflow-y: scroll;

        margin: auto;
        margin-top: 20px;
        padding: 0px 10px;

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

            background-color: transparent;
        }
    }

    .save-wrapper{
        text-align: center;
        margin-top: 20px;
    }

    .save-wrapper button, form button{
        border-radius: 10px;

        padding: 15px 35px;

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