<template>
  <div class="flex items-center justify-between p-2 border rounded hover:bg-gray-100">
    <div class="flex items-center gap-3">
      <span :class="['mdi', 'mdi-account', 'text-3xl']"></span>
      <span class="font-medium">USUARIO</span>
    </div>
    <div class="flex items-center gap-2 text-gray-600">
      <button class="mdi mdi-invoice-text-send-outline text-xl" @click="onSendFile(user)"></button>
      <button class="mdi mdi-chat-outline text-xl" @click="$emit('chat', user)"></button>
      <button class="mdi mdi-account-remove-outline text-xl" @click="$emit('kick', user)"></button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { OpenFile } from '../../wailsjs/go/main/App'
import { SendFile } from '../../wailsjs/go/streaming/FileStream'
import { streaming } from '../../wailsjs/go/models'
import { defineProps } from 'vue'

async function onSendFile(user: any) {
    
    const filePath = await OpenFile()
    if (filePath == "") {
      return
    }

    const parts = filePath.split('/');
    const filename = parts[parts.length - 1];
    
    const response = await fetch(filePath)
    const arrayBuffer = await response.arrayBuffer()
    const fileData = Array.from(new Uint8Array(arrayBuffer))


    SendFile(user.id, filename,fileData) 
  }

function onChat(user: any) {
  alert(`Chatear con ${user.name}`)
}

function onKick(user: any) {
  alert(`Expulsar a ${user.name}`)
}

const props = defineProps({
  user: Object as () => { id: number; name: string; icon: string }
})


</script>
