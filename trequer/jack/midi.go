package jack

import (
	"log"
	"runtime/cgo"
)

type CallbackHandler interface {
	OnRegisterPort()
	OnBeat(count uint32)
}

func NewClient(name string) (*Client, error) {
	client, err := clientOpen(name, NullOption)
	if err != nil {
		return nil, err
	}

	client.cbHandle = cgo.NewHandle(client)
	client.setPortRegistrationCallback(client.cbHandle)
	// client.setProgressCallback(client.cbHandle)

	channel0, err := client.portRegister("channel-0", midiType, portIsOutput)
	if err != nil {
		return nil, err
	}
	client.channel = channel0

	err = client.activate()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) Close() {
	err := c.clientClose()
	if err != nil {
		log.Println(err)
	}

	c.handle = nil
	c.cbHandle.Delete()
}

func (c *Client) SetCallbackHandler(handler CallbackHandler) {
	c.cbHandler = handler
}

func (c *Client) ListMidiOut() ([]string, error) {
	return c.getPorts("", "midi", portIsInput), nil
}
