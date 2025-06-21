<script setup lang="ts">
    import UserCard from '../../components/UserCard/UserCard.vue';
    import Modal from '../../components/Modal/Modal.vue';
    import EditPlayerName from '../../components/EditPlayerName/EditPlayerName.vue';
    import { ref } from 'vue';
    import { scoreBoardApi } from "../../api/api"

    let usernames = [
        "darien miller",
        "angela lourens",
        "mark rubens",
        "shaniqua blank",
        "Daniel Negroni",
        "Luis Hernandez",
        "Victoria Inuhauzo",
        "superdeeduperlongexmaple ofaververylongname"
    ]

    let showEditPlayerNameModal = ref<boolean>(false)
    let usernameToEdit = ref<string>("")
    
    const props = defineProps<{
        currentLocation: string
    }>()

    const addUserNameToEdit = (username: string) => {
        usernameToEdit.value = username
        showEditPlayerNameModal.value = true
    }

    const removePlayer = async (playerIndex: number) => {
        usernames = usernames.filter((_, index) => {
            return playerIndex != index
        })

        try {
            await scoreBoardApi.delete(`/remove-user-from-location/${props.currentLocation}`, { data: { username: usernameToEdit.value } })
        } catch (error) {
            console.log("err:", error);
        }
        
    }
</script>

<template>
    <div class="people">People at {{ currentLocation }}:</div>
    <div class="user-cards">
        <UserCard
            v-for="(username, index) in usernames"
            :key="index"
            :playerIndex="index"
            :username="username"
            :showModal="addUserNameToEdit"
            :removePlayer="removePlayer"
        />
    </div>

    <Modal 
        :modalHeader="'Edit Username'"
        :modalContent="EditPlayerName"
        :modalProps="{ 
            username: usernameToEdit,
            hideModal: () => showEditPlayerNameModal = false
        }"
        :show="showEditPlayerNameModal"
        :onHide="() => showEditPlayerNameModal = false"
    />
</template>

<style scoped lang="scss">
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
            padding: 0px 10px;
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