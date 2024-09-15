import { createRouter, createWebHistory } from 'vue-router'
import Register from './components/register.vue'
import Login from './components/Login.vue'
import Dashboard from './components/Dashboard.vue'
import ListBooks from "@/components/ListBooks.vue";
import Books from "@/components/Books.vue";
import BookDetails from "@/components/BookDetails.vue";
import Manga from "@/components/Manga.vue";
import Comics from "@/components/Comics.vue";

const routes = [
    { path: '/register', component: Register },
    { path: '/login', component: Login },
    { path: '/list', component: ListBooks, meta: { requiresAuth: true }  },
    { path: '/books', component: Books, meta: { requiresAuth: true }  },
    { path: '/dashboard', component: Dashboard, meta: { requiresAuth: true } },
    { path: '/books/:id', name: "BookDetails", component: BookDetails, meta: { requiresAuth: true } },
    { path: '/manga', component: Manga, meta: { requiresAuth: true } },
    { path: '/comic', component: Comics, meta: { requiresAuth: true } },
    { path: '/:catchAll(.*)', redirect: '/dashboard' },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

router.beforeEach((to, from, next) => {
    const token = localStorage.getItem('token')

    // Если пользователь авторизован, перенаправляем с login/register на dashboard
    if ((to.path === '/login' || to.path === '/register') && token) {
        next('/dashboard')
    }
    // Если требуется авторизация и токена нет, перенаправляем на login
    else if (to.matched.some(record => record.meta.requiresAuth) && !token) {
        next('/login')
    } else {
        next()
    }
})

export default router
