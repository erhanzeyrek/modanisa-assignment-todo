<template>
  <div>
    <input type="text" id="todo-item" :value="text" />
  </div>
</template>

<script>
export default {
  name: 'todo-input-component',
  props: ['inputText'],
  data: function () {
    return { text: this.inputText };
  },
  computed: {},
};
</script>
