import axios from "axios"
import { ENDPOINT } from '../url';

export const EntityService = {
    postEntry: async function(data) {
        return await axios.post(ENDPOINT+'/entity', {data})
    },

    deleteEntry: async function(id) {
        return axios.delete(ENDPOINT+'/entity/'+id)
    }
} 