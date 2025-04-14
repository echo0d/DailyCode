

<script>
export default {
//     mounted () {
//         // vue2.x 获取组件上的属性
//         console.log(this.$attrs.subData);  // 'some other data'
        
//         // vue2.x 获取slot插槽的节点
//         console.log(this.$slots);

//     },
//     methods: {
//         // vue2.x emit方法(子组件向父组件传递数据)
//         handleClick () {
//             this.$emit('subClick', 'vue2.x - this is subData')
//         },
//     },
    setup (props, context) {
        let { attrs, slots, emit } = context;

        // vue3.x 获取组件上的属性
        console.log(attrs.subData);  // 'some other data'

        // vue3.x 获取slot插槽的节点
        console.log(slots.default());

        // vue3.x emit方法(子组件向父组件传递数据)
        function handleClick () {
            emit('subClick', 'vue3.x - this is subData');
        }

        return { handleClick }
    }
}
</script>

<template>
    <!-- 父组件向子组件传递数据 -->
    <div @click="handleClick">Child</div>
</template>