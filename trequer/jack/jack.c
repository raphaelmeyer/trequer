#include <jack/jack.h>

#include "_cgo_export.h"

static void port_registration_callback(jack_port_id_t port, int reg,
                                       void *arg) {
  goPortRegistration(port, reg, (uintptr_t)arg);
}

jack_client_t *jack_client_open_go(const char *client_name,
                                   jack_options_t options,
                                   jack_status_t *status) {
  return jack_client_open(client_name, options, status);
}

int jack_set_port_registration_callback_go(jack_client_t *client,
                                           uintptr_t arg) {
  return jack_set_port_registration_callback(client, port_registration_callback,
                                             (void *)arg);
}
