<script setup lang="ts">
  import { onMounted } from "vue";
import NavBar from "./components/NavBar/NavBar.vue"

  onMounted(() => {
      console.log("app.vue")
  })
</script>

<template>
  <NavBar />
  <main>
    <RouterView />
  </main>
</template>