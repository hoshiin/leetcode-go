package leetcode

import "testing"

func TestReplaceWords(t *testing.T) {
	tests := []struct {
		dictionary []string
		sentence   string
		want       string
	}{
		{[]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery", "the cat was rat by the bat"},
		{[]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs", "a a b c"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ReplaceWords(tt.dictionary, tt.sentence); got != tt.want {
				t.Errorf("ReplaceWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
