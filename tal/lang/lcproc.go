package lang

import (
	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc/v2/parsing/stringParsing"
	"github.com/pt-main/lc/v2/parsing/stringParsing/parser3"
)

func NewLexer() *stringParsing.Lexer {
	return stringParsing.NewLexer([]stringParsing.LexerRule{
		{
			Type:    "CODE",
			Pattern: regexp2.MustCompile(`"([^"\\]|\\.)*"`, 0),
		},
		{
			// lc/v2 matches rules against the whole rune slice anchored at the
			// current position, so "^" is a real line start again and no longer
			// swallows the newline in front of a "-- #cmd" annotation. Emitting
			// newlines as their own token type keeps the previous token stream,
			// and the parser skips them as ignorable.
			Type:    "WHITESPACE",
			Pattern: regexp2.MustCompile(`\n+`, 0),
		},
		{
			Type:    "COMMAND",
			Pattern: regexp2.MustCompile(`(?m)^\s*--\s*\#(?<cmd>[^\s]+)\s*(?<args>.*?)$`, 0),
		},
		{
			Type:    "MAINBLOCK",
			Pattern: regexp2.MustCompile(`(?m)^\s*--\s*\@$`, 0),
		},
		{
			Type:    "GLOBALBLOCK",
			Pattern: regexp2.MustCompile(`(?m)^\s*--\s*\@!$`, 0),
		},
		{
			Type:    "BLOCK",
			Pattern: regexp2.MustCompile(`(?m)^\s*--\s*\@(?<name>[^\s]+)$`, 0),
		},
		{
			Type:    "CODE",
			Pattern: regexp2.MustCompile(`(?s).`, 0),
		},
	}, &stringParsing.LexerConfig{
		UseBracketBalance: false,
		Brackets:          [][2]string{},
	})
}

func NewParser() *parser3.Adapter {
	p := parser3.NewParser(NewLexer(), parser3.Grammar{
		"file": parser3.Rule{
			Name: "file",
			Expr: parser3.NodeExpr{
				NodeType: "file",
				Expr: parser3.RepeatExpr{
					Expr: parser3.ChoiceExpr{
						Alternatives: []parser3.Expr{
							parser3.NamedExpr{RuleName: "block"},
							parser3.NamedExpr{RuleName: "mainblock"},
							parser3.NamedExpr{RuleName: "globalblock"},
						},
					},
					Min: 1,
				},
			},
		},
		"mainblock": parser3.Rule{
			Name: "mainblock",
			Expr: parser3.NodeExpr{
				NodeType: "block",
				Expr: parser3.SequenceExpr{
					Exprs: []parser3.Expr{
						parser3.TokenExpr{TokenType: "MAINBLOCK"},
						parser3.NamedExpr{RuleName: "code"},
					},
				},
			},
		},
		"globalblock": parser3.Rule{
			Name: "globalblock",
			Expr: parser3.NodeExpr{
				NodeType: "block",
				Expr: parser3.SequenceExpr{
					Exprs: []parser3.Expr{
						parser3.TokenExpr{TokenType: "GLOBALBLOCK"},
						parser3.NamedExpr{RuleName: "code"},
					},
				},
			},
		},
		"block": parser3.Rule{
			Name: "block",
			Expr: parser3.NodeExpr{
				NodeType: "block",
				Expr: parser3.SequenceExpr{
					Exprs: []parser3.Expr{
						parser3.TokenExpr{TokenType: "BLOCK"},
						parser3.RepeatExpr{
							Expr: parser3.TokenExpr{TokenType: "COMMAND"},
							Min:  0,
						},
						parser3.NamedExpr{RuleName: "code"},
					},
				},
			},
		},
		"code": parser3.Rule{
			Name: "code",
			Expr: parser3.NodeExpr{
				NodeType: "code",
				Expr: parser3.RepeatExpr{
					Expr: parser3.TokenExpr{TokenType: "CODE"},
					Min:  0,
				},
			},
		},
	}, "file", []string{"WHITESPACE"})
	return &parser3.Adapter{Parser: p}
}
