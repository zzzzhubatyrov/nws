import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import { createRouter, createWebHistory } from 'vue-router'
import AuthView from './view/AuthView.vue'
import HomeView from './view/HomeView.vue'
import ChartsView from './view/ChartsView.vue'
import LastDataView from "./view/LastDataView.vue";
import InventoryView from "./view/InventoryView.vue";
import ReportsView from "./view/ReportsView.vue";
import ConfigurationView from "./view/ConfigurationView.vue";

const router = createRouter({
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView // AuthView
        },
        {
            path: '/charts',
            name: 'charts',
            component: ChartsView
        },
        {
            path: "/latest",
            name: 'latest',
            component: LastDataView,
        },
        {
            path: "/inventory",
            name: 'inventory',
            component: InventoryView,
        },
        {
            path: "/requests",
            name: 'requests',
            component: ReportsView,
        },
        {
            path: "/config",
            name: 'config',
            component: ConfigurationView,
        }
        // {
        //     path: '/index',
        //     name: 'home',
        //     component: HomeView
        // },
    ],
    history: createWebHistory(),
})


const app = createApp(App)
app.use(router)
app.mount('#app')
