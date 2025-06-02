<script setup lang="ts">
    import { SavedGame } from '../../types/types';
    import { useWindowSize } from "@vueuse/core"
    import star from "../../assets/star.png"

    const props = defineProps<SavedGame>()
    const { width } = useWindowSize();
</script>

<template>
    <div class="game">
        <div class="winner-wrapper">
            <img :src="star" alt="">
            <span class="winner"> Winner: {{ props.winner.username }} </span> 
            <img :src="star" alt="">
        </div>
        <div class="score">Score: {{ props.winner.score }}</div>

        <div class="view-players-wrapper">
            <button>
                View players
            </button>
        </div>

        <!--  -->
        <div class="location-date-wrapper">
            <div class="location">Location: 
                <span v-if="props.location">
                    {{ props.location.location_name }}
                </span>
            </div>
            <div class="date-played">Date played: {{ props.created_at }}</div>
        </div>

        <!--  -->
        <div class="points-wrapper">
            <div class="total-points">
                <span>{{ props.total_points }}</span><br>
                Total points 
            </div>
            <div class="average-points">
                <span>{{ props.average_points }}</span><br>
                <span v-if="width < 768">Avg. Points</span>
                <span v-else>Avgerage Points</span>
            </div>
        </div>
    </div>

    
</template>

<style scoped lang="scss">
    .game{
        border-radius: 10px;

        background-color: white;
        box-shadow: 8px 8px 5px var(--primary-color);
        transition: 0.3s;

        width: 80vw;
        text-align: center;
        font-size: 19px;

        margin: auto;
        margin-top: 30px;
        margin-bottom: 30px;
        padding: 10px;

        @media screen and (min-width: 768px) {
            width: fit-content;
            padding: 30px;
        }

        &:hover{
            box-shadow: 10px 10px 15px var(--primary-color);
            transform: translateY(-5px);
        }

        .winner-wrapper{
            display: flex;
            align-items: center;
            font-size: 30px;
            font-style: italic;

            .winner{
                @media screen and (min-width: 768px) {
                    margin: 0px 20px;
                }
            }

            img{
                height: 45px;
                width: auto;
                margin: 0px 10px;

                @media screen and (min-width: 768px) {
                    height: 60px;
                }
            }

            @media screen and (min-width: 768px) {
                margin: 0px 40px;
            }
        }

        .score{
            color: rgb(37, 199, 130);
            font-size: 35px;
        }        

        .view-players-wrapper{
            margin: 15px;

            button{
                padding: 10px 28px;
                
                background-color: var(--primary-color);
                color: aliceblue;
                font-size: 18px;

                border: none;
                border-radius: 8px;
                transition: 0.3s;

                &:hover{
                    cursor: pointer;
                    padding: 10px 40px;
                }
            }
        }

        .location-date-wrapper{
            color: grey;
        }

        .points-wrapper{
            display: flex;
            justify-content: space-between;
            margin-top: 15px;
            color: black;
        }
    }
</style>