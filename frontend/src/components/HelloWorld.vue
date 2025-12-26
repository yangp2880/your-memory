<script setup>
import { ref } from 'vue'

defineProps({
  msg: String
})

const count = ref(0)
</script>

<template>
  <h1>{{ msg }}</h1>
  <div>
	  这是一个组件
  </div>
</template>

<style scoped>
a {
  color: #42b983;
}
</style>
