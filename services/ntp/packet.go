// Copyright (C) 2018 Nomango - All Rights Reserved

package ntp

import "time"

const (
	ntpEpochOffset = 2208988800
)

// Packet NTP packet
// see more information about NTP:
// http://www.ntp.org/documentation.html
type Packet struct {
	Settings       uint8  // leap yr indicator, ver number and mode
	Stratum        uint8  // stratum of local clock
	Poll           int8   // poll exponent
	Precision      int8   // precision exponent
	RootDelay      uint32 // root delay
	RootDispersion uint32 // root dispersion
	ReferenceID    uint32 // reference id
	RefTimeSec     uint32 // reference timestamp sec
	RefTimeFrac    uint32 // reference timestamp fractional
	OrigTimeSec    uint32 // origin time secs
	OrigTimeFrac   uint32 // origin time fractional
	RxTimeSec      uint32 // receive time secs
	RxTimeFrac     uint32 // receive time frac
	TxTimeSec      uint32 // transmit time secs
	TxTimeFrac     uint32 // transmit time frac
}

// DefaultPacket returns a default NTP packet
// configure request settings by specifying the first byte as
// 00 011 011 (or 0x1B)
// |  |   +-- client mode (3)
// |  + ----- version (3)
// + -------- leap year indicator, 0 no warning
func DefaultPacket() *Packet {
	return &Packet{Settings: 0x1B}
}

// Parse read unix time from packet
func (p *Packet) Parse() time.Time {
	secs := float64(p.TxTimeSec) - ntpEpochOffset
	nanos := (int64(p.TxTimeFrac) * 1e9) >> 32
	return time.Unix(int64(secs), nanos)
}
