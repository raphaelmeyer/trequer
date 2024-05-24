import './assets/main.css';

import { createApp } from 'vue';
import { createPinia } from 'pinia';

import App from './App.vue';
import router from './router';

import { EventsOn } from '../wailsjs/runtime';

import { useMidiPortsStore } from './stores/midiport';

const app = createApp(App);

app.use(createPinia());
app.use(router);

app.mount('#app');

const midiports = useMidiPortsStore();
midiports.refresh();

EventsOn('ports-changed', () => midiports.refresh());
