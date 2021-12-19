<template>
  <div>
    <ul class="todos">
      <li v-for="todo in todos" :key="todo">{{ todo }}</li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'list-todos-component',
  props: ['todos'],
  data: () => {
    return {};
  },
};
</script>
