package files

type CodedNumber struct {
	Number     int
	Characters string
}

var (
	codedNumbers = []CodedNumber{
		{
			Number: 0,
			Characters: "" +
				" _ " +
				"| |" +
				"|_|",
		},
		{
			Number: 1,
			Characters: "" +
				"   " +
				"  |" +
				"  |",
		},
		{
			Number: 2,
			Characters: "" +
				" _ " +
				" _|" +
				"|_ ",
		},
		{
			Number: 3,
			Characters: "" +
				" _ " +
				" _|" +
				" _|",
		},
		{
			Number: 4,
			Characters: "" +
				"   " +
				"|_|" +
				"  |",
		},
		{
			Number: 5,
			Characters: "" +
				" _ " +
				"|_ " +
				" _|",
		},
		{
			Number: 6,
			Characters: "" +
				" _ " +
				"|_ " +
				"|_|",
		},
		{
			Number: 7,
			Characters: "" +
				" _ " +
				"  |" +
				"  |",
		},
		{
			Number: 8,
			Characters: "" +
				" _ " +
				"|_|" +
				"|_|",
		},
		{
			Number: 9,
			Characters: "" +
				" _ " +
				"|_|" +
				" _|",
		},
	}
)
