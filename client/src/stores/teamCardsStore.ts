import { defineStore } from 'pinia';
import { ref } from 'vue';
import { Team } from "../types/types"

export const teamCardsStore = defineStore("teamCards", () => {
    const teamCards = ref<Team[]>([])

    const setTeamCards = (cards: Team[]) => {
        teamCards.value = cards
    }

    const addTeamCard = (card: Team) => {
        teamCards.value = [...teamCards.value, card]
    }

    const removeTeamCard = (cardIndex: number) => {
        teamCards.value = teamCards.value.filter((_, index) => {
            return cardIndex != index
        })
    }

    const addPoints = (index: number, amountToAdd: number) => {
        scoreCards.value[index].score += amountToAdd

        if (scoreCards.value[index].score > 99999) {
            scoreCards.value[index].score = 99999
        }
    }

    const minusPoints = (index: number, amountToAdd: number) => {
        scoreCards.value[index].score -= amountToAdd

        if (scoreCards.value[index].score < 0) {
            scoreCards.value[index].score = 0
        }
    }

    const resetPoints = (index: number) => {
        scoreCards.value[index].score = 0
    }

    const getWinner = (): Card => { 
        let highestScore: Card = scoreCards.value[0]

        for (let i = 1; i < scoreCards.value.length; i++) {
            if (scoreCards.value[i].score > highestScore.score) {
                highestScore = scoreCards.value[i];
            }
        }

        return highestScore
    }

    const resetAllPoints = () => {
        scoreCards.value.forEach(card => {
            card.score = 0
        })
    }

    return { 
        scoreCards, 
        addScoreCard, 
        removeCard, 
        addPoints, 
        minusPoints, 
        resetPoints, 
        resetAllPoints, 
        setCards ,
        getWinner
    }
}, {
    persist: true
})