package ping

import (
	"bytes"
	"testing"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func BenchmarkProcessPacket(b *testing.B) {
	pinger, _ := NewPinger("127.0.0.1")

	pinger.ipv4 = true
	pinger.addr = "127.0.0.1"
	pinger.network = "ip4:icmp"
	pinger.id = 123
	pinger.Tracker = 456

	t := append(timeToBytes(time.Now()), intToBytes(pinger.Tracker)...)
	if remainSize := pinger.Size - timeSliceLength - trackerLength; remainSize > 0 {
		t = append(t, bytes.Repeat([]byte{1}, remainSize)...)
	}

	body := &icmp.Echo{
		ID:   pinger.id,
		Seq:  pinger.sequence,
		Data: t,
	}

	msg := &icmp.Message{
		Type: ipv4.ICMPTypeEchoReply,
		Code: 0,
		Body: body,
	}

	msgBytes, _ := msg.Marshal(nil)

	pkt := packet{
		nbytes: len(msgBytes),
		bytes:  msgBytes,
		ttl:    24,
	}

	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()

	for k := 0; k < b.N; k++ {
		pinger.processPacket(&pkt)
	}
}
