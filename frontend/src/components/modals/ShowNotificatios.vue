<template>
    <div v-if="visible" class="overlay">
        <div class="modal">
            <Notification ref="notifier" />

            <div class="p-6 space-x-4">
                <button
                    @click="() => notify('Operación completada', 'success')"
                    class="btn px-4 py-2 rounded text-sm"
                >
                    Success
                </button>
                <button
                    @click="() => notify('Ha ocurrido un error', 'error')"
                    class="btn"
                >
                    Error
                </button>
                <button
                    @click="() => notify('Advertencia del sistema', 'warning')"
                    class="btn"
                >
                    Warning
                </button>
                <button
                    @click="() => notify('Mensaje informativo', 'info')"
                    class="btn"
                >
                    Info
                </button>
                <button @click="$emit('cerrar')" class="btn">cerrar</button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from "vue";
import Notification from "../Notificacion.vue";
defineProps(["visible"]);
const notifier = ref(null);

function notify(message, type) {
    notifier.value?.showNotification(message, type);
}
</script>

<style scoped>
.btn {
    /* @apply bg-gray-200 hover:bg-gray-300 text-sm px-4 py-2 rounded; */
    background-color: gray;
}
.overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.4);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
}

.modal {
    background-color: white;
    padding: 20px;
    border-radius: 8px;
    width: 400px;
}

.cerrar {
    margin-top: 10px;
}

table {
    width: 100%;
    margin-top: 10px;
    border-collapse: collapse;
}

td,
th {
    border-bottom: 1px solid #ccc;
    padding: 8px;
}

button {
    margin-right: 5px;
}
</style>
