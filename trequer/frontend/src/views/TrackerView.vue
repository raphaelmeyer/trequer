<script setup lang="ts">
import SelectOutput from '@/components/SelectOutput.vue';
import PatternEditor from '@/components/PatternEditor.vue';
import PatternSelector from '@/components/PatternSelector.vue';
import SequenceEditor from '@/components/SequenceEditor.vue';
import { onMounted, onUnmounted } from 'vue';
import { useEditorStore } from '@/stores/editor';

const editor = useEditorStore();

function onKeyup(event: KeyboardEvent) {
  editor.handleKey(event);
}

onMounted(() => {
  document.addEventListener('keyup', onKeyup);
});

onUnmounted(() => {
  document.removeEventListener('keyup', onKeyup);
});
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
