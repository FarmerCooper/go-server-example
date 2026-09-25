<script setup lang="ts">
interface Todo {
  id: string
  title: string
  completed: boolean
  createdAt: string
}

interface TodoListResponse {
  todos: Todo[]
}

const config = useRuntimeConfig()

const title = ref('')

const {
  data,
  status,
  error,
  refresh,
} = await useFetch<TodoListResponse>('/todos', {
  baseURL: config.public.apiBase,

  default: () => ({
    todos: [],
  }),
})

async function createTodo() {
  const value = title.value.trim()

  if (!value) {
    return
  }

  await $fetch('/todos', {
    baseURL: config.public.apiBase,

    method: 'POST',

    body: {
      title: value,
      completed: false,
    },
  })

  title.value = ''

  await refresh()
}
</script>

<template>
  <main>
    <h1>Todos</h1>

    <form @submit.prevent="createTodo">
      <input
        v-model="title"
        type="text"
        placeholder="What needs to be done?"
      >

      <button type="submit">
        Add Todo
      </button>
    </form>

    <p v-if="status === 'pending'">
      Loading...
    </p>

    <p v-if="error">
      {{ error.message }}
    </p>

    <ul v-if="data">
      <li
        v-for="todo in data.todos"
        :key="todo.id"
      >
        {{ todo.title }}
      </li>
    </ul>
  </main>
</template>