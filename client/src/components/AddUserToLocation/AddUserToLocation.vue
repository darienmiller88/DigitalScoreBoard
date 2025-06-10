<script setup lang="ts">
    import { ref, onMounted } from 'vue'
    import { scoreCardsStore } from "../../stores/scoreCardsStore"
    import { storeToRefs } from 'pinia';
    import { Player } from "../../types/types"
    import { scoreBoardApi } from "../../api/api"
    import { selectedLocationStore } from '../../stores/selectedLocationStore';
    import { optionsStore } from "../../stores/optionsStore"
    import Loading from '../Loading/Loading.vue';
    import Select from '../Select/Select.vue';

    //Stateful variables
    const { scoreCards } = storeToRefs(scoreCardsStore())
    // const { selectedLocation } = storeToRefs(selectedLocationStore())
    const { allLocationOptions } = storeToRefs(optionsStore())

    //Stateful methods
    const { addScoreCard } = scoreCardsStore()

    const duplicateErrorMessage = ref<string>("")
    const firstName = ref<string>("")
    const lastName = ref<string>("")
    const locationModel = ref<string>("")
    let isLoading = ref<boolean>(false)

    const addUser = async () => {
        isLoading.value = true

        const newPlayer: Player = {
            username: firstName.value + " " + lastName.value,
            score: 0
        }

        if (scoreCards.value.some(card => card.username == newPlayer.username)) {
            duplicateErrorMessage.value = `${newCard.username} already exists! Please select another username.`
            console.log("duplicate user:", duplicateErrorMessage);
            
            setTimeout(() => {
                duplicateErrorMessage.value = ""
            }, 3000);
        }else{
            addScoreCard(newCard)

            try {
                const res = await scoreBoardApi.post(`/add-user-to-location/${locationModel}`, {"username": newCard.username})
                
                console.log("res", res.data)
            } catch (error) {
                console.log("err:", error)
            }

            firstName.value = ""
            lastName.value = ""
        }

        isLoading.value = false
    }

    onMounted(() => {
        // locationModel
    })
</script>

<template>
    <Select 
        :options="allLocationOptions"
        :selectModel="locationModel"
    />

    <form @submit.prevent="addUser">
        <div class="add-user-wrapper">
            <label for="add-firstname">First Name</label><br>
            <input 
                class="form-element"
                id="add-first-name" 
                v-model="firstName" 
                minlength="1"
                maxlength="15" 
                type="text" 
                name="addFirstName" 
                placeholder="Player's First Name " 
                required
            ><br><br>
            <label for="add-user">Last Name</label><br>
            <input 
                class="form-element"
                id="add-last-name" 
                v-model="lastName" 
                minlength="1"
                maxlength="15" 
                type="text" 
                name="addLastName" 
                placeholder="Player's Last Name" 
                required
            >
            <div class="error">{{ duplicateErrorMessage }}</div>
        </div>

        <button class="form-element" type="submit">
            <Loading :height="35" :usePrimary="false" v-if="isLoading" />
            <div v-else> Add Player To Location </div>
        </button>
    </form>
</template>

<style scoped lang="scss">
    form{
        text-align: center;
        margin-top: 20px;

        .form-element{
            width: 75vw;

            @media screen and (min-width: 768px) {
                width: 40vw;
            }
        }

        .add-user-wrapper{
            color: var(--primary-color);
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

                border: 2px solid var(--primary-color);
                border-radius: 5px;

                background-color: transparent;
            }

            .error{
                text-align: center;
                color: red;
                font-size: 25px;
            }
        }

        button{
            border-radius: 10px;
            padding: 15px 35px;

            transition: 0.3s;
            font-size: 25px;
            font-weight: 600;
            margin-top: 10px;   
            background-color: var(--primary-color);
            color: var(--bg-color);

            border: 2px solid var(--primary-color);

            &:hover{
                cursor: pointer;
                background-color: var(--primary-color-transparent);
            }

            &:active {
                transform: translateY(-5px);
            }
        }
    }
</style>