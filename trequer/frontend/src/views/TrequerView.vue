<script setup lang="ts">
import { useMidiPortsStore } from '@/stores/midiport';

import { computed } from 'vue';

const midiports = useMidiPortsStore();
const ports = computed(() => midiports.ports);
</script>

<template>
  <div>Trequer</div>
  <div v-for="port in ports" :key="port">{{ port }}</div>
</template>
