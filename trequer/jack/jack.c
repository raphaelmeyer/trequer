#include <jack/jack.h>
#include <jack/ringbuffer.h>

#include "_cgo_export.h"

#include <stdint.h>

static void port_registration_callback(jack_port_id_t port, int reg,
                                       void *arg) {
  goPortRegistration(port, reg, (uintptr_t)arg);
}

static int process_callback(jack_nframes_t nframes, void *arg) {
  struct ProcessContext *context = (struct ProcessContext *)(arg);
  uint32_t const old_beat = context->frames / context->frames_per_beat;
  uint32_t const next_beat =
      (context->frames + nframes) / context->frames_per_beat;
  context->frames += nframes;
  if (old_beat != next_beat) {
    size_t const written = jack_ringbuffer_write(
        context->beat, (char *)(&next_beat), sizeof(uint32_t));
    if (written != sizeof(uint32_t)) {
      return 1;
    }
  }

  return 0;
}

jack_client_t *jack_client_open_go(const char *client_name,
                                   jack_options_t options,
                                   jack_status_t *status) {
  return jack_client_open(client_name, options, status);
}

int jack_set_process_callback_go(jack_client_t *client,
                                 struct ProcessContext *arg) {
  return jack_set_process_callback(client, process_callback, arg);
}

int jack_set_port_registration_callback_go(jack_client_t *client,
                                           uintptr_t arg) {
  return jack_set_port_registration_callback(client, port_registration_callback,
                                             (void *)arg);
}

int jack_ringbuffer_read_go(jack_ringbuffer_t *rb, void *dest, size_t cnt) {
  return jack_ringbuffer_read(rb, (char *)dest, cnt);
}
