<template>
    <div v-if="visible" class="overlay">
        <div class="modal">
            <h2>Usuarios conectados</h2>

            <table>
                <thead>
                    <tr>
                        <th>
                            <input
                                type="checkbox"
                                v-model="allSelected"
                                @change="selectAll"
                            />
                            Seleccionar todos
                        </th>
                        <th>Nombre</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(user, index) in usuarios" :key="index">
                        <td>
                            <input
                                type="checkbox"
                                v-model="user.selected"
                                @change="changeSelectedAll"
                            />
                        </td>
                        <td>{{ user.nombre }}</td>
                    </tr>
                </tbody>
            </table>
            <span>{{ file ? file : "No hay archivo seleccionado" }} </span>
            <button @click="seleccionarArchivos">Seleccionar Archivos</button>
            <button @click="enviar">Enviar</button>
            <button class="cerrar" @click="$emit('cerrar')">Cerrar</button>
        </div>
    </div>
</template>

<script setup>
import { ref } from "vue";
import { OpenFile } from "../../../wailsjs/go/main/App";

defineProps(["visible"]);
const allSelected = ref(false);
const file = ref(null);

const usuarios = ref([
    { nombre: "Juan", selected: false },
    { nombre: "Ana", selected: false },
    { nombre: "Pedro", selected: false },
]);
function seleccionarArchivos() {
    OpenFile().then((result) => {
        console.log(result);
        file.value = result;
        file.value = file.value.split("/").pop();
    });
}
function enviar() {
    console.log("Enviando archivos...");
}
function eliminar(index) {
    usuarios.value.splice(index, 1);
}
function selectAll() {
    usuarios.value.forEach((user) => (user.selected = allSelected.value));
}
function changeSelectedAll() {
    allSelected.value = !usuarios.value.some((user) => !user.selected);
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
