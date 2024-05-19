<script setup lang="ts">
import { ListMidiOutputs } from '../../wailsjs/go/audiocontrol/AudioControl';

import { onMounted, ref } from 'vue';

const ports = ref<string[]>([]);

onMounted(() => {
  ListMidiOutputs().then((value) => {
    ports.value = value;
  });
});
</script>

<template>
  <div>Trequer</div>
  <div v-for="port in ports" :key="port">{{ port }}</div>
</template>
