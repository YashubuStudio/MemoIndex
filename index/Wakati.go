package index

import (
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

// Wakati は検索モードで分かち書きして返します
func Wakati(text string) (string, error) {
	// Kagomeトークナイザの初期化（IPA辞書使用、BOS/EOSトークン省略）
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return "", err
	}
	// 検索モードで形態素解析を実行
	tokens := t.Analyze(text, tokenizer.Search)
	// トークンから表層形を取り出してスペースで連結
	words := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			continue // BOS/EOSトークンの除外
		}
		words = append(words, token.Surface)
	}
	return strings.Join(words, " "), nil
}

func CreateKeywordIndexMapping() mapping.IndexMapping { // ← bleve. ではなく mapping. に変更
	fieldMapping := bleve.NewTextFieldMapping()
	fieldMapping.Analyzer = "keyword"

	docMapping := bleve.NewDocumentMapping()
	docMapping.AddFieldMappingsAt("body", fieldMapping)

	indexMapping := bleve.NewIndexMapping()
	indexMapping.DefaultAnalyzer = "keyword"
	indexMapping.DefaultMapping = docMapping

	return indexMapping
}

func createSimpleIndexMapping() mapping.IndexMapping {
	fieldMapping := bleve.NewTextFieldMapping()
	fieldMapping.Analyzer = "standard" // ← ここだけ変更！

	docMapping := bleve.NewDocumentMapping()
	docMapping.AddFieldMappingsAt("body", fieldMapping)

	indexMapping := bleve.NewIndexMapping()
	indexMapping.DefaultAnalyzer = "standard"
	indexMapping.DefaultMapping = docMapping

	return indexMapping
}
