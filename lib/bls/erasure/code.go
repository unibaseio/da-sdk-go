package erasure

import (
	"bytes"
	"fmt"
	"math/big"

	bls "github.com/unibaseio/da-sdk-go/lib/bls"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

var _ types.IErasure = (*RS)(nil)

type RS struct {
	Data  int // DataNum is the number of data row vectors.
	Total int // ParityNum is the number of parity row vectors.

	Matrix matrix
}

func NewRS(n, k int) (*RS, error) {
	if k <= 0 || n <= 0 || k > n {
		return nil, fmt.Errorf("illegal k/n")
	}
	r := &RS{
		Data:   k,
		Total:  n,
		Matrix: makeEncodeMatrix(n, k),
	}

	return r, nil
}

func (r *RS) Encode(buf [][]byte, need []int) ([][]byte, error) {
	if len(buf) != r.Data {
		return nil, fmt.Errorf("should align to %d", r.Data)
	}

	vec := make([]bls.Fr, r.Data)
	for i := 0; i < r.Data; i++ {
		vec[i].SetBytes(buf[i])
	}

	resvec, err := r.EncodeFr(vec, need)
	if err != nil {
		return nil, err
	}
	res := make([][]byte, len(resvec))
	for i := 0; i < len(resvec); i++ {
		res[i] = resvec[i].Marshal()
	}

	return res, nil
}

func (r *RS) EncodeFr(vec []bls.Fr, at []int) ([]bls.Fr, error) {
	var temp bls.Fr
	if len(vec) != r.Data {
		return nil, fmt.Errorf("should align to %d", r.Data)
	}

	if len(at) == 0 {
		at = make([]int, 0, r.Total-r.Data)
		for i := r.Data; i < r.Total; i++ {
			at = append(at, i)
		}
	}

	res := make([]bls.Fr, len(at))
	for i := 0; i < len(at); i++ {
		if at[i] >= r.Total {
			return res, fmt.Errorf("index %d should less than %d", at[i], r.Total)
		}

		if at[i] < r.Data {
			return res, fmt.Errorf("index %d should larger than %d", at[i], r.Data)
		}
		for j := 0; j < r.Data; j++ {
			temp.Mul(&r.Matrix[at[i]*r.Data+j], &vec[j])
			res[i].Add(&res[i], &temp)
		}
	}

	return res, nil
}

func (r *RS) Check(buf [][]byte, val [][]byte, needcheck []int) error {
	if len(buf) != r.Data {
		return fmt.Errorf("should align to %d", r.Data)
	}

	if len(needcheck) != len(val) {
		return fmt.Errorf("check val and index length should equal")
	}

	switch len(buf[0]) {
	case 32:
		vec := make([]bls.Fr, r.Data)
		for i := 0; i < r.Data; i++ {
			vec[i].SetBytes(buf[i])
		}

		var temp, res bls.Fr
		for i := 0; i < len(needcheck); i++ {
			if needcheck[i] >= r.Total {
				return fmt.Errorf("index %d should less than %d", needcheck[i], r.Total)
			}

			if needcheck[i] < r.Data {
				return fmt.Errorf("index %d should larger than %d", needcheck[i], r.Data)
			}
			res.SetZero()
			for j := 0; j < r.Data; j++ {
				temp.Mul(&r.Matrix[needcheck[i]*r.Data+j], &vec[j])
				res.Add(&res, &temp)
			}
			resb := res.Bytes()
			if !bytes.Equal(val[i], resb[:]) {
				return fmt.Errorf("unequal val at: %d", needcheck[i])
			}
		}

	case 48, 96:
		vec := make([]bls.G1, r.Data)
		for i := 0; i < r.Data; i++ {
			_, err := vec[i].SetBytes(buf[i])
			if err != nil {
				return err
			}
		}

		var temp bls.G1
		bigval := new(big.Int)
		for i, row := range needcheck {
			if row >= r.Total {
				return fmt.Errorf("index %d should less than %d", row, r.Total)
			}

			if row < r.Data {
				return fmt.Errorf("index %d should larger than %d", row, r.Data)
			}

			var res bls.G1
			for j := 0; j < r.Data; j++ {
				r.Matrix[row*r.Data+j].BigInt(bigval)
				temp.ScalarMultiplication(&vec[j], bigval)
				res.Add(&res, &temp)
			}
			var to bls.G1
			_, err := to.SetBytes(val[i])
			if err != nil {
				return err
			}
			if !to.Equal(&res) {
				return fmt.Errorf("unequal val at: %d", row)
			}
		}

	default:
		return fmt.Errorf("unsupported data length: %d", len(buf[0]))
	}
	return nil
}

var _ types.IEncode = (*Reconst)(nil)

type Reconst struct {
	rs       *RS
	survived []int
	matrix   matrix
}

func NewReconst(n, k int, survived []int) (types.IEncode, error) {
	r, err := NewRS(n, k)
	if err != nil {
		return nil, err
	}
	return r.NewReconst(survived)
}

func (r *RS) NewReconst(survived []int) (types.IEncode, error) {
	if len(survived) < r.Data {
		return nil, fmt.Errorf("survived at least %d", r.Data)
	}

	if len(survived) > r.Data {
		survived = survived[:r.Data]
	}
	max := -1
	for i := 0; i < len(survived); i++ {
		if survived[i] <= max {
			return nil, fmt.Errorf("survived should increase")
		}
		if survived[i] >= r.Total {
			return nil, fmt.Errorf("survived should less than %d", r.Total)
		}
		max = survived[i]
	}
	em, err := r.Matrix.makeEncMatrixForReconst(survived)
	if err != nil {
		return nil, err
	}

	return &Reconst{
		rs:       r,
		survived: survived,
		matrix:   em,
	}, nil
}

func (re *Reconst) Encode(buf [][]byte, need []int) ([][]byte, error) {
	if len(buf) != re.rs.Data {
		return nil, fmt.Errorf("should align to %d", re.rs.Data)
	}

	vec := make([]bls.Fr, re.rs.Data)
	for i := 0; i < re.rs.Data; i++ {
		vec[i].SetBytes(buf[i])
	}

	resvec, err := re.EncodeFr(vec, need)
	if err != nil {
		return nil, err
	}
	res := make([][]byte, len(resvec))
	for i := 0; i < len(resvec); i++ {
		res[i] = resvec[i].Marshal()
	}

	return res, nil
}

func (re *Reconst) EncodeFr(vec []bls.Fr, need []int) ([]bls.Fr, error) {
	if len(vec) < re.rs.Data {
		return nil, fmt.Errorf("data at least %d", re.rs.Data)
	}

	if len(vec) > re.rs.Data {
		vec = vec[:re.rs.Data]
	}

	if len(need) == 0 {
		need = make([]int, 0, re.rs.Total-re.rs.Data)
		for i := re.rs.Data; i < re.rs.Total; i++ {
			need = append(need, i)
		}
	}

	need1 := make([]int, 0, len(need))
	need2 := make([]int, 0, len(need))
	for _, nc := range need {
		if nc < re.rs.Data {
			need1 = append(need1, nc)
		} else {
			need2 = append(need2, nc)
		}
	}

	var temp bls.Fr
	res := make([]bls.Fr, len(need1))
	for i, row := range need1 {
		for j := 0; j < re.rs.Data; j++ {
			temp.Mul(&re.matrix[row*re.rs.Data+j], &vec[j])
			res[i].Add(&res[i], &temp)
		}
	}

	if len(need2) > 0 {
		// should recover all data
		before := -1
		datavec := make([]bls.Fr, re.rs.Data)
		for i := 0; i < len(re.survived); i++ {
			diff := re.survived[i] - before - 1
			if diff > 0 {
				for j := 1; j <= diff; j++ {
					if before+j == re.rs.Data {
						break
					}
					// recover it
					for k := 0; k < re.rs.Data; k++ {
						temp.Mul(&re.matrix[(before+j)*re.rs.Data+k], &vec[k])
						datavec[before+j].Add(&datavec[before+j], &temp)
					}
				}
			}
			if re.survived[i] >= re.rs.Data {
				break
			}
			// already has
			datavec[re.survived[i]].Set(&(vec[i]))
			before = re.survived[i]
		}

		// re-encode
		res2, err := re.rs.EncodeFr(datavec, need2)
		if err != nil {
			return nil, err
		}
		res = append(res, res2...)
	}

	return res, nil
}
