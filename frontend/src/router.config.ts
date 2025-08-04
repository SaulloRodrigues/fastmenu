import { createWebHistory, createRouter } from "vue-router";

import HomeView from './pages/Home.vue'
import NotFound from "./pages/NotFound.vue";

const routes = [
    { path: '/', component: HomeView },
    { path: '/:pathMatch(.*)*', component: NotFound }, // rota 404 corrigida
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router;
