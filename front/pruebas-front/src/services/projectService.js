import axios from 'axios';

const PROJECT_URL = 'http://localhost:8081/v1/projects';

class ProjectService {
    getProjectById(id) {
        const token = localStorage.getItem('token');

        const headers = {
        Authorization: `${token}`
        };

        return axios.get(`${PROJECT_URL}/${id}`, { headers });
    }

    createProject(data){
        const token = localStorage.getItem('token');

        const headers = {
        Authorization: `${token}`
        };

        return axios.post(`${PROJECT_URL}`, data, { headers });
    }

    getProjects() {
        const token = localStorage.getItem('token');

        const headers = {
        Authorization: `${token}`
        };

        return axios.get(PROJECT_URL, { headers });
    }
}

export default new ProjectService();