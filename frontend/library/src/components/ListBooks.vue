<template>
  <Navbar />
  <div class="container mt-5">
    <h3 class="mb-4">Прочитанные книги</h3>
    <ul class="list-group mb-5">
      <li
          v-for="book in readBooks"
          :key="book.ID"
          class="list-group-item d-flex justify-content-between align-items-center"
      >
        <router-link :to="{ name: 'BookDetails', params: { id: book.ID } }">
        {{ book.title}}
        </router-link>
        <div>
          <span class="badge bg-success rounded-pill me-2">Прочитано</span>
          <button class="btn btn-outline-danger btn-sm" @click="removeBook(book.ID, 'read')">
            Удалить
          </button>
        </div>
      </li>
    </ul>

    <h3 class="mb-4">Заброшенные книги</h3>
    <ul class="list-group">
      <li
          v-for="book in abandonedBooks"
          :key="book.id"
          class="list-group-item d-flex justify-content-between align-items-center"
      >
        <router-link :to="{ name: 'BookDetails', params: { id: book.ID } }">
        {{ book.title }}
        </router-link>
        <div>
          <span class="badge bg-danger rounded-pill me-2">Заброшено</span>
          <button class="btn btn-outline-danger btn-sm" @click="removeBook(book.ID, 'abandoned')">
            Удалить
          </button>
        </div>
      </li>
    </ul>
  </div>

</template>

<script setup>
import {ref, onMounted, onUnmounted} from 'vue'
import Navbar from "@/components/Navbar.vue";
import Footer from "@/components/Footer.vue";

const readBooks = ref([])
const abandonedBooks = ref([])

// Функция для получения токена и user_id
const token = localStorage.getItem('token')
let user_id = null
if (token) {
  const { user_id: id } = parseJwt(token)
  user_id = id
}

// Функция для получения книг
const fetchBooks = async () => {
  try {
    const token = localStorage.getItem('token')
    if (!token) {
      console.error('Token not found in localStorage')
      return
    }

    const headers = {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json'
    }

    const readResponse = await fetch(`http://localhost:8080/users/${user_id}/read-books`, {
      method: 'GET',
      headers: headers
    })
    if (!readResponse.ok) {
      throw new Error(`Failed to fetch read books, status: ${readResponse.status}`)
    }
    readBooks.value = await readResponse.json()

    const abandonedResponse = await fetch(`http://localhost:8080/users/${user_id}/abandoned-books`, {
      method: 'GET',
      headers: headers
    })
    if (!abandonedResponse.ok) {
      throw new Error(`Failed to fetch abandoned books, status: ${abandonedResponse.status}`)
    }
    abandonedBooks.value = await abandonedResponse.json()
  } catch (error) {
    console.error('Error:', error)
  }
}

// Функция для удаления книги из списка
const removeBook = async (bookId, listType) => {
  const endpoint = listType === 'read'
      ? `/users/${user_id}/read-books/${bookId}`
      : `/users/${user_id}/abandoned-books/${bookId}`

  try {
    const response = await fetch(`http://localhost:8080${endpoint}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
      }
    })

    if (response.ok) {
      if (listType === 'read') {
        readBooks.value = readBooks.value.filter(book => book.id !== bookId)
      } else {
        abandonedBooks.value = abandonedBooks.value.filter(book => book.id !== bookId)
      }
      console.log(`Книга успешно удалена из списка ${listType}`)
    } else {
      console.error(`Не удалось удалить книгу из списка ${listType}, статус: ${response.status}`)
    }
  } catch (error) {
    console.error('Ошибка при удалении книги:', error)
  }
}

const socket = new WebSocket('ws://localhost:8080/ws')


const setupWebSocket = () => {
  socket.onmessage = (event) => {
    const message = JSON.parse(event.data)
    if (message.action === 'add') {
      fetchBooks() // Обновляем список книг при добавлении
    } else if (message.action === 'remove') {
      fetchBooks() // Обновляем список книг при удалении
    }else if (message.action === 'drop') {
      fetchBooks() // Обновляем список книг при удалении
    }
  }
}

// Вызов fetchBooks при монтировании компонента

onMounted(() => {
  setupWebSocket()
  fetchBooks()
})
onUnmounted(() => {
  socket.close()
})

// Функция для разбора JWT токена
function parseJwt(token) {
  const base64Url = token.split('.')[1]
  const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/')
  const jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
    return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)
  }).join(''))

  return JSON.parse(jsonPayload)
}
</script>

<style>
.container {
  max-width: 800px;
}

h3 {
  font-size: 1.75rem;
  font-weight: bold;
}

.list-group-item {
  font-size: 1.25rem;
}
</style>
