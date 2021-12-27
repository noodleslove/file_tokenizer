# Work Report

## Name: <ins> Eddie Ho </ins>

## Features:

- Not implemented:
    - None

<br><br>

- Implemented:
    - FileTokenizer struct, all elements within are private
    - NewFileTokenizer function returns a pointer to a new FileTokenizer struct
    - FileTokenizer.NextToken function returns next token in the file if there
    is more
    - FileTokenizer.Pos function returns the current cursor position of the 
    whole file
    - FileTokenizer.BlockPos function returns the current cursor position of 
    the current block
    - FileTokenizer.More function returns true if cursor has reached the last 
    token of the whole file

<br><br>

- Partly implemented:
    - None

<br><br>

- Bugs:
    -  FileTokenizer breaks a file into a number of blocks. When a word is in 
    the middle of two blocks, it will break into two words. For example, in 
    main, "great" at the end of the third block. Because including the 
    whole word will exceed the buffer size, the program break it into "gr" 
    and "eat".
