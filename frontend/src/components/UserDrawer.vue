<template>
  <transition name="slide">
    <aside
      v-if="visible"
      class="fixed right-0 top-0 h-full w-72 bg-white shadow-lg border-l border-gray-300 p-4 overflow-auto"
    >
      <header class="flex justify-between items-center mb-4">
        <h2 class="text-xl font-semibold">Usuarios Conectados</h2>
        <button class="mdi mdi-close text-2xl" @click="$emit('update:visible', false)"></button>
      </header>
      <div class="space-y-2">
        <UserItem
          v-for="user in users"
          :key="user.id"
          :user="user"
          @send-file="onSendFile"
          @chat="onChat"
          @kick="onKick"
        />
      </div>
    </aside>
  </transition>
</template>

<script lang="ts" setup>
import { defineProps, defineEmits } from 'vue'
import UserItem from './UserItem.vue'

const props = defineProps({
  visible: Boolean
})
const emit = defineEmits(['update:visible'])

const users = [
  { id: 1, name: 'Ana Pérez', icon: 'mdi-account-circle' },
  { id: 2, name: 'Carlos Gómez', icon: 'mdi-account-circle' },
  { id: 3, name: 'Lucía Martínez', icon: 'mdi-account-circle' }
]

function onSendFile(user: any) {
  alert(`Enviar archivo a ${user.name}`)
}

function onChat(user: any) {
  alert(`Chatear con ${user.name}`)
}

function onKick(user: any) {
  alert(`Expulsar a ${user.name}`)
}
</script>

<style>
.slide-enter-active,
.slide-leave-active {
  transition: transform 0.3s ease;
}
.slide-enter-from,
.slide-leave-to {
  transform: translateX(100%);
}
</style>