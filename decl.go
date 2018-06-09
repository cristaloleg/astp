package astp

import "go/ast"

// IsFuncDecl returns true if a given ast.Node is a function declaration (*ast.FuncDecl).
func IsFuncDecl(node ast.Node) bool {
	_, ok := node.(*ast.FuncDecl)
	return ok
}

// IsGenDecl returns true if a given ast.Node is a generic declaration (*ast.GenDecl).
func IsGenDecl(node ast.Node) bool {
	_, ok := node.(*ast.GenDecl)
	return ok
}

// IsImportSpec returns true if a given ast.Node is an import declaration (*ast.ImportSpec).
func IsImportSpec(node ast.Node) bool {
	_, ok := node.(*ast.ImportSpec)
	return ok
}

// IsValueSpec returns true if a given ast.Node is a value declaration (*ast.ValueSpec).
func IsValueSpec(node ast.Node) bool {
	_, ok := node.(*ast.ValueSpec)
	return ok
}

// IsTypeSpec returns true if a given ast.Node is a type declaration (*ast.TypeSpec).
func IsTypeSpec(node ast.Node) bool {
	_, ok := node.(*ast.TypeSpec)
	return ok
}
