import axios from 'axios';

const LOGIN_URL = 'http://localhost:8081/v1/login';

class SessionService {
    getSession(data) {
        return axios.post(LOGIN_URL, data);
    }

    getUserById(id) {
        return axios.get(LOGIN_URL + id);
    }

    createUser(data) {
        return axios.post(LOGIN_URL, data);
    }

    updateUser(id, data) {
        return axios.put(LOGIN_URL + id, data);
    }

    deleteUser(id) {
        return axios.delete(LOGIN_URL + id);
    }
}

export default new SessionService();