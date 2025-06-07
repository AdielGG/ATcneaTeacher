<template>
  <nav class="w-16 bg-orange-600 text-white flex flex-col items-center py-4 space-y-6">
    <button class="mdi mdi-arrow-left-bold text-3xl" @click="$router.back()"></button>
    <button class="mdi mdi-plus text-3xl"></button>
    <button class="mdi mdi-download text-3xl"></button>
    <button class="mdi mdi-account-group-outline text-3xl" @click="$emit('toggle-users')"></button>
  </nav>
</template>

<script lang="ts" setup></script>