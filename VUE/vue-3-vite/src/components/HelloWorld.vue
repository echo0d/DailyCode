<template>
  <div>
    <p>{{ message }}</p>
    <p>{{ anotherMessage }}</p>
    <button @click="changeMessage">Change Message</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: 'Hello World!',
      anotherMessage: 'Welcome to Vue.js!'
    }
  },
  methods: {
    changeMessage() {
      this.message = 'Message Changed!';
    }
  }
}
</script>