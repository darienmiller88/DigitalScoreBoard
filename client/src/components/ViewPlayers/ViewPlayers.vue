<script setup lang="ts">
    const props = defineProps<{
        teamPlayers: string[]
        removePlayerFromTeam: (player: string) => void
    }>()
</script>

<!-- This component will be sent to the modal when the user clicks on "view players" on the team card -->
<template>
    <h2 v-if="!teamPlayers.length">No players added to team yet!</h2>
    <div v-else class="players">
        <div class="player" v-for="(player, index) in props.teamPlayers" :key="index">
            {{ player }}
            <button @click="() => props.removePlayerFromTeam(player)">Remove</button>
        </div>
    </div>
</template>

<style scoped lang="scss">
    h2{
        text-align: center;
    }

    .players{
        text-align: center;        
        max-height: 25vh;
        overflow-y: scroll;

        @media screen and (min-width: 992px) {
            max-height: 35vh;
        }

        .player{
            display: flex;
            align-items: center;
            justify-content: space-between;

            margin: 5px 0;
            padding: 10px;
            font-size: 16px;
            font-weight: bold;
            
            color: #333;
            background: #f8f9fa;
            
            border-radius: 8px;
            transition: all 0.3s ease-in-out;

            &:hover{
                background: #e9ecef;
            }

            @media screen and (min-width: 992px) {
                font-size: 26px;
            }
        }

        button{         
            background: rgb(220, 53, 69);
            color: white;
            
            border: none;
            border-radius: 6px;
            padding: 6px 12px;
            font-size: 16px;

            cursor: pointer;
            transition: 0.3s;

            &:hover{
                transform: scale(1.1);
            }
        }
    }
</style>