<script setup lang="ts">
    import Loading from '../Loading/Loading.vue';

    import { ref } from 'vue';
    import { scoreBoardApi } from '../../api/api';
    import { AxiosError } from 'axios';

    const props = defineProps<{
        currentLocation: string
        playerName: string
        playerIndex: number
        players: string[]
        hideModal: () => void
        editPlayerName: (playerIndex: number, newName: string) => void
    }>()

    let editFirstName = ref<string>("")
    let editLastName = ref<string>("")
    let isLoading = ref<boolean>(false)
    let errorMessage = ref<string>("")
    let isErrorMessage = ref<boolean>(false)

    const onSubmit = async () => {        
        try {
            isLoading.value = true

            //Call the following route, and send in the following body.
            await scoreBoardApi.put(`/change-player-name/${props.currentLocation}`, {
                new_player_name: editFirstName.value.trim() + " " + editLastName.value.trim(),
                old_player_name: props.playerName
            })
            
            //Edit the player name on the front end so the user can see it after it's been changed in the backend
            props.editPlayerName(props.playerIndex, editFirstName.value + " " + editLastName.value)
            
            //After changing the player name, close the modal!
            props.hideModal()
        } catch (error) {
            if (error instanceof AxiosError) {
                errorMessage.value = error.response?.data
                isErrorMessage.value = true

                setTimeout(() => {
                   isErrorMessage.value = false 
                }, 4000)
            } else {
                console.log("Unknown error:", error);
            }
        }    

        isLoading.value = false
        editFirstName.value = ""
        editLastName.value = ""
    }
</script>

<template>
    <div class="content-title">
        Currently Editing player "{{ playerName }}"...
    </div>
    <form @submit.prevent="onSubmit">
        <div class="input-wrapper">
            <label>First Name</label><br />
            <input 
                class="form-element"
                id="edit-first-name"
                v-model="editFirstName"
                name="edit-first-name"
                :maxlength="15"
                type="text" 
                placeholder="Edit First Name"
                required
            />
        </div>

        <div class="input-wrapper">
            <label>Last Name</label><br />
            <input 
                class="form-element"
                id="edit-last-name"
                v-model="editLastName"
                name="edit-last-name"
                :maxlength="15"
                type="text" 
                placeholder="Edit Last Name"
                required
            />
        </div>
        <div class="button-wrapper">
            <button>
                <Loading v-if="isLoading" :height="20" :usePrimary="true"/>
                <span v-else> Submit Name Change</span>  
            </button>
        </div>
    </form>
    <div class="error-message" v-if="isErrorMessage">{{ errorMessage }}</div>
</template>

<style scoped lang="scss">
    .error-message{
        text-align: center;
        color: red;        
        margin: auto;
        padding: 10px;
        font-size: 18px;

        @media (min-width: 768px) {
            font-size: 20px;
        }
    }

    .content-title{
        text-align: center;
        font-weight: 600;
        font-size: 24px;
        margin: 20px;
    }

    form{
        display: flex;
        flex-direction: column;
        width: fit-content;
        margin: auto;
        
        .input-wrapper{
            margin: 5px;

            label{
                font-size: 20px;
            }
    
            input{
                padding: 10px;
                font-size: 20px;
                width: 75vw;
    
                @media (min-width: 768px) {
                    width: 60vw;
                }
    
                @media (min-width: 1024px) {
                    width: 25vw;
                }
            }
        }

        .button-wrapper{
            text-align: center;
            margin: 15px;
        
            button{
                width: 100%;
                background-color: dodgerblue;
                color: aliceblue;
                border: none;
                border-radius: 10px;
                font-size: 18px;
                padding: 12px;
                transition: 0.3s;

                &:hover{
                    cursor: pointer;
                    background-color: grey;
                }
            }
        }
    }
</style>