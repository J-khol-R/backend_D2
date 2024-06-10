import axios from 'axios';

const URL = 'http://localhost:8081/v1/projects';

class ProjectService {
    createProject(data) {
        return axios.post(URL, data);
    }

    getProjects() {
        return axios.get(URL);
    }
}

export default new ProjectService();