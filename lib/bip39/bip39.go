// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

// Package bip39 is the Golang implementation of the BIP39 spec.
//
// The official BIP39 spec can be found at
// https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki
package bip39

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"hash"
	"io"
	"math/big"
	"strings"
	"sync"

	"github.com/zerjioang/gotools/lib/radix"

	"github.com/zerjioang/gotools/lib/bip39/wordlists"

	"github.com/zerjioang/gotools/lib/stack"

	"golang.org/x/crypto/pbkdf2"
)

// supported language names
const (
	ChineseSimplified  = "chinese-simplified"
	ChineseTraditional = "chinese-traditional"
	English            = "english"
	French             = "french"
	Italian            = "italian"
	Japanese           = "japanese"
	Korean             = "korean"
	Spanish            = "spanish"
)

var (
	// Some bitwise operands for working with big.Ints
	last11BitsMask  = big.NewInt(2047)
	shift11BitsMask = big.NewInt(2048)
	bigOne          = big.NewInt(1)
	bigTwo          = big.NewInt(2)

	//checksum mask values
	checksum12 = big.NewInt(15)
	checksum15 = big.NewInt(31)
	checksum18 = big.NewInt(63)
	checksum21 = big.NewInt(127)
	checksum24 = big.NewInt(255)

	// used to use only the desired x of 8 available checksum bits.
	// 256 bit (word length 24) requires all 8 bits of the checksum,
	// and thus no shifting is needed for it (we would get a divByZero crash if we did)
	wordLengthChecksumShiftMapping = map[int]*big.Int{
		12: big.NewInt(16),
		15: big.NewInt(8),
		18: big.NewInt(4),
		21: bigTwo,
	}

	// wordlist lock
	listLock sync.Mutex

	// wordList is the set of words to use
	wordList []string

	// wordMap is a reverse lookup map for wordList
	// wordMap map[string]int
	currentRadixtree *radix.Tree
)

var (
	errInvalidWord = stack.New("word not found in reverse map")
	// ErrInvalidMnemonic is returned when trying to use a malformed mnemonic.
	ErrInvalidMnemonic = stack.New("invalid mnenomic due to malformed input")

	// ErrEntropyLengthInvalid is returned when trying to use an entropy set with
	// an invalid size.
	ErrEntropyLengthInvalid = stack.New("entropy length must be [128, 256] and a multiple of 32")

	// ErrValidatedSeedLengthMismatch is returned when a validated seed is not the
	// same size as the given seed. This should never happen is present only as a
	// sanity assertion.
	ErrValidatedSeedLengthMismatch = stack.New("seed length does not match validated seed length")

	// ErrChecksumIncorrect is returned when entropy has the incorrect checksum.
	ErrChecksumIncorrect = stack.New("checksum incorrect")

	supportedWordlists map[string]*radix.Tree
)

func init() {
	// preload all supported wordlists
	supportedWordlists = map[string]*radix.Tree{
		ChineseSimplified:  initializeInternalWordlist(wordlists.ChineseSimplified),
		ChineseTraditional: initializeInternalWordlist(wordlists.ChineseTraditional),
		English:            initializeInternalWordlist(wordlists.English),
		French:             initializeInternalWordlist(wordlists.French),
		Italian:            initializeInternalWordlist(wordlists.Italian),
		Japanese:           initializeInternalWordlist(wordlists.Japanese),
		Korean:             initializeInternalWordlist(wordlists.Korean),
		Spanish:            initializeInternalWordlist(wordlists.Spanish),
	}
	// set default language to english
	currentRadixtree = initializeInternalWordlist(wordlists.English)
	// initialize list lock
	listLock = sync.Mutex{}
}

// SetWordList sets the list of words to use for mnemonics. Currently the list
// that is set is used package-wide.
func initializeInternalWordlist(list []string) *radix.Tree {
	wordList = list
	tree := radix.New()
	for i, v := range wordList {
		tree.Insert(v, i)
	}
	return tree
}

