import axios from 'axios';

const USER_URL = 'http://localhost:8081/v1/user';

class UserService {
    getAllUsers() {
        return axios.get(USER_URL);
    }

    getUserById(id) {
        return axios.get(USER_URL + id);
    }

    createUser(data) {
        return axios.post(USER_URL, data);
    }

    updateUser(id, data) {
        return axios.put(USER_URL + id, data);
    }

    deleteUser(id) {
        return axios.delete(USER_URL + id);
    }
}

export default new UserService();