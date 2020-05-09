// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package solc

import (
	"os/exec"
	"testing"
)

const (
	testSource = `
pragma solidity >0.0.0;
contract test {
   /// @notice Will multiply ` + "`a`" + ` by 7.
   function multiply(uint a) public returns(uint d) {
       return a * 7;
   }
}
`
)

func skipWithoutSolc(t *testing.T) {
	if _, err := exec.LookPath("solc"); err != nil {
		t.Skip(err)
	}
}

func TestSolidityCompiler(t *testing.T) {
	skipWithoutSolc(t)

	contracts, err := CompileSolidityString(testSource)
	if err != nil {
		t.Fatalf("error compiling source. result %v: %v", contracts, err)
	}
	if len(contracts) != 1 {
		t.Errorf("one contract expected, got %d", len(contracts))
	}
	c, ok := contracts["test"]
	if !ok {
		c, ok = contracts["<stdin>:test"]
		if !ok {
			t.Fatal("info for contract 'test' not present in result")
		}
	}
	if c.Code == "" {
		t.Error("empty code")
	}
	if c.Info.Source != testSource {
		t.Error("wrong source")
	}
	if c.Info.CompilerVersion == "" {
		t.Error("empty version")
	}
}

func TestSolidityCompileError(t *testing.T) {
	skipWithoutSolc(t)

	contracts, err := CompileSolidityString(testSource[4:])
	if err == nil {
		t.Errorf("error expected compiling source. got none. result %v", contracts)
	}
	t.Logf("error: %v", err)
}

func TestCompileSolidity(t *testing.T) {
	t.Run("compile-from-single-string-raw", func(t *testing.T) {
		raw := `pragma solidity >=0.4.0 <0.6.0;
contract Voting {
  // constructor to initialize candidates
  // vote for candidates
  // get count of votes for each candidates

  bytes32[] public candidateList;
  mapping (bytes32 => uint8) public votesReceived;

  constructor(bytes32[] memory candidateNames) public {
    // solidity requires that any 
    candidateList = candidateNames;
  }

  function voteForCandidate(bytes32 candidate) public {
    require(validCandidate(candidate));
    votesReceived[candidate] += 1;
  }  

  function totalVotesFor(bytes32 candidate) view public returns(uint8) {
    require(validCandidate(candidate));
    return votesReceived[candidate];
  }

  function validCandidate(bytes32 candidate) view public returns (bool) {
    for(uint i=0; i < candidateList.length; i++) {
      if (candidateList[i] == candidate) {
        return true;
      }
    }
    return false;
  }
}`
		contracts, err := CompileSolidityString(raw)
		if err != nil {
			t.Fatalf("error compiling source. result %v: %v", contracts, err)
		}
		if len(contracts) != 1 {
			t.Errorf("one contract expected, got %d", len(contracts))
		}
		c, ok := contracts["Voting"]
		if !ok {
			c, ok = contracts["<stdin>:Voting"]
			if !ok {
				t.Fatal("info for contract 'Voting' not present in result")
			}
		}
		if c.Code == "" {
			t.Error("empty code")
		}
		if c.Info.Source != raw {
			t.Error("wrong source")
		}
		if c.Info.CompilerVersion == "" {
			t.Error("empty version")
		}
	})
}
