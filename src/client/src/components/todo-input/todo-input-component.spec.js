import { mount } from '@cypress/vue';
import TodoInputComponent from './todo-input-component.vue';

describe('todo-component', () => {
  it('writes text to todo input', () => {
    const inputText = 'buy some milk';
    mount(TodoInputComponent, {
      propsData: {
        inputText,
      },
    });

    cy.get('[id="todo-item"]').should('have.value', inputText);
  });
});
