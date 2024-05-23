import { ref } from 'vue';
import { defineStore } from 'pinia';

import { ListMidiOut } from '../../wailsjs/go/main/App';

export const useMidiPortsStore = defineStore('midiport', () => {
  const ports = ref<string[]>([]);

  function refresh() {
    ListMidiOut().then((value) => {
      ports.value = value;
    });
  }

  return { ports, refresh };
});
