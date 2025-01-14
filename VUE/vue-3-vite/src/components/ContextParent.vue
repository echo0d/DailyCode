
<script>
import ContextSub from './ContextSub.vue'
export default {
    setup () {

        function subClick (e) {
            console.log(e); // 接收子组件传递过来的数据
        }
        return { subClick }
    },
    components: { ContextSub },
}
</script>

<template>
    <ContextSub subData="some other data" @subClick='subClick'>parent</ContextSub>
</template>
