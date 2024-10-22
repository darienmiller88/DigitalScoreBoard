import { defineStore } from 'pinia';
import { ref } from 'vue';
import { Card } from "../types/types"

export const scoreCardsStore = defineStore("scoreCardsStore", () => {
    const scoreCards = ref<Card[]>([])

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

        if (scoreCards.value[index].score < 0) {
            scoreCards.value[index].score = 0
        }
    }

    return { scoreCards, addScoreCard, removeCard, addPoints, minusPoints }
}, {
    persist: true
})