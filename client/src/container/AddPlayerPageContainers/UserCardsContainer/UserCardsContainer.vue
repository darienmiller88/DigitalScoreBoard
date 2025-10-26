<script setup lang="ts">
    //Components
    import UserCard from '../../../components/UserCard/UserCard.vue';
    import Modal from '../../../components/Modal/Modal.vue';
    import EditPlayerName from '../../../components/ModalComponents/AddPlayerPageModalComponents/EditPlayerName/EditPlayerName.vue';
    import Loading from '../../../components/Loading/Loading.vue';
    
    //Data/Data manipulation
    import { onMounted, ref, watch } from 'vue';
    import { scoreBoardApi } from "../../../api/api"
    import { PlayerCard } from "../../../types/types"
    import { optionsStore } from "../../../stores/optionsStore"
    import { AddPlayerPageStore } from '../../../stores/PageStores/AddPlayerPageStore';
    import { useToast } from "vue-toastification";
    import { storeToRefs } from 'pinia';

    const { currentLocation, players } = storeToRefs(AddPlayerPageStore())
    const { removePlayerFromArray, editPlayerName, setPlayers } = AddPlayerPageStore()

    let showEditPlayerNameModal = ref<boolean>(false)
    let playerNameToEdit = ref<string>("")
    let indexOfPlayerBeingEditted = ref<number>(0)
    let isLoading = ref<boolean>(true)
    
    const { allLocationOptions } = optionsStore()
    const toast = useToast()

    //Function for UserCard child to use edit the name currently on the card i.e. from mary rose to mary gold
    const addPlayerNameToEdit = (playerName: string, playerIndex: number) => {
        playerNameToEdit.value = playerName
        indexOfPlayerBeingEditted.value = playerIndex
        showEditPlayerNameModal.value = true
    }

    //Remove a player from a given ADAPT location.
    const removePlayerFromDB = async (playerIndex: number) => {
        const playerToRemove: string = players.value[playerIndex]

        //First, remove the player from the front end.
        removePlayerFromArray(playerIndex)

        try {
            //Afterwards, remove them from the back end
            await scoreBoardApi.delete(`/remove-user-from-location/${ currentLocation.value }`, { data: { player_name: playerToRemove } })
            
            toast.success(`${playerToRemove} removed!`, {
                timeout: 1000
            })
        } catch (error) {
            console.log("err:", error)
            toast.error(`Error: ${error}`, {
                timeout: 2000
            })
        }   
    }

    //Retrieve all players from a given ADAPT location.
    const getPlayers = async (locationName: string) => {
        try {            
            const playersResult = await scoreBoardApi.get<PlayerCard[]>(`/get-all-users/${locationName}`) 
            
            setPlayers(playersResult.data.map(player => player.username))
        } catch (error) {
            console.log("err:", error);
        }
    }
 
    //When the location is changed, retrieve the new list of players from that location.
    watch(() => currentLocation.value, async (newLocation) => {
        //When getting the new list of users, first display the loading spinner
        isLoading.value = true        

        //Get all of the players
        await getPlayers(newLocation)

        //Then remove the loading spinner
        isLoading.value = false
    })

    //On mount, get the first location from the list of all locations, and display the players from there.
    onMounted(async () => {
        await getPlayers(allLocationOptions[0])

        isLoading.value = false
    })
</script>

<template>
    <div v-if="isLoading" class="loading-wrapper">
        <Loading :height="100" :usePrimary="true"/>
    </div>
    <div v-else>
        <div class="no-players" v-if="players.length === 0">No Players at {{ currentLocation }}.</div>
        <div class="people" v-else>Players at {{ currentLocation }}:</div>
        <div class="user-cards">
            <UserCard
                v-for="(player, index) in players"
                :key="index"
                :playerIndex="index"
                :playerName="player"
                :showModalAndSetCurrentPlayer="addPlayerNameToEdit"
                :removePlayer="removePlayerFromDB"
            />
        </div> 
    </div>

    <Modal 
        :modalHeader="'Edit Username'"
        :modalContent="EditPlayerName"
        :modalProps="{ 
            playerName: playerNameToEdit,
            playerIndex: indexOfPlayerBeingEditted,
            hideModal: () => showEditPlayerNameModal = false,
            editPlayerName: editPlayerName,
            players: players,
            currentLocation: currentLocation
        }"
        :show="showEditPlayerNameModal"
        :onHide="() => showEditPlayerNameModal = false"
    />
</template>

<style scoped lang="scss">
    .loading-wrapper{
        margin: 40px;    
    }

    .no-players{
        text-align: center;
        font-size: 25px;
        margin: 25px;
        color: var(--primary-color);
    }

    .people{
        text-align: center;
        color: var(--primary-color);
        margin-top: 20px;
        font-size: 20px;
        
        @media (min-width: 768px) {
            font-size: 30px;
        }
    }

    .user-cards{
        display: grid;
        grid-template-columns: 1fr;    

        width: fit-content; 
        max-height: 50vh;
        max-width: 90vw;

        margin: auto;
        margin-top: 20px;
        margin-bottom: 20px;
        overflow-y: scroll;
        transition: 0.3s;

        &:hover{
            box-shadow: 0 4px 10px var(--primary-color);
        }

        // Show two columns of cards for tablets
        @media (min-width: 768px) and (max-width: 1024px){
            display: grid;
            grid-template-columns: repeat(2, 1fr);
            gap: 30px;
            max-height: 45vh;
            padding: 0px 20px;
        }

        //Show three columns for laptops less than 2K
        @media (min-width: 1025px){
            display: grid;
            grid-template-columns: repeat(3, 1fr);
            gap: 10px 50px;
            max-height: 75vh;
            padding: 0px 15px;
        }

        // 2k monitors
        @media (min-width: 2560px){
            display: grid;
            grid-template-columns: repeat(4, 1fr);
            gap: 30px;
            max-height: 75vh;
        }
    }
</style>