// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	// tips: 初期化が重要なことを伝えるには明示的な初期化を使い、初期値が問題にならない場合には暗黙的な初期化を使う
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		// ループを回すごとに文字列sは完全に再構築される。sの古い内容はガベージコレクションによって解放される。
		// 関係するデータの量が大きい場合にはこの方法は非効率のため、strings.Join関数を使うべき(ex: echo3/main.go)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

//!-
