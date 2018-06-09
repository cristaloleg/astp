package astp

import "go/ast"

// IsStmt returns true if a given ast.Node is a statement(ast.Stmt).
func IsStmt(node ast.Node) bool {
	return IsBadStmt(node) ||
		IsDeclStmt(node) ||
		IsEmptyStmt(node) ||
		IsLabeledStmt(node) ||
		IsExprStmt(node) ||
		IsSendStmt(node) ||
		IsIncDecStmt(node) ||
		IsAssignStmt(node) ||
		IsGoStmt(node) ||
		IsDeferStmt(node) ||
		IsReturnStmt(node) ||
		IsBranchStmt(node) ||
		IsBlockStmt(node) ||
		IsIfStmt(node) ||
		IsCaseClause(node) ||
		IsSwitchStmt(node) ||
		IsTypeSwitchStmt(node) ||
		IsCommClause(node) ||
		IsSelectStmt(node) ||
		IsForStmt(node) ||
		IsRangeStmt(node)
}

// IsBadStmt returns true if a given ast.Node is a XXX(*ast.BadStmt)
func IsBadStmt(node ast.Node) bool {
	_, ok := node.(*ast.BadStmt)
	return ok
}

// IsDeclStmt returns true if a given ast.Node is a XXX(*ast.DeclStmt)
func IsDeclStmt(node ast.Node) bool {
	_, ok := node.(*ast.DeclStmt)
	return ok
}

// IsEmptyStmt returns true if a given ast.Node is a XXX(*ast.EmptyStmt)
func IsEmptyStmt(node ast.Node) bool {
	_, ok := node.(*ast.EmptyStmt)
	return ok
}

// IsLabeledStmt returns true if a given ast.Node is a XXX(*ast.LabeledStmt)
func IsLabeledStmt(node ast.Node) bool {
	_, ok := node.(*ast.LabeledStmt)
	return ok
}

// IsExprStmt returns true if a given ast.Node is a XXX(*ast.ExprStmt)
func IsExprStmt(node ast.Node) bool {
	_, ok := node.(*ast.ExprStmt)
	return ok
}

// IsSendStmt returns true if a given ast.Node is a XXX(*ast.SendStmt)
func IsSendStmt(node ast.Node) bool {
	_, ok := node.(*ast.SendStmt)
	return ok
}

// IsIncDecStmt returns true if a given ast.Node is a XXX(*ast.IncDecStmt)
func IsIncDecStmt(node ast.Node) bool {
	_, ok := node.(*ast.IncDecStmt)
	return ok
}

// IsAssignStmt returns true if a given ast.Node is a XXX(*ast.AssignStmt)
func IsAssignStmt(node ast.Node) bool {
	_, ok := node.(*ast.AssignStmt)
	return ok
}

// IsGoStmt returns true if a given ast.Node is a XXX(*ast.GoStmt)
func IsGoStmt(node ast.Node) bool {
	_, ok := node.(*ast.GoStmt)
	return ok
}

// IsDeferStmt returns true if a given ast.Node is a XXX(*ast.DeferStmt)
func IsDeferStmt(node ast.Node) bool {
	_, ok := node.(*ast.DeferStmt)
	return ok
}

// IsReturnStmt returns true if a given ast.Node is a XXX(*ast.ReturnStmt)
func IsReturnStmt(node ast.Node) bool {
	_, ok := node.(*ast.ReturnStmt)
	return ok
}

// IsBranchStmt returns true if a given ast.Node is a XXX(*ast.BranchStmt)
func IsBranchStmt(node ast.Node) bool {
	_, ok := node.(*ast.BranchStmt)
	return ok
}

// IsBlockStmt returns true if a given ast.Node is a XXX(*ast.BlockStmt)
func IsBlockStmt(node ast.Node) bool {
	_, ok := node.(*ast.BlockStmt)
	return ok
}

// IsIfStmt returns true if a given ast.Node is a XXX(*ast.IfStmt)
func IsIfStmt(node ast.Node) bool {
	_, ok := node.(*ast.IfStmt)
	return ok
}

// IsCaseClause returns true if a given ast.Node is a XXX(*ast.CaseClause)
func IsCaseClause(node ast.Node) bool {
	_, ok := node.(*ast.CaseClause)
	return ok
}

// IsSwitchStmt returns true if a given ast.Node is a XXX(*ast.SwitchStmt)
func IsSwitchStmt(node ast.Node) bool {
	_, ok := node.(*ast.SwitchStmt)
	return ok
}

// IsTypeSwitchStmt returns true if a given ast.Node is a XXX(*ast.TypeSwitchStmt)
func IsTypeSwitchStmt(node ast.Node) bool {
	_, ok := node.(*ast.TypeSwitchStmt)
	return ok
}

// IsCommClause returns true if a given ast.Node is a XXX(*ast.CommClause)
func IsCommClause(node ast.Node) bool {
	_, ok := node.(*ast.CommClause)
	return ok
}

// IsSelectStmt returns true if a given ast.Node is a XXX(*ast.SelectStmt)
func IsSelectStmt(node ast.Node) bool {
	_, ok := node.(*ast.SelectStmt)
	return ok
}

// IsForStmt returns true if a given ast.Node is a XXX(*ast.ForStmt)
func IsForStmt(node ast.Node) bool {
	_, ok := node.(*ast.ForStmt)
	return ok
}

// IsRangeStmt returns true if a given ast.Node is a XXX(*ast.RangeStmt)
func IsRangeStmt(node ast.Node) bool {
	_, ok := node.(*ast.RangeStmt)
	return ok
}
