import axios from 'axios';
import { ENDPOINT } from '../url';

export const ProjectService = {
    getOneProject: async function(id) {
        return await axios.get(ENDPOINT+'projects/'+id)
    },

    getProjects: async function() {
        return await axios.get(ENDPOINT+'projects')
    },

    createProject: async function(data) {
        return await axios.post(ENDPOINT+'projects', {data})
    }
}