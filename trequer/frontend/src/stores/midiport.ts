import { ref } from 'vue';
import { defineStore } from 'pinia';
import { ListMidiOutputs } from '../../wailsjs/go/audiocontrol/AudioControl';

export const useMidiPortsStore = defineStore('midiport', () => {
  const ports = ref<string[]>([]);

  function refresh() {
    ListMidiOutputs().then((value) => {
      ports.value = value;
    });
  }

  return { ports, refresh };
});
