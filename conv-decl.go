package astp

import "go/ast"

// AsFuncDecl returns a given ast.Node as a function declaration (*ast.FuncDecl).
func AsFuncDecl(node ast.Node) *ast.FuncDecl {
	decl, ok := node.(*ast.FuncDecl)
	if !ok {
		panic("expected *ast.FuncDecl")
	}
	return decl
}

// AsGenDecl returns a given ast.Node as a generic declaration (*ast.GenDecl).
func AsGenDecl(node ast.Node) *ast.GenDecl {
	decl, ok := node.(*ast.GenDecl)
	if !ok {
		panic("expected *ast.GenDecl")
	}
	return decl
}

// AsImportSpec returns a given ast.Node as an import declaration (*ast.ImportSpec).
func AsImportSpec(node ast.Node) *ast.ImportSpec {
	decl, ok := node.(*ast.ImportSpec)
	if !ok {
		panic("expected *ast.ImportSpec")
	}
	return decl
}

// AsValueSpec returns a given ast.Node as a value declaration (*ast.ValueSpec).
func AsValueSpec(node ast.Node) *ast.ValueSpec {
	decl, ok := node.(*ast.ValueSpec)
	if !ok {
		panic("expected *ast.ValueSpec")
	}
	return decl
}

// AsTypeSpec returns a given ast.Node as a type declaration (*ast.TypeSpec).
func AsTypeSpec(node ast.Node) *ast.TypeSpec {
	decl, ok := node.(*ast.TypeSpec)
	if !ok {
		panic("expected *ast.TypeSpec")
	}
	return decl
}
