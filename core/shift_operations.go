package core

type ShiftFunc func (uint32, uint8) (uint32, bool)

/* Perform shift operation, updating condition codes */
func ShiftOp(regs *Registers, value uint32, shift_n uint8, S bool, do_shift ShiftFunc) uint32 {
	var result uint32
	var carry_out bool

	if shift_n == 0 {
		result, carry_out = value, regs.Apsr.C
	} else {
		result, carry_out = do_shift(value, shift_n)
	}

	if S {
		regs.Apsr.N = (result & 0x80000000) != 0
		regs.Apsr.Z = (result) == 0
		regs.Apsr.C = carry_out
	}

	return result
}

/* Perform LSL instruction, updating condition codes */
func LSL(regs *Registers, value uint32, shift_n uint8, S bool) uint32 {
	return ShiftOp(regs, value, shift_n, S, LSL_C)
}

/* Left shift value by a positive amount */
func LSL_C(value uint32, amount uint8) (uint32, bool) {
	extended := uint64(value)

	extended = extended << amount

	result := uint32(extended & 0xffffffff)
	carry_out := (extended & 0x100000000) != 0

	return result, carry_out
}
