package cliflag

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

func TestCliFlag(t *testing.T) {
	cliApp := cli.App{
		// Name: "testing",
		Action: func(ctx *cli.Context) error {
			// p := Provider(ctx, ".")
			p := New(ctx, ".")
			x, err := p.Read()
			require.NoError(t, err)
			require.NotEmpty(t, x)

			fmt.Printf("x: %v\n", x)

			return nil
		},
		Flags: []cli.Flag{
			cli.HelpFlag,
			cli.VersionFlag,
			&cli.StringFlag{
				Name:    "test",
				Usage:   "test flag",
				Value:   "test",
				Aliases: []string{"t"},
				EnvVars: []string{"TEST_FLAG"},
			},
			&cli.StringFlag{
				Name:     "lol",
				Usage:    "test flag",
				Value:    "test",
				Required: true,
				EnvVars:  []string{"TEST_FLAG"},
			},
			&cli.StringFlag{
				Name:     "ll",
				Usage:    "test flag",
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "x",
				Description: "yeah yeah testing",
				Action: func(ctx *cli.Context) error {
					// p := Provider(ctx, "")
					p := New(ctx, ".")
					x, err := p.Read()
					require.NoError(t, err)
					require.NotEmpty(t, x)
					fmt.Printf("x: %v\n", x)
					return nil
				},
				Flags: []cli.Flag{
					cli.HelpFlag,
					cli.VersionFlag,
					&cli.StringFlag{
						Name:     "lol",
						Usage:    "test flag",
						Value:    "test",
						Required: true,
						EnvVars:  []string{"TEST_FLAG"},
					},
					&cli.IntFlag{
						Name:     "ll",
						Usage:    "test flag",
						Required: true,
					},
				},
			},
		},
	}

	x := append([]string{"testing", "--lol", "loll", "--ll", "2024-10-22", "--test", "gf", "x", "--lol", "dsf", "--ll", "11"}, os.Args...)
	// x := append([]string{"--lol", "loll", "--ll", "2024-10-22", "--test", "gf", "x", "--lol", "dsf", "--ll", "11"}, os.Args...)

	err := cliApp.Run(x)
	// err := cliApp.Run([]string{"app", "--ll", "loooo"})

	require.NoError(t, err)
}

type genericType struct {
	s string
}

func (g *genericType) Set(value string) error {
	g.s = value
	return nil
}

func (g *genericType) String() string {
	return g.s
}

