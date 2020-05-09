// compilation instructions
// gcc -O3 -o parser parser.c

// to get assembly code, add -S flag
// gcc -O3 -S parser.c

// clang -S -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti -masm=intel -mno-red-zone -O3 -m64 -mavx -mavx2 -msse4.1 -Wall -Wextra -mstackrealign -mllvm -inline-threshold=1000
// c2goasm -a -c -s -f parser.s parser_amd64.s

// sources of implementation
// https://github.com/zserge/jsmn/blob/master/jsmn.h
// https://github.com/forkachild/C-Simple-JSON-Parser/blob/master/json.h
// https://github.com/forkachild/C-Simple-JSON-Parser/blob/master/json.c

#ifndef JSONBOOST_H
#define JSONBOOST_H

#include <stdio.h>

// definitions for c++ style code
#ifndef __cplusplus
    typedef char*                                   string;
    typedef unsigned char                           bool;
    #define true                                    (1)
    #define false                                   (0)
    #define TRUE                                    true
    #define FALSE                                   false
#endif

#ifdef JSONBOOST_STATIC
    #define JSONBOOST_API static
#else
    #define JSONBOOST_API extern
#endif

// helper definitions
#define new(x)                                  (x *) malloc(sizeof(x))
#define newWithSize(x, y)                       (x *) malloc(y * sizeof(x))
#define renewWithSize(x, y, z)                  (y *) realloc(x, z * sizeof(y))
#define isWhitespace(x)                         x == '\r' || x == '\n' || x == '\t' || x == ' '
#define removeWhitespace(x)                     while(isWhitespace(*x)) x++
#define removeWhitespaceCalcOffset(x, y)        while(isWhitespace(*x)) { x++; y++; }

typedef char                                    character;

// empty string
#define NONE ""
// empty json object
#define NONE_JSON_OBJ "{}"

// empty json array
#define NONE_JSON_ARRAY "[]"

// internal variables
#define BUFFER_SIZE 1024

// success code
#define SUCCESS 0

// error codes
#define INVALID_JSON_LEN 1
#define INVALID_KEY_LEN 2
#define INVALID_JSON_DATA_RECEIVED 3

enum jsonerr {
  /* Not enough tokens were provided */
  JSON_ERROR_NOMEM = -1,
  /* Invalid character inside JSON string */
  JSON_ERROR_INVAL = -2,
  /* The string is not a full JSON packet, more bytes expected */
  JSON_ERROR_PART = -3
};

// define json gramar
/**
 * JSON type identifier. Basic types are:
 * 	o Object
 * 	o Array
 * 	o String
 * 	o Other primitive: number, boolean (true/false) or null
 */
typedef enum {
  JSON_UNDEFINED = 0,
  JSON_OBJECT = 1,
  JSON_ARRAY = 2,
  JSON_STRING = 3,
  JSON_PRIMITIVE = 4
} jsontype_t;

/**
 * JSON token description.
 * type		type (object, array, string etc.)
 * start	start position in JSON data string
 * end		end position in JSON data string
 */
typedef struct {
    jsontype_t type; /* Token type */
    int start;       /* Token start position */
    int end;         /* Token end position */
    int size;        /* Number of child (nested) tokens */
    #ifdef JSON_PARENT_LINKS
      int parent;
    #endif
} jsontoken_t;

/**
 * JSON parser. Contains an array of token blocks available. Also stores
 * the string being parsed now and current position in that string.
 */
typedef struct {
  unsigned int pos;     /* offset in the JSON string */
  unsigned int toknext; /* next token to allocate */
  int toksuper;         /* superior token node, e.g. parent object or array */
} json_parser;

// this method will extract content from array given
// start, end and steps positions
int parts(char* src, char* dst, int start, int end, int steps){
     int i, j;
     // for (i = start , j = 0 ; i <= end ; i += steps , ++j)
     for (i = start , j = 0 ; i < end ; i += steps , ++j)
           dst[j] = src[i];
     //for (i = 0 ; i < (end - start)/steps + 1 ; ++i)
     //      dst[i] = temp[i];
    return (end - start)/steps + 1;
}

// validate will check for json string validity and return error if json
// is syntactically invalid
int validate(char json[], int jsonLen){
    // minimum valid json document is -> {}, []
    // we skip as valid single item json
    if (jsonLen < 2) {
       return INVALID_JSON_LEN;
    }
}

// iterates over a json string looking for specified key
// @returns status code. any value != 0 means error
int lookup(char json[], char key[], int jsonLen, int keyLen, char *result) {
    // minimum valid json document is -> {}, []
    if (jsonLen < 2) {
       result = NONE;
       return INVALID_JSON_LEN;
    }
    // minimum key is a single char
    if (keyLen < 1) {
       result = NONE;
       return INVALID_KEY_LEN;
    }
    // 1. Initial check! validate if json starts with [ or { and ends with ] or }
    int valid = (json[0] == '[' || json[0] == '{') && (json[jsonLen-1] == ']' || json[jsonLen-1] == '}');
    if(!valid){
        result = NONE;
        return INVALID_JSON_DATA_RECEIVED;
    }
    // allocate our buffer
    char dst[BUFFER_SIZE];

    // start processing input json
    for (int i = 0; i < jsonLen; ++i){

    }
    // extract parts
    int slice = parts(json, dst, 0, jsonLen, 5);
    return SUCCESS;
}

#endif /* JSONBOOST_H */