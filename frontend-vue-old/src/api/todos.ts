export interface Todo {
  id: number
  title: string
  completed: boolean
  createdAt: string
}

const API_URL = 'http://localhost:8080'

export async function listTodos(): Promise<Todo[]> {
  const response = await fetch(`${API_URL}/todos`)

  if (!response.ok) {
    throw new Error(
      `Failed to load todos: ${response.status}`,
    )
  }

  const data = await response.json()

  return data.todos
}

export async function createTodo(
  title: string,
): Promise<Todo> {
  const response = await fetch(`${API_URL}/todos`, {
    method: 'POST',

    headers: {
      'Content-Type': 'application/json',
    },

    body: JSON.stringify({
      title,
      completed: false,
    }),
  })

  if (!response.ok) {
    throw new Error(
      `Failed to create todo: ${response.status}`,
    )
  }

  return response.json()
}