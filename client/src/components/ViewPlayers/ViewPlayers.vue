<script setup lang="ts">
    const props = defineProps<{
        teamPlayers: string[]
    }>()
</script>

<template>
    <h2 v-if="!teamPlayers.length">No players added to team yet!</h2>
    <div v-else class="players">
        <div class="player" v-for="(player, index) in props.teamPlayers" :key="index">
            {{ player }}
        </div>
    </div>
</template>

<style scoped lang="scss">
    h2{
        text-align: center;
    }

    .players{
        text-align: center;
        // border: 2px solid red;
        
        max-height: 25vh;
        overflow-y: scroll;

        .player{
            margin: 10px;
            font-size: 18px;

            @media screen and (min-width: 992px) {
                font-size: 26px;
            }
        }
    }
</style>