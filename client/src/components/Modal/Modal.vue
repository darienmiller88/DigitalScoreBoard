<script setup lang="ts">
    const props = defineProps<{ 
        modalHeader: string
        show: boolean
        onHide: () => void
        modalContent: any
    }>()

    const closeModal = (e: Event) => {
        console.log(e);
        
        if(e.target == props.modalContent){
           props.onHide()
        }
    }
</script>

<template>
    <div :class="`modal ${show ? 'slidein' : 'slideout'}`" @click="closeModal" tabindex="0" role="button" >
        <div class="modal_body">
            <div class="modal_header">
                <div class="filler"></div>
                <div class="header">{{ props.modalHeader }}</div>
                <span class="close" @click="onHide" role="button" tabindex="0" >&times;</span>
            </div>
            <div class="modal_content">
                {{ props.modalContent }}
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
        width: 100%;
        height: 100%;

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
                    // border: 2px solid red;
                    font-size: 17px;
                    text-align: center;

                    @media screen and (min-width: 768px) {
                        font-size: 25px;
                    }
                }

                .close{
                    // border: 2px solid saddlebrown;
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
                // border: 2px solid yellow;
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

        @media screen and (min-width: 768px) {
            .modal_body{
                width: 80vw;
            }
        }

        @media screen and (min-width: 992px) {
            .modal_body{
                width: 50vw;
                // height: 100vh;
            }
        }
    }
</style>