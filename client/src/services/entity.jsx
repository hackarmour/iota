import axios from "axios"
import { ENDPOINT } from '../url';

export const EntityService = {
    getEntities: async function() {
        return await axios.get(ENDPOINT)
    },

    postEntry: async function(data) {
        return await axios.post(ENDPOINT, {data})
    },

    deleteEntry: async function(id) {
        return axios.delete(ENDPOINT+id)
    }
} 