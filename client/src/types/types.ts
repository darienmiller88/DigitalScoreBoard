export type PlayerCard = {
    username: string
    score: number
}

export type AvailablePlayer = {
    player_name: string
    isAddedToGame: boolean
}

export type Team = {
    team_name: string
    players: string[]
    score: number
}

export type Location = {
    id:            string
    users:         PlayerCard[]
    location_name: string
}

export type SavedGame = {
    id:             string
    winner:         PlayerCard
    location_name:  string
    players?:       PlayerCard[]
    teams?:         Team[]
    created_at:     string
    total_points:   number
    average_points: number
}