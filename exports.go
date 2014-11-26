// Copyright (c) 2014 Joseph D Poirier
// Distributable under the terms of The simplified BSD License
// that can be found in the LICENSE file.

// Package visa wraps National Instruments VISA (Virtual Instrument Software
// Architecture) driver. The driver allows a client application to communicate
// with most instrumentation buses including GPIB, USB, Serial, and Ethernet.
// VISA is an industry standard for instrument communications.
//
// The package is low level and, for the most part, is one-to-one with the
// exported C functions it wraps. Clients would typically build instrument
// drivers around the package but it can also be used directly.
package visa

/*
#include <visa.h>
*/
import "C"

//export go_cb
func go_cb(instr C.ViSession, etype C.ViEventType, eventContext C.ViEvent, userHandle C.ViAddr) {
	(*(PUserCallback(userHandle)))(Object(instr), uint32(etype), uint32(eventContext))
}
