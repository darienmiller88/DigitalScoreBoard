<script setup lang="ts">
    import { Team } from "../../types/types"
    
    const props = defineProps<Team>()
</script>

<template>
    <div class="team-scorecard">
      <div class="team-name">{{ props.team_name }}</div>
  
      <!-- Team Actions -->
      <div class="team-actions">
        <button @click="">Reset Points</button>
        <button @click="">Clear Team</button>
      </div>
  
      <!-- Points -->
      <div class="points-control">
        <button @click="">-</button>
        <span>{{ props.score }}</span>
        <button @click="">+</button>
      </div>
  
      <!-- Players List -->
      <div class="players-section">
        <h3>Players</h3>
        <ul>
          <li>
            Darien
            <button @click="">Remove</button>
          </li>
        </ul>
  
        <!-- Add Player -->
        <div class="add-player">
          <input
            placeholder="Enter player name"
          />
          <button @click="">Add Player</button>
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
    // padding: 20px;
    transition: 0.3s;

    border: var(--main-text-color) 2px solid;
    border-radius: 10px;
    width: fit-content;

    &:hover{
        box-shadow: 10px 10px 15px var(--main-text-color);
        transform: translateY(-5px);
    }

    .team-name{
      font-size: 35px;
    }

    .team-actions{
      button{
        border: var(--main-text-color) 2px solid;
        border-radius: 10px;
        background-color: transparent;
        color: var(--main-text-color);

        padding: 10px 15px;
        margin: 0px 5px;

      }
    }

    .points-control {
      display: flex;
      justify-content: center;
      align-items: center;
      gap: 10px;
      margin: 15px 0;
    
      button {
        background-color: #ff0000;
        color: white;
        border: none;
        padding: 10px 20px;
        border-radius: 5px;
        cursor: pointer;
      }
    }


  }
  
  .players-section {
    margin-top: 20px;
  }
  
  .players-section ul {
    list-style: none;
    padding: 0;
  }
  
  .players-section li {
    margin: 5px 0;
  }
  
  .add-player {
    margin-top: 15px;
  }
  
  .add-player input {
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 5px;
    margin-right: 10px;
  }
  
  .add-player button {
    padding: 8px 15px;
    background-color: #00bfff;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
</style>
  