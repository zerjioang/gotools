// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package bip39

/*

# BIP39: Mnemonic code for generating deterministic keys

BIPs : This BIP describes the implementation of a mnemonic code or mnemonic sentence - a group of easy to remember words - for the generation of deterministic wallets.
The English-language wordlist for the BIP39 standard has 2048 words, so if the phrase contained only 12 random words, the number of possible combinations would be 204⁸¹² = ²¹³² and the phrase would have 132 bits of security. Actual security of a 12-word BIP39 seed phrase is only 128 bits.
Generally a seed phrase only works with the same wallet software that created it. If storing for a long period of time it's a good idea to write the name of the wallet too.

BIP39: Mnemonic code for generating deterministic keys , BIP39 - used to manage your recovery seed and recovery words. Abstract. This BIP describes the implementation of a mnemonic code or mnemonic sentence - a group of easy to remember words - for the generation of deterministic wallets. It consists of two parts: generating the mnemonic, and converting it into a binary seed. Example: let bip39 = require("bip39");


# initial package performance

BenchmarkBip39/bip39-generate-4         	  200000	     10377 ns/op	   0.10 MB/s	    2512 B/op	      62 allocs/op

after some minor changes

BenchmarkBip39/bip39-generate-4         	  200000	      7832 ns/op	   0.13 MB/s	    2544 B/op	      63 allocs/op
BenchmarkBip39/is-valid-4   		      	  500000	      3797 ns/op	   0.26 MB/s	     384 B/op	       1 allocs/op

## internal BIP 39 functions initial performance

BenchmarkBip39/bip39-generate-4         	  					  200000	      6761 ns/op	   0.15 MB/s	    2544 B/op	      63 allocs/op
BenchmarkBip39/is-valid-4               	  					  500000	      3395 ns/op	   0.29 MB/s	     384 B/op	       1 allocs/op
BenchmarkBip39/initialize-lists/ChineseSimplified-4         	    1000	   1586709 ns/op	   0.00 MB/s	  361000 B/op	    9723 allocs/op
BenchmarkBip39/initialize-lists/ChineseTraditional-4        	    1000	   1662755 ns/op	   0.00 MB/s	  363368 B/op	    9939 allocs/op
BenchmarkBip39/initialize-lists/English-4                   	    2000	   1020722 ns/op	   0.00 MB/s	  391224 B/op	   11281 allocs/op
BenchmarkBip39/initialize-lists/French-4                    	    2000	   1048716 ns/op	   0.00 MB/s	  398760 B/op	   11602 allocs/op
BenchmarkBip39/initialize-lists/Italian-4                   	    2000	   1078969 ns/op	   0.00 MB/s	  395368 B/op	   11393 allocs/op
BenchmarkBip39/initialize-lists/Japanese-4                  	    1000	   1312863 ns/op	   0.00 MB/s	  402824 B/op	   11629 allocs/op
BenchmarkBip39/initialize-lists/Spanish-4                   	    2000	   1149723 ns/op	   0.00 MB/s	  393881 B/op	   11452 allocs/op
BenchmarkBip39/set-wordlists/ChineseSimplified-4            	    1000	   1630522 ns/op	   0.00 MB/s	  361000 B/op	    9723 allocs/op
BenchmarkBip39/set-wordlists/ChineseTraditional-4           	    1000	   1704186 ns/op	   0.00 MB/s	  363368 B/op	    9939 allocs/op
BenchmarkBip39/set-wordlists/English-4                      	    2000	   1135437 ns/op	   0.00 MB/s	  391224 B/op	   11281 allocs/op
BenchmarkBip39/set-wordlists/French-4                       	    2000	   1068540 ns/op	   0.00 MB/s	  398760 B/op	   11602 allocs/op
BenchmarkBip39/set-wordlists/Italian-4                      	    2000	   1077658 ns/op	   0.00 MB/s	  395368 B/op	   11393 allocs/op
BenchmarkBip39/set-wordlists/Japanese-4                     	    1000	   1228736 ns/op	   0.00 MB/s	  402824 B/op	   11629 allocs/op
BenchmarkBip39/set-wordlists/Spanish-4                      	    2000	   1287332 ns/op	   0.00 MB/s	  393880 B/op	   11452 allocs/op
BenchmarkBip39/get-word-list-4                                2000000000	         0.33 ns/op	3068.85 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/get-word-list-from-tree-4                       200000000	         6.65 ns/op	 150.45 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/has-word-4                                      300000000	         6.28 ns/op	 159.19 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/generate-secure-entropy-4                    	 2000000	       711 ns/op	   1.40 MB/s	      32 B/op	       1 allocs/op
BenchmarkBip39/new-entropy-4                                	 2000000	       707 ns/op	   1.41 MB/s	      32 B/op	       1 allocs/op
BenchmarkBip39/entropy-from-mnemonic-4                      	 3000000	       577 ns/op	   1.73 MB/s	     224 B/op	       2 allocs/op
BenchmarkBip39/resolve-mask/case-12-4                       	20000000	       104 ns/op	   9.58 MB/s	      80 B/op	       2 allocs/op
BenchmarkBip39/resolve-mask/case-15-4                       	10000000	       107 ns/op	   9.34 MB/s	      80 B/op	       2 allocs/op
BenchmarkBip39/resolve-mask/case-18-4                       	20000000	       106 ns/op	   9.36 MB/s	      80 B/op	       2 allocs/op
BenchmarkBip39/resolve-mask/case-21-4                       	20000000	       110 ns/op	   9.07 MB/s	      80 B/op	       2 allocs/op
BenchmarkBip39/resolve-mask/case-24-4                       	20000000	       102 ns/op	   9.77 MB/s	      80 B/op	       2 allocs/op
BenchmarkBip39/resolve-mask/case-invalid-4                     500000000	         3.12 ns/op	 320.43 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/new-mnemonic-4                               	  200000	      7074 ns/op	   0.14 MB/s	    2576 B/op	      63 allocs/op
BenchmarkBip39/new-mnemonic-to-byte-array-4                 	 2000000	       838 ns/op	   1.19 MB/s	     384 B/op	       2 allocs/op
BenchmarkBip39/NewSeedWithErrorChecking-4                   	 2000000	       849 ns/op	   1.18 MB/s	     384 B/op	       2 allocs/op
BenchmarkBip39/NewSeed-4                                    	     500	   2776058 ns/op	   0.00 MB/s	    1044 B/op	      10 allocs/op
BenchmarkBip39/IsMnemonicValid-4                            	 3000000	       466 ns/op	   2.14 MB/s	     192 B/op	       1 allocs/op
BenchmarkBip39/splitMnemonicWords-4                         	 5000000	       363 ns/op	   2.75 MB/s	     192 B/op	       1 allocs/op

Optimizations:

1. Declare resolve-mask function values as constants performance is increased as:

BenchmarkBip39/resolve-mask/case-12-4 					2000000000	         0.32 ns/op	3095.78 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/resolve-mask/case-15-4 					2000000000	         0.33 ns/op	2986.61 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/resolve-mask/case-18-4 					2000000000	         0.32 ns/op	3084.64 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/resolve-mask/case-21-4 					2000000000	         0.33 ns/op	3020.36 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/resolve-mask/case-24-4 					2000000000	         0.33 ns/op	3032.85 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/resolve-mask/case-invalid-4        	 	2000000000	         0.33 ns/op	3041.24 MB/s	       0 B/op	       0 allocs/op

No allocations anymore nor operations.
(warning: make sure it works for you in concurrency situations)

2. We have replace now, the algorithm that switch from one list to another improving context switching
when calling BIP39 from external endpoints such as our REST API. Following set-list method optimization is shown:

BenchmarkBip39/bip39-generate-4         	  200000	      6277 ns/op	   0.16 MB/s	    1320 B/op	      60 allocs/op
BenchmarkBip39/is-valid-4               	  500000	      3463 ns/op	   0.29 MB/s	     384 B/op	       1 allocs/op
BenchmarkBip39/initialize-lists/ChineseSimplified-4         	    1000	   1598048 ns/op	   0.00 MB/s	  361001 B/op	    9723 allocs/op
BenchmarkBip39/initialize-lists/ChineseTraditional-4        	    1000	   1708024 ns/op	   0.00 MB/s	  363369 B/op	    9939 allocs/op
BenchmarkBip39/initialize-lists/English-4                   	    1000	   1129950 ns/op	   0.00 MB/s	  391225 B/op	   11281 allocs/op
BenchmarkBip39/initialize-lists/French-4                    	    2000	   1275452 ns/op	   0.00 MB/s	  398761 B/op	   11602 allocs/op
BenchmarkBip39/initialize-lists/Italian-4                   	    1000	   1344074 ns/op	   0.00 MB/s	  395369 B/op	   11393 allocs/op
BenchmarkBip39/initialize-lists/Japanese-4                  	    1000	   1440368 ns/op	   0.00 MB/s	  402825 B/op	   11629 allocs/op
BenchmarkBip39/initialize-lists/Spanish-4                   	    2000	   1266276 ns/op	   0.00 MB/s	  393881 B/op	   11452 allocs/op
BenchmarkBip39/set-wordlists/ChineseSimplified-4            	50000000	        23.8 ns/op	  42.09 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/set-wordlists/ChineseTraditional-4           	50000000	        24.4 ns/op	  41.01 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/set-wordlists/English-4                      	50000000	        26.1 ns/op	  38.38 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/set-wordlists/French-4                       	50000000	        26.9 ns/op	  37.19 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/set-wordlists/Italian-4                      	50000000	        31.8 ns/op	  31.47 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/set-wordlists/Japanese-4                     	50000000	        27.4 ns/op	  36.54 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/set-wordlists/Spanish-4                      	50000000	        36.1 ns/op	  27.68 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/get-word-list-4                              	2000000000	         0.32 ns/op	3101.32 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/get-word-list-from-tree-4                    	500000000	         3.88 ns/op	 258.04 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/has-word-4                                   	500000000	         3.68 ns/op	 271.84 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/generate-secure-entropy-4                    	 2000000	       740 ns/op	   1.35 MB/s	      32 B/op	       1 allocs/op
BenchmarkBip39/new-entropy-4                                	 2000000	       710 ns/op	   1.41 MB/s	      32 B/op	       1 allocs/op
BenchmarkBip39/entropy-from-mnemonic-4                      	 2000000	       582 ns/op	   1.72 MB/s	     192 B/op	       1 allocs/op
BenchmarkBip39/resolve-mask/case-12-4                       	2000000000	         0.35 ns/op	2876.37 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/resolve-mask/case-15-4                       	2000000000	         0.33 ns/op	3032.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/resolve-mask/case-18-4                       	2000000000	         0.34 ns/op	2966.02 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/resolve-mask/case-21-4                       	2000000000	         0.36 ns/op	2805.59 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/resolve-mask/case-24-4                       	2000000000	         0.33 ns/op	3072.90 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/resolve-mask/case-invalid-4                  	2000000000	         0.33 ns/op	2997.77 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/new-mnemonic-4                               	  200000	      6842 ns/op	   0.15 MB/s	    1336 B/op	      60 allocs/op
BenchmarkBip39/new-mnemonic-to-byte-array-4                 	 1000000	      1009 ns/op	   0.99 MB/s	     384 B/op	       2 allocs/op
BenchmarkBip39/NewSeedWithErrorChecking-4                   	 2000000	       958 ns/op	   1.04 MB/s	     384 B/op	       2 allocs/op
BenchmarkBip39/NewSeed-4                                    	     500	   2725293 ns/op	   0.00 MB/s	    1044 B/op	      10 allocs/op
BenchmarkBip39/IsMnemonicValid-4                            	 3000000	       594 ns/op	   1.68 MB/s	     192 B/op	       1 allocs/op
BenchmarkBip39/addChecksum-4                                	2000000000	         0.31 ns/op	3188.26 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/computeChecksum-4                            	2000000000	         0.33 ns/op	3025.96 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/validateEntropyBitSize-4                     	2000000000	         0.34 ns/op	2919.95 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/padByteSlice-4                               	2000000000	         0.35 ns/op	2851.40 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/compareByteSlices-4                          	2000000000	         0.32 ns/op	3133.07 MB/s	       0 B/op	       0 allocs/op
BenchmarkBip39/splitMnemonicWords-4                         	 3000000	       442 ns/op	   2.26 MB/s	     192 B/op	       1 allocs/op

OLD: BenchmarkBip39/set-wordlists/ChineseSimplified-4            	    1000	   1630522 ns/op	   0.00 MB/s	  361000 B/op	    9723 allocs/op
NEW: BenchmarkBip39/set-wordlists/ChineseSimplified-4            	50000000	        23.8 ns/op	  42.09 MB/s	       0 B/op	       0 allocs/op

As you can see, switching is now x50000 times faster with no memory allocations
*/
