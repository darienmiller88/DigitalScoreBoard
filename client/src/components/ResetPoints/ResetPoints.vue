<script setup lang="ts">
    const props = defineProps<{
        resetAllPoints: () => void
    }>()
</script>

<template>
    <div class="reset-all-points-wrapper">
        <button type="button" @click="props.resetAllPoints" >
            Reset All Points
        </button>
    </div>
</template>

<style scoped lang="scss">
    .reset-all-points-wrapper{
        text-align: center;
        margin: 25px;

        button{
            border-radius: 10px;
            border: none;

            padding: 10px 18px;
            font-size: 16px;
            font-weight: 700;
            transition: 0.3s;
            
            background-color: red;
            color: aliceblue;

            &:hover{
                cursor: pointer;
                background-color: rgba($color: #e30909, $alpha: .8);
            }
            
            &:active{
                transform: translateY(-5px);
            }
        }
    }
</style>