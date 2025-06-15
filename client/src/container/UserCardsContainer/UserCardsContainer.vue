<script setup lang="ts">
    import UserCard from '../../components/UserCard/UserCard.vue';

    const usernames = [
        "darien miller",
        "angela lourens",
        "mark rubens",
        "shaniqua blank"
    ]
</script>

<template>
    <div class="user-cards">
        <UserCard
            v-for="username in usernames"
            :username="username"
        />
    </div>
</template>

<style scoped lang="scss">
    .user-cards{
        border: 2px solid black;
    }
</style>