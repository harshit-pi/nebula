//go:build (!linux || android) && !e2e_testing && !darwin
// +build !linux android
// +build !e2e_testing
// +build !darwin

package udp

import (
	"errors"
	"net/netip"
	"testing"

	"github.com/slackhq/nebula/test"
)

func TestGenericConnWriteToIPv6FromIPv4Listener(t *testing.T) {
	c, err := NewGenericListener(test.NewLogger(), Settings{Listen: netip.MustParseAddrPort("127.0.0.1:0")})
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	if err := c.WriteTo([]byte("hi"), netip.MustParseAddrPort("[::1]:4242")); !errors.Is(err, ErrInvalidIPv6RemoteForSocket) {
		t.Errorf("WriteTo IPv6 remote = %v, want ErrInvalidIPv6RemoteForSocket", err)
	}
	if err := c.WriteTo([]byte("hi"), netip.MustParseAddrPort("[::ffff:127.0.0.1]:4242")); err != nil {
		t.Errorf("WriteTo IPv4-mapped remote = %v, want nil", err)
	}
}
