package db

import (
	"errors"
	"fmt"

	"github.com/kr/pretty"
	"github.com/spf13/cobra"

	"ModelHub-server/models"
)

var listIndexesFlags struct {
	cluster  string
	db       string
	coll     string
	required bool
}

func init() {
	ListIndexesCmd.Flags().StringVar(&listIndexesFlags.cluster, "cluster", "hub", "which cluster")
	ListIndexesCmd.Flags().StringVar(&listIndexesFlags.db, "db", "hub", "which db")
	ListIndexesCmd.Flags().StringVar(&listIndexesFlags.coll, "coll", "", "which collection, empty means all in db")
	ListIndexesCmd.Flags().BoolVar(&listIndexesFlags.required, "required", false, "if true list required indexes, else list exist indexes")
}

var ListIndexesCmd = &cobra.Command{
	Use:   "listIndexes",
	Short: "List indexes",
	Long:  `List indexes.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if listIndexesFlags.required {
			if listIndexesFlags.db == "hub" {
				fmt.Printf("%# v\n", pretty.Formatter(models.HubDBIndexes))
			} else {
				return errors.New("unknown db")
			}
		} else {
			collNames, err := models.DBCollNames(listIndexesFlags.cluster, listIndexesFlags.db)
			if err != nil {
				return err
			}

			for _, collName := range collNames {
				if listIndexesFlags.coll == "" || listIndexesFlags.coll == collName {
					coll, err := models.NewMongoColl(listIndexesFlags.cluster, listIndexesFlags.db, collName)
					if err != nil {
						return err
					}

					indexes, err := coll.Indexes()
					if err != nil {
						return err
					}
					fmt.Printf("%# v\n", pretty.Formatter(indexes))
				}
			}
		}

		return nil
	},
}
