<template>
  <Navbar />
  <div class="container mt-5">
    <div v-if="book">
      <!-- Название книги -->
      <nav aria-label="breadcrumb">
        <ol class="breadcrumb">
          <li class="breadcrumb-item"><a :href="getTypeUrl(book.type)">{{ book.type }}</a></li>
          <li class="breadcrumb-item active" aria-current="page">{{ book.title }}</li>
        </ol>
      </nav>

      <div class="row">
        <!-- Левая колонка: изображение книги -->
        <div class="col-md-6">
          <img :src="book.cover" class="img-fluid" :alt="`Обложка книги ${book.title}`" />
        </div>

        <!-- Правая колонка: детали книги и покупка -->
        <div class="col-md-6">
          <h2>{{ book.title }}</h2>
          <p>Автор: <strong>{{ book.author }}</strong></p>
          <p>Год издания: <strong>{{ book.year }}</strong></p>

          <!-- Описание книги -->
          <div class="mt-4">
            <h5>Описание</h5>
            <p>{{ book.description }}</p>
          </div>


        </div>
      </div>
    </div>
    <div v-else>
      <p>Загрузка данных о книге...</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import Navbar from "@/components/Navbar.vue";

const route = useRoute()
const book = ref(null)

const fetchBook = async () => {
  try {
    const response = await fetch(`http://localhost:8080/books/${route.params.id}`)
    if (!response.ok) {
      throw new Error('Failed to fetch book details')
    }
    book.value = await response.json()
  } catch (error) {
    console.error('Error fetching book details:', error)
  }
}
const getTypeUrl = (type) => {
  const typeUrls = {
    "Манга": '/manga',
    "Комиксы": '/comic',
    "Книги": '/books',
  }
  return typeUrls[type] || '/books'
}

onMounted(fetchBook)
</script>

<style scoped>
.container {
  max-width: 800px;
}

.img-fluid {
  width: 100%; /* Изображение занимает всю ширину контейнера */
  max-width: 400px; /* Фиксированная максимальная ширина */
  max-height: 500px; /* Фиксированная максимальная высота */
  object-fit: cover; /* Сохраняет пропорции изображения */
  margin: 0 auto; /* Центровка изображения */
}

</style>
