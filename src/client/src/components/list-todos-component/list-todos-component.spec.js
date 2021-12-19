import { mount } from '@cypress/vue';
import ListTodosComponent from './list-todos-component.vue';

describe('todo-component', () => {
  it('writes text to todo input', () => {
    const todos = ['buy some milk'];

    mount(ListTodosComponent, {
      propsData: {
        todos,
      },
    });

    cy.get('.todos').eq(0).should('contain.text', 'buy some milk');
  });
});
