import { useEffect, useState } from "react";
import { fetchTodos } from "./todo.svc";
import { Todo } from "./Todo";
import { AddTodo } from "./AddTodo";
import Grid from "@material-ui/core/Grid";

export function TodoList() {
  const [todos, setTodos] = useState([]);

  const refresh = () => {
    fetchTodos().then((data) => setTodos(data));
  };

  useEffect(() => {
    refresh();
  }, []);

  return (
    <Grid container>
      {todos.map((todo) => (
        <Grid item xs={4} key={todo._id}>
          <Todo todo={todo} refresh={refresh} />
        </Grid>
      ))}
      <Grid item xs={4}>
        <AddTodo refresh={refresh}/>
      </Grid>
    </Grid>
  );
}
