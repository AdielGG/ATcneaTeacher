<script lang="ts" setup>
import { ref } from "vue";
import * as pdfjsLib from "pdfjs-dist/legacy/build/pdf";
import {
    BroadcastPDF,
    BroadcastPage,
} from "../../wailsjs/go/streaming/PDFStream";
import ActionBar from "../components/ActionBar.vue";

// Configura la URL pública del worker
pdfjsLib.GlobalWorkerOptions.workerSrc = "/pdf.worker.min.js";

const pdfCanvas = ref<HTMLCanvasElement | null>(null);
const currentPage = ref(1);
const scale = ref(1.0);
let loadingTask: any;
const totalPages = ref(0);
let pdf: any;

async function renderPDF(pageNum: number) {
    if (!pdfCanvas.value) return;

    BroadcastPage(pageNum);

    const page = await pdf.getPage(pageNum);
    const viewport = page.getViewport({ scale: scale.value });
    const canvas = pdfCanvas.value;
    const context = canvas.getContext("2d");

    canvas.width = viewport.width;
    canvas.height = viewport.height;

    await page.render({ canvasContext: context!, viewport }).promise;
}

function onFileChange(event: Event) {
    const input = event.target as HTMLInputElement;
    if (!input.files || input.files.length === 0) return;

    const file = input.files[0];
    const reader = new FileReader();

    reader.onload = async () => {
        if (reader.result instanceof ArrayBuffer) {
            loadingTask = pdfjsLib.getDocument({ data: reader.result });
            pdf = await loadingTask.promise;

            totalPages.value = pdf.numPages;
            renderPDF(1);
            BroadcastPDF(Array.from(new Uint8Array(reader.result)));
        }
    };
    reader.readAsArrayBuffer(file);
}

function nextPage() {
    if (currentPage.value < totalPages.value) {
        currentPage.value++;
        renderPDF(currentPage.value);
    }
}

function prevPage() {
    if (currentPage.value > 1) {
        currentPage.value--;
        renderPDF(currentPage.value);
    }
}

function zoomIn() {
    scale.value += 0.2;
    renderPDF(currentPage.value);
}

function zoomOut() {
    scale.value = Math.max(0.5, scale.value - 0.2);
    renderPDF(currentPage.value);
}
</script>

<template>
    <div class="flex flex-col h-screen bg-gray-100 space-y-4">
        <ActionBar />
        <RouterView />
    </div>
</template>
