package testdata

type Expected struct {
	Bytes int
	Lines int
	Words int
	Chars int
}

type TableData struct {
	Filepath string
	Expected Expected
}

var TableDrivenData = []TableData{
	{
		Filepath: "testdata/test1.txt",
		Expected: Expected{
			Bytes: 342190,
			Lines: 7145,
			Words: 58164,
			Chars: 339292,
		},
	},
	{
		Filepath: "testdata/test2.txt",
		Expected: Expected{
			Bytes: 2,
			Lines: 1,
			Words: 0,
			Chars: 2,
		},
	},
	{
		Filepath: "testdata/test3.txt",
		Expected: Expected{
			Bytes: 738,
			Lines: 9,
			Words: 113,
			Chars: 738,
		},
	},
	{
		Filepath: "testdata/test4.txt",
		Expected: Expected{
			Bytes: 44,
			Lines: 5,
			Words: 6,
			Chars: 44,
		},
	},
	{
		Filepath: "testdata/chineseRandom.txt",
		Expected: Expected{
			Bytes: 1613,
			Lines: 13,
			Words: 12,
			Chars: 561,
		},
	},
	{
		Filepath: "testdata/oneLongWord.txt",
		Expected: Expected{
			Bytes: 137,
			Lines: 0,
			Words: 1,
			Chars: 137,
		},
	},
}