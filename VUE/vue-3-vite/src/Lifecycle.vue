<script>
import { onBeforeMount, onMounted, onBeforeUpdate, onUpdated, onBeforeUnmount, onUnmounted, onRenderTracked, onRenderTriggered } from 'vue'
export default {
    setup () {
        console.log('setup');
        
        // 生命周期钩子(没有beforeCreate和created)
        onBeforeMount(() => { console.log('onBeforeMount'); })
        onMounted(() => { console.log('onMounted'); })
        onBeforeUpdate(() => { console.log('onBeforeUpdate'); })
        onUpdated(() => { console.log('onUpdated'); })
        onBeforeUnmount(() => { console.log('onBeforeUnmount'); })
        onUnmounted(() => { console.log('onUnmounted'); })

        // 新增的debug钩子  生产环境中会被忽略
        onRenderTracked(() => { console.log('onRenderTracked'); }) // 每次渲染后重新收集响应式依赖，在onMounted前触发，页面更新后也会触发
        onRenderTriggered(() => { console.log('onRenderTriggered'); }) // 每次触发页面重新渲染时的自动执行，在onBeforeUpdate之前触发
    },
    beforeCreate () {
        console.log('beforeCreate');
    },
    created () {
        console.log('created');
    }
}
</script>