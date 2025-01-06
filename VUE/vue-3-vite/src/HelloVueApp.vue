<script setup>
import { ref } from 'vue';

const message = ref('Hello Vue!!');

function sayHello() {
  alert(message.value);
}
</script>

<template>
  <div>{{ message }}</div>
  <button @click="sayHello">Say Hello</button>
</template>