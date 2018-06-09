package astp

import "go/ast"

// IsExpr returns true if a given ast.Node is an expression(ast.Expr).
func IsExpr(node ast.Node) bool {
	_, ok := node.(ast.Expr)
	return ok
}

// IsBadExpr returns true if a given ast.Node is a bad expression (*ast.IsBadExpr).
func IsBadExpr(node ast.Node) bool {
	_, ok := node.(*ast.BadExpr)
	return ok
}

// IsIdent returns true if a given ast.Node is an identifier (*ast.IsIdent).
func IsIdent(node ast.Node) bool {
	_, ok := node.(*ast.Ident)
	return ok
}

// IsEllipsis returns true if a given ast.Node is an `...` (ellipsis) (*ast.IsEllipsis).
func IsEllipsis(node ast.Node) bool {
	_, ok := node.(*ast.Ellipsis)
	return ok
}

// IsBasicLit returns true if a given ast.Node is a literal of basic type (*ast.IsBasicLit).
func IsBasicLit(node ast.Node) bool {
	_, ok := node.(*ast.BasicLit)
	return ok
}

// IsFuncLit returns true if a given ast.Node is a function literal (*ast.IsFuncLit).
func IsFuncLit(node ast.Node) bool {
	_, ok := node.(*ast.FuncLit)
	return ok
}

// IsCompositeLit returns true if a given ast.Node is a composite literal (*ast.IsCompositeLit).
func IsCompositeLit(node ast.Node) bool {
	_, ok := node.(*ast.CompositeLit)
	return ok
}

// IsParenExpr returns true if a given ast.Node is a parenthesized expression (*ast.IsParenExpr).
func IsParenExpr(node ast.Node) bool {
	_, ok := node.(*ast.ParenExpr)
	return ok
}

// IsSelectorExpr returns true if a given ast.Node is a selector expression (*ast.IsSelectorExpr).
func IsSelectorExpr(node ast.Node) bool {
	_, ok := node.(*ast.SelectorExpr)
	return ok
}

// IsIndexExpr returns true if a given ast.Node is an index expression (*ast.IsIndexExpr).
func IsIndexExpr(node ast.Node) bool {
	_, ok := node.(*ast.IndexExpr)
	return ok
}

// IsSliceExpr returns true if a given ast.Node is a slice expression (*ast.IsSliceExpr).
func IsSliceExpr(node ast.Node) bool {
	_, ok := node.(*ast.SliceExpr)
	return ok
}

// IsTypeAssertExpr returns true if a given ast.Node is a type assert expression (*ast.IsTypeAssertExpr).
func IsTypeAssertExpr(node ast.Node) bool {
	_, ok := node.(*ast.TypeAssertExpr)
	return ok
}

// IsCallExpr returns true if a given ast.Node is an expression followed by an argument list (*ast.IsCallExpr).
func IsCallExpr(node ast.Node) bool {
	_, ok := node.(*ast.CallExpr)
	return ok
}

// IsStarExpr returns true if a given ast.Node is a star expression(unary "*" or apointer) (*ast.IsStarExpr)
func IsStarExpr(node ast.Node) bool {
	_, ok := node.(*ast.StarExpr)
	return ok
}

// IsUnaryExpr returns true if a given ast.Node is a unary expression (*ast.IsUnaryExpr).
func IsUnaryExpr(node ast.Node) bool {
	_, ok := node.(*ast.UnaryExpr)
	return ok
}

// IsBinaryExpr returns true if a given ast.Node is a binary expression (*ast.IsBinaryExpr).
func IsBinaryExpr(node ast.Node) bool {
	_, ok := node.(*ast.BinaryExpr)
	return ok
}

// IsKeyValueExpr returns true if a given ast.Node is a (key:value) pair (*ast.IsKeyValueExpr).
func IsKeyValueExpr(node ast.Node) bool {
	_, ok := node.(*ast.KeyValueExpr)
	return ok
}

// IsArrayType returns true if a given ast.Node is an array or slice type (*ast.IsArrayType).
func IsArrayType(node ast.Node) bool {
	_, ok := node.(*ast.ArrayType)
	return ok
}

// IsStructType returns true if a given ast.Node is a struct type (*ast.IsStructType).
func IsStructType(node ast.Node) bool {
	_, ok := node.(*ast.StructType)
	return ok
}

// IsFuncType returns true if a given ast.Node is a function type (*ast.IsFuncType).
func IsFuncType(node ast.Node) bool {
	_, ok := node.(*ast.FuncType)
	return ok
}

// IsInterfaceType returns true if a given ast.Node is an interface type (*ast.IsInterfaceType).
func IsInterfaceType(node ast.Node) bool {
	_, ok := node.(*ast.InterfaceType)
	return ok
}

// IsMapType returns true if a given ast.Node is a map type (*ast.IsMapType).
func IsMapType(node ast.Node) bool {
	_, ok := node.(*ast.MapType)
	return ok
}

// IsChanType returns true if a given ast.Node is a channel type (*ast.IsChanType).
func IsChanType(node ast.Node) bool {
	_, ok := node.(*ast.ChanType)
	return ok
}
