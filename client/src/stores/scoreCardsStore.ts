import { defineStore } from 'pinia';
import { ref } from 'vue';
import { Card } from "../types/types"

export const scoreCardsStore = defineStore("scoreCardsStore", () => {
    const scoreCards = ref<Card[]>()

    return { scoreCards }
})