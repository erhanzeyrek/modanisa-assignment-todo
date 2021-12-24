import TodoService from '../../src/services/todo-service/todo-service';
import * as Pact from '@pact-foundation/pact';

describe('ToDo API', () => {
  const todoService = new TodoService('http://localhost:10000');

  describe('getAllTodos()', () => {
    beforeEach((done) => {
      const contentTypeJsonMatcher = Pact.Matchers.term({
        matcher: 'application\\/json; *charset=utf-8',
        generate: 'application/json; charset=utf-8',
      });

      global.provider
        .addInteraction({
          state: 'gets all todos',
          uponReceiving: 'a GET request to get all todos',
          withRequest: {
            method: 'GET',
            path: '/todos',
            headers: {
              Accept: 'application/json',
              'Content-Type': contentTypeJsonMatcher,
            },
            body: { message: 'buy some milk' },
          },
          willRespondWith: {
            status: 200,
            headers: {
              'Content-Type': contentTypeJsonMatcher,
            },
            body: Pact.Matchers.somethingLike([
              { id: 33, message: 'buy some milk' },
            ]),
          },
        })
        .then(() => done());
    });

    it('send request according to contract', (done) => {
      todoService
        .getAllTodos()
        .then((response) => {
          expect(response).somethingLike([
            { id: 33, message: 'buy some milk' },
          ]);
        })
        .then(() => {
          global.provider.verify().then(
            () => done(),
            (error) => {
              done.fail(error);
            }
          );
        });
    });
  });

  describe('addTodo()', () => {
    beforeEach((done) => {
      const contentTypeJsonMatcher = Pact.Matchers.term({
        matcher: 'application\\/json; *charset=utf-8',
        generate: 'application/json; charset=utf-8',
      });

      global.provider
        .addInteraction({
          state: 'creates a todo',
          uponReceiving: 'a POST request to create a todo',
          withRequest: {
            method: 'POST',
            path: '/addTodo',
            headers: {
              Accept: 'application/json',
              'Content-Type': contentTypeJsonMatcher,
            },
            body: { message: 'buy some milk' },
          },
          willRespondWith: {
            status: 200,
            headers: {
              'Content-Type': contentTypeJsonMatcher,
            },
            body: Pact.Matchers.somethingLike({ id: 33 }),
          },
        })
        .then(() => done());
    });

    it('send request according to contract', (done) => {
      todoService
        .addTodo('buy some milk')
        .then((response) => {
          expect(response.id).somethingLike(33);
        })
        .then(() => {
          global.provider.verify().then(
            () => done(),
            (error) => {
              done.fail(error);
            }
          );
        });
    });
  });
});
