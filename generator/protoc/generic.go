package protoc

import (
	"fmt"
	"log"
	"strings"

	"github.com/zeeraw/protogen/config"
)

// RunGeneric will run generation of a package in a generic way
func (p *Protoc) RunGeneric(pkg *config.Package, files ...string) error {
	log.Printf("protoc running generic for %q\n", pkg.Language)
	if _, err := p.CheckExtension(pkg.Language); err != nil {
		return err
	}
	args, err := p.BuildGeneric(pkg, files...)
	if err != nil {
		return err
	}
	return p.Exec(args...)
}

// BuildGeneric will construct the protoc command for a package in a generic way
func (p *Protoc) BuildGeneric(pkg *config.Package, files ...string) ([]string, error) {
	log.Printf("protoc building generic for %q\n", pkg.Language)
	flag := fmt.Sprintf("--%s_out", pkg.Language)
	lang := []string{}
	lang = append(lang, fmt.Sprintf("%s=", flag))
	lang = append(lang, pkg.Output)

	args := []string{}
	args = append(args, fmt.Sprintf("-I%s", pkg.Root()))
	args = append(args, strings.Join(lang, ""))
	args = append(args, files...)
	return args, nil
}
