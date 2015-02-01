// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package blockchain_test

import (
	"testing"

	"github.com/PointCoin/btcd/blockchain"
	"github.com/PointCoin/btcutil"
)

// TestMerkle tests the BuildMerkleTreeStore API.
func TestMerkle(t *testing.T) {
	block := btcutil.NewBlock(&Block100000)
	merkles := blockchain.BuildMerkleTreeStore(block.Transactions())
	calculatedMerkleRoot := merkles[len(merkles)-1]
	wantMerkle := &Block100000.Header.MerkleRoot
	if !wantMerkle.IsEqual(calculatedMerkleRoot) {
		t.Errorf("BuildMerkleTreeStore: merkle root mismatch - "+
			"got %v, want %v", calculatedMerkleRoot, wantMerkle)
	}
}
