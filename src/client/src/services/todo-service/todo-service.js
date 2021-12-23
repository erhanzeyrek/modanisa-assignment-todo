import axios from 'axios';

export default class TodoService {
  instance = axios.create({
    baseURL: process.ENV.API_URL,
    timeout: 5000,
    headers: {
      'Content-Type': 'application/json',
    },
  });

  getAllTodos() {
    return this.instance.get('/todos');
  }
}
