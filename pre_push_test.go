// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	// [memo]NewScanner スキャナを新しく作成
	input := bufio.NewScanner(os.Stdin) // [memo]os.Stdin: ユーザーからのキーボード入力や、他のプログラムからのパイプ入力を読み取ることができます。
	// [memo]Scanメソッドは、次のトークン(行)が存在する場合にtrueを返し、存在しない場合にfalseを返します。Textメソッドは、現在のトークンを文字列として返します。
	for input.Scan() {
		// ゼロ値とは、Go言語で変数が宣言されたときに自動的に割り当てられるデフォルトの値のことを指します。ゼロ値は、変数の型によって異なります。
		// mapが指定されたキーがなくてもその型のゼロ値を返すので、初期化が不要
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-

counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
counts := make(map[string]int)
// [memo]NewScanner スキャナを新しく作成
input := bufio.NewScanner(os.Stdin) // [memo]os.Stdin: ユーザーからのキーボード入力や、他のプログラムからのパイプ入力を読み取ることができます。
// [memo]Scanメソッドは、次のトークン(行)が存在する場合にtrueを返し、存在しない場合にfalseを返します。Textメソッドは、現在のトークンを文字列として返します。
for input.Scan() {
	// ゼロ値とは、Go言語で変数が宣言されたときに自動的に割り当てられるデフォルトの値のことを指します。ゼロ値は、変数の型によって異なります。
	// mapが指定されたキーがなくてもその型のゼロ値を返すので、初期化が不要
	counts[input.Text()]++
}
// NOTE: ignoring potential errors from input.Err()
for line, n := range counts {
	if n > 1 {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
}

//!-	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	counts := make(map[string]int)
	// [memo]NewScanner スキャナを新しく作成
	input := bufio.NewScanner(os.Stdin) // [memo]os.Stdin: ユーザーからのキーボード入力や、他のプログラムからのパイプ入力を読み取ることができます。
	// [memo]Scanメソッドは、次のトークン(行)が存在する場合にtrueを返し、存在しない場合にfalseを返します。Textメソッドは、現在のトークンを文字列として返します。
	for input.Scan() {
		// ゼロ値とは、Go言語で変数が宣言されたときに自動的に割り当てられるデフォルトの値のことを指します。ゼロ値は、変数の型によって異なります。
		// mapが指定されたキーがなくてもその型のゼロ値を返すので、初期化が不要
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-
