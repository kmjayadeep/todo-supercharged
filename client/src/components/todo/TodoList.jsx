import { useEffect, useState } from "react";
import { fetchTodos } from './todo.svc';


export function TodoList() {
  const [ todos, setTodos ] = useState([])

  useEffect(()=>{
    fetchTodos().then(data=>setTodos(todos));
  }, []);

  return ""
}

