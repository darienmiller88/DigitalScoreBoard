<script setup lang="ts">
  import NavBar from "./components/NavBar/NavBar.vue"
</script>

<template>
  <NavBar v-if="$route.path !== '/signup'"/>
  <main>
    <RouterView />
  </main>
</template>