<script setup lang="ts">
    import { ref } from 'vue';
    import { useToast } from 'vue-toastification';
    import { SavedGame } from '../../../../types/types';
    import { scoreBoardApi } from '../../../../api/api';    
    import Loading from '../../../Loading/Loading.vue';

    let isLoading = ref<boolean>(false)
    const toast = useToast()
    const props = defineProps<{
        gameIndex: number
        game: SavedGame
        hideModal: () => void
        removeGameFromArray: (gameIndex: number) => void
    }>()

    const deleteGame = async () => {
        try {
            props.removeGameFromArray(props.gameIndex)

            const res = await scoreBoardApi.delete(`/delete-save-game/${props.game.id}`)

            console.log("res:", res.data);
            props.hideModal()
            toast.success("Game successfully deleted!", { timeout: 2000 })
        } catch (error) {
            console.log(error)
        }
    }
</script>

<template>
    <div class="warning">
        Are you sure you want to delete this game?
    </div>
    <div class="button-wrapper">
        <button @click="deleteGame">
            <Loading v-if="isLoading" :usePrimary="false" :height="30"/>
            <span v-else>Delete Game</span>
        </button>
    </div>
</template>

<style scoped lang="scss">
    .warning{
        text-align: center;
        font-size: 22px;
    }

    .button-wrapper{
        text-align: center;
        margin: 10px;

        button{
            padding: 10px 15px;
            background-color: red;
            font-size: 18px;
            color: white;
            border: none;
            border-radius: 10px;
            transition: 0.3s;
    
            &:hover{
                cursor: pointer;
                transform: scale(1.1);
            }
        }
    }
</style>