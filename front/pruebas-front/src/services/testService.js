import axios from 'axios';

const TEST_URL = 'http://localhost:8081/v1/test';

class TestService {
    createTest(data){
        const token = localStorage.getItem('token');

        const headers = {
        Authorization: `${token}`
        };

        return axios.post(`${TEST_URL}`, data, { headers });
    }
}

export default new TestService();