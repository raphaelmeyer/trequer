<script setup lang="ts">
const props = defineProps<{ label: string }>();
</script>

<template>
  <div class="flex flex-col bg-sky-950 border-sky-600 border-4 rounded-xl p-1 gap-y-2">
    <div class="text-center">{{ props.label }}</div>
    <slot></slot>
  </div>
</template>
