<script setup lang="ts">
import { keyName } from '@/models/midi';
import type { Note } from '@/models/song';
import { computed } from 'vue';

const props = defineProps<{
  note: Note | undefined;
  selected: boolean;
}>();

const name = computed(() => {
  if (props.note !== undefined) {
    return keyName(props.note.key);
  }
  return '...';
});
</script>

<template>
  <div
    class="flex px-2 rounded-md"
    :class="{
      '[&:nth-child(4n+1)]:bg-sky-900': !props.selected,
      'bg-sky-200 text-sky-900': props.selected,
    }"
  >
    <div class="w-10">{{ name }}</div>
    <div class="w-8 text-right">{{ props.note?.volume ?? '..' }}</div>
    <div class="w-8 text-right">{{ props.note?.length ?? '..' }}</div>
  </div>
</template>
