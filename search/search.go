package search

import (
	"fmt"
	"log"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/spf13/cobra"
	"ykvario.com/MemoIndex/config"
)

var SearchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "全文検索を行います（上位3件を表示）",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		queryText := strings.Join(args, " ")

		indexPath := config.AppConfig.IndexPath
		if indexPath == "" {
			indexPath = "./memoindex.bleve"
		}

		index, err := bleve.Open(indexPath)
		if err != nil {
			log.Fatalf("インデックスの読み込みに失敗しました: %v", err)
		}
		defer index.Close()

		q := bleve.NewMatchQuery(queryText)
		search := bleve.NewSearchRequestOptions(q, 3, 0, false)
		search.Highlight = bleve.NewHighlight()

		result, err := index.Search(search)
		if err != nil {
			log.Fatalf("検索に失敗しました: %v", err)
		}

		if result.Total == 0 {
			fmt.Println("検索結果がありません。")
			return
		}

		for i, hit := range result.Hits {
			fmt.Printf("%d. %s\n", i+1, hit.ID)
			if fragments, ok := hit.Fragments["body"]; ok && len(fragments) > 0 {
				fmt.Printf("   ...%s...\n", fragments[0])
			}
		}
	},
}
