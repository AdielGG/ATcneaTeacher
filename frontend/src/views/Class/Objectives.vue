<template>
    <div class="p-6 max-w-2xl mx-auto bg-white rounded-xl shadow-md space-y-6">
        <div>
            <h2 class="text-xl font-bold text-gray-800">Objetivos</h2>
            <ul class="list-disc list-inside text-gray-700 mt-2">
                <li v-for="(item, index) in objectives" :key="index">
                    {{ item }}
                </li>
            </ul>
        </div>

        <div>
            <h2 class="text-xl font-bold text-gray-800">Bibliografía</h2>
            <ul class="list-decimal list-inside text-gray-700 mt-2">
                <li v-for="(ref, index) in bibliography" :key="index">
                    <a
                        v-if="isURL(ref)"
                        :href="ref"
                        class="text-blue-600 underline"
                        target="_blank"
                        rel="noopener noreferrer"
                        >{{ ref }}</a
                    >
                    <span v-else>{{ ref }}</span>
                </li>
            </ul>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { defineProps } from "vue";

interface Props {
    objectives: string[];
    bibliography: string[];
}

const props = defineProps<Props>();

const isURL = (str: string): boolean => {
    try {
        const url = new URL(str);
        return ["http:", "https:"].includes(url.protocol);
    } catch (_) {
        return false;
    }
};
</script>
