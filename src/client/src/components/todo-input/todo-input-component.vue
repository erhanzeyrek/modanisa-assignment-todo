<template>
  <div>
    <input
      type="text"
      id="todo-item"
      v-bind:value="inputText"
      v-on:input="$emit('input', $event.target.value)"
    />
  </div>
</template>

<script>
export default {
  name: 'todo-input-component',
  props: ['inputText'],
};
</script>
