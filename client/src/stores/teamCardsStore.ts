import { defineStore } from 'pinia';
import { ref } from 'vue';
import { PlayerCard, Team } from "../types/types"

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
        teamCards.value[index].score += amountToAdd

        if (teamCards.value[index].score > 99999) {
            teamCards.value[index].score = 99999
        }
    }

    const minusPoints = (index: number, amountToAdd: number) => {
        teamCards.value[index].score -= amountToAdd
    }

    const resetPointsForTeam = (index: number) => {
        teamCards.value[index].score = 0
    }

    const getWinningTeam = (): PlayerCard => { 
        let highestScoringTeam: Team = teamCards.value[0]

        for (let i = 1; i < teamCards.value.length; i++) {
            if (teamCards.value[i].score > highestScoringTeam.score) {
                highestScoringTeam = teamCards.value[i];
            }
        }

        return {
            username: highestScoringTeam.team_name,
            score: highestScoringTeam.score
        }
    }

    const getTotalPoints = (): number => {
        let totalPoints: number = 0

        teamCards.value.forEach(team => {
            totalPoints += team.score
        })

        return totalPoints
    }

    const totalPoints = (): number => {
        return teamCards.value.reduce((accumulator, team) => accumulator + team.score, 0)
    }

    const addPlayerToTeam = (index: number, player: string) => {
        teamCards.value[index].players.push(player)
    }

    const getPlayers = (): PlayerCard[] => {
        let players: PlayerCard[] = []

        teamCards.value.forEach(team => {
            players = [...players, ...team.players.map(player => ({
                username: player,
                score: 0 
            }))]
        })

        return players;
    }

    const getAveragePoints = (): number => {
        return getTotalPoints() / teamCards.value.length
    }

    const resetAllPoints = () => {
        teamCards.value.forEach(card => {
            card.score = 0
        })
    }

    return { 
        teamCards, 
        addTeamCard, 
        addPlayerToTeam,
        removeTeamCard, 
        addPoints, 
        minusPoints, 
        resetPointsForTeam, 
        resetAllPoints, 
        setTeamCards,
        getWinningTeam,
        totalPoints,
        getAveragePoints, 
        getPlayers
    }
}, {
    persist: true
})