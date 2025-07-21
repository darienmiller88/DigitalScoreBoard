<script setup lang="ts">
    //libraries
    import { ref, onMounted } from 'vue'
    import { storeToRefs } from 'pinia';
    import { scoreBoardApi } from "../../../api/api"
    import { useToast } from "vue-toastification";
    import { optionsStore } from "../../../stores/optionsStore"

    //Components
    import Loading from '../../../components/Loading/Loading.vue';
    import Select from '../../../components/Select/Select.vue'

    //Stateful variables
    const { allLocationOptions } = storeToRefs(optionsStore())

    let isLoading = ref<boolean>(false)    
    let isDuplicatePlayer = ref<boolean>(false)
    const duplicateErrorMessage = ref<string>("")
    const firstName = ref<string>("")
    const lastName = ref<string>("")
    const toast = useToast()

    //Props
    const props = defineProps<{
        currentLocation: string
        changeLocation: (location: string) => void
        players: string[]
        addNewPlayerToArray: (playerName: string) => void
    }>()

    const addUser = async () => {
        const newPlayer: string = firstName.value.trim() + " " + lastName.value.trim()
        
        isLoading.value = true
        if (props.players.some(player => player.toLowerCase() === newPlayer.toLowerCase())) {
            duplicateErrorMessage.value = `"${newPlayer}" already exists! Please select another name.`
            isDuplicatePlayer.value = true

            setTimeout(() => {
                duplicateErrorMessage.value = ""
                isDuplicatePlayer.value = false
            }, 3000)
        }else{
            props.addNewPlayerToArray(newPlayer)

            try {
                const res = await scoreBoardApi.post(`/add-user-to-location/${props.currentLocation}`, { player_name: newPlayer })
                
                toast.success(`${newPlayer} successfully added!`, {
                    timeout: 2000
                })
                console.log("res", res.data)
            } catch (error) {
                console.log("err:", error)
                toast.success(`Error occurred: ${error}`, {
                    timeout: 2000
                })
            }

            firstName.value = ""
            lastName.value = ""
        }

        isLoading.value = false
    }

    const onChangeSelect = (event: Event) => {
        props.changeLocation((event.target as HTMLSelectElement).value)
    }

    // when this component is mounted, load the current location with the first location so the select tag isn't blank
    onMounted(() => {
        props.changeLocation(allLocationOptions.value[0])
    })
</script>

<template>
    <div class="select-wrapper">
        <Select 
            :options="allLocationOptions"
            :selectModel="currentLocation"
            :onChange="onChangeSelect"
        />
    </div>

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
        </div>
        
        <button class="form-element" type="submit" >
            <Loading :height="35" :usePrimary="false" v-if="isLoading" />
            <div v-else> Add Player To Location </div>
        </button>
        <div class="error" v-if="isDuplicatePlayer">{{ duplicateErrorMessage }}</div>
    </form>
</template>

<style scoped lang="scss">
    .error{
        text-align: center;
        color: red;
        font-size: 25px;
        font-weight: 600;
        margin: 20px;
    }

    .select-wrapper{
        text-align: center;
        
        @media (min-width: 786px) {
            margin: 20px;
        }
    }

    form{
        text-align: center;
        margin-top: 20px;

        .form-element{
            width: 75vw;

            @media screen and (min-width: 768px) {
                width: 50vw;
            }

            @media screen and (min-width: 992px) {
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
                padding: 15px;

                border: 2px solid var(--primary-color);
                border-radius: 5px;

                background-color: transparent;
                color: var(--primary-color);

                &:focus{
                    border: 2px solid var(--primary-color);
                    outline: none;
                }
            }
        }

        button{
            border: 2px solid var(--primary-color);
            border-radius: 10px;
            padding: 15px 10px;
            margin-top: 15px;   

            transition: 0.3s;
            font-size: 20px;
            font-weight: 600;

            background-color: var(--primary-color);
            color: var(--bg-color);
            
            @media (min-width: 768px) {
                font-size: 28px;
                padding: 10px 14px;
            }

            @media (min-width: 992px) {
                font-size: 25px;
                padding: 15px 35px;
            }

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