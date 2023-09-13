<script setup lang="ts">
import { keyName } from '@/models/midi';

const pattern = Array.from({ length: 64 }, () => '... ...');

pattern[0] = keyName(69);
pattern[4] = keyName(73);
pattern[8] = keyName(21);
pattern[16] = keyName(127);
</script>

<template>
  <div v-for="(note, index) in pattern" :key="index">{{ note }}</div>
</template>
