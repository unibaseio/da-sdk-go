package erasure

import (
	"bytes"
	"math/rand"
	"testing"

	bls "github.com/unibaseio/da-sdk-go/lib/bls"
)

func TestMatrix(t *testing.T) {
	n := 19
	k := 6
	rs, err := NewRS(n, k)
	if err != nil {
		t.Fatal(err)
	}

	data := make([]bls.Fr, k)
	for i := 0; i < len(data); i++ {
		data[i].SetRandom()
	}

	res := make([]bls.Fr, n)
	var temp bls.Fr
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			temp.Mul(&rs.Matrix[i*k+j], &data[j])
			res[i].Add(&res[i], &temp)
		}
		if i < k && !res[i].Equal(&data[i]) {
			t.Fatal("unequal data")
		}
	}

	suvpos := rand.Intn(n - k)
	if suvpos == 0 {
		suvpos = 1
	}
	survived := make([]int, k)
	for i := 0; i < k; i++ {
		survived[i] = (suvpos + i) % n
	}

	if suvpos > k {
		suvpos = k
	}
	needReconst := make([]int, suvpos)
	for i := 0; i < len(needReconst); i++ {
		needReconst[i] = i
	}

	em, err := rs.Matrix.makeEncMatrixForReconst(survived)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(em))

	vs := make([]bls.Fr, len(survived))
	for i, row := range survived {
		vs[i].Set(&res[row])
	}

	rescon := make([]bls.Fr, len(needReconst))
	for i, row := range needReconst {
		for j := 0; j < k; j++ {
			temp.Mul(&em[row*k+j], &vs[j])
			rescon[i].Add(&rescon[i], &temp)
		}
		if !rescon[i].Equal(&res[i]) {
			t.Fatal("recons fail at:", i, row)
		}
	}
}

func TestEncode(t *testing.T) {
	n := 17
	k := 6
	rs, err := NewRS(n, k)
	if err != nil {
		t.Fatal(err)
	}

	data := make([]bls.Fr, k)
	for i := 0; i < len(data); i++ {
		data[i].SetRandom()
	}
	parity, err := rs.EncodeFr(data, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(parity) != n-k {
		t.Fatal("parity fail")
	}

	data = append(data, parity...)

	for i := k; i < n; i++ {
		pres, err := rs.EncodeFr(data[:k], []int{i})
		if err != nil {
			t.Fatal(err)
		}
		if parity[i-k].Cmp(&pres[0]) != 0 {
			t.Fatal("encode fail")
		}
	}

	suvpos := k
	survived := make([]int, k)
	surdata := make([][]byte, k)
	for i := 0; i < k; i++ {
		survived[i] = (suvpos + i) % n
		surdata[i] = data[survived[i]].Marshal()
	}

	re, err := rs.NewReconst(survived)
	if err != nil {
		t.Fatal(err)
	}

	need := []int{1, 2, 4, 13, 16}
	needdata, err := re.Encode(surdata, need)
	if err != nil {
		t.Fatal(err)
	}

	if len(needdata) != len(need) {
		t.Fatal("re length fail")
	}

	for i, row := range need {
		if !bytes.Equal(needdata[i], data[row].Marshal()) {
			t.Log(survived)
			t.Log(need)
			t.Fatal("re fail")
		}
	}

}

func BenchmarkMatrixGenerate(b *testing.B) {
	for j := 0; j < b.N; j++ {
		n := 96
		k := 64
		_, err := NewRS(n, k)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMatrix(b *testing.B) {
	for j := 0; j < b.N; j++ {
		n := 96
		k := 64
		rs, err := NewRS(n, k)
		if err != nil {
			b.Fatal(err)
		}

		suvpos := rand.Intn(n - k)
		survived := make([]int, k)
		for i := 0; i < k; i++ {
			survived[i] = (suvpos + i) % n
		}
		needReconst := make([]int, suvpos)
		for i := 0; i < len(needReconst); i++ {
			needReconst[i] = i
		}

		_, err = rs.Matrix.makeEncMatrixForReconst(survived)
		if err != nil {
			b.Fatal(err)
		}
	}
}
