import { ref } from 'vue';
import { defineStore } from 'pinia';

import { ListMidiOutputs } from '../../wailsjs/go/main/App';

export const useMidiPortsStore = defineStore('midiport', () => {
  const ports = ref<string[]>([]);

  function refresh() {
    ListMidiOutputs().then((value) => {
      ports.value = value;
    });
  }

  return { ports, refresh };
});
