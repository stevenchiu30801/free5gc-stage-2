package nrf_handler_test

import (
	"gofree5gc/src/nrf/nrf_handler"
	"gofree5gc/src/nrf/nrf_handler/nrf_message"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	go nrf_handler.Handle()
	msg := nrf_message.ChannelMessage{}
	msg.Event = nrf_message.EventNFDiscovery
	//msg.Value = ngapMsg
	nrf_handler.SendMessage(msg)

	time.Sleep(100 * time.Millisecond)

}
