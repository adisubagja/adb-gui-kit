package backend

import "testing"

func TestQuoteShellArgEscapesSingleQuote(t *testing.T) {
	input := "/sdcard/My Folder/O'Reilly"
	quoted := quoteShellArg(input)

	expected := "'/sdcard/My Folder/O'\"'\"'Reilly'"
	if quoted != expected {
		t.Fatalf("unexpected quote result. expected %s, got %s", expected, quoted)
	}
}
