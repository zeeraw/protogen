package cli

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/generator"
)

func (r *Runner) cmdGenerate() cli.Command {
	const usage = "performs code generation based on the protogen configuration"
	const description = `Your source will be retrieved, checked out and have protoc run for
the relevant files on every release tag you've specified in your
.protogen file`

	return cli.Command{
		Name:        "generate",
		Usage:       usage,
		Description: description,
		Action:      r.generate,
		Flags: []cli.Flag{
			r.flagProtogenFile(),
		},
	}
}

func (r *Runner) generate(cc *cli.Context) error {
	cfg, err := ReadConfigFromFilePath(r.protogenFile)
	if err != nil {
		log.Println(err)
		fmt.Printf("Cannot read protogen file %q: %v\n", r.protogenFile, err)
		return err
	}
	cfg.General = &config.General{
		Verbose: r.verbose,
	}
	if err := generator.Generate(cfg); err != nil {
		fmt.Printf("Cannot generate code: %v\n", err)
		return err
	}
	return nil
}
