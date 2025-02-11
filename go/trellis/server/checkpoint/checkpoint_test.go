package checkpoint

import (
	"testing"

	"github.com/31333337/bmrng/go/trellis/crypto"
)

func TestDecryptionLogic(t *testing.T) {
	numShares := 10
	groupSecret, groupKey := crypto.NewDHKeyPair()
	userSecret, userKey := crypto.NewDHKeyPair()
	groupShares := crypto.AdditiveShares(&groupSecret, numShares)

	targetKey := userSecret.Mul(&groupKey)

	sum := crypto.ZeroPoint()
	for i := range groupShares {
		partialKey := groupShares[i].Mul(&userKey)
		sum.Accumulate(&partialKey)
	}

	if !targetKey.Equals(sum) {
		t.Fail()
	}
}
