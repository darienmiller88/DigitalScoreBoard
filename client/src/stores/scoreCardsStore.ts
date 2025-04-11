import { defineStore } from 'pinia';
import { ref } from 'vue';
import { Card } from "../types/types"

export const scoreCardsStore = defineStore("scoreCards", () => {
    const scoreCards = ref<Card[]>([])

    const setUserCards = (cards: Card[]) => {
        scoreCards.value = cards
    }

    const addScoreCard = (card: Card) => {
        scoreCards.value = [...scoreCards.value, card]
    }

    const removeCard = (cardIndex: number) => {
        scoreCards.value = scoreCards.value.filter((_, index) => {
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
    }

    const resetPoints = (index: number) => {
        scoreCards.value[index].score = 0
    }

    const totalPoints = (): number => {
        return scoreCards.value.reduce((accumulator, currentValue) => accumulator + currentValue.score, 0)
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
        setUserCards,
        getWinner, 
        totalPoints
    }
}, {
    persist: true
})