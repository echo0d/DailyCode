<template>
    <!-- 父组件向子组件传递数据 -->
    <Sub :name="name" :age="age" />
</template>

<script>
import { ref } from 'vue'
import Sub from './Sub.vue'
export default {
    setup () {
        const name = ref('张三');
        const age = ref(20)

        return { name, age }
    },
    components: { Sub },
}
</script>