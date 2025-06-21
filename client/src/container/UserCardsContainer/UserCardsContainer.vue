<script setup lang="ts">
    import UserCard from '../../components/UserCard/UserCard.vue';
    import Modal from '../../components/Modal/Modal.vue';
    import EditPlayerName from '../../components/EditPlayerName/EditPlayerName.vue';
    import { onMounted, ref, watch } from 'vue';
    import { scoreBoardApi } from "../../api/api"
    import { PlayerCard } from "../../types/types"
    import { optionsStore } from "../../stores/optionsStore"

    const { allLocationOptions } = optionsStore()
    let players = ref<string[]>([])
    let showEditPlayerNameModal = ref<boolean>(false)
    let playerNameToEdit = ref<string>("")
    
    const props = defineProps<{
        currentLocation: string
    }>()

    //Function for UserCard child to use edit the name currently on the card i.e. from mary rose to mary gold
    const addPlayerNameToEdit = (username: string) => {
        playerNameToEdit.value = username
        showEditPlayerNameModal.value = true
    }

    const removePlayer = async (playerIndex: number) => {
        players.value = players.value.filter((_, index) => {
            return playerIndex != index
        })

        try {
            await scoreBoardApi.delete(`/remove-user-from-location/${props.currentLocation}`, { data: { username: playerNameToEdit.value } })
        } catch (error) {
            console.log("err:", error);
        }   
    }

    //When the location is changed, retrieve the new list of players from that location.
    watch(() => props.currentLocation, async (newLocation) =>{
        try {            
            const playersResult = await scoreBoardApi.get<PlayerCard[]>(`/get-all-users/${newLocation}`) 
            
            players.value = playersResult.data.map(player => player.username)
        } catch (error) {
            console.log("err:", error);
        }
    })

    //On mount, get the first location from the list of all locations, and display the players from there.
    onMounted(async () => {
        try {            
            const playersResult = await scoreBoardApi.get<PlayerCard[]>(`/get-all-users/${allLocationOptions[0]}`) 
            
            players.value = playersResult.data.map(player => player.username)
        } catch (error) {
            console.log("err:", error);
        }
    })
</script>

<template>
    <div class="no-players" v-if="players.length === 0">No Players at {{ currentLocation }}.</div>
    <div class="people" v-else>Players at {{ currentLocation }}:</div>
    <div class="user-cards">
        <UserCard
            v-for="(player, index) in players"
            :key="index"
            :playerIndex="index"
            :playerName="player"
            :showModal="addPlayerNameToEdit"
            :removePlayer="removePlayer"
        />
    </div>

    <Modal 
        :modalHeader="'Edit Username'"
        :modalContent="EditPlayerName"
        :modalProps="{ 
            playerName: playerNameToEdit,
            hideModal: () => showEditPlayerNameModal = false
        }"
        :show="showEditPlayerNameModal"
        :onHide="() => showEditPlayerNameModal = false"
    />
</template>

<style scoped lang="scss">
    .no-players{
        text-align: center;
        font-size: 25px;
        margin: 25px;
        color: var(--primary-color);
    }

    .people{
        text-align: center;
        font-size: 30px;
        color: var(--primary-color);
        margin-top: 20px;
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
        @media only screen and (min-width: 1025px){
            display: grid;
            grid-template-columns: repeat(3, 1fr);
            gap: 10px 50px;
            max-height: 75vh;
            padding: 0px 15px;
        }

        // 2k monitors
        @media only screen and (min-width: 2560px){
            display: grid;
            grid-template-columns: repeat(3, 1fr);
            gap: 30px;
            max-height: 75vh;
        }
    }
</style>