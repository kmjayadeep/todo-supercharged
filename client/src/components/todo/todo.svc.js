const BASE_URL = "http://localhost:8080/v1";

export function fetchTodos() {
  return fetch(`${BASE_URL}/todo`).then((response) => response.json());
}

export function markAsDone(id) {
  return fetch(`${BASE_URL}/todo/${id}/done`, {
    method: "PUT",
  });
}

export function deleteTodo(id) {
  return fetch(`${BASE_URL}/todo/${id}`, {
    method: "DELETE",
  });
}

export function addTodo(title, description) {
  const body = JSON.stringify({
    title,
    description
  });
  return fetch(`${BASE_URL}/todo`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body
  });
}
