<script setup lang="ts">
import { ref } from 'vue'

defineProps<{ msg: string }>()

const count = ref(0)
</script>

<template>
  <h1 class="text-3xl font-bold text-center text-green-500">{{ msg }}</h1>

</template>

<style scoped></style>
