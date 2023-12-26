package parser

import (
	"fmt"
	"slices"

	"github.com/tsingbx/interpreter/token"
)

func (p *Paser) skipToToken(s ...token.TokenType) {
	notTokenTypeFun := func(e token.TokenType) bool {
		return e != p.curToken.Type
	}
	for {
		contains := slices.ContainsFunc(s, notTokenTypeFun)
		if contains {
			break
		}
		p.nextToken()
	}
}

func (p *Paser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Paser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Paser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Paser) Errors() []string {
	return p.errors
}

func (p *Paser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Paser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