// SetWordList sets the list of words to use for mnemonics. Currently the list
// that is set is used package-wide.
func SetWordList(language string) {
	tree, ok := supportedWordlists[language]
	if ok {
		listLock.Lock()
		currentRadixtree = tree
		listLock.Unlock()
	}
}

// GetWordList gets the list of words to use for mnemonics.
func GetWordList() []string {
	return wordList
}

// GetWordIndex gets word index in wordMap.
func GetWordIndexFromTree(word string) (int, bool) {
	idx, ok := currentRadixtree.Get(word)
	if ok {
		return idx.(int), ok
	}
	return 0, false
}

func HasWord(word string) bool {
	_, ok := currentRadixtree.Get(word)
	return ok
}

// NewEntropy will create random entropy bytes
// so long as the requested size bitSize is an appropriate size.
//
// bitSize is the size of entropy bytes requested
func GenerateSecureEntropy(entropyBits uint16) ([]byte, error) {
	entropyBytes := entropyBits / 8
	raw := make([]byte, entropyBytes)
	_, rErr := io.ReadFull(rand.Reader, raw)
	return raw, rErr
}

// NewEntropy will create random entropy bytes
// so long as the requested size bitSize is an appropriate size.
//
// bitSize has to be a multiple 32 and be within the inclusive range of {128, 256}
func NewEntropy(bitSize int) ([]byte, stack.Error) {
	err := validateEntropyBitSize(bitSize)
	if err.Occur() {
		return nil, err
	}

	entropy := make([]byte, bitSize/8)
	_, rErr := rand.Read(entropy)
	return entropy, stack.Ret(rErr)
}

// EntropyFromMnemonic takes a mnemonic generated by this library,
// and returns the input entropy used to generate the given mnemonic.
// An stack is returned if the given mnemonic is invalid.
func EntropyFromMnemonic(mnemonic string) ([]byte, stack.Error) {
	mnemonicSlice, isValid := splitMnemonicWords(mnemonic)
	if !isValid {
		return nil, ErrInvalidMnemonic
	}

	// Decode the words into a big.Int.
	b := big.NewInt(0)
	for _, v := range mnemonicSlice {
		index, found := GetWordIndexFromTree(v)
		if found == false {
			return nil, errInvalidWord
		}
		var wordBytes [2]byte
		binary.BigEndian.PutUint16(wordBytes[:], uint16(index))
		b = b.Mul(b, shift11BitsMask)
		b = b.Or(b, big.NewInt(0).SetBytes(wordBytes[:]))
	}

	// Build and add the checksum to the big.Int.
	checksum := big.NewInt(0)
	checksumMask := resolveChecksumMask(len(mnemonicSlice))
	if checksumMask == nil {
		return nil, ErrInvalidMnemonic
	}
	checksum = checksum.And(b, checksumMask)
	b.Div(b, big.NewInt(0).Add(checksumMask, bigOne))

	// The entropy is the underlying bytes of the big.Int. Any upper bytes of
	// all 0's are not returned so we pad the beginning of the slice with empty
	// bytes if necessary.
	entropy := b.Bytes()
	entropy = padByteSlice(entropy, len(mnemonicSlice)/3*4)

	// Generate the checksum and compare with the one we got from the mneomnic.
	entropyChecksumBytes := computeChecksum(entropy)
	entropyChecksum := big.NewInt(int64(entropyChecksumBytes[0]))
	if l := len(mnemonicSlice); l != 24 {
		checksumShift := wordLengthChecksumShiftMapping[l]
		entropyChecksum.Div(entropyChecksum, checksumShift)
	}

	if checksum.Cmp(entropyChecksum) != 0 {
		return nil, ErrChecksumIncorrect
	}

	return entropy, stack.Nil()
}

// used to isolate the checksum bits from the entropy+checksum byte array
func resolveChecksumMask(value int) *big.Int {
	switch value {
	case 12:
		return checksum12
	case 15:
		return checksum15
	case 18:
		return checksum18
	case 21:
		return checksum21
	case 24:
		return checksum24
	}
	return nil
}

