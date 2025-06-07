<template>
    <div class="fixed bottom-4 right-4 z-50 flex flex-col gap-3">
        <div
            v-for="(notif, index) in notifications"
            :key="index"
            :class="[
                'rounded-lg shadow-md px-4 py-3 min-w-[250px] relative flex items-start gap-2',
                typeStyles[notif.type],
            ]"
        >
            <span class="text-xl">{{ typeIcons[notif.type] }}</span>
            <div class="flex-1 text-sm text-gray-800">{{ notif.message }}</div>
            <button
                class="text-gray-400 hover:text-gray-600 text-sm ml-2"
                @click="remove(index)"
            >
                ✖
            </button>
        </div>
    </div>
</template>

<script setup>
import { ref } from "vue";

const notifications = ref([]);

// Colores según tipo
const typeStyles = {
    info: "bg-blue-100 border-l-4 border-blue-500",
    success: "bg-green-100 border-l-4 border-green-500",
    error: "bg-red-100 border-l-4 border-red-500",
    warning: "bg-yellow-100 border-l-4 border-yellow-500",
};

// Íconos según tipo
const typeIcons = {
    info: "ℹ️",
    success: "✅",
    error: "❌",
    warning: "⚠️",
};

function showNotification(message, type = "info", duration = 5000) {
    const id = Date.now();
    notifications.value.push({ message, type, id });

    setTimeout(() => {
        const index = notifications.value.findIndex((n) => n.id === id);
        if (index !== -1) remove(index);
    }, duration);
}

function remove(index) {
    notifications.value.splice(index, 1);
}

defineExpose({ showNotification });
</script>
