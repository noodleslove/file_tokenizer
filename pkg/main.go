package main

import (
	"fmt"
	"os"

	. "fake.com/file_tokenizer/pkg/f_tokenizer"
	"fake.com/string_tokenizer/pkg/tokenizer"
)

func main() {
	ftk := NewFileTokenizer("solitude.txt")
	var t *tokenizer.Token = nil

	f, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	f.WriteString(fmt.Sprintf(
		"%9s %25s %10s %10s %10s\n",
		"Index", "Word", "Length", "BlockPos", "Pos",
	))

	tokenCount := 0
	t = ftk.Tokenize()
	for ftk.More() {
		if t.TypeStr() == "ALPHA" {
			s := fmt.Sprintf(
				"%8d: %25s %10d %10d %10d\n",
				tokenCount, t.TokenStr(), len(t.TokenStr()), ftk.BlockPos(), ftk.Pos(),
			)
			_, err := f.WriteString(s)
			if err != nil {
				panic(err)
			}
			tokenCount++
		}
		t = ftk.Tokenize()
	}

	f.Sync()
	fmt.Printf("Word count: %d\n", tokenCount)
}
