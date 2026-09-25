<script setup lang="ts">
import { onMounted, ref } from 'vue'
import {
  createTodo,
  listTodos,
  type Todo,
} from './api/todos'

const todos = ref<Todo[]>([])
const title = ref('')
const loading = ref(false)

async function loadTodos() {
  loading.value = true

  try {
    todos.value = await listTodos()
  } finally {
    loading.value = false
  }
}

async function addTodo() {
  if (!title.value.trim()) {
    return
  }

  await createTodo(title.value)

  title.value = ''

  await loadTodos()
}

onMounted(loadTodos)
</script>

<template>
  <main>
    <h1>Todos</h1>

    <form @submit.prevent="addTodo">
      <input
        v-model="title"
        placeholder="What needs to be done?"
      />

      <button type="submit">
        Add
      </button>
    </form>

    <p v-if="loading">
      Loading...
    </p>

    <ul v-else>
      <li
        v-for="todo in todos"
        :key="todo.id"
      >
        {{ todo.title }}
      </li>
    </ul>
  </main>
</template>