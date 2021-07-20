import { useEffect, useState } from "react";
import { fetchTodos } from './todo.svc';
import { Todo } from './Todo';
import Grid from '@material-ui/core/Grid';


export function TodoList() {
  const [ todos, setTodos ] = useState([])

  useEffect(()=>{
    fetchTodos().then(data=>setTodos(data));
  }, []);

  return (
    <Grid container>
    {todos.map((todo) => (
      <Grid item xs={4} key={todo._id}>
        <Todo todo={todo}/>
      </Grid>
      ))}
    </Grid>
  )
}

