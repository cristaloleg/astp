package astp

import "go/ast"

// AsBadExpr returns a given node as a *ast.BadExpr.
func AsBadExpr(node ast.Node) *ast.BadExpr {
	expr, ok := node.(*ast.BadExpr)
	if !ok {
		panic("expected *ast.BadExpr")
	}
	return expr
}

// AsIdent returns a given node as a *ast.Ident.
func AsIdent(node ast.Node) *ast.Ident {
	expr, ok := node.(*ast.Ident)
	if !ok {
		panic("expected *ast.Ident")
	}
	return expr
}

// AsEllipsis returns a given node as a *ast.Ellipsis.
func AsEllipsis(node ast.Node) *ast.Ellipsis {
	expr, ok := node.(*ast.Ellipsis)
	if !ok {
		panic("expected *ast.Ellipsis")
	}
	return expr
}

// AsBasicLit returns a given node as a *ast.BasicLit.
func AsBasicLit(node ast.Node) *ast.BasicLit {
	expr, ok := node.(*ast.BasicLit)
	if !ok {
		panic("expected *ast.BasicLit")
	}
	return expr
}

// AsFuncLit returns a given node as a *ast.FuncLit.
func AsFuncLit(node ast.Node) *ast.FuncLit {
	expr, ok := node.(*ast.FuncLit)
	if !ok {
		panic("expected *ast.FuncLit")
	}
	return expr
}

// AsCompositeLit returns a given node as a *ast.CompositeLit.
func AsCompositeLit(node ast.Node) *ast.CompositeLit {
	expr, ok := node.(*ast.CompositeLit)
	if !ok {
		panic("expected *ast.CompositeLit")
	}
	return expr
}

// AsParenExpr returns a given node as a *ast.ParenExpr.
func AsParenExpr(node ast.Node) *ast.ParenExpr {
	expr, ok := node.(*ast.ParenExpr)
	if !ok {
		panic("expected *ast.ParenExpr")
	}
	return expr
}

// AsSelectorExpr returns a given node as a *ast.SelectorExpr.
func AsSelectorExpr(node ast.Node) *ast.SelectorExpr {
	expr, ok := node.(*ast.SelectorExpr)
	if !ok {
		panic("expected *ast.SelectorExpr")
	}
	return expr
}

// AsIndexExpr returns a given node as a *ast.IndexExpr.
func AsIndexExpr(node ast.Node) *ast.IndexExpr {
	expr, ok := node.(*ast.IndexExpr)
	if !ok {
		panic("expected *ast.IndexExpr")
	}
	return expr
}

// AsSliceExpr returns a given node as a *ast.SliceExpr.
func AsSliceExpr(node ast.Node) *ast.SliceExpr {
	expr, ok := node.(*ast.SliceExpr)
	if !ok {
		panic("expected *ast.SliceExpr")
	}
	return expr
}

// AsTypeAssertExpr returns a given node as a *ast.TypeAssertExpr.
func AsTypeAssertExpr(node ast.Node) *ast.TypeAssertExpr {
	expr, ok := node.(*ast.TypeAssertExpr)
	if !ok {
		panic("expected *ast.TypeAssertExpr")
	}
	return expr
}

// AsCallExpr returns a given node as a *ast.CallExpr.
func AsCallExpr(node ast.Node) *ast.CallExpr {
	expr, ok := node.(*ast.CallExpr)
	if !ok {
		panic("expected *ast.CallExpr")
	}
	return expr
}

// AsStarExpr returns a given node as a *ast.StarExpr.
func AsStarExpr(node ast.Node) *ast.StarExpr {
	expr, ok := node.(*ast.StarExpr)
	if !ok {
		panic("expected *ast.StarExpr")
	}
	return expr
}

// AsUnaryExpr returns a given node as a *ast.UnaryExpr.
func AsUnaryExpr(node ast.Node) *ast.UnaryExpr {
	expr, ok := node.(*ast.UnaryExpr)
	if !ok {
		panic("expected *ast.UnaryExpr")
	}
	return expr
}

// AsBinaryExpr returns a given node as a *ast.BinaryExpr.
func AsBinaryExpr(node ast.Node) *ast.BinaryExpr {
	expr, ok := node.(*ast.BinaryExpr)
	if !ok {
		panic("expected *ast.BinaryExpr")
	}
	return expr
}

// AsKeyValueExpr returns a given node as a *ast.KeyValueExpr.
func AsKeyValueExpr(node ast.Node) *ast.KeyValueExpr {
	expr, ok := node.(*ast.KeyValueExpr)
	if !ok {
		panic("expected *ast.KeyValueExpr")
	}
	return expr
}

// AsArrayType returns a given node as a *ast.ArrayType.
func AsArrayType(node ast.Node) *ast.ArrayType {
	expr, ok := node.(*ast.ArrayType)
	if !ok {
		panic("expected *ast.ArrayType")
	}
	return expr
}

// AsStructType returns a given node as a *ast.StructType.
func AsStructType(node ast.Node) *ast.StructType {
	expr, ok := node.(*ast.StructType)
	if !ok {
		panic("expected *ast.StructType")
	}
	return expr
}

// AsFuncType returns a given node as a *ast.FuncType.
func AsFuncType(node ast.Node) *ast.FuncType {
	expr, ok := node.(*ast.FuncType)
	if !ok {
		panic("expected *ast.FuncType")
	}
	return expr
}

// AsInterfaceType returns a given node as a *ast.InterfaceType.
func AsInterfaceType(node ast.Node) *ast.InterfaceType {
	expr, ok := node.(*ast.InterfaceType)
	if !ok {
		panic("expected *ast.InterfaceType")
	}
	return expr
}

// AsMapType returns a given node as a *ast.MapType.
func AsMapType(node ast.Node) *ast.MapType {
	expr, ok := node.(*ast.MapType)
	if !ok {
		panic("expected *ast.MapType")
	}
	return expr
}

// AsChanType returns a given node as a *ast.ChanType.
func AsChanType(node ast.Node) *ast.ChanType {
	expr, ok := node.(*ast.ChanType)
	if !ok {
		panic("expected *ast.ChanType")
	}
	return expr
}
