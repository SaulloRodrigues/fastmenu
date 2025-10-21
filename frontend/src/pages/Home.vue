<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { Button } from '@/components/ui/button'

interface Product {
    id: number;
    name: string;
    description: string;
    image_url: string;
    price: number;
    categories?: Category[];
}

interface Category {
    id: number;
    name: string;
}

const products = ref<Product[]>([])
const loading = ref(false)

// Buscar todos os produtos
const fetchAllProducts = async () => {
    loading.value = true;
    try {
        const res = await fetch('/api/products');
        if (!res.ok) {
            console.error('Erro ao buscar produtos:', res.status);
            return;
        }
        const data = await res.json();
        products.value = data;
    } catch (err) {
        console.error('Erro inesperado ao buscar produtos:', err);
    } finally {
        loading.value = false;
    }
}

// Agrupar produtos por categoria
const productsByCategory = computed(() => {
    const grouped: { [key: string]: Product[] } = {};
    
    products.value.forEach(product => {
        if (product.categories && product.categories.length > 0) {
            product.categories.forEach(category => {
                if (!grouped[category.name]) {
                    grouped[category.name] = [];
                }
                grouped[category.name].push(product);
            });
        } else {
            // Produtos sem categoria vão para "Outros"
            if (!grouped['Outros']) {
                grouped['Outros'] = [];
            }
            grouped['Outros'].push(product);
        }
    });
    
    return grouped;
});

onMounted(async () => {
    await fetchAllProducts();
})
</script>

<template>
    <div class="max-w-7xl mx-auto py-4">
        <!-- Loading State -->
        <div v-if="loading" class="flex justify-center items-center py-12">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-black"></div>
        </div>

        <!-- Produtos Agrupados por Categoria -->
        <div v-else class="space-y-12">
            <div v-for="(categoryProducts, categoryName) in productsByCategory" :key="categoryName" class="space-y-6">
                <!-- Título da Categoria -->
                <div class="text-start">
                    <h2 class="text-3xl font-bold text-gray-900 mb-2">{{ categoryName.charAt(0).toUpperCase() + categoryName.slice(1) }}</h2>
                </div>

                <!-- Grid de Produtos da Categoria -->
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
                    <div class="group relative bg-white rounded-2xl shadow-lg hover:shadow-xl transition-all duration-300 overflow-hidden border border-gray-200 hover:border-gray-300 hover:-translate-y-1" 
                         v-for="product in categoryProducts" :key="product.id">
                        
                        <!-- Image Container -->
                        <div class="relative h-48 bg-gray-50 overflow-hidden">
                            <img :src="product.image_url" alt="Product Image"
                                class="w-full h-full object-contain transition-transform duration-300 group-hover:scale-105" />
                            
                            <!-- Price Badge -->
                            <div class="absolute top-3 right-3 bg-white/90 backdrop-blur-sm rounded-full px-3 py-1.5 shadow-md">
                                <span class="text-sm font-bold text-gray-900">{{ product.price }} R$</span>
                            </div>
                        </div>

                        <!-- Content Container -->
                        <div class="p-5 space-y-3">
                            <!-- Product Title -->
                            <h3 class="text-lg font-semibold text-gray-900 leading-tight group-hover:text-blue-600 transition-colors duration-200">
                                {{ product.name }}
                            </h3>
                            
                            <!-- Description -->
                            <p class="text-sm text-gray-600 leading-relaxed line-clamp-3 group-hover:text-gray-700 transition-colors duration-200">
                                {{ product.description }}
                            </p>

                            <!-- Action Button -->
                            <div class="pt-2">
                                <Button 
                                    variant="default" 
                                    size="lg" 
                                    class="w-full bg-black hover:bg-gray-800 text-white font-medium transition-colors duration-200"
                                >
                                    Adicionar ao Carrinho
                                </Button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
