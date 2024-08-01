package parser

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/vishnu/glox/token"
)

type Expr interface {
	accept(visitor IExprVisitor) string
}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

type Grouping struct {
	Expression Expr
}

type Literal struct {
	Value any
}

type Unary struct {
	Operator token.Token
	Right    Expr
}

func (binary *Binary) accept(visitor IExprVisitor) string {
	return visitor.visitBinary(*binary)
}
func (unary *Unary) accept(visitor IExprVisitor) string {
	return visitor.visitUnary(*unary)
}
func (grouping *Grouping) accept(visitor IExprVisitor) string {
	return visitor.visitGrouping(*grouping)
}
func (literal *Literal) accept(visitor IExprVisitor) string {
	return visitor.visitLiteral(*literal)
}

type IExprVisitor interface {
	visitBinary(binary Binary) string
	visitGrouping(group Grouping) string
	visitLiteral(literal Literal) string
	visitUnary(unary Unary) string
}

type ExprVisitor struct {
}

func (ExprVisitor ExprVisitor) visitBinary(expr Binary) string {
	return paranthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}
func (ExprVisitor ExprVisitor) visitGrouping(expr Grouping) string {
	return paranthesize("group", expr.Expression)
}
func (ExprVisitor ExprVisitor) visitLiteral(expr Literal) string {
	if expr.Value == nil {
		return "nil"
	}
	// fmt.Println("type: %t", expr.Value.(type))
	// return fmt.Sprintf(expr.Value)
	switch expr.Value.(type) {
	case string:
		return expr.Value.(string)
	case int:
		return strconv.Itoa(expr.Value.(int))
	case float64:
		val, _ := expr.Value.(float64)
		return fmt.Sprint(val)
	}
	return ""
}
func (ExprVisitor ExprVisitor) visitUnary(expr Unary) string {
	return paranthesize(expr.Operator.Lexeme, expr.Right)
}

func paranthesize(name string, expr ...Expr) string {
	var buffer bytes.Buffer

	_, _ = buffer.WriteString("(")
	_, _ = buffer.WriteString(name)
	for _, x := range expr {
		_, _ = buffer.WriteString(" ")
		_, _ = buffer.WriteString(x.accept(ExprVisitor{}))
	}
	buffer.WriteString(")")
	return buffer.String()
}

func AstPrinter(expr Expr) string {
	return expr.accept(ExprVisitor{})
}
