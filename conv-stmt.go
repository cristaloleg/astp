package astp

import "go/ast"

// AsBadStmt returns a given node as a *ast.BadStmt.
func AsBadStmt(node ast.Node) *ast.BadStmt {
	stmt, ok := node.(*ast.BadStmt)
	if !ok {
		panic("expected *ast.BadStmt")
	}
	return stmt
}

// AsDeclStmt returns a given node as a *ast.DeclStmt.
func AsDeclStmt(node ast.Node) *ast.DeclStmt {
	stmt, ok := node.(*ast.DeclStmt)
	if !ok {
		panic("expected *ast.DeclStmt")
	}
	return stmt
}

// AsEmptyStmt returns a given node as a *ast.EmptyStmt.
func AsEmptyStmt(node ast.Node) *ast.EmptyStmt {
	stmt, ok := node.(*ast.EmptyStmt)
	if !ok {
		panic("expected *ast.EmptyStmt")
	}
	return stmt
}

// AsLabeledStmt if a node as isXXX(*ast.LabeledStmt.
func AsLabeledStmt(node ast.Node) *ast.LabeledStmt {
	stmt, ok := node.(*ast.LabeledStmt)
	if !ok {
		panic("expected *ast.LabeledStmt")
	}
	return stmt
}

// AsExprStmt returns a given node as a *ast.ExprStmt.
func AsExprStmt(node ast.Node) *ast.ExprStmt {
	stmt, ok := node.(*ast.ExprStmt)
	if !ok {
		panic("expected *ast.ExprStmt")
	}
	return stmt
}

// AsSendStmt returns a given node as a *ast.SendStmt.
func AsSendStmt(node ast.Node) *ast.SendStmt {
	stmt, ok := node.(*ast.SendStmt)
	if !ok {
		panic("expected *ast.SendStmt")
	}
	return stmt
}

// AsIncDecStmt returns a given node as a *ast.IncDecStmt.
func AsIncDecStmt(node ast.Node) *ast.IncDecStmt {
	stmt, ok := node.(*ast.IncDecStmt)
	if !ok {
		panic("expected *ast.IncDecStmt")
	}
	return stmt
}

// AsAssignStmt returns a given node as a *ast.AssignStmt.
func AsAssignStmt(node ast.Node) *ast.AssignStmt {
	stmt, ok := node.(*ast.AssignStmt)
	if !ok {
		panic("expected *ast.AssignStmt")
	}
	return stmt
}

// AsGoStmt returns a given node as a *ast.GoStmt.
func AsGoStmt(node ast.Node) *ast.GoStmt {
	stmt, ok := node.(*ast.GoStmt)
	if !ok {
		panic("expected *ast.GoStmt")
	}
	return stmt
}

// AsDeferStmt returns a given node as a *ast.DeferStmt.
func AsDeferStmt(node ast.Node) *ast.DeferStmt {
	stmt, ok := node.(*ast.DeferStmt)
	if !ok {
		panic("expected *ast.DeferStmt")
	}
	return stmt
}

// AsReturnStmt returns a given node as a *ast.ReturnStmt.
func AsReturnStmt(node ast.Node) *ast.ReturnStmt {
	stmt, ok := node.(*ast.ReturnStmt)
	if !ok {
		panic("expected *ast.ReturnStmt")
	}
	return stmt
}

// AsBranchStmt returns a given node as a *ast.BranchStmt.
func AsBranchStmt(node ast.Node) *ast.BranchStmt {
	stmt, ok := node.(*ast.BranchStmt)
	if !ok {
		panic("expected *ast.BranchStmt")
	}
	return stmt
}

// AsBlockStmt returns a given node as a *ast.BlockStmt.
func AsBlockStmt(node ast.Node) *ast.BlockStmt {
	stmt, ok := node.(*ast.BlockStmt)
	if !ok {
		panic("expected *ast.BlockStmt")
	}
	return stmt
}

// AsIfStmt returns a given node as a *ast.IfStmt.
func AsIfStmt(node ast.Node) *ast.IfStmt {
	stmt, ok := node.(*ast.IfStmt)
	if !ok {
		panic("expected *ast.IfStmt")
	}
	return stmt
}

// AsCaseClause returns a given node as a *ast.CaseClause.
func AsCaseClause(node ast.Node) *ast.CaseClause {
	stmt, ok := node.(*ast.CaseClause)
	if !ok {
		panic("expected *ast.CaseClause")
	}
	return stmt
}

// AsSwitchStmt returns a given node as a *ast.SwitchStmt.
func AsSwitchStmt(node ast.Node) *ast.SwitchStmt {
	stmt, ok := node.(*ast.SwitchStmt)
	if !ok {
		panic("expected *ast.SwitchStmt")
	}
	return stmt
}

// AsTypeSwitchStmt if a node as isXXX(*ast.TypeSwitchStmt.
func AsTypeSwitchStmt(node ast.Node) *ast.TypeSwitchStmt {
	stmt, ok := node.(*ast.TypeSwitchStmt)
	if !ok {
		panic("expected *ast.TypeSwitchStmt")
	}
	return stmt
}

// AsCommClause returns a given node as a *ast.CommClause.
func AsCommClause(node ast.Node) *ast.CommClause {
	stmt, ok := node.(*ast.CommClause)
	if !ok {
		panic("expected *ast.CommClause")
	}
	return stmt
}

// AsSelectStmt returns a given node as a *ast.SelectStmt.
func AsSelectStmt(node ast.Node) *ast.SelectStmt {
	stmt, ok := node.(*ast.SelectStmt)
	if !ok {
		panic("expected *ast.SelectStmt")
	}
	return stmt
}

// AsForStmt returns a given node as a *ast.ForStmt.
func AsForStmt(node ast.Node) *ast.ForStmt {
	stmt, ok := node.(*ast.ForStmt)
	if !ok {
		panic("expected *ast.ForStmt")
	}
	return stmt
}

// AsRangeStmt returns a given node as a *ast.RangeStmt.
func AsRangeStmt(node ast.Node) *ast.RangeStmt {
	stmt, ok := node.(*ast.RangeStmt)
	if !ok {
		panic("expected *ast.RangeStmt")
	}
	return stmt
}
