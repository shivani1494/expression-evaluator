package Expression_Evaluator

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func SetupTest() ExpressionEvaluator {
	var expr ExpressionEvaluator
	expr.NewExpressionEvaluator()
	return expr
}

type buildPrintTestTest struct {
	input    string
	expected string
}

var buildPrintTestTests = []buildPrintTestTest{
	{"2 3 + 4 5 + * 100.0 + 20 8 * +", "2.0 3.0 + 4.0 5.0 + * 100.0 + 20.0 8.0 * + "},
	{"2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + " +
		"4 5 + * 100.0 + 20 8 * + + +",
		"2.0 3.0 + 4.0 5.0 + * 100.0 + 20.0 8.0 * + 2.0 3.0 + 4.0 5.0 + * 100.0 + 20.0 8.0 * + + 2.0 3.0 + " +
			"4.0 5.0 + * 100.0 + 20.0 8.0 * + 2.0 3.0 + 4.0 5.0 + * 100.0 + 20.0 8.0 * + + + "},
	{"", ""},
}

type evaluationTest struct {
	input    string
	expected float64
}

var evaluationTests = []evaluationTest{
	{"2 3 + 4 5 + * 100.0 + 20 8 * +", 305.0},

	{"", math.Inf(-1)},

	{"2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + +", 610},

	{"2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + " +
		"4 5 + * 100.0 + 20 8 * + + +", 1220},

	{"2111 3 + 4 5 + * 100.0 + 20321112111 8 * + 2111321112111 3 + 4 5 + * 100.0 + 20321112111 8 * + + 21112111 3" +
		" + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 " +
		"+ 20 8 * + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + + + 2111 3 + 4 5 + * 100.0 + " +
		"20321112111 8 * + 2111321112111 3 + 4 5 + * 100.0 + 20321112111 8 * + + 21112111 3 + 4 5 + * 100.0 + " +
		"20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + " +
		"20 8 * + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + + + +", 38654435665678},
	{"10.5 2 log 10.5 2 log + 10.5 2 log + 10.5 2 log + 10.5 2 log + 10.5 2 log + 10.5 2 log + 10.5 2 log + 10.5 2 log + " +
		"10.5 2 log + 10.5 2 log + 10.5 2 log +", 40.707809073345125},
}

type levelOrderTraversalTest struct {
	input string
}

var levelOrderTraversalTests = []levelOrderTraversalTest{
	{"2 3 + 4 5 + * 100.0 + 20 8 * +"},

	{"2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + +"},

	{"2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + " +
		"* 100.0 + 20 8 * + + +"},

	{"2111 3 + 4 5 + * 100.0 + 20321112111 8 * + 2111321112111 3 + 4 5 + * 100.0 + 20321112111 8 * + + 21112111 3" +
		" + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 " +
		"+ 20 8 * + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + + + 2111 3 + 4 5 + * 100.0 + " +
		"20321112111 8 * + 2111321112111 3 + 4 5 + * 100.0 + 20321112111 8 * + + 21112111 3 + 4 5 + * 100.0 + " +
		"20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + " +
		"20 8 * + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + + + +"},
	{"10.5 2 log 10.5 2 log + 10.5 2 log + 10.5 2 log + 10.5 2 log + 10.5 2 log + 10.5 2 log + 10.5 2 log + 10.5 2 log + " +
		"10.5 2 log + 10.5 2 log + 10.5 2 log +"},
}

func Test_BuildAndPrintExpressionTree(t *testing.T) {

	for _, tt := range buildPrintTestTests {

		expr := SetupTest()

		_, err := expr.BuildExpressionTree(tt.input)
		if err != nil {
			assert.Error(t, err, "could not build the tree")
		}
		binExprTree := expr.PrintExpressionTree()
		//fmt.Println(binExprTree)
		assert.Equal(t, binExprTree, tt.expected)
	}
}

// Note- could do setup only once and modify the root value to nil
// across iterations and reuse the sync.map across calls to reuse state
// but a sub-optimal approach
func Test_EvaluateExpressionTree(t *testing.T) {

	for _, tt := range evaluationTests {

		expr := SetupTest()
		_, err := expr.BuildExpressionTree(tt.input)
		if err != nil {
			assert.Error(t, err, "could not build the tree")
		}
		evaluatedExpr, err := expr.EvaluateExpressionTree()
		assert.Equal(t, evaluatedExpr, tt.expected)
	}
}

func Test_binaryTreeLevelOrderTraversal(t *testing.T) {

	for _, tt := range levelOrderTraversalTests {

		expr := SetupTest()
		root, err := expr.BuildExpressionTree(tt.input)
		if err != nil {
			assert.Error(t, err, "could not build the tree")
		}
		arrStacks := binaryTreeLevelOrderTraversal(root)

		for i := 0; i < len(arrStacks); i++ {
			for j := 0; j < len(arrStacks[i]); j++ {

				if arrStacks[i][j].op == "" {
					fmt.Printf(" %+v ", arrStacks[i][j].data)
				} else {
					fmt.Printf(" %+v ", arrStacks[i][j].op)
				}

			}
			fmt.Println()
		}
	}
}

func Test_isOperator(t *testing.T) {

	//construct an operand set
	operandSet := mapset.NewSet()
	operandSet.Add("+")
	operandSet.Add("-")
	operandSet.Add("/")
	operandSet.Add("*")
	operandSet.Add("^")
	operandSet.Add("log")

	present := isOperator("log", operandSet)
	assert.Equal(t, true, present)

	present = isOperator("q", operandSet)
	assert.Equal(t, false, present)

	present = isOperator("..", operandSet)
	assert.Equal(t, false, present)

}

func Test_compute(t *testing.T) {

	computedNum, err := compute("log", 10.7, 5.8)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, 1.3483704899018047, computedNum)

	computedNum, err = compute("^", 3.7, 4.7)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, 468.3258209785275, computedNum)

	computedNum, err = compute("+", 109090909.7, 533333.8)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, 1.096242435e+08, computedNum)

	computedNum, err = compute("-", 5.634563266e+07, 51111.8)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, 5.629452086e+07, computedNum)

	computedNum, err = compute("*", 105555.7, 533.8)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, 5.634563266e+07, computedNum)

	computedNum, err = compute("/", -109999.7, 987129349)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, -0.00011143392718637525, computedNum)

	computedNum, err = compute("q", 10.7, 5.8)
	if err != nil {
		assert.Error(t, err, "operator not found")
	}
}
