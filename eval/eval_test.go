package eval_test

import (
	"testing"

	"github.com/EclesioMeloJunior/monkey-lang/eval"
	"github.com/EclesioMeloJunior/monkey-lang/lexer"
	"github.com/EclesioMeloJunior/monkey-lang/object"
	"github.com/EclesioMeloJunior/monkey-lang/parser"
)

func TestEvaluationLiteralObjects(t *testing.T) {
	testcases := []struct {
		input    string
		expected interface{}
	}{
		{"5;", 5},
		{"10;", 10},

		{"true;", true},
		{"false;", false},
		{"!true;", false},
		{"!false;", true},
		{"!!true;", true},
		{"!!false;", false},
		{"!5;", false},
		{"!!5;", true},
	}

	for _, tt := range testcases {
		evaluated := testEval(tt.input)
		testEvaluatedObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Representation {
	l := lexer.New(input)
	p := parser.New(l)

	prog := p.ParseProgram()
	return eval.Eval(prog)
}

func testEvaluatedObject(t *testing.T, r object.Representation, expected interface{}) {
	switch exp := expected.(type) {
	case int:
		testIntegerObject(t, r, int64(exp))
	case int64:
		testIntegerObject(t, r, int64(exp))
	case bool:
		testBooleanObject(t, r, exp)
	}
}

func testIntegerObject(t *testing.T, r object.Representation, expected int64) {
	result, ok := r.(*object.Integer)
	if !ok {
		t.Fatalf("expected *object.Integer. got=%T (%+v)", r, r)
	}

	if result.Value != expected {
		t.Fatalf("expected %d. got=%d", expected, result.Value)
	}
}

func testBooleanObject(t *testing.T, r object.Representation, expected bool) {
	result, ok := r.(*object.Boolean)
	if !ok {
		t.Fatalf("expected *object.Boolean. got=%T (%+v)", r, r)
	}

	if result.Value != expected {
		t.Fatalf("expected %t. got=%t", expected, result.Value)
	}
}
