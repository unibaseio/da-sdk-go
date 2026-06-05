package erasure

import (
	"fmt"
	"math/big"

	bls "github.com/unibaseio/da-sdk-go/lib/bls"
)

type matrix []bls.Fr

func makeEncodeMatrix(n, k int) (matrix []bls.Fr) {
	matrix = make([]bls.Fr, n*k)
	// identity matrix
	for i := 0; i < k; i++ {
		matrix[i*k+i].SetOne()
	}

	// cachy matrix: (1/(i + j), 0 <= j < k, k <= i < 2*k)
	off := k * k // Skip the identity matrix.
	for i := k; i < n; i++ {
		for j := 0; j < k; j++ {
			matrix[off].SetInt64(int64(i + j))
			matrix[off].Inverse(&matrix[off])
			res := new(big.Int)
			matrix[off].BigInt(res)
			off++
		}
	}
	return matrix
}

// makeEncMatrixForReconst makes an encoding matrix by calculating
// the inverse matrix of survived encoding matrix.
func (m matrix) makeEncMatrixForReconst(survived []int) (em matrix, err error) {
	k := len(survived)
	msurvive := make([]bls.Fr, k*k)
	for i, l := range survived {
		copy(msurvive[i*k:i*k+k], m[l*k:l*k+k])
	}
	em, err = matrix(msurvive).invert(len(survived))
	if err != nil {
		return
	}
	return
}

var ErrNotSquare = fmt.Errorf("not a square matrix")
var ErrSingularMatrix = fmt.Errorf("matrix is singular")

// invert calculates m's inverse matrix,
// and return it or any error.
func (m matrix) invert(k int) (inv matrix, err error) {
	if k*k != len(m) {
		err = ErrNotSquare
		return
	}

	mm := make([]bls.Fr, 2*k*k)
	left := mm[:k*k]
	copy(left, m) // Copy m, avoiding side effect.

	// Make an identity matrix.
	inv = mm[k*k:]
	for i := 0; i < k; i++ {
		inv[i*k+i].SetOne()
	}

	for i := 0; i < k; i++ {
		// Pivoting.
		if left[i*k+i].IsZero() {
			// Find a row with non-zero in current column and swap.
			// If there is no one, means it's a singular matrix.
			var j int
			for j = i + 1; j < k; j++ {
				if !left[j*k+i].IsZero() {
					break
				}
			}
			if j == k {
				return nil, ErrSingularMatrix
			}

			matrix(left).swap(i, j, k)
			inv.swap(i, j, k)
		}

		if !left[i*k+i].IsOne() {
			var v bls.Fr
			v.Inverse(&left[i*k+i]) // 1/pivot
			// Scale row.
			for j := 0; j < k; j++ {
				left[i*k+j].Mul(&left[i*k+j], &v)
				inv[i*k+j].Mul(&inv[i*k+j], &v)
			}
		}

		// Use elementary transformation to
		// make all elements(except pivot) in the left matrix
		// become 0.
		for j := 0; j < k; j++ {
			if j == i {
				continue
			}

			v := left[j*k+i]
			if !v.IsZero() {
				var temp bls.Fr
				for l := 0; l < k; l++ {
					temp.Mul(&v, &left[i*k+l])
					left[j*k+l].Sub(&left[j*k+l], &temp)

					temp.Mul(&v, &inv[i*k+l])
					inv[j*k+l].Sub(&inv[j*k+l], &temp)
				}
			}
		}
	}

	return
}

// swap square matrix row[i] & row[j], col = n
func (m matrix) swap(i, j, n int) {
	for k := 0; k < n; k++ {
		m[i*n+k], m[j*n+k] = m[j*n+k], m[i*n+k]
	}
}
