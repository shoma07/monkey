package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

// 文
type Statement interface {
	Node
	statementNode()
}

// 式
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// let文
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// return文
type ReturnStatement struct {
	Token token.Token // token.RETURN
	Value Expression
}

func (n *ReturnStatement) statementNode() {}
func (n *ReturnStatement) TokenLiteral() string {
	return n.Token.Literal
}
func (n *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(n.TokenLiteral() + " ")

	if n.Value != nil {
		out.WriteString(n.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// 式文
type ExpressionStatement struct {
	Token      token.Token // 式の最初のトークン
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// 識別子
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

// 整数リテラル
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (n *IntegerLiteral) expressionNode() {}
func (n *IntegerLiteral) TokenLiteral() string {
	return n.Token.Literal
}
func (n *IntegerLiteral) String() string {
	return n.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token // 前置トークン ex) !
	Operator string
	Right    Expression
}

func (n *PrefixExpression) expressionNode() {}
func (n *PrefixExpression) TokenLiteral() string {
	return n.Token.Literal
}
func (n *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(n.Operator)
	out.WriteString(n.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token // 中値トークン ex) +
	Left     Expression
	Operator string
	Right    Expression
}

func (n *InfixExpression) expressionNode() {}
func (n *InfixExpression) TokenLiteral() string {
	return n.Token.Literal
}
func (n *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(n.Left.String())
	out.WriteString(" " + n.Operator + " ")
	out.WriteString(n.Right.String())
	out.WriteString(")")

	return out.String()
}

// 真偽
type Boolean struct {
	Token token.Token
	Value bool
}

func (n *Boolean) expressionNode() {}
func (n *Boolean) TokenLiteral() string {
	return n.Token.Literal
}
func (n *Boolean) String() string {
	return n.Token.Literal
}

type BlockStatement struct {
	Token      token.Token // { トークン
	Statements []Statement
}

func (n *BlockStatement) statementNode() {}
func (n *BlockStatement) TokenLiteral() string {
	return n.Token.Literal
}
func (n *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range n.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type IfExpression struct {
	Token       token.Token // 'if' トークン
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (n *IfExpression) expressionNode() {}
func (n *IfExpression) TokenLiteral() string {
	return n.Token.Literal
}
func (n *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(n.Condition.String())
	out.WriteString(" ")
	out.WriteString(n.Consequence.String())

	if n.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(n.Alternative.String())
	}

	return out.String()
}

// 関数リテラル
type FunctionLiteral struct {
	Token      token.Token // 'fn' トークン
	Parameters []*Identifier
	Body       *BlockStatement
}

func (n *FunctionLiteral) expressionNode() {}
func (n *FunctionLiteral) TokenLiteral() string {
	return n.Token.Literal
}
func (n *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range n.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(n.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(n.Body.String())

	return out.String()
}

// 関数呼び出し
type CallExpression struct {
	Token     token.Token // '('トークン
	Function  Expression
	Arguments []Expression
}

func (n *CallExpression) expressionNode() {}
func (n *CallExpression) TokenLiteral() string {
	return n.Token.Literal
}
func (n *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range n.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(n.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
