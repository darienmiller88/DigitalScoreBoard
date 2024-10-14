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

    return { scoreCards, addScoreCard, removeCard }
}, {
    persist: true
})