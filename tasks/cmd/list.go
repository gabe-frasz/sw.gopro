package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/gabe-frasz/gopro/tasks/internal/app/entity"
	"github.com/gabe-frasz/gopro/tasks/internal/app/usecases"
	"github.com/gabe-frasz/gopro/tasks/internal/infra/database/repository"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var (
	allFlag bool
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list tasks",
		Run: func(cmd *cobra.Command, args []string) {
			db, err := sql.Open("sqlite3", "tasks.db")
			if err != nil {
				panic(err)
			}
			defer db.Close()

			taskRepository := repository.NewSqlTaskRepository(db)
			var tasks []*entity.Task

			if allFlag {
				usecase := usecases.NewListAllTasks(taskRepository)
				tasks, err = usecase.Execute()
				if err != nil {
					panic(err)
				}
			} else {
				usecase := usecases.NewListUncompletedTasks(taskRepository)
				tasks, err = usecase.Execute()
				if err != nil {
					panic(err)
				}
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

			fmt.Fprintln(w, "DESCRIPTION\t DONE\t CREATED AT")

			for _, task := range tasks {
				desc := task.Description
				done := strconv.FormatBool(task.Done)
        createdAt := timediff.TimeDiff(task.CreatedAt)

				fmt.Fprintln(w, desc+"\t "+done+"\t "+createdAt)
			}

			w.Flush()
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "list all tasks")
}
