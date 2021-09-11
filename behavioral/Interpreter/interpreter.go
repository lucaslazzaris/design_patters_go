package main

import "fmt"

type Context struct {
	cont map[VariableExp]bool
}

func (context *Context) lookup(exp VariableExp) bool {
	return context.cont[exp]
}

func (context *Context) assign(exp VariableExp, assigned bool) {
	context.cont[exp] = assigned
}

type BooleanExp interface {
	evaluate(Context) bool
	replace(string, BooleanExp) BooleanExp
	copy() BooleanExp
}

type VariableExp struct {
	name string	
}

func newVariableExp(name string) *VariableExp {
	return &VariableExp{name: name}
}

func (exp *VariableExp) evaluate(context Context) bool {
	return context.lookup(*exp)
}

func (exp *VariableExp) copy() BooleanExp {
	return newVariableExp(exp.name)
}

func (exp *VariableExp) replace(newName string, boolExp BooleanExp) BooleanExp {
	if (exp.name == newName){
		return boolExp.copy()
	} else {
		return newVariableExp(exp.name)
	}
}

type AndExp struct {
	op1 BooleanExp
	op2 BooleanExp
}

func (exp *AndExp) evaluate(context Context) bool {
	return exp.op1.evaluate(context) && exp.op2.evaluate(context)
}

func (exp *AndExp) copy() BooleanExp {
	return &AndExp{op1: exp.op1.copy(), op2: exp.op2.copy()}
}

func (exp *AndExp) replace(name string, boolExp BooleanExp) BooleanExp {
	return &AndExp{
		op1: exp.op1.replace(name, exp),
		op2: exp.op2.replace(name, exp),
	}
}

type OrExp struct {
	op1 BooleanExp
	op2 BooleanExp
}

func (exp *OrExp) evaluate(context Context) bool {
	return exp.op1.evaluate(context) || exp.op2.evaluate(context)
}

func (exp *OrExp) copy() BooleanExp {
	return &OrExp{op1: exp.op1.copy(), op2: exp.op2.copy()}
}

func (exp *OrExp) replace(name string, boolExp BooleanExp) BooleanExp {
	return &OrExp{
		op1: exp.op1.replace(name, exp),
		op2: exp.op2.replace(name, exp),
	}
}

type NotExp struct {
	op BooleanExp
}

func (exp *NotExp) evaluate(context Context) bool {
	return !exp.op.evaluate(context)
}

func (exp *NotExp) copy() BooleanExp {
	return &NotExp{op: exp.op.copy()}
}

func (exp *NotExp) replace(name string, boolExp BooleanExp) BooleanExp {
	return &NotExp{
		op: exp.op.replace(name, exp),
	}
}

func main() {
	context := Context{cont: make(map[VariableExp]bool)}

	x := VariableExp{name: "X"}
	y := VariableExp{name: "Y"}

	expression := OrExp{
		op1: &AndExp{op1: &y, op2: &x},
		op2: &AndExp{op1: &y, op2: &NotExp{op: &x}},
	}

	context.assign(x, false)
	context.assign(y, true)

	fmt.Println(expression.evaluate(context))
}
