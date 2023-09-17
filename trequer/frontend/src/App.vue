<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';

import AppTitle from './components/AppTitle.vue';
import { useSongStore } from './stores/song';

const song = useSongStore();
song.patterns = {
  0: {
    tracks: {
      0: {
        length: 64,
        notes: {
          0: { key: 57, volume: 64, length: 2 },
          4: { key: 64, volume: 64, length: 2 },
          8: { key: 72, volume: 64, length: 2 },
          12: { key: 64, volume: 64, length: 2 },
          16: { key: 57, volume: 64, length: 2 },
          20: { key: 64, volume: 64, length: 2 },
          24: { key: 72, volume: 64, length: 2 },
          28: { key: 64, volume: 64, length: 2 },
          32: { key: 57, volume: 64, length: 2 },
          36: { key: 64, volume: 64, length: 2 },
          40: { key: 72, volume: 64, length: 2 },
          44: { key: 64, volume: 64, length: 2 },
          48: { key: 57, volume: 64, length: 2 },
          52: { key: 64, volume: 64, length: 2 },
          56: { key: 72, volume: 64, length: 2 },
          60: { key: 64, volume: 64, length: 2 },
        },
      },
    },
  },
};
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
