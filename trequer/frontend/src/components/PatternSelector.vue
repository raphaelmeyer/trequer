<script setup lang="ts">
import { mdiArrowLeftCircle, mdiArrowRightCircle } from '@mdi/js';
import LabelBox from './LabelBox.vue';

const pattern = 0;
</script>

<template>
  <LabelBox label="Pattern" class="w-32">
    <div class="flex justify-between">
      <svg class="w-6 h-6" viewBox="0 0 24 24">
        <path :d="mdiArrowLeftCircle" class="fill-sky-200" />
      </svg>
      <div class="font-bold">{{ pattern }}</div>
      <svg class="w-6 h-6" viewBox="0 0 24 24">
        <path :d="mdiArrowRightCircle" class="fill-sky-200" />
      </svg>
    </div>
  </LabelBox>
</template>
