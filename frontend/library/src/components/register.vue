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
              <h2 class="fs-6 fw-normal text-center text-secondary mb-4">Create your account</h2>
              <form @submit.prevent="register">
                <div class="row gy-2 overflow-hidden">
                  <div class="col-12">
                    <div class="form-floating mb-3">
                      <input v-model="username" type="text" placeholder="Username" class="form-control" :class="{'is-invalid': usernameError}">
                      <label for="username" class="form-label">Username</label>
                      <div v-if="usernameError" class="invalid-feedback">{{ usernameError }}</div>
                    </div>
                  </div>
                  <div class="col-12">
                    <div class="form-floating mb-3">
                      <input v-model="email" type="email" placeholder="Email" class="form-control" :class="{'is-invalid': emailError}">
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
                    <div class="form-floating mb-3">
                      <input v-model="confirmPassword" type="password" placeholder="Confirm Password" class="form-control" :class="{'is-invalid': confirmPasswordError}">
                      <label for="confirmPassword" class="form-label">Confirm Password</label>
                      <div v-if="confirmPasswordError" class="invalid-feedback">{{ confirmPasswordError }}</div>
                    </div>
                  </div>

                  <div class="col-12">
                    <div class="d-grid my-3">
                      <button class="btn btn-primary btn-lg" type="submit">Register</button>
                    </div>
                  </div>
                  <div class="col-12">
                    <p class="m-0 text-secondary text-center">Already have an account? <a href="/login" class="link-primary text-decoration-none">Log in</a></p>
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

const username = ref('')
const email = ref('')
const password = ref('')
const usernameError = ref(null)
const emailError = ref(null)
const passwordError = ref(null)
const serverError = ref(null)
const router = useRouter()

const confirmPassword = ref('')
const confirmPasswordError = ref(null)

const validateForm = () => {
  usernameError.value = username.value.length < 3 ? 'Username must be at least 3 characters' : null
  emailError.value = !/^\S+@\S+\.\S+$/.test(email.value) ? 'Invalid email address' : null
  passwordError.value = password.value.length < 6 ? 'Password must be at least 6 characters' : null
  confirmPasswordError.value = confirmPassword.value !== password.value ? 'Passwords do not match' : null

  return !usernameError.value && !emailError.value && !passwordError.value && !confirmPasswordError.value
}



const register = async () => {
  if (!validateForm()) {
    return
  }

  try {
    const response = await fetch('http://localhost:8080/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username: username.value, email: email.value, password: password.value })
    })
    const data = await response.json()
    if (response.ok) {
      router.push('/login') // перенаправление на страницу логина после успешной регистрации
    } else {
      serverError.value = data.message
    }
  } catch (error) {
    console.error('Error:', error)
    serverError.value = 'Something went wrong. Please try again.'
  }
}
</script>
