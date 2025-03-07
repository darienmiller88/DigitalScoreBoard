<script setup lang="ts">
    import { ref, watchEffect } from 'vue';

    let selectedPlayerName = ref<string>("heloo")
    
    const props = defineProps<{
        teamPlayers: string[]
        locationName: string
    }>()

    watchEffect(() => {
        if (props.teamPlayers.length > 0) {
            console.log("AddNewPlayer props:", props.teamPlayers, "and", props.locationName);
            selectedPlayerName.value = props.teamPlayers[0];
        }
    })

    watchEffect(() => {
      
    })
</script>

<template>
    <h2 class="players-location">Players in {{ props.locationName }} avaible to add:</h2>
    <div class="select-wrapper">
        <select v-model="selectedPlayerName" name="team-players" id="team-players">
            <option v-for="(player, index) in props.teamPlayers" :value="player" :key="index">
                {{ player }}
            </option>
        </select>
        <button>Add Player</button>
    </div>
</template>

<style scoped lang="scss">
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