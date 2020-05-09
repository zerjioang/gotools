#include "textflag.h"

// void *memcpy(void *dst, const void *src, size_t n)
// DI = dst, SI = src, DX = size
TEXT ·_memcpy(SB), NOSPLIT|NOFRAME, $16-0
	PUSHQ R8
	PUSHQ CX
	XORQ  CX, CX // clear register

// void *memset(void *str, int c, size_t n)
// DI = str, SI = c, DX = size
TEXT ·_memset(SB), NOSPLIT|NOFRAME, $16-0
	PUSHQ CX
    LONG $0x0101f669; WORD $0x0101 // imul esi, 0x1010101
    MOVQ SI, CX
    ROLQ $32, CX
    ORQ CX, SI
	XORQ CX, CX // clear register
