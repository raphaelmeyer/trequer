<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';

import AppTitle from './components/AppTitle.vue';
</script>

<template>
  <div class="bg-sky-700 text-sky-200 flex flex-col">
    <AppTitle />
    <div class="flex flex-row gap-2 p-2">
      <RouterLink to="/tracker">Tracker</RouterLink>
      <RouterLink to="/settings">Settings</RouterLink>
    </div>
    <RouterView />
  </div>
</template>
