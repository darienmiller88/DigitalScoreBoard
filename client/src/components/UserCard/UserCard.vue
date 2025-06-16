<script setup lang="ts">
    defineProps<{
        username: string
    }>()
</script>

<template>
    <div class="user-card">
        <div class="username">{{ username }}</div>
        <div class="divider"></div>
        <div class="remove-wrapper">
            <button>Remove User</button>
        </div>
        <div class="edit-wrapper">
            <button>Edit Name</button>
        </div>
    </div>
</template>

<style scoped lang="scss">
    .user-card{
        border: 2px solid var(--primary-color);
        border-radius: 10px;

        width: 85vw;
        box-shadow: 8px 8px 5px rgba(173, 216, 230, 0.548);
        margin: 20px 0px; 

        transition: 0.3s;
        text-align: center;

        @media screen and (min-width: 768px) {
            width: 40vw;
            margin: 20px 0px; 
        }

        @media screen and (min-width: 992px) {
            width: 25vw;
            margin: 20px 0px; 
        }

        @media screen and (min-width: 1400px) {
            width: 15vw;
            margin: 20px 0px; 
        }

        .username{
            color: var(--primary-color);
            font-size: 30px;
        }

        .divider{
            width: 99%;
            border: 2px solid var(--primary-color);
            margin-top: 5px;
            margin-bottom: 20px;
        }
    }
</style>