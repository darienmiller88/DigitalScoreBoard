<script setup lang="ts">
    import { useWindowSize } from "@vueuse/core"

    const { width } = useWindowSize();
    const availablePlayers = ["darien miller niller nuiller", "vicky", "sean", "kash", "darien", "vicky", "sean", "kash"]
</script>

<template>
    <div class="available-players">
        <div v-for="player in availablePlayers" class="available-player">
            <span>{{ player }}</span>
            <button v-if="width > 768">Add To Game</button>
            <button v-else>Add</button>
        </div>
    </div>
</template>

<style scoped lang="scss">
    .available-players{
        display: grid;
        // flex-direction: column;
        // align-items: center;
        // border: 2px solid black;
        width: fit-content;
        margin: auto;
        margin-bottom: 20px;
        max-height: 20vh;
        overflow-y: scroll;

        .available-player{
            display: grid;
            grid-template-columns: 75% auto;
            border: 2px salmon solid;
            width: 75vw;
            
            span{
                border: 2px solid red;
                // margin-right: 10px;
                padding: 8px;
                white-space: nowrap;
                font-size: 24px;
                overflow-x: scroll;
            }

            button{
                border: 2px solid green;
                font-size: 24px;   
            }
        }
    }
</style>