// NewMnemonic will return a string consisting of the mnemonic words for
// the given entropy.
// If the provide entropy is invalid, an stack error will be returned.
func NewMnemonic(entropy []byte) (string, stack.Error) {
	// Compute some lengths for convenience.
	entropyBitLength := len(entropy) * 8
	checksumBitLength := entropyBitLength / 32
	sentenceLength := (entropyBitLength + checksumBitLength) / 11

	// Validate that the requested size is supported.
	err := validateEntropyBitSize(entropyBitLength)
	if err.Occur() {
		return "", err
	}

	// Add checksum to entropy.
	entropy = addChecksum(entropy)

	// Break entropy up into sentenceLength chunks of 11 bits.
	// For each word AND mask the rightmost 11 bits and find the word at that index.
	// Then bitshift entropy 11 bits right and repeat.
	// Add to the last empty slot so we can work with LSBs instead of MSB.

	// Entropy as an int so we can bitmask without worrying about bytes slices.
	entropyInt := big.NewInt(0).SetBytes(entropy)

	// Slice to hold words in.
	words := make([]string, sentenceLength)

	// Throw away big.Int for AND masking.
	word := big.NewInt(0)

	for i := sentenceLength - 1; i >= 0; i-- {
		// Get 11 right most bits and bitshift 11 to the right for next time.
		word.And(entropyInt, last11BitsMask)
		entropyInt.Div(entropyInt, shift11BitsMask)

		// Get the bytes representing the 11 bits as a 2 byte slice.
		wordBytes := padByteSlice(word.Bytes(), 2)
		// Convert bytes to an index and add that word to the list.
		words[i] = wordList[binary.BigEndian.Uint16(wordBytes)]
	}

	return strings.Join(words, " "), stack.Nil()
}

// MnemonicToByteArray takes a mnemonic string and turns it into a byte array
// suitable for creating another mnemonic.
// An error is returned if the mnemonic is invalid.
func MnemonicToByteArray(mnemonic string, raw ...bool) ([]byte, stack.Error) {
	var (
		mnemonicSlice    = strings.Split(mnemonic, " ")
		entropyBitSize   = len(mnemonicSlice) * 11
		checksumBitSize  = entropyBitSize % 32
		fullByteSize     = (entropyBitSize-checksumBitSize)/8 + 1
		checksumByteSize = fullByteSize - (fullByteSize % 4)
	)

	// Pre validate that the mnemonic is well formed and only contains words that
	// are present in the word list.
	if !IsMnemonicValid(mnemonic) {
		return nil, ErrInvalidMnemonic
	}

	// Convert word indices to a big.Int representing the entropy.
	checksummedEntropy := big.NewInt(0)
	modulo := big.NewInt(2048)
	for _, v := range mnemonicSlice {
		value, _ := GetWordIndexFromTree(v)
		index := big.NewInt(int64(value))
		checksummedEntropy.Mul(checksummedEntropy, modulo)
		checksummedEntropy.Add(checksummedEntropy, index)
	}

	// Calculate the unchecksummed entropy so we can validate that the checksum is
	// correct.
	checksumModulo := big.NewInt(0).Exp(bigTwo, big.NewInt(int64(checksumBitSize)), nil)
	rawEntropy := big.NewInt(0).Div(checksummedEntropy, checksumModulo)

	// Convert big.Ints to byte padded byte slices.
	rawEntropyBytes := padByteSlice(rawEntropy.Bytes(), checksumByteSize)
	checksummedEntropyBytes := padByteSlice(checksummedEntropy.Bytes(), fullByteSize)

	// Validate that the checksum is correct.
	newChecksummedEntropyBytes := padByteSlice(addChecksum(rawEntropyBytes), fullByteSize)
	if !compareByteSlices(checksummedEntropyBytes, newChecksummedEntropyBytes) {
		return nil, ErrChecksumIncorrect
	}

	if len(raw) > 0 && raw[0] {
		return rawEntropyBytes, stack.Nil()
	}

	return checksummedEntropyBytes, stack.Nil()
}

