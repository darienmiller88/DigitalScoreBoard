export type Player = {
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
    users:         Player[]
    location_name: string
}

export type SavedGame = {
    id:             string
    winner:         Player
    location?:      Location
    teams?:         Team[]
    created_at:     string
    total_points:   number
    average_points: number
}