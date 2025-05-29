<script setup lang="ts">
    import { teamCardsStore } from "../../stores/teamCardsStore";
    import { optionsStore } from "../../stores/optionsStore";
    import { Team } from "../../types/types"
    import { storeToRefs } from "pinia";
    import { selectedTeamStore } from '../../stores/selectedLocationStore';
        
    const props = defineProps<{
      team: Team
      cardIndex: number
      pointValue: number
      openAddTeamPlayerModal: () => void
      openViewTeamPlayers: () => void
    }>()
    
    const { selectedTeam } = storeToRefs(selectedTeamStore())
    const { remainingLocationOptions } = storeToRefs(optionsStore())
    const { addOptionToRemainingLocationOptions } = optionsStore()
    const { addPoints, minusPoints, resetPoints, removeTeamCard } = teamCardsStore()

    const resetOptions = (index: number) => {
        removeTeamCard(index)
        addOptionToRemainingLocationOptions(props.team.team_name)

        //If after ALL the locations have been added to a team game, and one team card is removed, add
        //that name to "selectedLocationName" variable so the select tag isn't empty.
        if (remainingLocationOptions.value.length === 1) {
            selectedTeam.value.team_name = remainingLocationOptions.value[0]
        }
    }
</script>

<template>
    <div class="team-scorecard">
      <div class="team-name">{{ props.team.team_name }}</div>
  
      <!-- Team Actions -->
      <div class="team-actions">
        <button class="reset" @click="() => resetPoints(props.cardIndex)">Reset Points</button>
        <button class="clear" @click="() => resetOptions(props.cardIndex)">Clear Team</button>
      </div>
  
      <!-- Points -->
      <div class="points-control">
        <button class="minus" @click="() => minusPoints(props.cardIndex, props.pointValue)">-</button>
        <span>{{ props.team.score }}</span>
        <button class="plus" @click="() => addPoints(props.cardIndex, props.pointValue)">+</button>
      </div>
  
      <!-- Players List -->
      <div class="players-section">
        <div class="view-players">
          <button @click="openViewTeamPlayers">View Players</button>
        </div>
  
        <!-- Add Player -->
        <div class="add-player">
          <button @click="openAddTeamPlayerModal">Add Player To Team</button>
        </div>
      </div>
    </div>
</template>
  
<style scoped lang="scss">
  .team-scorecard {
    /* background-color: #10141e;
    */
    box-shadow: 8px 8px 5px rgba(173, 216, 230, 0.548);
    color: var(--main-text-color);
    
    text-align: center; 
    padding: 10px 10px;
    margin: 18px 0px; 
    transition: 0.3s;

    border: var(--main-text-color) 2px solid;
    border-radius: 10px;
    width: fit-content;

    // No margins on tablets and above to allow "gap" to work
    @media only screen and (min-width: 768px){
      margin: 0px 0px; 
    }

    &:hover{
        box-shadow: 10px 10px 15px var(--main-text-color);
        transform: translateY(-5px);
    }

    .team-name{
      font-size: 35px;
    }

    .team-actions{ 
      display: flex;
      justify-content: center;
      margin: 5px 0px;

      .clear{
        background-color: rgb(154, 154, 154);  
        
        &:hover{
          background-color: rgba(154, 154, 154, 0.7);  
        }
      }

      .reset{
        background-color: rgb(16, 103, 15);

        &:hover{
          background-color:  rgba(16, 103, 15, 0.8) 
        }
      }

      button{
        border: none;
        border-radius: 10px;
        color: aliceblue;
        font-weight: bold;
        font-size: 18px;

        @media only screen and (min-width: 992px){
          font-size: 15px;   
        }

        padding: 10px 15px;
        margin: 0px 5px;
        transition: 0.3s;

        &:hover{
          cursor: pointer;
        }
      }
    }

    .points-control {
      display: flex;
      justify-content: center;
      align-items: center;
      gap: 10px;
      margin: 25px 0;

      .minus{
        background-color: #ff0000;
      }

      span{
        color: var(--main-text-color);
        font-size: 20px;
      }

      .plus{
        background-color: var(--main-text-color);
      }
    
      button {
        color: black;
        border: none;
        padding: 10px 25px;
        border-radius: 5px;
        cursor: pointer;
      }
    }
  }

  .players-section {
    margin-top: 20px;

     .view-players{
        button{
          border: 2px solid var(--main-text-color);
          border-radius: 5px;
          padding: 10px 22px;

          background-color: transparent;
          color: var(--main-text-color);
          transition: 0.3s;

          &:hover{
            cursor: pointer;
            font-size: 17px;
          }
        }
      }
  
  }
  
  .add-player {
    margin-top: 15px;
    
    button {
      padding: 8px 20px;
      background-color: #00bfff;
      color: white;

      border: none;
      border-radius: 5px;

      font-weight: bold;
      font-size: 18px;
      cursor: pointer;

      @media only screen and (min-width: 992px){
        font-size: 15px;   
      }
    }
  }
  
</style>
  