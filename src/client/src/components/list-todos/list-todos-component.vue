<template>
  <div>
    <ul class="todos">
      <li v-for="todo in todoList" :key="todo">{{ todo }}</li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'list-todos-component',
  props: ['todos'],

  computed: {
    todoList() {
      return this.todos;
    },
  },
};
</script>
<style>
.todos {
  height: 100px;
}
</style>
