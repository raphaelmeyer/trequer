<script setup lang="ts">
import SelectOutput from '@/components/SelectOutput.vue';
import PatternEditor from '@/components/PatternEditor.vue';
import PatternSelector from '@/components/PatternSelector.vue';
</script>

<template>
  <div class="p-2">
    <PatternSelector></PatternSelector>
    <SelectOutput></SelectOutput>
    <PatternEditor></PatternEditor>
  </div>
</template>
