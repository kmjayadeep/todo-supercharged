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
