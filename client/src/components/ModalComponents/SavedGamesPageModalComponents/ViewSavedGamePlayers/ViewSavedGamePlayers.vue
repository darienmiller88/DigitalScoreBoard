<script setup lang="ts">
    import { PlayerCard } from '../../../../types/types';

    const props = defineProps<{
        playersInSavedGame: PlayerCard[]
    }>()
</script>

<template>
    <div class="players">
        <div class="player" v-for="player in props.playersInSavedGame">
            {{ player.username }} scored: <span> {{ player.score }} </span> 
        </div>
    </div>
</template>

<style scoped lang="scss">
    .players{
        height: 40vh;
        overflow-y: scroll;

        .player{
            text-align: center;
            transition: 0.3s;
            
            font-size: 24px;
            margin-bottom: 1rem;
            padding: 0.8rem 1rem;
    
            background: white;
            border-radius: 6px;
            box-shadow: 0 1px 3px rgba(5, 5, 5, 0.6);
    
            span{
                color: rgb(17, 175, 28);
            }
    
            &:hover{
                box-shadow: 0 3px 6px rgba(15, 15, 15, 0.9);
            }
        }
    }
    
</style>