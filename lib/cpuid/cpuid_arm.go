// +build arm arm64 arm64be armbe

package cpuid

const (
	notImplemented = 0
)

// cpuid executes the CPUID instruction to obtain processor identification and
// feature information.
// cpuid executes the CPUID instruction with the given EAX, ECX inputs.
func cpuid(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32) {
	return notImplemented, notImplemented, notImplemented, notImplemented
}

// xgetbv executes the XGETBV instruction.
func xgetbv() (eax, edx uint32) {
	return notImplemented, notImplemented
}
