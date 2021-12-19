<template>
  <div id="app">
    <todo-input-component></todo-input-component>
    <button>Add</button>
    <list-todos-component></list-todos-component>
  </div>
</template>

<script>
import TodoInputComponent from './components/todo-input/todo-input-component.vue';
import ListTodosComponent from './components/list-todos/list-todos-component.vue';

export default {
  name: 'App',
  props: ['todos'],
  components: {
    TodoInputComponent,
    ListTodosComponent,
  },
};
</script>
