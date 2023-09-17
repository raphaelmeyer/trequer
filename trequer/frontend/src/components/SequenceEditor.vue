<script setup lang="ts">
import LabelBox from './LabelBox.vue';

const sequence = [0, 0, 1, 2, 4, 1, 2, 3, 2, 2, 5];
</script>

<template>
  <LabelBox label="Sequence" class="flex-grow">
    <div class="flex gap-1">
      <div v-for="(id, index) in sequence" :key="index" class="px-2 rounded-md bg-sky-900">
        {{ id }}
      </div>
    </div>
  </LabelBox>
</template>
