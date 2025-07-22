package search

import (
	"fmt"
	"log"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/spf13/cobra"
	"ykvario.com/MemoIndex/config"
	"ykvario.com/MemoIndex/i18n"
	idx "ykvario.com/MemoIndex/index"
)

// Result represents a single search hit.
type Result struct {
	Path     string
	Fragment string
}

// ExecuteSearch performs a search against the index and returns up to limit results.
func ExecuteSearch(queryText string, limit int) ([]Result, error) {
	indexPath := config.AppConfig.IndexPath
	if indexPath == "" {
		indexPath = "./memoindex.bleve"
	}

	index, err := bleve.Open(indexPath)
	if err != nil {
		// インデックスが存在しない場合は新規作成
		index, err = bleve.New(indexPath, idx.CreateKeywordIndexMapping())
		if err != nil {
			return nil, fmt.Errorf("インデックスの読み込みに失敗しました: %w", err)
		}
	}
	defer index.Close()

	// 🔽 修正：分かち書きしてから検索クエリ作成
	wakatiQuery, err := idx.Wakati(queryText)
	if err != nil {
		return nil, fmt.Errorf("検索語の分かち書きに失敗: %w", err)
	}
	fmt.Println("【検索クエリ】", wakatiQuery)

	q := bleve.NewQueryStringQuery(wakatiQuery) // ← ここを MatchQuery → QueryStringQuery に変更

	searchReq := bleve.NewSearchRequestOptions(q, limit, 0, false)
	searchReq.Highlight = bleve.NewHighlight()

	result, err := index.Search(searchReq)
	if err != nil {
		return nil, fmt.Errorf("検索に失敗しました: %w", err)
	}

	hits := make([]Result, 0, len(result.Hits))
	for _, hit := range result.Hits {
		frag := ""
		if fragments, ok := hit.Fragments["body"]; ok && len(fragments) > 0 {
			frag = fragments[0]
		}
		hits = append(hits, Result{Path: hit.ID, Fragment: frag})
	}
	return hits, nil
}

var SearchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "全文検索を行います（上位3件を表示）",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		queryText := strings.Join(args, " ")

		results, err := ExecuteSearch(queryText, 3)
		if err != nil {
			log.Fatalf("%v", err)
		}

		if len(results) == 0 {
			fmt.Println(i18n.T("no_results", nil))
			return
		}

		for i, hit := range results {
			fmt.Printf("%d. %s\n", i+1, hit.Path)
			if hit.Fragment != "" {
				fmt.Printf("   ...%s...\n", hit.Fragment)
			}
		}
	},
}
