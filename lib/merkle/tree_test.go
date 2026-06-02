package merkle

import (
	"bytes"
	"crypto/sha256"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/utils"
)

var segSize = 32
var depth = 5

func TestMerkelTree(t *testing.T) {
	for nc := 1 << depth; nc < (1 << (depth + 1)); nc++ {
		for j := 0; j < nc; j++ {
			tree := New(sha256.New())

			tree.SetIndex(uint64(j))

			for i := 0; i < nc; i++ {
				b := utils.RandomBytes(segSize)
				tree.Push(b)
			}

			merkleRoot, merkleProof, pindex, numLeaves := tree.Prove()
			if pindex != uint64(j) {
				t.Fatal("wrong proof index")
			}

			if numLeaves != uint64(nc) {
				t.Fatal("wrong node count")
			}

			if !bytes.Equal(tree.Root(), merkleRoot) {
				t.Fatal("wrong node root")
			}

			t.Logf("prooflen: %d\n", len(merkleProof))
			//t.Log("proot:", hex.EncodeToString(merkleRoot))

			verified := VerifyProof(sha256.New(), merkleRoot, merkleProof, pindex)
			if !verified {
				t.Fatal("wrong merkle proof at: ", j, nc)
			}
		}
	}
}

func TestMerkelProof(t *testing.T) {
	var buf bytes.Buffer

	var numNodes = 1<<5 + 7
	var proofIndex = 1<<5 + 1
	t.Logf("nodes: %d, proof index: %d\n", numNodes, proofIndex)

	for i := 0; i < numNodes; i++ {
		b := utils.RandomBytes(segSize)
		buf.Write(b)
	}

	merkleRoot, merkleProof, _, err := BuildReaderProof(&buf, sha256.New(), segSize, uint64(proofIndex))
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("prooflen: %d\n", len(merkleProof))

	verified := VerifyProof(sha256.New(), merkleRoot, merkleProof, uint64(proofIndex))
	if !verified {
		t.Fatal("wrong merkle proof")
	}
}

func TestMerkelRoot(t *testing.T) {
	var numNodes = 1<<6 + 1<<4

	t1 := New(sha256.New())
	t2 := NewRTree(sha256.New())

	//segSize = t1.hash.BlockSize()
	t.Logf("nodes: %d, segseize: %d\n", numNodes, segSize)

	for i := 0; i < numNodes; i++ {
		b := utils.RandomBytes(segSize)
		t1.Push(b)
		t2.Push(b)

		if !bytes.Equal(t1.Root(), t2.Root()) {
			t.Fatal("unequal root at: ", i)
		}
	}

	if !bytes.Equal(t1.Root(), t1.Root()) {
		t.Fatal("unequal root self")
	}

	if !bytes.Equal(t1.Root(), t2.Root()) {
		t.Fatal("unequal root")
	}
}

func TestMerkelAux(t *testing.T) {
	numNodes := 1 << 5
	subshard := 1 << 3

	for ; numNodes < 1<<6; numNodes++ {
		t.Logf("total %d, sub: %d, count: %d", numNodes, subshard, 1+(numNodes-1)/subshard)

		h := sha256.New()
		b := utils.RandomBytes(segSize * numNodes)
		root, err := ReaderRoot(bytes.NewReader(b), h, segSize)
		if err != nil {
			t.Fatal(err)
		}

		aroot, aux, err := ReaderRootAux(bytes.NewReader(b), h, segSize, subshard)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("aux length: ", len(aux)/h.Size())

		if !bytes.Equal(root, aroot) {
			t.Fatal("unequal root")
		}
	}
}

func TestMerkelAuxProof(t *testing.T) {
	numNodes := 1 << 5
	subshard := 1 << 3
	for ; numNodes < 1<<6; numNodes++ {
		t.Logf("total %d, sub: %d, count: %d", numNodes, subshard, 1+(numNodes-1)/subshard)

		h := sha256.New()
		b := utils.RandomBytes(segSize * numNodes)
		root, err := ReaderRoot(bytes.NewReader(b), h, segSize)
		if err != nil {
			t.Fatal(err)
		}

		aroot, aux, err := ReaderRootAux(bytes.NewReader(b), h, segSize, subshard)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("aux length: ", len(aux)/h.Size())

		if !bytes.Equal(root, aroot) {
			t.Fatal("unequal root")
		}

		for i := 0; i < numNodes; i++ {
			start := (i / subshard) * subshard
			end := start + subshard
			if end > numNodes {
				end = numNodes
			}
			nroot, proofset, err := BuildDataAuxProof(b[start*segSize:end*segSize], aux, h, segSize, subshard, uint64(i))
			if err != nil {
				t.Log(i)
				t.Fatal(err)
			}
			if !bytes.Equal(root, nroot) {
				t.Fatal("unequal root at: ", i)
			}
			ok := VerifyProof(h, root, proofset, uint64(i))
			if !ok {
				t.Fatal("wrong proof at: ", i)
			}
		}
	}
}

func TestBits(t *testing.T) {
	pindex := 13

	for i := 0; i < 10; i++ {
		pi := ((pindex & (1 << i)) >> i)
		t.Log(i, pi)
	}
}
