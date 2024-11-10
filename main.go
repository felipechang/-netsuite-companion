package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"netsuite-companion/file"
	"netsuite-companion/store"
	"netsuite-companion/util"
	"os"
)

func main() {
	baseStore := store.NewBaseStore()
	tree := file.CreateTree()
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "initialize global and working directory settings",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "force",
						Usage:   "force global initialization",
						Aliases: []string{"f"},
					},
				},
				Action: func(cCtx *cli.Context) error {
					force := cCtx.Bool("force")
					err := baseStore.CreateGlobal(force)
					if err != nil {
						return err
					}
					err = tree.Build()
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "add",
				Usage: "Add a project, or create script, module, or TS type files",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Please select one of the following options:")
					for _, option := range util.GetOptions() {
						fmt.Printf("  - %s\n", option)
					}
					return nil
				},
				Subcommands: []*cli.Command{
					{
						Name:  "project",
						Usage: "Add a new project to organize your code",
						Action: func(cCtx *cli.Context) error {
							err := baseStore.CreateProject()
							if err != nil {
								return err
							}
							global, err := baseStore.RetrieveGlobal()
							if err != nil {
								return err
							}
							project, err := baseStore.RetrieveProject()
							if err != nil {
								return err
							}
							err = tree.CreateProjectFolder(global, project)
							if err != nil {
								return err
							}
							err = tree.CreateManifest(project)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "bundle",
						Usage: "Add a new bundle script",
						Action: func(cCtx *cli.Context) error {
							instruct := cCtx.String("instruct")
							global, err := baseStore.RetrieveGlobal()
							if err != nil {
								return err
							}
							project, err := baseStore.RetrieveProject()
							if err != nil {
								return err
							}
							err = tree.CreateBundle(global, project, instruct)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "client",
						Usage: "Add a new client script",
						Action: func(cCtx *cli.Context) error {
							instruct := cCtx.String("instruct")
							global, err := baseStore.RetrieveGlobal()
							if err != nil {
								return err
							}
							project, err := baseStore.RetrieveProject()
							if err != nil {
								return err
							}
							err = tree.CreateClient(global, project, instruct)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "formclient",
						Usage: "Add a new form client script",
						Action: func(cCtx *cli.Context) error {
							instruct := cCtx.String("instruct")
							global, err := baseStore.RetrieveGlobal()
							if err != nil {
								return err
							}
							project, err := baseStore.RetrieveProject()
							if err != nil {
								return err
							}
							err = tree.CreateFormClient(global, project, instruct)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "mapreduce",
						Usage: "Add a new map reduce script",
						Action: func(cCtx *cli.Context) error {
							instruct := cCtx.String("instruct")
							global, err := baseStore.RetrieveGlobal()
							if err != nil {
								return err
							}
							project, err := baseStore.RetrieveProject()
							if err != nil {
								return err
							}
							err = tree.CreateMapReduce(global, project, instruct)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "massupdate",
						Usage: "Add a new mass update script",
						Action: func(cCtx *cli.Context) error {
							instruct := cCtx.String("instruct")
							global, err := baseStore.RetrieveGlobal()
							if err != nil {
								return err
							}
							project, err := baseStore.RetrieveProject()
							if err != nil {
								return err
							}
							err = tree.CreateMassUpdate(global, project, instruct)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "portlet",
						Usage: "Add a new portlet script",
						Action: func(cCtx *cli.Context) error {
							instruct := cCtx.String("instruct")
							global, err := baseStore.RetrieveGlobal()
							if err != nil {
								return err
							}
							project, err := baseStore.RetrieveProject()
							if err != nil {
								return err
							}
							err = tree.CreatePortlet(global, project, instruct)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "restlet",
						Usage: "Add a new restlet script",
						Action: func(cCtx *cli.Context) error {
							instruct := cCtx.String("instruct")
							global, err := baseStore.RetrieveGlobal()
							if err != nil {
								return err
							}
							project, err := baseStore.RetrieveProject()
							if err != nil {
								return err
							}
							err = tree.CreateRestlet(global, project, instruct)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "scheduled",
						Usage: "Add a new scheduled script",
						Action: func(cCtx *cli.Context) error {
							instruct := cCtx.String("instruct")
							global, err := baseStore.RetrieveGlobal()
							if err != nil {
								return err
							}
							project, err := baseStore.RetrieveProject()
							if err != nil {
								return err
							}
							err = tree.CreateScheduled(global, project, instruct)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "suitelet",
						Usage: "Add a new suitelet script",
						Action: func(cCtx *cli.Context) error {
							instruct := cCtx.String("instruct")
							global, err := baseStore.RetrieveGlobal()
							if err != nil {
								return err
							}
							project, err := baseStore.RetrieveProject()
							if err != nil {
								return err
							}
							err = tree.CreateSuitelet(global, project, instruct)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "userevent",
						Usage: "Add a new user event script",
						Action: func(cCtx *cli.Context) error {
							instruct := cCtx.String("instruct")
							global, err := baseStore.RetrieveGlobal()
							if err != nil {
								return err
							}
							project, err := baseStore.RetrieveProject()
							if err != nil {
								return err
							}
							err = tree.CreateUserEvent(global, project, instruct)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "workflowaction",
						Usage: "Add a new workflow action script",
						Action: func(cCtx *cli.Context) error {
							instruct := cCtx.String("instruct")
							global, err := baseStore.RetrieveGlobal()
							if err != nil {
								return err
							}
							project, err := baseStore.RetrieveProject()
							if err != nil {
								return err
							}
							err = tree.CreateWorkflowAction(global, project, instruct)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "module",
						Usage: "Add a module file",
						Action: func(cCtx *cli.Context) error {
							instruct := cCtx.String("instruct")
							global, err := baseStore.RetrieveGlobal()
							if err != nil {
								return err
							}
							project, err := baseStore.RetrieveProject()
							if err != nil {
								return err
							}
							err = tree.CreateModule(global, project, instruct)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "type",
						Usage: "Add a TypeScript type",
						Action: func(cCtx *cli.Context) error {
							instruct := cCtx.String("instruct")
							global, err := baseStore.RetrieveGlobal()
							if err != nil {
								return err
							}
							project, err := baseStore.RetrieveProject()
							if err != nil {
								return err
							}
							err = tree.CreateType(global, project, instruct)
							if err != nil {
								return err
							}
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
