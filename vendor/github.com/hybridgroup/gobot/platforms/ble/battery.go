package ble

import (
	"bytes"

	"github.com/hybridgroup/gobot"
)

var _ gobot.Driver = (*BLEBatteryDriver)(nil)

type BLEBatteryDriver struct {
	name       string
	connection gobot.Connection
	gobot.Eventer
}

// NewBLEBatteryDriver creates a BLEBatteryDriver by name
func NewBLEBatteryDriver(a *BLEClientAdaptor, name string) *BLEBatteryDriver {
	n := &BLEBatteryDriver{
		name:       name,
		connection: a,
		Eventer:    gobot.NewEventer(),
	}

	return n
}
func (b *BLEBatteryDriver) Connection() gobot.Connection { return b.connection }
func (b *BLEBatteryDriver) Name() string                 { return b.name }

// adaptor returns BLE adaptor
func (b *BLEBatteryDriver) adaptor() *BLEClientAdaptor {
	return b.Connection().(*BLEClientAdaptor)
}

// Start tells driver to get ready to do work
func (b *BLEBatteryDriver) Start() (errs []error) {
	return
}

// Halt stops battery driver (void)
func (b *BLEBatteryDriver) Halt() (errs []error) { return }

func (b *BLEBatteryDriver) GetBatteryLevel() (level uint8) {
	var l uint8
	c, _ := b.adaptor().ReadCharacteristic("180f", "2a19")
	buf := bytes.NewBuffer(c)
	val, _ := buf.ReadByte()
	l = uint8(val)
	return l
}
