package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	lengthOfProgram := len(program.Statements)

	if lengthOfProgram != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", lengthOfProgram)
	}

	tests := []struct {
		expectedIdetifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]

		if !testLetStatement(t, statement, tt.expectedIdetifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {

	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral no 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStatment, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStatment.Name.Value != name {
		t.Errorf("letStatment.Name.Value not '%s'. got=%s", name, letStatment.Name.Value)
		return false
	}

	if letStatment.Name.TokenLiteral() != name {
		t.Errorf("letStatment.Name.TokenLiteral() not '%s'. got=%s",
			name, letStatment.Name.TokenLiteral())
		return false
	}

	return true
}
