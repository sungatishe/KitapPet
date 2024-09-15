<script setup>
// Initialization for ES Users
import { Dropdown, Collapse, initMDB } from "mdb-ui-kit";
import 'bootstrap';
import router from "@/router.js";

initMDB({ Dropdown, Collapse });

const logout = () => {
  localStorage.removeItem('token');
  router.push('/login');
}
</script>

<template>
  <!-- Navbar -->
  <nav class="navbar navbar-expand-lg navbar-light bg-body-tertiary">
    <!-- Container wrapper -->
    <div class="container-fluid">
      <!-- Toggle button -->
      <button
          data-mdb-collapse-init
          class="navbar-toggler"
          type="button"
          data-mdb-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent"
          aria-expanded="false"
          aria-label="Toggle navigation"
      >
        <i class="fas fa-bars"></i>
      </button>

      <!-- Collapsible wrapper -->
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <!-- Navbar brand -->
        <a class="navbar-brand mt-2 mt-lg-0" href="/dashboard">
          <img
              src="@/assets/imgs/book.png"
              height="30"
              alt="MDB Logo"
              loading="lazy"
          />
        </a>
        <!-- Left links -->
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <a class="nav-link" href="/books">Книги</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="/manga">Манга</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="/comic">Комиксы</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="/list">Список литературы</a>
          </li>
        </ul>
        <!-- Left links -->
      </div>
      <!-- Collapsible wrapper -->

      <!-- Right elements -->
      <div class="d-flex align-items-center">
        <!-- Avatar -->
        
        <div>
          <a
              class="d-flex align-items-center"
              href="#"
              role="button"
              aria-expanded="false"
              @click="logout"
          >
            <img
                src="@/assets/imgs/logout.png"
                height="25"
                alt="Black and White Portrait of a Man"
                loading="lazy"
            />
          </a>

        </div>
      </div>
      <!-- Right elements -->
    </div>
    <!-- Container wrapper -->
  </nav>
  <!-- Navbar -->
</template>
