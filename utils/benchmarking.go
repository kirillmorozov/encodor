package utils

// EncodeBecnhmarks is a standard set of inputs that's used to measure encoding
// performance.
var EncodeBenchmarks = []struct {
	Name string
	Text []byte
}{
	{
		Name: "Multiline",
		Text: []byte("word1.1 word1.2 @username1 #hashtag1\nword2.1 word2.2 @username2 #hashtag2"),
	},
	{
		Name: "Lorem Ipsum",
		Text: []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
	},
}
