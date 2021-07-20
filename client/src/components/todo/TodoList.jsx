import { useEffect, useState } from "react";
import { fetchTodos } from './todo.svc';
import { Todo } from './Todo';


export function TodoList() {
  const [ todos, setTodos ] = useState([])

  useEffect(()=>{
    fetchTodos().then(data=>setTodos(data));
  }, []);

  return (
    <>
    {todos.map((todo) => (
      <Todo todo={todo}/>
      ))}
    </>
  )
}

