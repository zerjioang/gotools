package udp

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"syscall"
	"testing"

	"golang.org/x/net/ipv4"
)

const UDP_HEADER_LEN = 8

type UdpHeader struct {
	SourcePort      uint16
	DestinationPort uint16
	Length          uint16
	Checksum        uint16
}

// manual checksum calculation
func csum(sum int, data []byte) int {
	for i, b := range data {
		if i&1 == 0 {
			sum += (int(b) << 8)
		} else {
			sum += int(b)
		}
	}
	for sum > 0xffff {
		sum += (sum >> 16)
		sum &= 0xffff
	}
	return sum
}

func TestSyscallMode(t *testing.T) {
	destIP := net.ParseIP("127.0.0.1")
	srcIP := net.ParseIP("127.0.0.2")
	proto := 17 // UDP
	data := []byte("testdata")

	udpheaderT := UdpHeader{
		SourcePort:      1111,
		DestinationPort: 2222,
		Length:          uint16(UDP_HEADER_LEN + len(data)),
	}

	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, &udpheaderT); err != nil {
		log.Fatal(err)
	}

	udpHeader := buf.Bytes()
	dataWithHeader := append(udpHeader, data...)

	h := &ipv4.Header{
		Version:  ipv4.Version,
		Len:      ipv4.HeaderLen,
		TotalLen: ipv4.HeaderLen + UDP_HEADER_LEN + len(data),
		ID:       12345,
		Protocol: proto,
		TTL:      64,
		Dst:      destIP.To4(),
		Src:      srcIP.To4(),
	}
	ipH, _ := h.Marshal()
	ipH[2], ipH[3] = ipH[3], ipH[2] // no idea why this is required, but these bytes are swapped after marshalling...
	sum := csum(0, ipH)
	sum ^= 0xffff
	h.Checksum = sum

	// for UDP checksum, calculate over IP pseudoheader
	sum = csum(0, ipH[12:20]) // src and dest IP
	sum = csum(sum, []byte{0, byte(proto), 0, byte(UDP_HEADER_LEN + len(data))})
	sum = csum(sum, dataWithHeader)
	sum ^= 0xffff
	// update checksum in marshalled stream
	dataWithHeader[6] = byte(sum >> 8)
	dataWithHeader[7] = byte(sum)

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	if err != nil {
		log.Fatal(err)
	}
	if err = syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1); err != nil {
		log.Fatal(err)
	}

	out, err := h.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	packet := append(out, dataWithHeader...)

	// destination address a second time, but here it is irrelevant
	addr := syscall.SockaddrInet4{}
	if err = syscall.Sendto(fd, packet, 0, &addr); err != nil {
		log.Fatal(err)
	}
	syscall.Close(fd)
}
