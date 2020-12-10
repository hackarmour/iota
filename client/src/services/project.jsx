import axios from 'axios';
import { ENDPOINT } from '../url';

export const ProjectService = {
    getProject: async function(id) {
        return await axios.get(ENDPOINT+'projects/'+id)
    },

    createProject: async function(data) {
        return await axios.post(ENDPOINT+'projects/'+id, {data})
    }
}