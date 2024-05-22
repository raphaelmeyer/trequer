import './assets/main.css';

import { createApp } from 'vue';
import { createPinia } from 'pinia';

import App from './App.vue';
import router from './router';

const app = createApp(App);

app.use(createPinia());
app.use(router);

app.mount('#app');

import { EventsOn } from '../wailsjs/runtime';

import { useMidiPortsStore } from './stores/midiport';

const midiports = useMidiPortsStore();
midiports.refresh();

EventsOn('ports-changed', () => midiports.refresh());
