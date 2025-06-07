<template>
    <div v-if="visible" class="overlay">
        <div class="modal">
            <h2>Usuarios conectados</h2>

            <table>
                <thead>
                    <tr>
                        <th>Nombre</th>
                        <th>Acciones</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(user, index) in usuarios" :key="index">
                        <td>{{ user.nombre }}</td>
                        <td>
                            <button @click="editar(user)">Editar</button>
                            <button @click="eliminar(index)">Eliminar</button>
                        </td>
                    </tr>
                </tbody>
            </table>

            <button class="cerrar" @click="$emit('cerrar')">Cerrar</button>
        </div>
    </div>
</template>

<script setup>
import { ref } from "vue";

defineProps(["visible"]);

const usuarios = ref([
    { nombre: "Juan" },
    { nombre: "Ana" },
    { nombre: "Pedro" },
]);

function editar(user) {
    alert(`Editar: ${user.nombre}`);
}

function eliminar(index) {
    usuarios.value.splice(index, 1);
}
</script>

<style scoped>
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
