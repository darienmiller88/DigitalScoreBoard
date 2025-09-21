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

        button{
            border-radius: 10px;
            border: none;

            padding: 10px 18px;
            font-size: 16px;
            font-weight: 700;
            transition: 0.3s;
            
            background-color: red;
            color: aliceblue;

            @media (min-width: 3840px) {
                font-size: 60px;
                padding: 20px 36px;
            }

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