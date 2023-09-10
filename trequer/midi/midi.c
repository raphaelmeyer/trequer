#include <alsa/asoundlib.h>

snd_seq_client_info_t * midi_snd_seq_client_info_malloc() {
  snd_seq_client_info_t * cinfo = NULL;
  if(snd_seq_client_info_malloc(&cinfo) == 0) {
    return cinfo;
  }

  return NULL;
}

snd_seq_port_info_t * midi_snd_seq_port_info_malloc() {
  snd_seq_port_info_t * pinfo = NULL;
  if(snd_seq_port_info_malloc(&pinfo) == 0) {
    return pinfo;
  }

  return NULL;
}
