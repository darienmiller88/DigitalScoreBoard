<script setup lang="ts">
    import Loading from '../Loading/Loading.vue';

    defineProps<{
        endAndSaveGame: () => void
        isLoading: boolean
    }>()
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
    button{
        border-radius: 10px;
        border: none;
        padding: 15px 35px;
        background-color: var(--primary-color);
        color: aliceblue;

        transition: 0.5s;
        font-size: 25px;
        font-weight: 600;
        width: 65vw;

        @media (min-width: 768px) {
            width: unset;
        }

        @media (min-width: 3840px) {
            font-size: 60px;
        }

        &:hover{
            cursor: pointer;
            background-color: var(--primary-color-transparent);
        }

        &:active {
            transform: translateY(-5px);
        }
    }

    // .save-wrapper button:active {
    //     transform: translateY(-5px);
    // }
</style>