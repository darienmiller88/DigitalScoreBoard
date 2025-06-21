<script setup lang="ts">
    import { ref } from 'vue';

    const props = defineProps<{
        username: string
        hideModal: () => void
    }>()

    let editFirstName = ref<string>("")
    let editLastName = ref<string>("")

    const onSubmit = () => {
        console.log("name:", editFirstName.value, editLastName.value)
        
        //Close modal after submitting
        props.hideModal()
    }
</script>

<template>
    <!-- <div class="content-title">
        Editing player "{{ username }}"...
    </div> -->
    <form @submit.prevent="onSubmit">
        <label>First Name</label><br />
        <input 
            class="form-element"
            id="edit-first-name"
            v-model="editFirstName"
            name="edit-first-name"
            :maxlength="15"
            type="text" 
            placeholder="Edit First Name"
            required
        /><br /><br/>
        <label>Last Name</label><br />
        <input 
            class="form-element"
            id="edit-last-name"
            v-model="editLastName"
            name="edit-last-name"
            :maxlength="15"
            type="text" 
            placeholder="Edit Last Name"
            required
        /><br />
        <div class="button-wrapper">
            <button>Submit Name Change</button>
        </div>
    </form>
</template>

<style scoped lang="scss">
    .content-title{
        text-align: center;
        font-size: 24px;
        margin-bottom: 20px;
    }

    form{
        width: fit-content;
        margin: auto;

        label{
            font-size: 20px;
        }

        input{
            padding: 10px;
            font-size: 20px;
            width: 75vw;

            @media (min-width: 768px) {
                width: 25vw;
            }
        }

        .button-wrapper{
            text-align: center;
            margin: 15px;
        
            button{
                width: 100%;
                background-color: dodgerblue;
                color: aliceblue;
                border: none;
                border-radius: 10px;
                font-size: 18px;
                padding: 12px;
                transition: 0.3s;

                &:hover{
                    cursor: pointer;
                    background-color: grey;
                }
            }
        }
    }
</style>