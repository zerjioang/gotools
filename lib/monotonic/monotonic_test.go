//
// Copyright Helix Distributed Ledger. All Rights Reserved.
// SPDX-License-Identifier: GNU GPL v3
//

package monotonic

import "testing"

func TestNow(t *testing.T) {
	for i := 0; i < 100; i++ {
		t1 := Now()
		t2 := Now()
		if t1 > t2 {
			t.Error(t1, "must be slower than", t2)
		}
	}
}

func TestSince(t *testing.T) {
	for i := 0; i < 100; i++ {
		t1 := Now()
		elapsed := Since(t1)
		if elapsed == 0 {
			t.Error(t1, "elapsed", elapsed)
		}
	}
}
