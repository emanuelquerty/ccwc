package main

type Expected struct {
	bytes int
	lines int
	words int
	chars int
}

type TestData struct {
	filepath string
	expected Expected
}

var testData = []TestData{
	{
		filepath: "testfiles/test.txt",
		expected: Expected{
			bytes: 342190,
			lines: 7145,
			words: 58164,
			chars: 339292,
		},
	},
	{
		filepath: "testfiles/test2.txt",
		expected: Expected{
			bytes: 2,
			lines: 1,
			words: 0,
			chars: 2,
		},
	},
	{
		filepath: "testfiles/test3.txt",
		expected: Expected{
			bytes: 738,
			lines: 9,
			words: 113,
			chars: 738,
		},
	},
	{
		filepath: "testfiles/test4.txt",
		expected: Expected{
			bytes: 44,
			lines: 5,
			words: 6,
			chars: 44,
		},
	},
	{
		filepath: "testfiles/chineseRandom.txt",
		expected: Expected{
			bytes: 1613,
			lines: 13,
			words: 12,
			chars: 561,
		},
	},
	{
		filepath: "testfiles/oneLongWord.txt",
		expected: Expected{
			bytes: 137,
			lines: 0,
			words: 1,
			chars: 137,
		},
	},
}