<script lang="ts" setup>
import { ref, computed } from "vue";

const props = defineProps<{ mode: "subjects" | "classes" }>();

interface Class {
    id: number;
    title: string;
    date: string;
}

interface Subject {
    id: number;
    name: string;
    classes: Class[];
}

// Datos ejemplo
const subjects: Subject[] = [
    {
        id: 1,
        name: "Matemáticas",
        classes: [
            { id: 1, title: "Álgebra Básica", date: "2024-03-01" },
            { id: 2, title: "Geometría", date: "2024-03-03" },
        ],
    },
    {
        id: 2,
        name: "Historia",
        classes: [{ id: 3, title: "Edad Media", date: "2024-03-02" }],
    },
    {
        id: 3,
        name: "Física",
        classes: [
            { id: 4, title: "Física I", date: "2024-03-04" },
            { id: 5, title: "Física II", date: "2024-03-05" },
        ],
    },
];

const searchTerm = ref("");

const filteredSubjects = computed(() => {
    if (!searchTerm.value) return subjects;
    const term = searchTerm.value.toLowerCase();
    return subjects.filter((s) => s.name.toLowerCase().includes(term));
});

const allClasses = computed(() =>
    subjects.flatMap((subject) =>
        subject.classes.map((c) => ({
            ...c,
            subjectName: subject.name,
        })),
    ),
);

const filteredClasses = computed(() => {
    if (!searchTerm.value) return allClasses.value;
    const term = searchTerm.value.toLowerCase();
    return allClasses.value.filter(
        (c) =>
            c.title.toLowerCase().includes(term) ||
            c.subjectName.toLowerCase().includes(term),
    );
});
</script>

<template>
    <div class="p-4">
        <h1 class="text-2xl font-bold mb-4">
            {{
                props.mode === "classes" ? "Clases disponibles" : "Asignaturas"
            }}
        </h1>

        <input
            v-model="searchTerm"
            type="text"
            placeholder="Buscar..."
            class="mb-6 w-full p-2 border rounded"
        />

        <div
            v-if="props.mode === 'subjects'"
            class="grid grid-cols-2 gap-4 bg-white"
        >
            <div
                v-for="subject in filteredSubjects"
                :key="subject.id"
                class="p-4 bg-amber-200 shadow rounded cursor-pointer hover:bg-gray-100"
                @click="$router.push('/class')"
            >
                {{ subject.name }}
            </div>
        </div>

        <div v-else class="space-y-4 grid gap-4">
            <div
                v-for="cls in filteredClasses"
                :key="cls.id"
                class="p-4 bg-amber-200 shadow rounded cursor-pointer hover:bg-gray-100"
                @click="$router.push('/classroom')"
            >
                <h2 class="font-semibold mb-1">{{ cls.title }}</h2>
                <p class="text-sm text-gray-600">
                    Asignatura: {{ cls.subjectName }}
                </p>
                <p class="text-sm text-gray-600">Fecha: {{ cls.date }}</p>
            </div>
        </div>
    </div>
</template>
