<script setup lang="ts">
  import { onMounted } from "vue";
  import NavBar from "./components/NavBar/NavBar.vue"
  import { getAllLocations } from "./utils/utils"

  onMounted(async () => {
      await getAllLocations()
  })
</script>

<template>
  <NavBar />
  <main>
    <RouterView />
  </main>
</template>