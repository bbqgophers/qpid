package i2c

import (
	"testing"
	"time"

	"github.com/hybridgroup/gobot/gobottest"
)

// --------- HELPERS
func initTestWiichuckDriver() (driver *WiichuckDriver) {
	driver, _ = initTestWiichuckDriverWithStubbedAdaptor()
	return
}

func initTestWiichuckDriverWithStubbedAdaptor() (*WiichuckDriver, *i2cTestAdaptor) {
	adaptor := newI2cTestAdaptor("adaptor")
	return NewWiichuckDriver(adaptor, "bot"), adaptor
}

// --------- TESTS

func TestNewWiichuckDriver(t *testing.T) {
	// Does it return a pointer to an instance of WiichuckDriver?
	var bm interface{} = NewWiichuckDriver(newI2cTestAdaptor("adaptor"), "bot")
	_, ok := bm.(*WiichuckDriver)
	if !ok {
		t.Errorf("NewWiichuckDriver() should have returned a *WiichuckDriver")
	}
}

func TestWiichuckDriver(t *testing.T) {
	wii := initTestWiichuckDriver()
	gobottest.Assert(t, wii.Name(), "bot")
	gobottest.Assert(t, wii.Connection().Name(), "adaptor")
	gobottest.Assert(t, wii.interval, 10*time.Millisecond)

	wii = NewWiichuckDriver(newI2cTestAdaptor("adaptor"), "bot", 100*time.Millisecond)
	gobottest.Assert(t, wii.interval, 100*time.Millisecond)
}

func TestWiichuckDriverStart(t *testing.T) {
	sem := make(chan bool)
	wii, adaptor := initTestWiichuckDriverWithStubbedAdaptor()

	adaptor.i2cReadImpl = func() ([]byte, error) {
		return []byte{1, 2, 3, 4, 5, 6}, nil
	}

	numberOfCyclesForEvery := 3

	wii.interval = 1 * time.Millisecond
	gobottest.Assert(t, len(wii.Start()), 0)

	go func() {
		for {
			<-time.After(time.Duration(numberOfCyclesForEvery) * time.Millisecond)
			if (wii.joystick["sy_origin"] == float64(44)) &&
				(wii.joystick["sx_origin"] == float64(45)) {
				sem <- true
			}
		}
	}()

	select {
	case <-sem:
	case <-time.After(100 * time.Millisecond):
		t.Errorf("origin not read correctly")
	}

}

func TestWiichuckDriverHalt(t *testing.T) {
	wii := initTestWiichuckDriver()

	gobottest.Assert(t, len(wii.Halt()), 0)
}

func TestWiichuckDriverCanParse(t *testing.T) {
	wii := initTestWiichuckDriver()

	// ------ When value is not encrypted
	decryptedValue := []byte{1, 2, 3, 4, 5, 6}
	wii.update(decryptedValue)

	// - This should be done by WiichuckDriver.parse
	gobottest.Assert(t, wii.data["sx"], float64(45))
	gobottest.Assert(t, wii.data["sy"], float64(44))
	gobottest.Assert(t, wii.data["z"], float64(0))
	gobottest.Assert(t, wii.data["c"], float64(0))
}

func TestWiichuckDriverCanAdjustOrigins(t *testing.T) {
	wii := initTestWiichuckDriver()

	// ------ When value is not encrypted
	decryptedValue := []byte{1, 2, 3, 4, 5, 6}
	wii.update(decryptedValue)

	// - This should be done by WiichuckDriver.adjustOrigins
	gobottest.Assert(t, wii.joystick["sx_origin"], float64(45))
	gobottest.Assert(t, wii.joystick["sy_origin"], float64(44))
}

func TestWiichuckDriverCButton(t *testing.T) {
	wii := initTestWiichuckDriver()

	// ------ When value is not encrypted
	decryptedValue := []byte{1, 2, 3, 4, 5, 6}
	wii.update(decryptedValue)

	// - This should be done by WiichuckDriver.updateButtons
	done := make(chan bool)

	wii.On(wii.Event(C), func(data interface{}) {
		gobottest.Assert(t, data, true)
		done <- true
	})

	wii.update(decryptedValue)

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Errorf("Did not recieve 'C' event")
	}
}

func TestWiichuckDriverZButton(t *testing.T) {
	wii := initTestWiichuckDriver()

	// ------ When value is not encrypted
	decryptedValue := []byte{1, 2, 3, 4, 5, 6}
	wii.update(decryptedValue)

	done := make(chan bool)

	wii.On(wii.Event(Z), func(data interface{}) {
		gobottest.Assert(t, data, true)
		done <- true
	})

	wii.update(decryptedValue)

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Errorf("Did not recieve 'Z' event")
	}
}

func TestWiichuckDriverUpdateJoystick(t *testing.T) {
	wii := initTestWiichuckDriver()

	// ------ When value is not encrypted
	decryptedValue := []byte{1, 2, 3, 4, 5, 6}

	// - This should be done by WiichuckDriver.updateJoystick
	expectedData := map[string]float64{
		"x": float64(0),
		"y": float64(0),
	}

	done := make(chan bool)

	wii.On(wii.Event(Joystick), func(data interface{}) {
		gobottest.Assert(t, data, expectedData)
		done <- true
	})

	wii.update(decryptedValue)

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Errorf("Did not recieve 'Joystick' event")
	}
}

