<script setup lang="ts">
import SelectOutput from '@/components/SelectOutput.vue';
import PatternEditor from '@/components/PatternEditor.vue';
import PatternSelector from '@/components/PatternSelector.vue';
import SequenceEditor from '@/components/SequenceEditor.vue';
</script>

<template>
  <SelectOutput></SelectOutput>
  <div class="flex px-2 gap-2">
    <PatternSelector></PatternSelector>
    <SequenceEditor></SequenceEditor>
  </div>
  <div class="p-2">
    <PatternEditor></PatternEditor>
  </div>
</template>
