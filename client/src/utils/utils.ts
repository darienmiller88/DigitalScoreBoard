import { storeToRefs } from "pinia"
import { scoreBoardApi } from "../api/api"
import { optionsStore } from "../stores/optionsStore"
import { Location } from "../types/types"

const { setAllLocationOptions } = optionsStore()
const { allLocationOptions } = storeToRefs(optionsStore())

export const getAllLocations = async () => {
    if (!allLocationOptions.value.length) {
        const locationsResponse = await scoreBoardApi.get<Location[]>("/get-all-locations")
            
        //Load all the locations into the following ref, storing it into local storage for faster access.
        setAllLocationOptions(locationsResponse.data.map(location => {          
            return location.location_name
        }))
    }
}