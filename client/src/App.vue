<script setup lang="ts">
  import NavBar from "./components/NavBar/NavBar.vue"
</script>

<template>
  <NavBar />
  <div class="title">Jeopardy Score Board</div>
  <main>
    <RouterView />
  </main>
</template>

<style scoped lang="scss">
  .title{
    color: var(--main-text-color);
    text-align: center;
    font-size: 40px;
    margin: 40px 0px;

    @media screen and (min-width: 768px) {
        font-size: 60px;
        margin: 10px 0px;
    }
  }
</style>