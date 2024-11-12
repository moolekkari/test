package cliflag

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

// FlagV2 implements the koanf.FlagV2 interface for urfave/cli/v2
// with support for deeply nested commands and hierarchical flags
type FlagV2 struct {
	ctx   *cli.Context
	delim string
}

func New(ctx *cli.Context, delim string) *FlagV2 {
	if delim == "" {
		delim = "."
	}

	return &FlagV2{
		ctx:   ctx,
		delim: delim,
	}
}

func (p *FlagV2) ReadBytes() ([]byte, error) {
	return nil, fmt.Errorf("cli FlagV2 does not support ReadBytes")
}

// Read parses CLI flags and returns a nested map structure
func (p *FlagV2) Read() (map[string]interface{}, error) {
	out := make(map[string]interface{})

	// Process global flags first
	// p.processFlags(p.ctx.App.Flags, "global", out)

	// Get command lineage (from root to current command)
	lineage := p.ctx.Lineage()
	if len(lineage) > 0 {
		// Build command path and process flags for each level
		var cmdPath []string
		for i := len(lineage) - 1; i >= 0; i-- {
			cmd := lineage[i]
			if cmd.Command == nil {
				continue
			}
			cmdPath = append(cmdPath, cmd.Command.Name)
			prefix := strings.Join(cmdPath, p.delim)
			p.processFlags(cmd.Command.Flags, prefix, out)
		}

		// Process current command's flags
		if p.ctx.Command != nil {
			cmdPath = append(cmdPath, p.ctx.Command.Name)
			prefix := strings.Join(cmdPath, p.delim)
			p.processFlags(p.ctx.Command.Flags, prefix, out)
		}
	}

	return out, nil
}

// processFlags handles flag processing for a specific command level
func (p *FlagV2) processFlags(flags []cli.Flag, prefix string, out map[string]interface{}) {
	for _, flag := range flags {
		name := flag.Names()[0]
		if p.ctx.IsSet(name) {
			value := p.getFlagValue(name)
			if value != nil {
				// Build the full path for the flag
				fullPath := name
				if prefix != "global" {
					fullPath = prefix + p.delim + name
				}

				p.setNestedValue(fullPath, value, out)
			}
		}
	}
}

// setNestedValue sets a value in the nested configuration structure
func (p *FlagV2) setNestedValue(path string, value interface{}, out map[string]interface{}) {
	parts := strings.Split(path, p.delim)
	current := out

	// Navigate/create the nested structure
	for i := 0; i < len(parts)-1; i++ {
		if _, exists := current[parts[i]]; !exists {
			current[parts[i]] = make(map[string]interface{})
		}
		current = current[parts[i]].(map[string]interface{})
	}

	// Set the final value
	current[parts[len(parts)-1]] = value
}

// getFlagValue extracts the typed value from the flag
func (p *FlagV2) getFlagValue(name string) interface{} {
	switch {
	case p.ctx.IsSet(name):
		switch {
		case p.ctx.String(name) != "":
			return p.ctx.String(name)
		case p.ctx.StringSlice(name) != nil:
			return p.ctx.StringSlice(name)
		case p.ctx.Int(name) != 0:
			return p.ctx.Int(name)
		case p.ctx.Int64(name) != 0:
			return p.ctx.Int64(name)
		case p.ctx.IntSlice(name) != nil:
			return p.ctx.IntSlice(name)
		case p.ctx.Float64(name) != 0:
			return p.ctx.Float64(name)
		case p.ctx.Bool(name):
			return p.ctx.Bool(name)
		case p.ctx.Duration(name).String() != "0s":
			return p.ctx.Duration(name)
		default:
			return p.ctx.Generic(name)
		}
	}
	return nil
}
