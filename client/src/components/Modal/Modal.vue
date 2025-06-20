<script setup lang="ts">
    import { ref } from 'vue';

    let props = defineProps<{ 
        modalHeader: string
        show: boolean
        onHide: () => void
        modalContent: any
        modalProps: Object
    }>()

    const modalRef = ref<HTMLElement | null>(null);
    const closeModal = (e: Event) => {
        if(e.target == modalRef.value){
            props.onHide()       
        }
    }
</script>

<template>
    <div :class="`modal ${show ? 'slidein' : 'slideout'}`" ref="modalRef"  @click="closeModal" >
        <div class="modal_body">
            <div class="modal_header">
                <div class="filler"></div>
                <div class="header">{{ modalHeader }}</div>
                <span class="close" @click="onHide" role="button" tabindex="0" >&times;</span>
            </div>
            <div class="modal_content">
                <component :is="modalContent" v-bind="modalProps" />
            </div>
            <div class="modal_footer" >
                <button @click="onHide">Close</button>
            </div>
        </div>
    </div>
</template>

<style lang="scss">
    $borderColor: rgb(200, 200, 200);
    $bright_blue: rgb(0, 150, 255);
    $border_radius: 5px;

    .slidein{
        left: 0;
    }

    .slideout{
        left: 100vw;
    }

    .modal{
        display: grid;
        position: fixed;
        top: 0;
        z-index: 2;
        width: 100vw;
        height: 100vh;

        background-color: rgba($color: #000000, $alpha: 0.4);
        transition: 0.3s;

        .modal_body{
            position: relative;
            margin: auto;
            width: 95vw;
            border-radius: 10px;

            .modal_header{
                display: grid;
                grid-template-columns: auto 90% auto;
                // justify-content: space-between;
                // border: 2px red solid;
                border-bottom: 1px solid $borderColor;
                background-color: white;
                padding: 10px 15px;
                border-top-left-radius: $border_radius;
                border-top-right-radius: $border_radius;

                .header{
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    font-size: 22px;
                    text-align: center;

                    @media (min-width: 768px) {
                        font-size: 30px;
                    }
                }

                .close{
                    text-align: right;
                    transition: 0.5s;
                    font-size: 35px;
                }

                .close:hover{
                    cursor: pointer;
                    color: gray;
                }
            }

            .modal_content{
                background-color: white;
                padding: 10px;
                border-bottom: 1px solid $borderColor;
            }

            .modal_footer{
                display: flex;
                justify-content: center;
                padding: 15px 10px;
                background-color: white;
                border-bottom-left-radius: $border_radius;
                border-bottom-right-radius: $border_radius;

                button{
                    background-color: $bright_blue;
                    border: none;
                    border-radius: 5px;
                    padding: 10px 25px;
                    color: white;
                    font-size: 15px;
                    transition: 0.5s;
                }

                button:hover{
                    cursor: pointer;
                    background-color: grayscale($color: $bright_blue);
                }
            }
        }

        @media (min-width: 768px) {
            .modal_body{
                width: 80vw;
            }
        }

        @media (min-width: 992px) {
            .modal_body{
                width: 50vw;
            }
        }
    }
</style>