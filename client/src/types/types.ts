export type Card = {
    username: string
    score: number
}

export type Team = {
    team_name: string
    players: string[]
    score: number
}

export type Location = {
    id:            string
    users:         Card[]
    location_name: string
}

export type SavedGame = {
    id:             string
    winner:         Card
    location?:      Location
    teams?:         Location[]
    created_at:     string
    total_points:   number
    average_points: number
}