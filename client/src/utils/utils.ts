import { storeToRefs } from "pinia"
import { scoreBoardApi } from "../api/api"
import { optionsStore } from "../stores/optionsStore"
import { Location } from "../types/types"

export const getAllLocations = async () => {
    const { setAllLocationOptions } = optionsStore()
    const { allLocationOptions } = storeToRefs(optionsStore())

    //Only load the all the location names if they're NOT in location storage. As a result, this function will be called
    //most one time. This is done to reduce database calls.
    try {
        if (!allLocationOptions.value.length) {
            const locationsResponse = await scoreBoardApi.get<Location[]>("/get-all-locations")
                
            //Load all the locations into the following ref, storing it into local storage for faster access.
            setAllLocationOptions(locationsResponse.data.map(location => {          
                return location.location_name
            }))
        }
    } catch (error) {
        console.log("err:", error);
    }
}