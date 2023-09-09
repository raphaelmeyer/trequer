<script setup lang="ts">
const pattern = Array.from({ length: 64 }, () => '... ...');
</script>

<template>
  <div v-for="(note, index) in pattern" :key="index">{{ note }}</div>
</template>
