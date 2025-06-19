<script setup lang="ts">
    import { ref } from 'vue';

    defineProps<{
        username: string
    }>()

    let editFirstName = ref<string>("")
    let editLastName = ref<string>("")

    const onSubmit = () => {
        console.log(editFirstName, editLastName)
    }
</script>

<template>
    <div>
        Editing player "{{ username }}"...
    </div>
    <form @submit.prevent="onSubmit">
        <input 
            class="form-element"
            id="edit-first-name"
            v-model="editFirstName"
            name="edit-first-name"
            type="text" 
            placeholder="Edit First Name"
        /><br />
        <input 
            class="form-element"
            id="edit-last-name"
            v-model="editLastName"
            name="edit-last-name"
            type="text" 
            placeholder="Edit Last Name"
        /><br />
        <button>Submit</button>
    </form>
</template>

<style scoped lang="scss">

</style>