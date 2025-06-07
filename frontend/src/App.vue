<template>
    <div class="flex h-screen overflow-hidden p-0">
        <UsuariosModal
            @cerrar="mostrarListaUsuarios = false"
            :visible="mostrarListaUsuarios"
        />
        <FileSendModal
            @cerrar="mostrarEnviarArchivo = false"
            :visible="mostrarEnviarArchivo"
        />
        <ShowNotificatios
            @cerrar="mostrarNotificacion = false"
            :visible="mostrarNotificacion"
        />
        <Sidebar @toggle-users="toggleDrawer" />
        <main class="flex-1 flex flex-col">
            <header
                class="flex items-center justify-between bg-orange-600 text-white px-4 py-2"
            >
                <div class="flex items-center gap-3">
                    <img
                        src="./assets/images/atcnea_w.svg"
                        alt="Logo"
                        class="h-8 w-8"
                    />
                    <span class="font-bold text-lg">Aula Tecnológica</span>
                </div>
                <!-- <div class="flex items-center gap-4">
          <router-link to="/" class="mr-4 text-blue-600 underline">Asignaturas</router-link>
          <router-link to="/class" class="text-blue-600 underline">Clases</router-link>
        </div> -->
                <div class="flex items-center gap-4">
                    <span>Profesor</span>
                    <button
                        @click="toggleDrawer"
                        class="mdi mdi-account-multiple-outline text-2xl"
                    ></button>
                </div>
            </header>
            <router-view></router-view>
        </main>
        <UserDrawer v-model:visible="drawerVisible" />
        <Notification ref="notifier" />
    </div>
</template>

<script setup>
import { ref, provide } from "vue";

import Sidebar from "./components/Sidebar.vue";
import UserDrawer from "./components/UserDrawer.vue";
import UsuariosModal from "./components/modals/UsuariosModal.vue";
import FileSendModal from "./components/modals/FileSendModal.vue";
import ShowNotificatios from "./components/modals/ShowNotificatios.vue";
import Notification from "./components/Notificacion.vue";

const notifier = ref(null);
const drawerVisible = ref(false);
const mostrarListaUsuarios = ref(false);
const mostrarEnviarArchivo = ref(false);
const mostrarNotificacion = ref(false);

const cambiarMostrarNotificacion = () => {
    mostrarNotificacion.value = !mostrarNotificacion.value;
};

const cambiarMostrarEnviarArchivo = () => {
    console.log("cambiarMostrarEnviarArchivo");
    mostrarEnviarArchivo.value = !mostrarEnviarArchivo.value;
};

const cambiarMostrarListaUsuarios = () => {
    mostrarListaUsuarios.value = !mostrarListaUsuarios.value;
};

provide("mostrarEnviarArchivo", cambiarMostrarEnviarArchivo);
provide("mostrarListaUsuarios", cambiarMostrarListaUsuarios);
provide("mostrarNotificacion", cambiarMostrarNotificacion);

function toggleDrawer() {
    drawerVisible.value = !drawerVisible.value;
}

function notify(message, type) {
    notifier.value?.showNotification(message, type);
}
</script>
