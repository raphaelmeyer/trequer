<script setup lang="ts">
import { GetMidiPorts } from '../../wailsjs/go/trequer/Trequer';

import { onMounted, ref } from 'vue';

const ports = ref<string[]>([]);

onMounted(() => {
  GetMidiPorts().then((value) => {
    ports.value = value.map((port) => port.address + ' | ' + port.client + ' | ' + port.port);
  });
});
</script>

<template>
  <div v-for="port in ports" :key="port">{{ port }}</div>
</template>
