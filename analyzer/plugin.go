package analyzer

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("cerret", func(conf any) (register.LinterPlugin, error) {
		return &cerretPlugin{}, nil
	})
}

type cerretPlugin struct{}

func (p *cerretPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{Analyzer}, nil
}

func (p *cerretPlugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