func TestWiichuckDriverEncrypted(t *testing.T) {
	wii := initTestWiichuckDriver()

	// ------ When value is encrypted
	wii = initTestWiichuckDriver()
	encryptedValue := []byte{1, 1, 2, 2, 3, 3}

	wii.update(encryptedValue)

	gobottest.Assert(t, wii.data["sx"], float64(0))
	gobottest.Assert(t, wii.data["sy"], float64(0))
	gobottest.Assert(t, wii.data["z"], float64(0))
	gobottest.Assert(t, wii.data["c"], float64(0))

	gobottest.Assert(t, wii.joystick["sx_origin"], float64(-1))
	gobottest.Assert(t, wii.joystick["sy_origin"], float64(-1))
}

func TestWiichuckDriverSetJoystickDefaultValue(t *testing.T) {
	wii := initTestWiichuckDriver()

	gobottest.Assert(t, wii.joystick["sy_origin"], float64(-1))

	wii.setJoystickDefaultValue("sy_origin", float64(2))

	gobottest.Assert(t, wii.joystick["sy_origin"], float64(2))

	// when current default value is not -1 it keeps the current value
	wii.setJoystickDefaultValue("sy_origin", float64(20))

	gobottest.Assert(t, wii.joystick["sy_origin"], float64(2))
}

func TestWiichuckDriverCalculateJoystickValue(t *testing.T) {
	wii := initTestWiichuckDriver()

	gobottest.Assert(t, wii.calculateJoystickValue(float64(20), float64(5)), float64(15))
	gobottest.Assert(t, wii.calculateJoystickValue(float64(1), float64(2)), float64(-1))
	gobottest.Assert(t, wii.calculateJoystickValue(float64(10), float64(5)), float64(5))
	gobottest.Assert(t, wii.calculateJoystickValue(float64(5), float64(10)), float64(-5))
}

func TestWiichuckDriverIsEncrypted(t *testing.T) {
	wii := initTestWiichuckDriver()

	encryptedValue := []byte{1, 1, 2, 2, 3, 3}
	gobottest.Assert(t, wii.isEncrypted(encryptedValue), true)

	encryptedValue = []byte{42, 42, 24, 24, 30, 30}
	gobottest.Assert(t, wii.isEncrypted(encryptedValue), true)

	decryptedValue := []byte{1, 2, 3, 4, 5, 6}
	gobottest.Assert(t, wii.isEncrypted(decryptedValue), false)

	decryptedValue = []byte{1, 1, 2, 2, 5, 6}
	gobottest.Assert(t, wii.isEncrypted(decryptedValue), false)

	decryptedValue = []byte{1, 1, 2, 3, 3, 3}
	gobottest.Assert(t, wii.isEncrypted(decryptedValue), false)
}

func TestWiichuckDriverDecode(t *testing.T) {
	wii := initTestWiichuckDriver()

	gobottest.Assert(t, wii.decode(byte(0)), float64(46))
	gobottest.Assert(t, wii.decode(byte(100)), float64(138))
	gobottest.Assert(t, wii.decode(byte(200)), float64(246))
	gobottest.Assert(t, wii.decode(byte(254)), float64(0))
}

func TestWiichuckDriverParse(t *testing.T) {
	wii := initTestWiichuckDriver()

	gobottest.Assert(t, wii.data["sx"], float64(0))
	gobottest.Assert(t, wii.data["sy"], float64(0))
	gobottest.Assert(t, wii.data["z"], float64(0))
	gobottest.Assert(t, wii.data["c"], float64(0))

	// First pass
	wii.parse([]byte{12, 23, 34, 45, 56, 67})

	gobottest.Assert(t, wii.data["sx"], float64(50))
	gobottest.Assert(t, wii.data["sy"], float64(23))
	gobottest.Assert(t, wii.data["z"], float64(1))
	gobottest.Assert(t, wii.data["c"], float64(2))

	// Second pass
	wii.parse([]byte{70, 81, 92, 103, 204, 205})

	gobottest.Assert(t, wii.data["sx"], float64(104))
	gobottest.Assert(t, wii.data["sy"], float64(93))
	gobottest.Assert(t, wii.data["z"], float64(1))
	gobottest.Assert(t, wii.data["c"], float64(0))
}

func TestWiichuckDriverAdjustOrigins(t *testing.T) {
	wii := initTestWiichuckDriver()

	gobottest.Assert(t, wii.joystick["sy_origin"], float64(-1))
	gobottest.Assert(t, wii.joystick["sx_origin"], float64(-1))

	// First pass
	wii.parse([]byte{1, 2, 3, 4, 5, 6})
	wii.adjustOrigins()

	gobottest.Assert(t, wii.joystick["sy_origin"], float64(44))
	gobottest.Assert(t, wii.joystick["sx_origin"], float64(45))

	// Second pass
	wii = initTestWiichuckDriver()

	wii.parse([]byte{61, 72, 83, 94, 105, 206})
	wii.adjustOrigins()

	gobottest.Assert(t, wii.joystick["sy_origin"], float64(118))
	gobottest.Assert(t, wii.joystick["sx_origin"], float64(65))
}
