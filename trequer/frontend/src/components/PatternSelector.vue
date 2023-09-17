<script setup lang="ts">
import { mdiArrowLeftCircle, mdiArrowRightCircle } from '@mdi/js';

const pattern = 0;
</script>

<template>
  <div class="flex justify-between gap-2 bg-sky-950 border-sky-600 border-4 rounded-xl p-1 w-32">
    <svg class="w-6 h-6" viewBox="0 0 24 24">
      <path :d="mdiArrowLeftCircle" class="fill-sky-200" />
    </svg>
    <div class="font-bold">{{ pattern }}</div>
    <svg class="w-6 h-6" viewBox="0 0 24 24">
      <path :d="mdiArrowRightCircle" class="fill-sky-200" />
    </svg>
  </div>
</template>
