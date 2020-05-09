package cpuid

type CpuFeatures struct {
	HasBMI2     bool `json:"bmi_2"`
	HasOSXSAVE  bool `json:"osxsave"`
	HasAES      bool `json:"aes"`
	HasVAES     bool `json:"vaes"`
	HasAVX      bool `json:"avx"`
	HasAVX512F  bool `json:"avx_512_f"`
	HasAVX512VL bool `json:"avx_512_vl"`
	HasAVX512DQ bool `json:"avx_512_dq"`

	HasSSE2      bool `json:"sse_2"`
	HasSSE3      bool `json:"sse_3"`
	HasPCLMULQDQ bool `json:"pclmulqdq"`
	HasSSSE3     bool `json:"ssse_3"`
	HasFMA       bool `json:"fma"`
	HasSSE41     bool `json:"sse_41"`
	HasSSE42     bool `json:"sse_42"`
	HasPOPCNT    bool `json:"popcnt"`
	HasBMI1      bool `json:"bmi_1"`
	HasAVX2      bool `json:"avx_2"`
	HasERMS      bool `json:"erms"`
	HasADX       bool `json:"adx"`

	EnabledAVX    bool `json:"enabled_avx"`
	EnabledAVX512 bool `json:"enabled_avx_512"`
}

var (
	// cpu contains feature flags relevant to selecting a Meow implementation.
	cpuFeatures CpuFeatures
)

// depetermine cpu features
// init determines whether to use assembly version by performing CPU feature
// check.
func init() {
	determineCPUFeatures()

	switch {
	case cpuFeatures.HasVAES && cpuFeatures.HasAVX512F && cpuFeatures.EnabledAVX512:
		break
	case cpuFeatures.HasAES && cpuFeatures.HasAVX && cpuFeatures.EnabledAVX:
		// AVX required for VEX-encoded AES instruction, which allows non-aligned memory addresses.
		break
	}
}

func GetCpuFeatures() CpuFeatures {
	return cpuFeatures
}
func GetCpuFeaturesPtr() *CpuFeatures {
	return &cpuFeatures
}

// determineCPUFeatures populates flags in global cpu variable by querying CPUID.
func determineCPUFeatures() {
	maxID, _, _, _ := cpuid(0, 0)
	if maxID < 1 {
		return
	}

	_, _, ecx1, edx1 := cpuid(1, 0)

	cpuFeatures.HasSSE2 = isSet(edx1, 26)
	c1 := isSet(ecx1, 0)
	cpuFeatures.HasSSE3 = c1
	cpuFeatures.HasPCLMULQDQ = c1
	cpuFeatures.HasSSSE3 = isSet(ecx1, 9)
	cpuFeatures.HasFMA = isSet(ecx1, 12)
	cpuFeatures.HasSSE41 = isSet(ecx1, 19)
	cpuFeatures.HasSSE42 = isSet(ecx1, 20)
	cpuFeatures.HasPOPCNT = isSet(ecx1, 23)
	cpuFeatures.HasAES = isSet(ecx1, 25)
	cpuFeatures.HasOSXSAVE = isSet(ecx1, 27)
	cpuFeatures.HasAVX = isSet(ecx1, 28)

	osSupportsAVX := false
	if cpuFeatures.HasOSXSAVE {
		eax, _ := xgetbv()
		// Check if XMM and YMM registers have OS support.
		cpuFeatures.EnabledAVX = (eax & 0x6) == 0x6
		cpuFeatures.EnabledAVX512 = (eax & 0xe0) == 0xe0
		osSupportsAVX = isSet(eax, 1) && isSet(eax, 2)
	}

	if maxID < 7 {
		return
	}

	_, ebx7, ecx7, _ := cpuid(7, 0)
	cpuFeatures.HasVAES = isSet(ecx7, 9)
	cpuFeatures.HasAVX512F = isSet(ebx7, 16)
	cpuFeatures.HasAVX512VL = isSet(ebx7, 31)
	cpuFeatures.HasAVX512DQ = isSet(ebx7, 17)

	cpuFeatures.HasBMI1 = isSet(ebx7, 3)
	cpuFeatures.HasAVX2 = osSupportsAVX && isSet(ebx7, 5)
	cpuFeatures.HasBMI2 = isSet(ebx7, 8)
	cpuFeatures.HasERMS = isSet(ebx7, 9)
	cpuFeatures.HasADX = isSet(ebx7, 19)
}

// isSet determines if bit i of x is set.
func isSet(x uint32, i uint) bool {
	return (x>>i)&1 == 1
}
