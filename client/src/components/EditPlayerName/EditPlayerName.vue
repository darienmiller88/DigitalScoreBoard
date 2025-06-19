<script setup lang="ts">
    defineProps<{
        username: string
    }>()
</script>

<template>
    <div>
        Editing {{ username }}
    </div>
    <form @submit.prevent="">
        <input type="text" placeholder="Edit First Name">
        <input type="text" placeholder="Edit Last Name"/>
    </form>
</template>

<style scoped lang="scss">

</style>