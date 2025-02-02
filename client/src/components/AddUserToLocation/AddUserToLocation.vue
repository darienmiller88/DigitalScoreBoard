<script setup lang="ts">
    import { ref } from 'vue';
    import { buttonActiveStore, ButtonState} from "../../stores/buttonActiveStore"
    import { scoreCardsStore } from "../../stores/scoreCardsStore"
    import { darkModeStore } from "../../stores/darkModeStore"
    import { storeToRefs } from 'pinia';
    import { Card } from "../../types/types"
    import { scoreBoardApi } from "../../api/api"
    import { selectedLocationStore } from '../../stores/selectedLocationStore';

    //Stateful variables
    const { currentButtonGroupState } = storeToRefs(buttonActiveStore())
    const { isDarkMode } = storeToRefs(darkModeStore())
    const { scoreCards } = storeToRefs(scoreCardsStore())
    const { selectedLocation } = storeToRefs(selectedLocationStore())

    //Stateful methods
    const { addScoreCard } = scoreCardsStore()

    const duplicateErrorMessage = ref<string>("")
    const username = ref<string>("")

    const addUser = async () => {
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

            try {
                const res = await scoreBoardApi.post(`/add-user-to-location/${selectedLocation.value}`, {"username": username.value})
                
                console.log("res", res.data)
            } catch (error) {
                console.log("err:", error)
            }

            username.value = ""
        }
    }
</script>

<template>
    <form v-if="currentButtonGroupState !== ButtonState.CREATE_NEW_TEAM_GAME" @submit.prevent="addUser">
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
            Add User To Location
        </button>
    </form>
</template>


<style scoped lang="scss">
    .dark-mode{
        border: 2px var(--primary-color) solid;
        color: var(--main-text-color);
        background-color: transparent;
    }

    .light-mode{
        color: var(--primary-color);
        background-color: var(--main-text-color)
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

        button{
            border-radius: 10px;
            padding: 15px 35px;

            transition: 0.5s;
            font-size: 25px;
            margin-top: 10px;   

            &:hover{
                cursor: pointer;
                background-color: var(--main-text-color-transparent);
            }

            &:active {
                transform: translateY(-5px);
            }
        }
    }

</style>