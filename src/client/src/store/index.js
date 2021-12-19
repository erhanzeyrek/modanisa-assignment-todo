import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    todos: [],
  },
  mutations: {
    addTodo(state, payload) {
      state.todos.push(payload.todo);
    },
  },
  getters: {
    GET_TODOS(state) {
      return state.todos;
    },
  },
});

export default store;
