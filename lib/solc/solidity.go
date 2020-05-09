// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package solc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

var (
	ErrVersionParse  = errors.New("can't parse solc version information")
	ErrEmptySource   = errors.New("solc: empty source string")
	ErrNoSourceFiles = errors.New("solc: no source files")
)

// Solidity contains information about the solidity compiler.
type Solidity struct {
	Path        string `json:"path"`
	Version     string `json:"version"`
	FullVersion string `json:"full_version"`
	Major       int    `json:"major"`
	Minor       int    `json:"minor"`
	Patch       int    `json:"patch"`
}

var (
	compiler        *Solidity
	versionErr      error
	compilationArgs []string
	compilerPath    string
)

func init() {
	//load solidity version data once, the first time
	compiler, versionErr = solidityVersion()
	//preload compilation args
	if versionErr == nil {
		compilerPath = compiler.Path
		compilationArgs = append(compiler.makeArgs(), "--")
	}
}

// --combined-output format
type solcOutput struct {
	Contracts map[string]struct {
		BinRuntime                                  string `json:"bin-runtime"`
		SrcMapRuntime                               string `json:"srcmap-runtime"`
		Bin, SrcMap, Abi, Devdoc, Userdoc, Metadata string
	}
	Version string `json:"version"`
}

func (s Solidity) makeArgs() []string {
	p := []string{
		"--combined-json", "bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc",
		"--optimize", // code optimizer switched on
	}
	if s.Major > 0 || s.Minor > 4 || s.Patch > 6 {
		p[1] += ",metadata"
	}
	return p
}

// SolidityVersion runs solc and parses its version output.
func SolidityVersion() (*Solidity, error) {
	return compiler, versionErr
}

// solidityVersion runs solc and parses its version output.
func solidityVersion() (*Solidity, error) {
	var solc = "solc"
	var out bytes.Buffer
	cmd := exec.Command(solc, "--version")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	matches := versionRegexp.FindStringSubmatch(out.String())
	if len(matches) != 4 {
		return nil, ErrVersionParse
	}
	s := &Solidity{Path: cmd.Path, FullVersion: out.String(), Version: matches[0]}
	if s.Major, err = strconv.Atoi(matches[1]); err != nil {
		return nil, err
	}
	if s.Minor, err = strconv.Atoi(matches[2]); err != nil {
		return nil, err
	}
	if s.Patch, err = strconv.Atoi(matches[3]); err != nil {
		return nil, err
	}
	return s, nil
}

// CompileSolidityString builds and returns all the contracts contained within a source string.
func CompileSolidityString(source string) (map[string]*Contract, error) {
	if len(source) == 0 {
		return nil, ErrEmptySource
	}
	s, err := solidityVersion()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(s.Path, append(compilationArgs, "-")...)
	cmd.Stdin = strings.NewReader(source)
	return s.run(cmd, source)
}

// CompileSolidity compiles all given Solidity source files.
func CompileSolidity(sourcefiles ...string) (map[string]*Contract, error) {
	if len(sourcefiles) == 0 {
		return nil, ErrNoSourceFiles
	}
	source, err := slurpFileBytes(sourcefiles)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(compilerPath, append(compilationArgs, sourcefiles...)...)
	return compiler.run(cmd, source)
}

// CompileSolidity compiles all given Solidity source files.
func CompileSolidityFileBytes(sourcefiles []string) (map[string]*Contract, error) {
	var source string
	var err error

	if len(sourcefiles) == 0 {
		return nil, ErrNoSourceFiles
	} else if len(sourcefiles) == 1 {
		source = sourcefiles[0]
		err = nil
	} else {
		source, err = slurpFileBytes(sourcefiles)
	}
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(compilerPath, append(compilationArgs, "-")...)
	cmd.Stdin = strings.NewReader(source)
	return compiler.run(cmd, source)
}

func (s *Solidity) run(cmd *exec.Cmd, source string) (map[string]*Contract, error) {
	var stderr, stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("solc: %v\n%s", err, stderr.Bytes())
	}

	return ParseCombinedJSON(stdout.Bytes(), source, s.Version, s.Version, strings.Join(s.makeArgs(), " "))
}

// ParseCombinedJSON takes the direct output of a solc --combined-output run and
// parses it into a map of string contract name to Contract structs. The
// provided source, language and compiler version, and compiler options are all
// passed through into the Contract structs.
//
// The solc output is expected to contain ABI, source mapping, user docs, and dev docs.
//
// Returns an error if the JSON is malformed or missing data, or if the JSON
// embedded within the JSON is malformed.
func ParseCombinedJSON(combinedJSON []byte, source string, languageVersion string, compilerVersion string, compilerOptions string) (map[string]*Contract, error) {
	var output solcOutput
	if err := json.Unmarshal(combinedJSON, &output); err != nil {
		return nil, err
	}

	// Compilation succeeded, assemble and return the contracts.
	contracts := make(map[string]*Contract)
	for name, info := range output.Contracts {
		// Parse the individual compilation results.
		var abi interface{}
		if err := json.Unmarshal([]byte(info.Abi), &abi); err != nil {
			return nil, fmt.Errorf("solc: error reading abi definition (%v)", err)
		}
		var userdoc interface{}
		if err := json.Unmarshal([]byte(info.Userdoc), &userdoc); err != nil {
			return nil, fmt.Errorf("solc: error reading user doc: %v", err)
		}
		var devdoc interface{}
		if err := json.Unmarshal([]byte(info.Devdoc), &devdoc); err != nil {
			return nil, fmt.Errorf("solc: error reading dev doc: %v", err)
		}
		contracts[name] = &Contract{
			Code:        "0x" + info.Bin,
			RuntimeCode: "0x" + info.BinRuntime,
			Info: ContractInfo{
				Source:          source,
				Language:        "Solidity",
				LanguageVersion: languageVersion,
				CompilerVersion: compilerVersion,
				CompilerOptions: compilerOptions,
				SrcMap:          info.SrcMap,
				SrcMapRuntime:   info.SrcMapRuntime,
				AbiDefinition:   abi,
				UserDoc:         userdoc,
				DeveloperDoc:    devdoc,
				Metadata:        info.Metadata,
			},
		}
	}
	return contracts, nil
}
