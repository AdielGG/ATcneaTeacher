<template>
    <div class="p-6 w-full h-full bg-white shadow-md space-y-4">
        <h1 class="text-2xl font-bold text-gray-800">{{ title }}</h1>
        <h2 class="text-lg font-semibold text-gray-700">Objetivo</h2>
        <p class="text-gray-600">{{ objective }}</p>
        <h2 class="text-lg font-semibold text-gray-700">Descripción</h2>
        <p class="text-gray-600">{{ description }}</p>
        <RouterLink to="/obj" class="bg-blue-700 rounded-4xl decoration-0"
            >Next</RouterLink
        >
    </div>
</template>

<script lang="ts" setup>
import { defineProps } from "vue";
import router from "../../router/router";

interface Props {
    title: string;
    objective: string;
    description: string;
}

function Next() {}

const props = defineProps<Props>();
</script>
