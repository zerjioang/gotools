// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package bip39

import (
	"encoding/hex"
	"testing"

	"github.com/zerjioang/gotools/lib/bip39/wordlists"
)

func BenchmarkBip39(b *testing.B) {
	b.Run("bip39-generate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		entropy, _ := hex.DecodeString("066dca1a2bb7e8a1db2832148ce9933eea0f3ac9548d793112d9a95c9407efad")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = NewMnemonic(entropy)
		}
	})
	b.Run("is-valid", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsMnemonicValid("all hour make first leader extend hole alien behind guard gospel lava path output census museum junior mass reopen famous sing advance salt reform")
		}
	})
	b.Run("initialize-lists", func(b *testing.B) {
		b.Run("ChineseSimplified", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				initializeInternalWordlist(wordlists.ChineseSimplified)
			}
		})
		b.Run("ChineseTraditional", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				initializeInternalWordlist(wordlists.ChineseTraditional)
			}
		})
		b.Run("English", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				initializeInternalWordlist(wordlists.English)
			}
		})
		b.Run("French", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				initializeInternalWordlist(wordlists.French)
			}
		})
		b.Run("Italian", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				initializeInternalWordlist(wordlists.Italian)
			}
		})
		b.Run("Japanese", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				initializeInternalWordlist(wordlists.Japanese)
			}
		})
		b.Run("Spanish", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				initializeInternalWordlist(wordlists.Spanish)
			}
		})
	})
	b.Run("set-wordlists", func(b *testing.B) {
		b.Run("ChineseSimplified", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SetWordList("chinese-simplified")
			}
		})
		b.Run("ChineseTraditional", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SetWordList("chinese-traditional")
			}
		})
		b.Run("English", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SetWordList("english")
			}
		})
		b.Run("French", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SetWordList("french")
			}
		})
		b.Run("Italian", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SetWordList("italian")
			}
		})
		b.Run("Japanese", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SetWordList("japanese")
			}
		})
		b.Run("Spanish", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SetWordList("spanish")
			}
		})
	})
	b.Run("get-word-list", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			GetWordList()
		}
	})
	b.Run("get-word-list-from-tree", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			GetWordIndexFromTree("")
		}
	})
	b.Run("has-word", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			HasWord("")
		}
	})
	b.Run("generate-secure-entropy", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = GenerateSecureEntropy(256)
		}
	})
	b.Run("new-entropy", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = NewEntropy(256)
		}
	})
	b.Run("entropy-from-mnemonic", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = EntropyFromMnemonic("letter advice cage absurd amount doctor acoustic avoid letter advice cage above")
		}
	})
	b.Run("resolve-mask", func(b *testing.B) {
		b.Run("case-12", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				resolveChecksumMask(12)
			}
		})
		b.Run("case-15", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				resolveChecksumMask(15)
			}
		})
		b.Run("case-18", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				resolveChecksumMask(18)
			}
		})
		b.Run("case-21", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				resolveChecksumMask(21)
			}
		})
		b.Run("case-24", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				resolveChecksumMask(24)
			}
		})
		b.Run("case-invalid", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				resolveChecksumMask(99)
			}
		})
	})
	b.Run("new-mnemonic", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		entropy, _ := NewEntropy(256)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = NewMnemonic(entropy)
		}
	})
	b.Run("new-mnemonic-to-byte-array", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = MnemonicToByteArray("letter advice cage absurd amount doctor acoustic avoid letter advice cage above")
		}
	})
	b.Run("NewSeedWithErrorChecking", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = NewSeedWithErrorChecking("letter advice cage absurd amount doctor acoustic avoid letter advice cage above", "passw0rd-h3r3")
		}
	})
	b.Run("NewSeed", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewSeed("letter advice cage absurd amount doctor acoustic avoid letter advice cage above", "passw0rd-h3r3")
		}
	})
	b.Run("IsMnemonicValid", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsMnemonicValid("letter advice cage absurd amount doctor acoustic avoid letter advice cage above")
		}
	})
	b.Run("addChecksum", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
		}
	})

	b.Run("computeChecksum", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
		}
	})

	b.Run("validateEntropyBitSize", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = validateEntropyBitSize(128)
		}
	})

	b.Run("padByteSlice", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
		}
	})

	b.Run("compareByteSlices", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
		}
	})

	// BenchmarkBip39/splitMnemonicWords-4         	 3000000	       393 ns/op	   2.54 MB/s	     192 B/op	       1 allocs/op
	b.Run("splitMnemonicWords", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = splitMnemonicWords("letter advice cage absurd amount doctor acoustic avoid letter advice cage above")
		}
	})
}
