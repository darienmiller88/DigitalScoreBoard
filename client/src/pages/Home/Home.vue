<script setup lang="ts">
    import ScoreCard from '../../components/ScoreCard/ScoreCard.vue';
    import { Card } from "../../types/types"
    import { onMounted, ref } from 'vue';
    import { darkModeStore } from "../../stores/darkModeStore"
    import { scoreCardsStore } from "../../stores/scoreCardsStore"
    import { selectedLocationStore } from '../../stores/selectedLocationStore';
    // import { locationsStore } from "../../stores/locationsStore"
    import { storeToRefs } from 'pinia';
    import { scoreBoardApi } from "../../api/api"
    import { Location, SavedGame } from "../../types/types"
    import { Icon } from "@iconify/vue"
    import { useWindowSize } from "@vueuse/core"

    const { width } = useWindowSize();

    const username = ref<string>("")
    const options = ref<string[]>([])
    const isLoading = ref<boolean>(true)
    let locations: Location[] = []

    //Stateful methods
    const { addScoreCard, setCards, resetAllPoints, getWinner } = scoreCardsStore()
    const { setSelectedLocation } = selectedLocationStore()
    // const { addLocation } = locationsStore()

    //Stateful variables
    const { isDarkMode } = storeToRefs(darkModeStore())
    const { scoreCards } = storeToRefs(scoreCardsStore())
    const { selectedLocation } = storeToRefs(selectedLocationStore())
    // const { locationsFromLocalStorage } = storeToRefs(locationsStore())

    const duplicateErrorMessage = ref<string>("")

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

    const optionClicked = (event: Event) => {
        const selectedValue = (event.target as HTMLSelectElement).value;

        locations.forEach(location => {
            if(location.location_name === selectedValue){
                setCards(location.users)                    
            }
        })

        setSelectedLocation(selectedValue)
    }

    const addSavedGame = async () => {
        try {
            const savedGame: SavedGame = {
                id: "",
                winner: getWinner(),
                location: {
                    id: "",
                    location_name: selectedLocation.value,
                    users: scoreCards.value
                },
                total_points: 0,
                average_points: 0.0,
                created_at: new Date().toLocaleDateString()
            }

            const res = await scoreBoardApi.post("/save-game", savedGame)

            console.log("res: ", res.data)
        } catch (error) {
            console.log("err:", error);
        }
    }

    onMounted(async () => {
        // if (locationsFromLocalStorage.value.length) {
        //     locations = locationsFromLocalStorage.value
        //     options.value = locations.map(location => {
        //         return location.location_name
        //     })   
        // } else {
        // }
        try {
            const locationsResponse = await scoreBoardApi.get<Location[]>("/get-all-locations")
    
            locations = locationsResponse.data
            options.value = locations.map(location => {
                //Cache the locations by adding them to local storage.
                // addLocation(location)
                return location.location_name
            })     

            if (!selectedLocation.value) {
                setSelectedLocation(options.value[0])
            }

            //If there are no current score cards set in local storage, retrieve them from the database
            // locationsResponse.data.forEach((location, i) => {
            //     if (location.location_name === selectedLocation.value) {
            //         setCards(locationsResponse.data[i].users)
            //     }
            // })
            // if (!scoreCards) {
            // } else{
            //     console.log("card:", scoreCards.value);
                
            // }         
        } catch (error) {
            console.log("err:", error);
        }
        
        console.log("locations:", locations, "option:", options, "host:", window.location.hostname);
        console.log("api:", scoreBoardApi.getUri());
        
        isLoading.value = false
    })
</script>

<template>
    <div class="title">Digital Score Board</div>
    <div :class="`location ${isDarkMode ? 'dark-mode-text' : 'light-mode-text'}`">Current Location: 
        <Icon icon="svg-spinners:180-ring" v-if="isLoading"/>
        <select 
            v-else
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
    <div class="button-group">
        <button :class="`${isDarkMode ? 'dark-mode-button-group' : 'light-mode-button-group'}`" >Create new game</button>
        <span v-if="width >= 768" :class="`${isDarkMode ? 'dark-mode-span' : 'light-mode-span'}`"></span>
        <button :class="`${isDarkMode ? 'dark-mode-button-group' : 'light-mode-button-group'}`">Add new user</button>
        <span v-if="width >= 768" :class="`${isDarkMode ? 'dark-mode-span' : 'light-mode-span'}`"></span>
        <button :class="`${isDarkMode ? 'dark-mode-button-group' : 'light-mode-button-group'}`">Create new team game</button>
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
            Add User To Location
        </button>
    </form>

    <div class="reset-all-points-wrapper">
        <button type="button" @click="resetAllPoints" >
            Reset All Points
        </button>
    </div>

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
    <div class="save-wrapper">
        <button type="button" @click="addSavedGame" :class="`${isDarkMode ? 'dark-mode' : 'light-mode'}`">
            Save Game
        </button>
    </div>
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

    .button-group{
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: stretch;
        
        padding: 0px 15px;
        margin-top: 30px;

        @media screen and (min-width: 768px) {
            flex-direction: row;
        }

        span{
            margin: 0px 15px;
            transition: 0.3s;
        }

        .dark-mode-span{
            border: 1px solid aliceblue;
        }

        .light-mode-span{
            border: 1px solid black;
        }

        .active{
            background-color: dodgerblue;
        }

        button{
            background-color: black;
            color: aliceblue;
            border: none;
            border-radius: 8px;

            padding: 15px 28px;
            margin-bottom: 18px;

            font-size: 20px;
            transition: 0.3s;

            box-shadow: 0px 4px 8px rgba(100, 100, 100, 0.25), /* Bottom shadow */
                0px -4px 8px rgba(100, 100, 100, 0.25), /* Top shadow */
                4px 0px 8px rgba(100, 100, 100, 0.25), /* Right shadow */
                -4px 0px 8px rgba(100, 100, 100, 0.25);

            @media screen and (min-width: 768px) {
                margin-bottom: 0px;

                &:hover{
                    cursor: pointer;
                    padding: 15px 38px;
                    font-size: 28px;
                    background-color: rgba(30, 144, 255, 0.7);
                }
            }
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
    }

    .reset-all-points-wrapper{
        text-align: center;
        margin: 15px;

        button{
            border-radius: 10px;
            border: none;

            padding: 10px 18px;
            font-size: 16px;
            font-weight: 700;
            transition: 0.3s;
            
            background-color: red;
            color: aliceblue;

            &:hover{
                cursor: pointer;
                background-color: rgba($color: #e30909, $alpha: .8);
            }
            
            &:active{
                transform: translateY(-5px);
            }
        }
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
        padding: 0px 10px;

        @media only screen and (min-width: 768px){
            display: grid;
            grid-template-columns: repeat(4, 1fr);
            gap: 30px;
            max-height: 75vh;
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
        margin: 40px;
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

    .save-wrapper button:active, form button:active {
        transform: translateY(-5px);
    }
</style>