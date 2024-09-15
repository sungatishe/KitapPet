<template>
  <section class="py-3 py-md-5">
    <div class="container">
      <div class="row justify-content-center">
        <div class="col-12 col-sm-10 col-md-8 col-lg-6 col-xl-5 col-xxl-4">
          <div class="card border border-light-subtle rounded-3 shadow-sm">
            <div class="card-body p-3 p-md-4 p-xl-5">
              <div class="text-center mb-3">
                <a href="#!">
                  <img src="@/assets/imgs/book.png" alt="BootstrapBrain Logo" width="70" height="57">
                </a>
              </div>
              <h2 class="fs-6 fw-normal text-center text-secondary mb-4">Sign in to your account</h2>
              <form @submit.prevent="login">
                <div class="row gy-2 overflow-hidden">
                  <div class="col-12">
                    <div class="form-floating mb-3">
                      <input v-model="email" type="text" placeholder="Email" class="form-control" :class="{'is-invalid': emailError}">
                      <label for="email" class="form-label">Email</label>
                      <div v-if="emailError" class="invalid-feedback">{{ emailError }}</div>
                    </div>
                  </div>
                  <div class="col-12">
                    <div class="form-floating mb-3">
                      <input v-model="password" type="password" placeholder="Password" class="form-control" :class="{'is-invalid': passwordError}">
                      <label for="password" class="form-label">Password</label>
                      <div v-if="passwordError" class="invalid-feedback">{{ passwordError }}</div>
                    </div>
                  </div>
                  <div class="col-12">
                    <div class="d-grid my-3">
                      <button class="btn btn-primary btn-lg" type="submit">Log in</button>
                    </div>
                  </div>
                  <div class="col-12">
                    <p class="m-0 text-secondary text-center">Don't have an account? <a href="/register" class="link-primary text-decoration-none">Sign up</a></p>
                  </div>
                  <div class="col-12">
                    <p v-if="serverError" class="text-danger text-center">{{ serverError }}</p>
                  </div>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import 'bootstrap';

const email = ref('')
const password = ref('')
const emailError = ref('')
const passwordError = ref('')
const serverError = ref('')
const router = useRouter()

const validateForm = () => {
  let isValid = true

  if (!email.value) {
    emailError.value = 'Email is required'
    isValid = false
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.value)) {
    emailError.value = 'Invalid email format'
    isValid = false
  } else {
    emailError.value = ''
  }

  if (!password.value) {
    passwordError.value = 'Password is required'
    isValid = false
  } else if (password.value.length < 6) {
    passwordError.value = 'Password must be at least 6 characters'
    isValid = false
  } else {
    passwordError.value = ''
  }

  return isValid
}

const login = async () => {
  serverError.value = ''

  if (!validateForm()) return

  try {
    const response = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email: email.value, password: password.value })
    })
    const data = await response.json()

    if (response.ok) {
      localStorage.setItem('token', data.token)
      router.push('/dashboard')
    } else {
      serverError.value = data.message || 'Login failed'
    }
  } catch (error) {
    serverError.value = 'An error occurred. Please try again later.'
    console.error('Error:', error)
  }
}
</script>
