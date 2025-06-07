<template>
    <div class="action-bar">
        <button
            v-for="(accion, index) in acciones"
            :key="index"
            class="action-btn"
            :title="accion.nombre"
            @click="() => ejecutar(index)"
        >
            <i :class="`mdi mdi-${accion.icono}`" class="icono"></i>
        </button>
    </div>
</template>

<script setup>
import { inject } from "vue";

const mostrarListaUsuarios = inject("mostrarListaUsuarios");
const mostrarEnviarArchivo = inject("mostrarEnviarArchivo");
const mostrarNotificacion = inject("mostrarNotificacion");

const acciones = [
    {
        nombre: "Informacion",
        icono: "book-information-variant",
        accion: function () {
            console.log("Informacion");
        },
    },
    {
        nombre: "Lista de usuarios",
        icono: "account-file-text",
        accion: function () {
            mostrarListaUsuarios();
            console.log("Lista de usuarios");
        },
    },
    {
        nombre: "Lista de PDF",
        icono: "presentation-play",
        accion: function () {
            console.log("Lista de PDF");
        },
    },
    {
        nombre: "Enviar archivo",
        icono: "file-upload-outline",
        accion: function () {
            mostrarEnviarArchivo();
            console.log("Enviar archivo");
        },
    },
    {
        nombre: "Enviar tarea",
        icono: "clipboard-text",
        accion: function () {
            console.log("Enviar tarea");
        },
    },
    {
        nombre: "Pizarra",
        icono: "developer-board",
        accion: function () {
            console.log("Pizarra");
        },
    },
    {
        nombre: "Videos",
        icono: "video",
        accion: function () {
            console.log("Videos");
        },
    },
    {
        nombre: "Chat General",
        icono: "chat",
        accion: function () {
            mostrarNotificacion();
            console.log("Chat General");
        },
    },
];

function ejecutar(index) {
    acciones[index].accion();
}
</script>

<style scoped>
.action-bar {
    display: flex;
    background-color: #f5f5f5; /* blanco hueso */
    border-bottom: 1px solid #ddd;
    align-items: center;
    gap: 16px;
}

.action-btn {
    display: flex;
    align-items: center;
    gap: 6px;
    background: none;
    border: none;
    cursor: pointer;
    color: #ff6d00;
    font-size: 14px;
    padding: 6px 10px;
    border-radius: 4px;
    transition: background 0.2s;
}

.action-btn:hover {
    background-color: rgba(255, 109, 0, 0.1);
}

.icono {
    font-size: 20px;
}
</style>
