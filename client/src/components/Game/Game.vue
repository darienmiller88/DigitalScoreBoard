<script setup lang="ts">
    import { SavedGame } from '../../types/types';
    import star from "../../assets/star.png"

    const props = defineProps<SavedGame>()
</script>

<template>
    <div class="game">
        <div class="winner">
            <img :src="star" alt="">
            Winner: {{ props.winner.username }}
            <img :src="star" alt="">
        </div>
        <div class="score">Score: {{ props.winner.score }}</div>

        <!--  -->
        <div class="location-date-wrapper">
            <div class="location">Location: 
                <span>
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
                Average Points
            </div>
        </div>
    </div>
</template>

<style scoped lang="scss">
    .game{
        border-radius: 10px;

        // color: var(--main-text-color);
        background-color: white;
        box-shadow: 8px 8px 5px rgba(173, 216, 230, 0.548);
        transition: 0.3s;

        width: fit-content;
        text-align: center;
        font-size: 19px;

        margin: auto;
        margin-top: 30px;
        margin-bottom: 30px;
        padding: 30px;

        &:hover{
            box-shadow: 10px 10px 15px var(--main-text-color);
            transform: translateY(-5px);
        }

        .winner{
            display: flex;
            align-items: center;
            font-size: 40px;
            font-style: italic;
            // border: 2px solid var(--main-text-color);

            img{
                height: 60px;
                width: auto;
                margin: 0px 10px;
            }
        }

        .score{
            color: rgb(37, 199, 130);
            font-size: 35px;
        }        

        .location-date-wrapper{
            margin-top: 15px;
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