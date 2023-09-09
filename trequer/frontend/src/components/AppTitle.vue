<script setup lang="ts">
import { mdiMusic } from '@mdi/js';
</script>

<template>
  <div class="flex gap-4 p-2">
    <svg class="w-8 h-8" viewBox="0 0 24 24">
      <path :d="mdiMusic" class="fill-sky-200" />
    </svg>
    <div class="text-xl font-bold">Trequer</div>
  </div>
</template>
