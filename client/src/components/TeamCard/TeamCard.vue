<script setup lang="ts">
    import { Team } from "../../types/types"
    // import Modal from "../Modal/Modal.vue";
    // import AddNewPlayer from "../AddNewPlayer/AddNewPlayer.vue";
    
    const props = defineProps<Team>()
    // let show = true

    // const handleHideModal = () => {
    //   console.log('Modal is being hidden!');
    //   show = false
    // };
</script>

<template>
    <div class="team-scorecard">
      <div class="team-name">{{ props.team_name }}</div>
  
      <!-- Team Actions -->
      <div class="team-actions">
        <button class="reset" @click="">Reset Points</button>
        <button class="clear" @click="">Clear Team</button>
      </div>
  
      <!-- Points -->
      <div class="points-control">
        <button class="minus" @click="">-</button>
        <span>{{ props.score }}</span>
        <button class="plus" @click="">+</button>
      </div>
  
      <!-- Players List -->
      <div class="players-section">
        <!-- <ul>
          <li>
            Darien
            <button @click="">Remove</button>
          </li>
        </ul> -->
  
        <!-- Add Player -->
        <div class="add-player">
          <!-- <input
            placeholder="Enter player name"
          /> -->
          <button @click="">Add Player To Team</button>
        </div>
      </div>
    </div>

    <!-- <Modal 
      :modal-header="'add team member'"
      :show="show"
      :modal-content="AddNewPlayer"
      :onHide="handleHideModal"
    /> -->
</template>
  
<style scoped lang="scss">
  .team-scorecard {
    /* background-color: #10141e;
    */
    box-shadow: 8px 8px 5px rgba(173, 216, 230, 0.548);
    color: var(--main-text-color);
    
    text-align: center; 
    padding: 10px 10px;
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
      margin: 5px 0px;

      .clear{
        // border: var(--main-text-color) solid 2px;
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
    font-weight: bold;
    cursor: pointer;
  }
</style>
  