package analyzer

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "Erret",
	Doc:      "Find leaking internal errors in your rpc",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		funcDecl := node.(*ast.FuncDecl)

		if funcDecl.Recv == nil {
			return
		}

		params := funcDecl.Type.Params.List
		if len(params) != 2 {
			return
		}

		firstParamType, ok := params[0].Type.(*ast.SelectorExpr)
		if !ok {
			return
		}

		pkgIdent, ok := firstParamType.X.(*ast.Ident)
		if !ok {
			return
		}
		if pkgIdent.Name != "context" || firstParamType.Sel.Name != "Context" {
			return
		}

		secondParamNames := params[1].Names
		if len(secondParamNames) != 1 || secondParamNames[0].Name != "in" {
			return
		}

		fmt.Println("Confirmed RPC", funcDecl.Name)

		if funcDecl.Body == nil {
			return
		}
	})

	return nil, nil
}
