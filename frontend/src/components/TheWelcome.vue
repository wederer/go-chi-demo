<script setup lang="ts">
import { reactive } from 'vue'
import * as Books from '@/models/books';

const state = reactive<{ books: Array<Books.Model>, error: null | Error }>({
  books: [],
  error: null
})

function fetchBooks() {
  Books.getAll()
      .then(books => state.books = books)
      .catch(err => state.error = err)
}


function deleteBook(id: string) {
  Books.deleteOne(id)
      .then(res => fetchBooks())
      .catch(err => state.error = err)
}

// fetch once at first render
fetchBooks()
</script>

<template>
  <h1>Test Change</h1>
  <h3 v-if="state.error">Error: {{state.error}}</h3>
  <div v-else>
    <v-table v-if="state.books.length > 0">
      <thead>
      <tr>
        <th class="text-left">
          Title
        </th>
        <th class="text-left">
          Number of Pages
        </th>
        <th>
        </th>
      </tr>
      </thead>
      <tbody>
      <tr
          v-for="item in state.books"
          :key="item.title"
      >
        <td>{{ item.title }}</td>
        <td>{{ item.no_pages }}</td>
        <td>
          <button @click="deleteBook(item.id)">Delete</button>
        </td>
      </tr>
      </tbody>
    </v-table>
    <h2 v-else>No data available</h2>
  </div>
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>
