package search

import (
	"fmt"
	"log"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/spf13/cobra"
)

var indexPath = "memoindex.bleve"

var SearchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "全文検索を行います（上位3件を表示）",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		queryText := strings.Join(args, " ")

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

		for i, hit := range result.Hits {
			fmt.Printf("%d. %s\n", i+1, hit.ID)
			if fragments, ok := hit.Fragments["body"]; ok && len(fragments) > 0 {
				fmt.Printf("   ...%s...\n", fragments[0])
			}
		}
	},
}
