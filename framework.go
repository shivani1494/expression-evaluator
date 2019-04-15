package Expression_Evaluator

import (
	"github.com/deckarep/golang-set"
	"sync"
)

const (
	defaultWorkerThreads = 3
)

type node struct {
	data  float64
	op    string
	left  *node
	right *node
}

type ExpressionEvaluator struct {
	root          *node
	stateMap      sync.Map
	operandSet    mapset.Set
	workerThreads int
}

func NewNode(data float64, op string) *node {
	node := new(node)
	node.data = data
	node.op = op
	node.left = nil
	node.right = nil
	return node
}

func (e *ExpressionEvaluator) NewExpressionEvaluator() {

	e.root = nil
	e.workerThreads = defaultWorkerThreads
	e.operandSet = mapset.NewSet()

	e.operandSet.Add("+")
	e.operandSet.Add("-")
	e.operandSet.Add("/")
	e.operandSet.Add("*")
	e.operandSet.Add("^")
	e.operandSet.Add("log")
}
