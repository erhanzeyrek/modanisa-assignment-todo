/// <reference types="cypress" />
describe('Home', () => {
  beforeEach(() => {
    // Because we're only testing the homepage
    // in this test file, we can run this command
    // before each individual test instead of
    // repeating it in every test.
    cy.visit('/');
  });

  it('Should display the add todo item text box, add button and todo list', () => {
    cy.get('[id="todo-item"]').should('be.visible');
    cy.get('button').contains('Add').should('be.visible');
    cy.get('.todos').should('be.visible');
  });

  it('Write "buy some milk" in the text box, click add button and see the item "buy some milk" in the list', () => {
    cy.get('[id="todo-item"]').type('buy some milk');
    cy.get('button').contains('Add', { timeout: 15000 }).click();
    cy.get('.todos').eq(0).should('contain.text', 'buy some milk');
  });
});
