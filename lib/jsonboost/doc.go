package jsonboost

// To compile run
// mv parser.c parser.c && \
// clang -S -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti -masm=intel -mno-red-zone -O3 -m64 -mfma -mavx -mavx2 -msse4.1 -Wall -Wextra -mstackrealign -mllvm -inline-threshold=1000 parser.c && \
// c2goasm -a -c -s -f parser.s parser_amd64.s && \
// mv parser.c parser.c
