<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import Navbar from "@/components/Navbar.vue"
import { parseJwt } from "@/utils/jwtUtils"
import Footer from "@/components/Footer.vue"
import Swal from 'sweetalert2'
import {fetchBooksAll} from "@/utils/bookService.js"; // SweetAlert2 для уведомлений

const books = ref([])
const router = useRouter()

const fetchBooks = async () => {
  const fetchedBooks = await fetchBooksAll('Книги');
  if (fetchedBooks.error) {
    console.error('Ошибка при получении манги:', fetchedBooks.error);
    return;
  }
  books.value = fetchedBooks;
}

const addToList = async (bookId, listType) => {
  const token = localStorage.getItem('token');

  if (!token) {
    router.push('/login');
    return;
  }

  const { user_id } = parseJwt(token);

  if (!bookId || !user_id) {
    console.error("Не удалось получить корректный user_id или bookId");
    return;
  }

  const addEndpoint = listType === 'read' ? `/users/${user_id}/read-books/${bookId}` : `/users/${user_id}/abandoned-books/${bookId}`;
  const removeEndpoint = listType === 'read' ? `/users/${user_id}/abandoned-books/${bookId}` : `/users/${user_id}/read-books/${bookId}`;

  try {
    await fetch(`http://localhost:8080${removeEndpoint}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      }
    });

    const addResponse = await fetch(`http://localhost:8080${addEndpoint}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      }
    });

    if (addResponse.ok) {
      books.value = books.value.map(book => {
        if (book.ID === bookId) {
          if (listType === 'read') {
            book.read = true;
            book.abandoned = false;
          } else if (listType === 'abandoned') {
            book.abandoned = true;
            book.read = false;
          }
        }
        return book;
      });

      // Показываем уведомление
      Swal.fire({
        title: listType === 'read' ? 'Добавлено в прочитанные!' : 'Добавлено в заброшенные!',
        icon: 'success',
        confirmButtonText: 'ОК'
      });
    } else {
      console.error(`Не удалось добавить книгу в список ${listType}`);
    }
  } catch (error) {
    console.error('Ошибка при обновлении книги:', error);
  }
};

onMounted(() => {
  fetchBooks()
})
</script>

<template>
  <Navbar />
  <div class="container mt-5">
    <h2 class="mb-4">Список книг</h2>
    <div class="row">
      <div class="col-md-4 mb-4" v-for="book in books" :key="book.id">
        <div class="card">
          <router-link :to="{ name: 'BookDetails', params: { id: book.ID } }">
            <img :src="book.cover" class="card-img-top" :alt="`Обложка книги ${book.title}`" />
          </router-link>
          <div class="card-body">
            <h5 class="card-title">{{ book.title }}</h5>
            <p class="card-text">Автор: {{ book.author }}</p>

            <!-- Кнопка "Добавить в прочитанные" -->
            <button
                @click="addToList(book.ID, 'read')"
                :disabled="book.read"
                class="btn btn-success"
            >
              Прочитано
            </button>

            <!-- Кнопка "Добавить в заброшенные" -->
            <button
                @click="addToList(book.ID, 'abandoned')"
                :disabled="book.abandoned"
                class="btn btn-danger"
            >
              Заброшено
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <Footer />
</template>

<style scoped>
.card {
  transition: transform 0.2s;
}
.card:hover {
  transform: scale(1.05);
}
.card-img-top {
  max-height: 300px;
  object-fit: cover;
}
</style>
