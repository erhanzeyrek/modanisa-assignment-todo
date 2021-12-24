import axios from 'axios';
import adapter from 'axios/lib/adapters/http';

export default class TodoService {
  constructor(apiUrl) {
    this.apiUrl = apiUrl;
  }

  getAllTodos() {
    return axios.request(
      {
        method: 'GET',
        url: `/todos`,
        baseURL: `${this.apiUrl}`,
        headers: {
          Accept: 'application/json; charset=utf-8',
          'Content-Type': 'application/json; charset=utf-8',
        },
      },
      adapter
    );
  }

  addTodo(payload) {
    return axios.request(
      {
        method: 'POST',
        url: `/addTodo`,
        baseURL: `${this.apiUrl}`,
        headers: {
          Accept: 'application/json; charset=utf-8',
          'Content-Type': 'application/json; charset=utf-8',
        },
        data: payload,
      },
      adapter
    );
  }
}
