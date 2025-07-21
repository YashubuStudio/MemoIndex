package index

import (
	"strings"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

// Wakati は分かち書きして返します
func Wakati(text string) (string, error) {
	// IPA辞書を渡してトークナイザ生成。BOS/EOSトークンは省略
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return "", err
	}

	words := t.Wakati(text)
	return strings.Join(words, " "), nil
}
