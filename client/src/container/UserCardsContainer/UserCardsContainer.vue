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
    <div class="people">People at {{ "Lawrence:" }}</div>
    <div class="user-cards">
        <UserCard
            v-for="username in usernames"
            :username="username"
        />
    </div>
</template>

<style scoped lang="scss">
    .people{
        text-align: center;
        font-size: 30px;
        color: var(--primary-color);
        margin-top: 20px;
    }

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

        // Show two columns of cards for tablets
        @media only screen and (min-width: 768px){
            display: grid;
            grid-template-columns: repeat(2, 1fr);
            gap: 30px;
            max-height: 75vh;
        }

        //Show three columns for laptops less than 2K
        @media only screen and (min-width: 992px){
            display: grid;
            grid-template-columns: repeat(3, 1fr);
            gap: 10px 50px;
            padding: 0px 10px;
            max-height: 75vh;
        }

        // 2k monitors
        @media only screen and (min-width: 2560px){
            display: grid;
            grid-template-columns: repeat(3, 1fr);
            gap: 30px;
            max-height: 75vh;
        }
    }
</style>