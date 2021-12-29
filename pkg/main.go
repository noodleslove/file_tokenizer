package main

import (
	"fmt"
	"os"

	. "github.com/noodleslove/file_tokenizer/pkg/f_tokenizer"
	"github.com/noodleslove/string_tokenizer/pkg/str_tokenizer"
)

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}

	ftk := NewFileTokenizer("solitude.txt", f)
	var t *str_tokenizer.Token = nil

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

	f.WriteString(fmt.Sprintf("Word count: %d\n", tokenCount))
	f.Sync()
}
