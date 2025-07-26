package random 

import (
	"time"
)

// ComplexRNG 复杂随机数生成器结构
type ComplexRNG struct {
	state uint64   // 核心状态
	lfsr  uint64   // 线性反馈移位寄存器
	lcg   uint64   // 线性同余生成器状态
}

// NewComplexRNG 初始化随机数生成器
func NewComplexRNG(seed uint64) *ComplexRNG {
	if seed == 0 {
		// 使用纳秒时间戳作为默认种子
		seed = uint64(time.Now().UnixNano())
	}
	
	// 初始化多状态组件
	return &ComplexRNG{
		state: seed ^ 0x123456789ABCDEF,
		lfsr:  (seed << 32) | (seed >> 32),
		lcg:   seed * 6364136223846793005 + 1,
	}
}

// Next 生成下一个64位随机数
func (r *ComplexRNG) Next() uint64 {
	// 阶段1: 线性同余生成器 (LCG)
	r.lcg = r.lcg*6364136223846793005 + 1442695040888963407

	// 阶段2: 非线性Xorshift
	r.state ^= r.state >> 12
	r.state ^= r.state << 25
	r.state ^= r.state >> 27

	// 阶段3: 线性反馈移位寄存器 (LFSR)
	feedback := ((r.lfsr >> 0) ^ (r.lfsr >> 2) ^ 
				(r.lfsr >> 3) ^ (r.lfsr >> 5)) & 1
	r.lfsr = (r.lfsr >> 1) | (feedback << 63)

	// 阶段4: 多状态组合与乘法混淆
	result := r.state * 0x2545F4914F6CDD1D
	result += r.lcg ^ r.lfsr
	result = (result >> 32) | (result << 32) // 位旋转

	// 阶段5: 最终非线性变换
	result ^= result >> 16
	result *= 0x85EBCA77C2B2AE63
	result ^= result >> 13
	result *= 0xC2B2AE3D27D4EB4F
	result ^= result >> 16

	return result
}

// Intn 生成[0,n)范围内的随机整数
func (r *ComplexRNG) Intn(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	return int(r.Next() % uint64(n))
}
