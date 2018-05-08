package main

import "testing"

func TestDecompress1(t *testing.T) {
	var compressed = []rune("3[abc]4[ab]c")
	var expected = []rune("abcabcabcababababc")
	testDecompress(t, compressed, expected)
}

func TestDecompress2(t *testing.T) {
	var compressed = []rune("15[x]4[yz]")
	var expected = []rune("xxxxxxxxxxxxxxxyzyzyzyz")
	testDecompress(t, compressed, expected)
}

func TestDecompress3(t *testing.T) {
	var compressed = []rune("2[3[a]b]")
	var expected = []rune("aaabaaab")
	testDecompress(t, compressed, expected)
}

func TestDecompress4(t *testing.T) {
	var compressed = []rune("0[abc]")
	var expected = []rune("")
	testDecompress(t, compressed, expected)
}

func TestDecompress5(t *testing.T) {
	var compressed = []rune("3[]")
	var expected = []rune("")
	testDecompress(t, compressed, expected)
}

func TestDecompress6(t *testing.T) {
	var compressed = []rune("31[a]0[3[x]]0[y]11[b3[d[0[c]]]")
	var expected = []rune("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabdddbdddbdddbdddbdddbdddbdddbdddbdddbdddbddd")
	testDecompress(t, compressed, expected)
}

func testDecompress(t *testing.T, compressed, expected []rune) {
	decompressed := Decompress(compressed)
	if !isEqual(decompressed, expected) {
		t.Errorf("Decompress failed for '%s', got: '%s', want: '%s'.", string(compressed), string(decompressed), string(expected))
	}
}

// test if two arrays are equal
func isEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