func TestNestedFlag(t *testing.T) {

	app := &cli.App{
		Name:     "kənˈtrīv",
		Version:  "v19.99.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Example Human",
				Email: "human@example.com",
			},
		},
		Copyright: "(c) 1999 Serious Enterprise",
		HelpName:  "contrive",
		Usage:     "demonstrate available API",
		UsageText: "contrive - demonstrating the available API",
		ArgsUsage: "[args and such]",
		Commands: []*cli.Command{
			&cli.Command{
				Name:        "doo",
				Aliases:     []string{"do"},
				Category:    "motion",
				Usage:       "do the doo",
				UsageText:   "doo - does the dooing",
				Description: "no really, there is a lot of dooing to be done",
				ArgsUsage:   "[arrgh]",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "forever", Aliases: []string{"forevvarr"}},
				},
				Subcommands: []*cli.Command{
					{
						Name: "wop",
						Action: func(cCtx *cli.Context) error {
							fmt.Fprintf(cCtx.App.Writer, ":wave: over here, eh\n")
							return nil
						},
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "whenever", Aliases: []string{"forevvarr"}},
						},
					},
				},
				SkipFlagParsing: false,
				HideHelp:        false,
				Hidden:          false,
				HelpName:        "doo!",
				BashComplete: func(cCtx *cli.Context) {
					fmt.Fprintf(cCtx.App.Writer, "--better\n")
				},
				Before: func(cCtx *cli.Context) error {
					fmt.Fprintf(cCtx.App.Writer, "brace for impact\n")
					return nil
				},
				After: func(cCtx *cli.Context) error {
					fmt.Fprintf(cCtx.App.Writer, "did we lose anyone?\n")
					return nil
				},
				Action: func(cCtx *cli.Context) error {
					cCtx.Command.FullName()
					cCtx.Command.HasName("wop")
					cCtx.Command.Names()
					cCtx.Command.VisibleFlags()
					fmt.Fprintf(cCtx.App.Writer, "dodododododoodododddooooododododooo\n")
					if cCtx.Bool("forever") {
						cCtx.Command.Run(cCtx)
					}
					return nil
				},
				OnUsageError: func(cCtx *cli.Context, err error, isSubcommand bool) error {
					fmt.Fprintf(cCtx.App.Writer, "for shame\n")
					return err
				},
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "fancy"},
			&cli.BoolFlag{Value: true, Name: "fancier"},
			&cli.DurationFlag{Name: "howlong", Aliases: []string{"H"}, Value: time.Second * 3},
			&cli.Float64Flag{Name: "howmuch"},
			&cli.GenericFlag{Name: "wat", Value: &genericType{}},
			&cli.Int64Flag{Name: "longdistance"},
			&cli.Int64SliceFlag{Name: "intervals"},
			&cli.IntFlag{Name: "distance"},
			&cli.IntSliceFlag{Name: "times"},
			&cli.StringFlag{Name: "dance-move", Aliases: []string{"d"}},
			&cli.StringSliceFlag{Name: "names", Aliases: []string{"N"}},
			&cli.UintFlag{Name: "age"},
			&cli.Uint64Flag{Name: "bigage"},
		},
		EnableBashCompletion: true,
		HideHelp:             false,
		HideVersion:          false,
		BashComplete: func(cCtx *cli.Context) {
			fmt.Fprintf(cCtx.App.Writer, "lipstick\nkiss\nme\nlipstick\nringo\n")
		},
		Before: func(cCtx *cli.Context) error {
			fmt.Fprintf(cCtx.App.Writer, "HEEEERE GOES\n")
			return nil
		},
		After: func(cCtx *cli.Context) error {
			fmt.Fprintf(cCtx.App.Writer, "Phew!\n")
			return nil
		},
		CommandNotFound: func(cCtx *cli.Context, command string) {
			fmt.Fprintf(cCtx.App.Writer, "Thar be no %q here.\n", command)
		},
		OnUsageError: func(cCtx *cli.Context, err error, isSubcommand bool) error {
			if isSubcommand {
				return err
			}

			fmt.Fprintf(cCtx.App.Writer, "WRONG: %#v\n", err)
			return nil
		},
		Action: func(cCtx *cli.Context) error {
			cli.DefaultAppComplete(cCtx)
			cli.HandleExitCoder(errors.New("not an exit coder, though"))
			cli.ShowAppHelp(cCtx)
			cli.ShowCommandCompletions(cCtx, "nope")
			cli.ShowCommandHelp(cCtx, "also-nope")
			cli.ShowCompletions(cCtx)
			cli.ShowSubcommandHelp(cCtx)
			cli.ShowVersion(cCtx)

			fmt.Printf("%#v\n", cCtx.App.Command("doo"))
			if cCtx.Bool("infinite") {
				cCtx.App.Run([]string{"app", "doo", "wop"})
			}

			if cCtx.Bool("forevar") {
				cCtx.App.RunAsSubcommand(cCtx)
			}

			cCtx.App.Setup()
			fmt.Printf("%#v\n", cCtx.App.VisibleCategories())
			fmt.Printf("%#v\n", cCtx.App.VisibleCommands())
			fmt.Printf("%#v\n", cCtx.App.VisibleFlags())

			fmt.Printf("%#v\n", cCtx.Args().First())
			if cCtx.Args().Len() > 0 {
				fmt.Printf("%#v\n", cCtx.Args().Get(1))
			}
			fmt.Printf("%#v\n", cCtx.Args().Present())
			fmt.Printf("%#v\n", cCtx.Args().Tail())

			set := flag.NewFlagSet("contrive", 0)
			nc := cli.NewContext(cCtx.App, set, cCtx)

			fmt.Printf("%#v\n", nc.Args())
			fmt.Printf("%#v\n", nc.Bool("nope"))
			fmt.Printf("%#v\n", !nc.Bool("nerp"))
			fmt.Printf("%#v\n", nc.Duration("howlong"))
			fmt.Printf("%#v\n", nc.Float64("hay"))
			fmt.Printf("%#v\n", nc.Generic("bloop"))
			fmt.Printf("%#v\n", nc.Int64("bonk"))
			fmt.Printf("%#v\n", nc.Int64Slice("burnks"))
			fmt.Printf("%#v\n", nc.Int("bips"))
			fmt.Printf("%#v\n", nc.IntSlice("blups"))
			fmt.Printf("%#v\n", nc.String("snurt"))
			fmt.Printf("%#v\n", nc.StringSlice("snurkles"))
			fmt.Printf("%#v\n", nc.Uint("flub"))
			fmt.Printf("%#v\n", nc.Uint64("florb"))

			fmt.Printf("%#v\n", nc.FlagNames())
			fmt.Printf("%#v\n", nc.IsSet("wat"))
			fmt.Printf("%#v\n", nc.Set("wat", "nope"))
			fmt.Printf("%#v\n", nc.NArg())
			fmt.Printf("%#v\n", nc.NumFlags())
			fmt.Printf("%#v\n", nc.Lineage()[1])
			nc.Set("wat", "also-nope")

			ec := cli.Exit("ohwell", 86)
			fmt.Fprintf(cCtx.App.Writer, "%d", ec.ExitCode())
			fmt.Printf("made it!\n")
			return ec
		},
		Metadata: map[string]interface{}{
			"layers":          "many",
			"explicable":      false,
			"whatever-values": 19.99,
		},
	}

	app.Run(os.Args)
}

