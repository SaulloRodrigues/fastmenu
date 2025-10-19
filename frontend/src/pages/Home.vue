<script setup lang="ts">
import { ref, onMounted } from 'vue'

interface Product {
    id: number;
    name: string;
    description: string;
    image_url: string;
    price: number;
}


const products = ref<Product[]>([])

onMounted(async () => {
    try {
        const res = await fetch('/api/products');
        if (!res.ok) {
            const errText = await res.text();
            console.error('Erro da API:', res.status, errText);
            return;
        }
        const data = await res.json();
        products.value = data;
    } catch (err) {
        console.error('Erro inesperado:', err);
    }
})
</script>

<template>
    <div class="max-w-6xl mx-auto p-4 h-fit bg-black/10 rounded-2xl">
        <div class="flex flex-row flex-wrap justify-between gap-4">
            <div class="flex flex-col space-y-1 w-80 h-fit p-5 bg-black/5 rounded-2xl" v-for="product in products"
                :key="product.id">
                <div class="flex items-center justify-center w-full h-fit">
                    <img :src="product.image_url" alt="Product Image"
                        class="w-40 h-fit object-cover mb-4" />
                </div>
                <span class="font-bold line-clamp-1">{{ product.name }}</span>
                <div class="w-full h-fit text-sm font-medium text-gray-700 text-justify max-h-[12.5rem] overflow-hidden">
                    <p class="line-clamp-6">{{ product.description }}</p>
                </div>
                <p>{{ product.price }} R$</p>
            </div>
        </div>
    </div>
</template>