// NewSeedWithErrorChecking creates a hashed seed output given the mnemonic string and a password.
// An stack is returned if the mnemonic is not convertible to a byte array.
func NewSeedWithErrorChecking(mnemonic string, password string) ([]byte, stack.Error) {
	_, err := MnemonicToByteArray(mnemonic)
	if err.Occur() {
		return nil, err
	}
	return NewSeed(mnemonic, password), stack.Nil()
}

// NewSeed creates a hashed seed output given a provided string and password.
// No checking is performed to validate that the string provided is a valid mnemonic.
func NewSeed(mnemonic string, password string) []byte {
	return NewSeedWithParams([]byte(mnemonic), []byte("mnemonic"+password), 2048, 64, sha512.New)
}

// NewSeed creates a hashed seed output given a provided string and password.
// No checking is performed to validate that the string provided is a valid mnemonic.
func NewSeedWithParams(mnemonic []byte, password []byte, iterations int, keyLength int, hashf func() hash.Hash) []byte {
	return pbkdf2.Key(mnemonic, password, iterations, keyLength, hashf)
}

// IsMnemonicValid attempts to verify that the provided mnemonic is valid.
// Validity is determined by both the number of words being appropriate,
// and that all the words in the mnemonic are present in the word list.
func IsMnemonicValid(mnemonic string) bool {
	// Create a list of all the words in the mnemonic sentence
	words := strings.Fields(mnemonic)

	// Get word count
	wordCount := len(words)

	// The number of words should be 12, 15, 18, 21 or 24
	if wordCount%3 != 0 || wordCount < 12 || wordCount > 24 {
		return false
	}

	// Check if all words belong in the wordlist
	for _, word := range words {
		ok := HasWord(word)
		if !ok {
			return false
		}
	}

	return true
}

// Appends to data the first (len(data) / 32)bits of the result of sha256(data)
// Currently only supports data up to 32 bytes
func addChecksum(data []byte) []byte {
	// Get first byte of sha256
	checksum := computeChecksum(data)
	firstChecksumByte := checksum[0]

	// len() is in bytes so we divide by 4
	checksumBitLength := uint(len(data) / 4)

	// For each bit of check sum we want we shift the data one the left
	// and then set the (new) right most bit equal to checksum bit at that index
	// staring from the left
	dataBigInt := new(big.Int).SetBytes(data)
	for i := uint(0); i < checksumBitLength; i++ {
		// Bitshift 1 left
		dataBigInt.Mul(dataBigInt, bigTwo)

		// Set rightmost bit if leftmost checksum bit is set
		if uint8(firstChecksumByte&(1<<(7-i))) > 0 {
			dataBigInt.Or(dataBigInt, bigOne)
		}
	}

	return dataBigInt.Bytes()
}

func computeChecksum(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}

// validateEntropyBitSize ensures that entropy is the correct size for being a
// mnemonic.
func validateEntropyBitSize(bitSize int) stack.Error {
	if (bitSize%32) != 0 || bitSize < 128 || bitSize > 256 {
		return ErrEntropyLengthInvalid
	}
	return stack.Nil()
}

// padByteSlice returns a byte slice of the given size with contents of the
// given slice left padded and any empty spaces filled with 0's.
func padByteSlice(slice []byte, length int) []byte {
	offset := length - len(slice)
	if offset <= 0 {
		return slice
	}
	newSlice := make([]byte, length)
	copy(newSlice[offset:], slice)
	return newSlice
}

// compareByteSlices returns true of the byte slices have equal contents and
// returns false otherwise.
func compareByteSlices(a, b []byte) bool {
	//return bytes.Compare(a, b)==0
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func splitMnemonicWords(mnemonic string) ([]string, bool) {
	// Create a list of all the words in the mnemonic sentence
	var words []string
	//words = strings.Fields(mnemonic)
	words = strings.Split(mnemonic, " ")

	// Get num of words
	numOfWords := len(words)

	// The number of words should be 12, 15, 18, 21 or 24
	invalidCount := numOfWords < 12 || numOfWords%3 != 0 || numOfWords > 24
	return words, !invalidCount
}
