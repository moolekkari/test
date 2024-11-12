// Package cliflag implements a koanf.Provider that reads commandline
// parameters as conf maps using ufafe/cli flag.
package cliflag

import (
	"errors"
	"fmt"

	"github.com/knadh/koanf/maps"
	"github.com/urfave/cli/v2"
)

// CliFlag implements a cli.Flag command line provider.
type CliFlag struct {
	ctx   *cli.Context
	delim string
}

// Provider returns a commandline flags provider that returns
// a nested map[string]interface{} of environment variable where the
// nesting hierarchy of keys are defined by delim. For instance, the
// delim "." will convert the key `parent.child.key: 1`
// to `{parent: {child: {key: 1}}}`.
func Provider(f *cli.Context, delim string) *CliFlag {
	return &CliFlag{
		ctx:   f,
		delim: delim,
	}
}

// Read reads the flag variables and returns a nested conf map.
func (p *CliFlag) Read() (map[string]interface{}, error) {
	mp := make(map[string]interface{})

	fmt.Printf("p.flagset.FlagNames(): %v\n", p.ctx.FlagNames())
	fmt.Printf("p.ctx.Args(): %v\n", p.ctx.Args())

	fmt.Printf("p.ctx.String(\"lol\"): %v\n", p.ctx.String("lol"))
	fmt.Printf("p.ctx.Int(\"ll\"): %v\n", p.ctx.Int("ll"))
	fmt.Printf("p.ctx.String(\"ll\"): %v\n", p.ctx.String("ll"))
	fmt.Printf("p.ctx.App.VisibleCommands(): %v\n", p.ctx.App.VisibleCommands())
	fmt.Printf("p.ctx.App.VisibleFlags(): %v\n", p.ctx.App.VisibleFlags())
	fmt.Printf("p.ctx.Args().First(): %v\n", p.ctx.Args().Slice())
	fmt.Printf("p.ctx.Args().Present(): %v\n", p.ctx.Args().Present())

	cli.NewContext(p.ctx.App, nil, nil)

	for _, v := range p.ctx.FlagNames() {
		mp[v] = p.ctx.Value(v)
	}

	if p.delim == "" {
		return mp, nil
	}

	return maps.Unflatten(mp, p.delim), nil
}

// ReadBytes is not supported by the cliflag provider.
func (p *CliFlag) ReadBytes() ([]byte, error) {
	return nil, errors.New("cliflag provider does not support this method")
}

// Watch is not supported by cliFlag.
func (p *CliFlag) Watch(cb func(event interface{}, err error)) error {
	return errors.New("posflag provider does not support this method")
}
