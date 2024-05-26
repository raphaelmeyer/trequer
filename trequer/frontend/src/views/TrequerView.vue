<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref } from 'vue';

import { useMidiPortsStore } from '@/stores/midiport';

import { EventsOn, EventsOff } from '../../wailsjs/runtime/runtime';
import { Play, Stop } from '../../wailsjs/go/main/App';

const midiports = useMidiPortsStore();
const ports = computed(() => midiports.ports);

const currentBeat = ref(0);

onMounted(() => {
  EventsOn('on-beat', (beat) => {
    currentBeat.value = beat;
  });
});

onUnmounted(() => {
  EventsOff('on-beat');
});
</script>

<template>
  <div>Trequer</div>
  <div v-for="port in ports" :key="port">{{ port }}</div>

  <div>
    <button type="button" @click="Play">Play</button>
    <button type="button" @click="Stop">Stop</button>
  </div>

  <div>beat {{ currentBeat }}</div>
</template>
