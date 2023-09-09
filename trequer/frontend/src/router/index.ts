import { createRouter, createWebHashHistory } from 'vue-router';

import TrackerView from '../views/TrackerView.vue';
import SettingsView from '../views/SettingsView.vue';

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: '/', redirect: '/tracker' },
    { path: '/tracker', name: 'tracker', component: TrackerView },
    {
      path: '/settings',
      name: 'settings',
      component: SettingsView,
    },
  ],
});

export default router;
