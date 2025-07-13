<script setup lang="ts">
    // import { scoreBoardApi } from '../../api/api';
    import Loading from '../Loading/Loading.vue';

    //Stateful variables
    // let isLoading = ref<boolean>(false)

    defineProps<{
        endAndSaveGame: () => void
        isLoading: boolean
    }>()

    // const addSavedGame = async () => {
    //     isLoading.value = true

    //     try {
    //         const savedGame: SavedGame = props
            
    //         console.log("game:", savedGame);
            
    //         // const res = await scoreBoardApi.post("/save-game", savedGame)
    //         // console.log("res: ", res.data)
    //     } catch (error) {
    //         console.log("err:", error);
    //     }

    //     isLoading.value = false
    // }
</script>

<template>
   <div class="save-wrapper">
        <button type="button" @click="endAndSaveGame" >
            <Loading :height="30" :usePrimary="false" v-if="isLoading"/>
            <div v-else> Save Game </div>
        </button>
    </div>
</template>

<style scoped lang="scss">
     .save-wrapper{
        text-align: center;
        margin: 40px;

        button{
            border-radius: 10px;
            border: none;
            padding: 15px 35px;
            background-color: var(--primary-color);
            color: aliceblue;
    
            transition: 0.5s;
            font-size: 25px;

            &:hover{
                cursor: pointer;
                background-color: var(--primary-color-transparent);
            }

            &:active {
                transform: translateY(-5px);
            }
        }
    }

    .save-wrapper button:active {
        transform: translateY(-5px);
    }
</style>