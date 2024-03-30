// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package detect_deadcode

import (
	_ "embed"
	"go/ast"
	"go/token"
	"go/types"
	"os"
	"slices"
	"strings"

	"golang.org/x/tools/go/callgraph/rta"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func DetectDeadCodeRanges(path string) [][2]int {
	cfg := &packages.Config{
		Mode: packages.LoadAllSyntax | packages.NeedModule,
	}
	initial, _ := packages.Load(cfg, path)
	codebyte, _ := os.ReadFile(path)
	code := strings.Split(string(codebyte), "\n")

	prog, pkgs := ssautil.AllPackages(initial, ssa.InstantiateGenerics)
	prog.Build()

	main := ssautil.MainPackages(pkgs)[0]

	var sourceFuncs []*ssa.Function
	packages.Visit(initial, nil, func(p *packages.Package) {
		for _, file := range p.Syntax {
			for _, decl := range file.Decls {
				if decl, ok := decl.(*ast.FuncDecl); ok {
					obj := p.TypesInfo.Defs[decl.Name].(*types.Func)
					fn := prog.FuncValue(obj)
					sourceFuncs = append(sourceFuncs, fn)
				}
			}
		}
	})

	res := rta.Analyze([]*ssa.Function{main.Func("init"), main.Func("main")}, false)

	reachablePosn := make(map[token.Position]bool)
	for fn := range res.Reachable {
		if fn.Pos().IsValid() || fn.Name() == "init" {
			reachablePosn[prog.Fset.Position(fn.Pos())] = true
		}
	}

	result := [][2]int{}
	for _, fn := range sourceFuncs {
		posn := prog.Fset.Position(fn.Pos())

		if !reachablePosn[posn] {
			reachablePosn[posn] = true // suppress dups with same pos

			if fn.Pkg.Pkg.Path() == "command-line-arguments" {
				startLine := posn.Line - 1
				var endLine int
				if slices.Contains([]byte(code[startLine]), '}') {
					endLine = startLine
				} else {
					endLine = startLine
					for code[endLine] != "}" {
						endLine++
					}
				}
				result = append(result, [2]int{startLine, endLine})
			}
		}
	}
	return result
}
