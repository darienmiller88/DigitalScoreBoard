<script setup lang="ts">
  import { onMounted } from "vue";
  import NavBar from "./components/NavBar/NavBar.vue"
  import { getAllLocations } from "./utils/utils"
  import { optionsStore } from "./stores/optionsStore"
  import { storeToRefs } from "pinia";
  
  const { allLocationOptions } = storeToRefs(optionsStore())

  onMounted(async () => {
      await getAllLocations()

      console.log("locations:", allLocationOptions.value);
      
  })
</script>

<template>
  <NavBar />
  <main>
    <RouterView />
  </main>
</template>