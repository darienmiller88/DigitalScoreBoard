<script setup lang="ts">
    import { ref, watchEffect } from 'vue';

    let selectedPlayerName = ref<string>("")
    
    const props = defineProps<{
        teamPlayersAvailableToAdd: string[]
        currentTeam: string[]
        locationName: string
        addPlayerToTeam: (player: string) => void
    }>()

    watchEffect(() => {
        if (props.teamPlayersAvailableToAdd.length > 0) {
            selectedPlayerName.value = props.teamPlayersAvailableToAdd[0];
        }
    })
</script>

<template>
    <h2 v-if="props.teamPlayersAvailableToAdd.length === 0" class="no-users">No users at {{ props.locationName }} left to add. </h2>
    <div v-else>
        <h2 class="players-location">Players in {{ props.locationName }} avaible to add:</h2>
        <div class="select-wrapper">
            <select v-model="selectedPlayerName" name="team-players" id="team-players">
                <option v-for="(player, index) in props.teamPlayersAvailableToAdd" :value="player" :key="index">
                    {{ player }}
                </option>
            </select>
            <button @click="() => props.addPlayerToTeam(selectedPlayerName)">Add Player</button>
        </div>
    </div>
</template>

<style scoped lang="scss">
    .no-users{
        text-align: center;
    }

    .select-wrapper{
        text-align: center;

        select{
            font-size: 22px;
            padding: 2px 5px;
            transition: 0.5s;
        }

        button{
            margin: 0px 10px;
            padding: 10px 18px;
            font-size: 18px;
            transition: 0.3s;
            
            border: none;
            border-radius: 10px;
            color: aliceblue;
            background-color: dodgerblue;

            &:hover{
                cursor: pointer;
                background-color: rgba(30, 144, 255, 0.7);
            }
        }
    }

    .players-location{
        text-align: center;
    }

    select{
        margin: auto;
    }
</style>