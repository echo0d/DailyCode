package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 1 // random length up to 24
	charcter := []rune(",.!@#$%^&*()_+{}[]|\\:;\"'<>,.?/~`")
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	// 随机一个位置
	m := rng.Intn((n + 1) / 2)
	c := charcter[rng.Intn(len(charcter))]
	// 从标点符号中随机一个出来，替换
	runes[m] = c
	runes[n-1-m] = c
	return string(runes)
}

// 生成一个非回文串
func randomNotPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 2
	runes := make([]rune, n)
	for {
		for i := range runes {
			// result[i] = charset[rand.Intn(len(charset))]
			runes[i] = rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		}
		if !IsPalindrome(string(runes)) {
			return string(runes)
		}
	}

}

// 回文字符串测试
func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}

	}
}

// 非回文字符串测试
func TestRandomNotPalindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		q := randomNotPalindrome(rng)
		if IsPalindrome(q) {
			t.Errorf("IsPalindrome(%q) = true", q)
		}
	}

}
