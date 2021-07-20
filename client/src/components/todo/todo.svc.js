const BASE_URL = "http://localhost:8080/v1";

export function fetchTodos() {
  return fetch(`${BASE_URL}/todo`)
    .then((response) => response.json())
}
