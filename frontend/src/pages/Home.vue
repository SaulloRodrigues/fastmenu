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
    <div class="max-w-5xl mx-auto p-4 h-fit bg-black/10">
        <div class="flex flex-row gap-4">
            <div class="w-fit h-fit py-4 bg-black-85" v-for="product in products" :key="product.id">
                <img :src="product.image_url" alt="Product Image" class="w-40 h-40 object-cover mb-2">
                <p>{{ product.name }}</p>
                <p>{{ product.description }}</p>
                <p>{{ product.price }} R$</p>
            </div>
        </div>
    </div>
</template>