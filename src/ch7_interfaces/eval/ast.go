package eval

//Expr expression interface
type Expr interface {
	Eval(env Env) float64
}

//Var indentifies variable
type Var string

type literal float64

type unary struct {
	op rune
	x  Expr
}

type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string
	args []Expr
}
