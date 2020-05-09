// +build 386 amd64

package cpuid

// NOTE: assembler functions

// cpuid executes the CPUID instruction to obtain processor identification and
// feature information.
// cpuid executes the CPUID instruction with the given EAX, ECX inputs.
//go:noescape
func cpuid(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32)

// xgetbv executes the XGETBV instruction.
//go:noescape
func xgetbv() (eax, edx uint32)
