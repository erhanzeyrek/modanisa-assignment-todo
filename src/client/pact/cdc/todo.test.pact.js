import TodoService from '../../src/services/todo-service/todo-service';

describe('ToDo API', () => {
  const todoService = new TodoService('http://localhost:8989');

  const expected_body_get_todos = [{ id: 1, message: 'buy some milk' }];
  const expected_body_add_todo = { id: 1 };
  const expected_body_getTodoById = { id: 1, message: 'buy some milk' };
  const sending_body_add_todo = { message: 'buy some milk' };

  describe('addTodo()', () => {
    beforeEach((done) => {
      global.provider
        .addInteraction({
          state: 'creates a todo',
          uponReceiving: 'a POST request to create a todo',
          withRequest: {
            method: 'POST',
            path: '/addTodo',
            headers: {
              Accept: 'application/json; charset=utf-8',
            },
            body: sending_body_add_todo,
          },
          willRespondWith: {
            status: 201,
            headers: {
              'Content-Type': 'application/json; charset=utf-8',
            },
            body: expected_body_add_todo,
          },
        })
        .then(() => done());
    });

    it('send request according to contract', (done) => {
      todoService
        .addTodo(sending_body_add_todo)
        .then((response) => {
          expect(response.status).toEqual(201);
          expect(response.headers['content-type']).toEqual(
            'application/json; charset=utf-8'
          );
          expect(response.data).toEqual(expected_body_add_todo);
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

  describe('getTodoById()', () => {
    beforeEach((done) => {
      global.provider
        .addInteraction({
          state: 'gets a todo by id',
          uponReceiving: 'a GET request to get a todo',
          withRequest: {
            method: 'GET',
            path: '/todos/1',
            headers: {
              Accept: 'application/json; charset=utf-8',
            },
            // body: sending_body_add_todo,
          },
          willRespondWith: {
            status: 200,
            headers: {
              'Content-Type': 'application/json; charset=utf-8',
            },
            body: expected_body_getTodoById,
          },
        })
        .then(() => done());
    });

    it('send request according to contract', (done) => {
      todoService
        .getTodoById(1)
        .then((response) => {
          expect(response.status).toEqual(200);
          expect(response.headers['content-type']).toEqual(
            'application/json; charset=utf-8'
          );
          expect(response.data).toEqual(expected_body_getTodoById);
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

  describe('getAllTodos()', () => {
    beforeEach((done) => {
      global.provider
        .addInteraction({
          state: 'gets all todos',
          uponReceiving: 'a GET request to get all todos',
          withRequest: {
            method: 'GET',
            path: '/todos',
            headers: {
              Accept: 'application/json; charset=utf-8',
            },
          },
          willRespondWith: {
            status: 200,
            headers: {
              'Content-Type': 'application/json; charset=utf-8',
            },
            body: expected_body_get_todos,
          },
        })
        .then(() => done());
    });

    it('send request according to contract', (done) => {
      todoService
        .getAllTodos()
        .then((response) => {
          expect(response.status).toEqual(200);
          expect(response.headers['content-type']).toEqual(
            'application/json; charset=utf-8'
          );
          expect(response.data).toEqual(expected_body_get_todos);
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
