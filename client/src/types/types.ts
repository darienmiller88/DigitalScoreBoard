export type Card = {
    username: string
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
    location:       Location
    total_points:   number
    average_points: number
}