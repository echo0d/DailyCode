<script>
import { toRefs } from 'vue';
export default {
    props: {
        name: String,
        age: Number
    },
    // mounted () {
    //     // vue2.x 的写法
    //     console.log(this.name); // 张三
    //     console.log(this.age); // 20
    // },
    setup (props) {
        // vue3.x 的写法
        console.log(props.name); // 张三
        console.log(props.age); // 20
        
        // let { name ,age } = props;  // 不能直接解构
        let { name, age } = toRefs(props);
        console.log(name.value, age.value); // 张三 20
    }
}
</script>

<template>
    <div>{{name}}{{age}}</div>
</template>