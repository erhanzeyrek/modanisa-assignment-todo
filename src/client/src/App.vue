<template>
  <div id="app">
    <todo-input-component v-model="todo"></todo-input-component>
    <button id="add-todo" @click="addTodo(todo)">Add</button>
    <list-todos-component :todos="todos"></list-todos-component>
  </div>
</template>

<script>
import TodoInputComponent from './components/todo-input/todo-input-component.vue';
import ListTodosComponent from './components/list-todos/list-todos-component.vue';

export default {
  name: 'App',
  components: {
    TodoInputComponent,
    ListTodosComponent,
  },
  computed: {
    todos() {
      return this.$store.getters['GET_TODOS'];
    },
  },
  data: () => {
    return { todo: null };
  },
  methods: {
    addTodo(value) {
      this.$store.commit('addTodo', { todo: value });
    },
  },
};
</script>