func TestAnotherNestedFlag(t *testing.T) {

}

// Example usage with deeply nested commands and flags
func Example() {
	app := &cli.App{
		Name: "myapp",
		Flags: []cli.Flag{
			// Global flags
			&cli.StringFlag{
				Name:  "config",
				Value: "config.yaml",
				Usage: "global configuration file",
			},
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "enable debug mode globally",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Value: "localhost",
						Usage: "server host",
					},
					&cli.IntFlag{
						Name:  "port",
						Value: 8080,
						Usage: "server port",
					},
				},
				Subcommands: []*cli.Command{
					{
						Name: "database",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "driver",
								Value: "postgres",
								Usage: "database driver",
							},
						},
						Subcommands: []*cli.Command{
							{
								Name: "migrate",
								Flags: []cli.Flag{
									&cli.StringFlag{
										Name:  "direction",
										Value: "up",
										Usage: "migration direction",
									},
									&cli.IntFlag{
										Name:  "steps",
										Value: 1,
										Usage: "number of migration steps",
									},
								},
							},
							{
								Name: "backup",
								Flags: []cli.Flag{
									&cli.StringFlag{
										Name:  "format",
										Value: "sql",
										Usage: "backup format",
									},
									&cli.StringFlag{
										Name:  "output",
										Value: "backup.sql",
										Usage: "output file",
									},
								},
							},
						},
					},
					{
						Name: "http",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{
								Name:  "cors.origins",
								Usage: "allowed CORS origins",
							},
							&cli.DurationFlag{
								Name:  "timeout",
								Value: time.Duration(30),
								Usage: "http timeout",
							},
						},
					},
				},
			},
		},
		Action: func(c *cli.Context) error {
			k := koanf.New(".")
			provider := New(c, ".")

			if err := k.Load(provider, nil); err != nil {
				return err
			}

			// Access nested configuration examples:
			fmt.Printf("Global config: %s\n", k.String("config"))
			fmt.Printf("Server host: %s\n", k.String("server.host"))
			fmt.Printf("DB driver: %s\n", k.String("server.database.driver"))
			fmt.Printf("Migration direction: %s\n", k.String("server.database.migrate.direction"))
			fmt.Printf("Backup format: %s\n", k.String("server.database.backup.format"))
			fmt.Printf("HTTP CORS: %v\n", k.Strings("server.http.cors.origins"))

			return nil
		},
	}

	// Example command:
	// ./myapp --config=custom.yaml --debug server --host=example.com --port=9090 database --driver=mysql migrate --direction=up --steps=3
	app.Run(os.Args)
}
