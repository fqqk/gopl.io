// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	// 変数が明示的に初期化されない場合には、暗黙的に変数の型に対するゼロ値に初期化される
	// GOのゼロ値について: https://qiita.com/tenntenn/items/c55095585af64ca28ab5
	// int,int8,int16,int32,int64: 0
	// uint,uint8,uint16,uint32,uint64: 0
	// byte,rune,uintptr: 0
	// float32,float64: 0
	// complex64,complex128: 0
	// string: ""(空文字列)
	// bool: false
	// error: nil
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

//!-
