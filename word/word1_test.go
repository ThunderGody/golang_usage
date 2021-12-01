package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Errorf(`TestIsPalindrome("detartrated") = false`)
	}

	if !IsPalindrome("kayak") {
		t.Errorf(`TestIsPalindrome("kayak") = false`)
	}
}

func TestNoPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Errorf(`TestIsPalindrome("palindrome") = true`)
	}
}


func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Errorf(`TestIsPalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}