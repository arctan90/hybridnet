/*
 Copyright 2022 The Hybridnet Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package utils

import (
	"net"
	"reflect"
	"testing"
)

func TestCIDR_NextIP(t *testing.T) {
	testCases := []struct {
		ip     net.IP
		nextIP net.IP
	}{
		{
			[]byte{192, 0, 2},
			nil,
		},
		{
			net.ParseIP("192.168.0.1"),
			net.IPv4(192, 168, 0, 2).To4(),
		},
		{
			net.ParseIP("192.168.0.255"),
			net.IPv4(192, 168, 1, 0).To4(),
		},
		{
			net.ParseIP("0.1.0.5"),
			net.IPv4(0, 1, 0, 6).To4(),
		},
		{
			net.ParseIP("AB12::123"),
			net.ParseIP("AB12::124"),
		},
		{
			net.ParseIP("AB12::FFFF"),
			net.ParseIP("AB12::1:0"),
		},
		{
			net.ParseIP("0::123"),
			net.ParseIP("0::124"),
		},
	}

	for _, test := range testCases {
		ip := NextIP(test.ip)
		if !ip.Equal(test.nextIP) {
			t.Errorf("expect next IP %s but got %s", test.nextIP, ip)
		}
	}
}

func TestCIDR_PrevIP(t *testing.T) {
	testCases := []struct {
		ip     net.IP
		prevIP net.IP
	}{
		{
			[]byte{192, 0, 2},
			nil,
		},
		{
			net.ParseIP("192.168.0.2"),
			net.IPv4(192, 168, 0, 1).To4(),
		},
		{
			net.ParseIP("192.168.1.0"),
			net.IPv4(192, 168, 0, 255).To4(),
		},
		{
			net.ParseIP("0.1.0.5"),
			net.IPv4(0, 1, 0, 4).To4(),
		},
		{
			net.ParseIP("AB12::123"),
			net.ParseIP("AB12::122"),
		},
		{
			net.ParseIP("AB12::1:0"),
			net.ParseIP("AB12::FFFF"),
		},
		{
			net.ParseIP("0::124"),
			net.ParseIP("0::123"),
		},
	}

	for _, test := range testCases {
		ip := PrevIP(test.ip)
		if !ip.Equal(test.prevIP) {
			t.Errorf("expect prev IP %s but got %s", test.prevIP, ip)
		}
	}
}

func TestCmpIP(t *testing.T) {
	testCases := []struct {
		a      net.IP
		b      net.IP
		result int
	}{
		{
			net.ParseIP("192.168.0.2"),
			nil,
			-2,
		},
		{
			net.ParseIP("192.168.0.2"),
			[]byte{192, 168, 5},
			-2,
		},
		{
			net.ParseIP("192.168.0.2"),
			net.ParseIP("AB12::123"),
			-2,
		},
		{
			net.ParseIP("192.168.0.2"),
			net.ParseIP("192.168.0.5"),
			-1,
		},
		{
			net.ParseIP("192.168.0.2"),
			net.ParseIP("192.168.0.5").To4(),
			-1,
		},
		{
			net.ParseIP("192.168.0.10"),
			net.ParseIP("192.168.0.5"),
			1,
		},
		{
			net.ParseIP("192.168.0.10"),
			net.ParseIP("192.168.0.10"),
			0,
		},
		{
			net.ParseIP("192.168.0.10"),
			net.ParseIP("192.168.0.10").To4(),
			0,
		},
		{
			net.ParseIP("AB12::122"),
			net.ParseIP("AB12::123"),
			-1,
		},
		{
			net.ParseIP("AB12::210"),
			net.ParseIP("AB12::123"),
			1,
		},
		{
			net.ParseIP("AB12::210"),
			net.ParseIP("AB12::210"),
			0,
		},
	}

	for _, test := range testCases {
		outcome := Cmp(test.a, test.b)
		if outcome != test.result {
			t.Errorf("expect result %d but got %d", test.result, outcome)
		}
	}
}

func TestCapacity(t *testing.T) {
	testCases := []struct {
		a      net.IP
		b      net.IP
		result int64
	}{
		{
			net.ParseIP("192.168.0.1"),
			net.ParseIP("192.168.0.2"),
			2,
		},
		{
			net.ParseIP("192.168.0.1"),
			net.ParseIP("192.168.0.1"),
			1,
		},
		{
			net.ParseIP("192.168.0.1"),
			net.ParseIP("AB12::210"),
			0,
		},
		{
			net.ParseIP("192.168.0.2"),
			[]byte{192, 168, 5},
			0,
		},
		{
			net.ParseIP("AB12::123"),
			net.ParseIP("AB12::210"),
			238,
		},
		{
			net.ParseIP("AB12::210"),
			net.ParseIP("AB12::123"),
			238,
		},
		{
			net.ParseIP("AB12::123"),
			net.ParseIP("AB12::123"),
			1,
		},
	}
	for _, test := range testCases {
		outcome := Capacity(test.a, test.b)
		if outcome != test.result {
			t.Errorf("expect result %d but got %d", test.result, outcome)
		}
	}
}

func TestNetwork(t *testing.T) {
	testCases := []struct {
		ipNet  *net.IPNet
		result *net.IPNet
	}{
		{
			nil,
			nil,
		},
		{
			&net.IPNet{
				IP:   nil,
				Mask: net.IPv4Mask(255, 255, 255, 0),
			},
			nil,
		},
		{
			&net.IPNet{
				IP:   net.IPv4(192, 168, 0, 1),
				Mask: nil,
			},
			nil,
		},
		{
			&net.IPNet{
				IP:   net.ParseIP("AB12::123"),
				Mask: net.IPv4Mask(255, 255, 255, 0),
			},
			nil,
		},
		{
			&net.IPNet{
				IP:   net.IPv4(192, 168, 0, 100).To4(),
				Mask: net.CIDRMask(120, 128),
			},
			&net.IPNet{
				IP:   net.IPv4(192, 168, 0, 0).To4(),
				Mask: net.CIDRMask(120, 128),
			},
		},
		{
			&net.IPNet{
				IP:   net.IPv4(192, 168, 0, 100),
				Mask: net.CIDRMask(24, 32),
			},
			&net.IPNet{
				IP:   net.IPv4(192, 168, 0, 0).To4(),
				Mask: net.CIDRMask(24, 32),
			},
		},
		{
			&net.IPNet{
				IP:   net.ParseIP("AB12::123"),
				Mask: net.CIDRMask(120, 128),
			},
			&net.IPNet{
				IP:   net.ParseIP("AB12::100"),
				Mask: net.CIDRMask(120, 128),
			},
		},
	}

	for _, test := range testCases {
		ipNet := Network(test.ipNet)
		if !reflect.DeepEqual(ipNet, test.result) {
			t.Errorf("expect result %s but got %s", test.result, ipNet)
		}
	}
}
