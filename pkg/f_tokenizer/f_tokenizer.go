/**
 * File: f_tokenizer.go
 * Author: Eddie Ho
 * Date: 2021-12-27
 * Project: File Tokenzier
 * Purpose: Declare FileTokenizer struct and implement functionalities.
 */

package f_tokenizer

import (
	"fmt"
	"io"
	"os"

	"github.com/noodleslove/string_tokenizer/pkg/str_tokenizer"
)

const MAX_BUFFER int = 199

type FileTokenizer struct {
	inFile   *os.File // file being tokenized
	outFile  *os.File // file being outputed
	pos      int      // Current position in the file
	blockPos int      // Current position in the current block
	more     bool     // false if last token of the last block
	//  has been processed and now we are at
	//  the end of the last block.
	strToken *str_tokenizer.StrTokenizer // The StrTokenizer object to tokenize current block
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * Postcondition: Construct a new FileTokenizer object
 *
 * @param fname
 */
func NewFileTokenizer(inName string, out *os.File) *FileTokenizer {
	in, err := os.OpenFile(inName, os.O_RDONLY, 0644)
	check(err)

	p := FileTokenizer{
		inFile:   in,
		outFile:  out,
		pos:      0,
		blockPos: 0,
		more:     false,
		strToken: str_tokenizer.NewStrTokenizer(),
	}
	p.GetNewBlock()
	return &p
}

/**
 * Postcondition: The return value is true if there are more token to get in
 *      the text file. Return false otherwise, meaning it has already extracted
 *      the last token from the text file.
 *
 * @return true
 * @return false
 */
func (f *FileTokenizer) More() bool {
	return f.more
}

/**
 * Postcondition: The return value is the cursor position of the whole file.
 *
 * @return int
 */
func (f *FileTokenizer) Pos() int {
	return f.pos + f.blockPos
}

/**
 * Postcondition: The return value is the cursor position of the current block.
 *
 * @return int
 */
func (f *FileTokenizer) BlockPos() int {
	return f.blockPos
}

/**
 * Postcondition: The return value is the next token in the block. If the block
 *      reaches its end, then it get a new block and get the next token. When
 *      there is no more block, it set _more to false and return a emtpy token.
 *
 * @return Token
 */
func (f *FileTokenizer) NextToken() *str_tokenizer.Token {
	var t *str_tokenizer.Token = nil
	if f.strToken.More() {
		t = f.strToken.Tokenize()
	} else if f.GetNewBlock() {
		t = f.strToken.Tokenize()
	} else {
		f.more = false
	}

	f.blockPos = f.strToken.Pos()
	return t
}

/**
 * Precondition: _more is true
 * Postcondition: Extract token from text file, and modify cursor position.
 *
 * @param f
 * @return *Token
 */
func (f *FileTokenizer) Tokenize() *str_tokenizer.Token {
	if !f.More() {
		panic("Reach end of file")
	}

	t := f.NextToken()
	return t
}

/**
 * Postcondition: The return value is false if there is no more blocks to get.
 *      Otherwise, it get a new block and update STokenizer's buffer and cursor
 *      position.
 *
 * @return true
 * @return false
 */
func (f *FileTokenizer) GetNewBlock() bool {
	buffer := make([]byte, MAX_BUFFER)
	n, err := f.inFile.Read(buffer)
	if err != nil {
		if err == io.EOF {
			err = f.inFile.Close()
			check(err)
			f.outFile.WriteString("*** END OF FILE ***\n")
			return false
		} else {
			panic(err)
		}
	}

	f.outFile.WriteString(
		fmt.Sprintf("----- New Block ---------------------[%d] bytes\n", n),
	)
	f.pos += f.blockPos
	f.strToken.SetString(string(buffer))
	f.more = true
	return true
}
