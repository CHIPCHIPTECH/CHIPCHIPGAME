package random 

import (
	"time"
)

// ComplexRNG 
type ComplexRNG struct {
	state uint64   // 
	lfsr  uint64   // 
	lcg   uint64   // 
}

// NewComplexRNG 
func NewComplexRNG(seed uint64) *ComplexRNG {
	if seed == 0 {
		// 
		seed = uint64(time.Now().UnixNano())
	}
	
	// 
	return &ComplexRNG{
		state: seed ^ 0x123456789ABCDEF,
		lfsr:  (seed << 32) | (seed >> 32),
		lcg:   seed * 6364136223846793005 + 1,
	}
}

// Next 
func (r *ComplexRNG) Next() uint64 {
	// 
	r.lcg = r.lcg*6364136223846793005 + 1442695040888963407

	// 
	r.state ^= r.state >> 12
	r.state ^= r.state << 25
	r.state ^= r.state >> 27

	// 
	feedback := ((r.lfsr >> 0) ^ (r.lfsr >> 2) ^ 
				(r.lfsr >> 3) ^ (r.lfsr >> 5)) & 1
	r.lfsr = (r.lfsr >> 1) | (feedback << 63)

	// 
	result := r.state * 0x2545F4914F6CDD1D
	result += r.lcg ^ r.lfsr
	result = (result >> 32) | (result << 32) // 

	// 
	result ^= result >> 16
	result *= 0x85EBCA77C2B2AE63
	result ^= result >> 13
	result *= 0xC2B2AE3D27D4EB4F
	result ^= result >> 16

	return result
}

// 
func (r *ComplexRNG) Intn(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	return int(r.Next() % uint64(n))
}
