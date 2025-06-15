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
        display: grid;
        grid-template-columns: 1fr;    

        width: fit-content; 
        max-height: 50vh;
        max-width: 90vw;

        margin: auto;
        margin-top: 20px;
        margin-bottom: 20px;
        overflow-y: scroll;

        border: 2px solid black;

        @media only screen and (min-width: 768px){
            display: grid;
            grid-template-columns: repeat(2, 1fr);
            gap: 30px;
            max-height: 75vh;
        }
    }
</style>