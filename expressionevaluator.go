package Expression_Evaluator

import (

	"github.com/golang-collections/collections/stack"
	"strconv"
	"fmt"
	"github.com/golang/glog"
	"strings"
	"bytes"
	"math"
	"errors"
)

func ( e *ExpressionEvaluator ) BuildExpressionTree(currExpr string) (*node, error) {

	opDataNodes := stack.New()
	tokens := strings.Split(currExpr, " ")

	for i := 0; i < len(tokens); i++ {

		if isOperator(tokens[i], e.operandSet) {

			curr := NewNode(-1, string(tokens[i]))

			a := opDataNodes.Pop()
			curr.right = a.(*node)
			b := opDataNodes.Pop()
			curr.left = b.(*node)

			opDataNodes.Push(curr);

			e.root = curr

		} else {
			num, err := strconv.ParseFloat(tokens[i], 64)

			if err != nil {
				return nil, err
			}
			opDataNodes.Push(NewNode(num, ""))
		}
	}

	return e.root, nil
}


func ( e *ExpressionEvaluator ) PrintExpressionTree() string {

	if e.root == nil {
		return ""
	}

	var binExprTree *bytes.Buffer
	binExprTree = bytes.NewBufferString("")

	postOrder(e.root, binExprTree)

	return binExprTree.String()
}

func ( e* ExpressionEvaluator ) evaluateExpressionPerLevel(id int, stacks <-chan *node, done chan<- int) {
	glog.Info("Worker-", id, "started")

	for j := range stacks {

		if j.op == "end" {
			break
		}

		if j.op == "" {
			continue
		}

		key := fmt.Sprintf("%f", j.left.data) + j.op + fmt.Sprintf("%f", j.right.data)
		glog.Info("Printing key- ", key)

		val, present := e.stateMap.Load(key)

		if !present {

			v, err := compute(j.op, j.left.data, j.right.data)

			if err != nil {
				panic(err)
			}

			e.stateMap.Store(key, v)
			val = v
		}

		j.data = val.(float64)
	}

	glog.Info("Worker-", id, "finished")
	done <- 1
}


func( e *ExpressionEvaluator ) EvaluateExpressionTree() (float64, error) {

	if e.root == nil {
		return math.Inf(-1), errors.New("root not found - expression tree is empty")
	}

	vecStacks := binaryTreeLevelOrderTraversal(e.root)

	for i := len(vecStacks)-1; i >= 0; i-- {

		jobs := make(chan *node, len(vecStacks[i])+ e.workerThreads)
		done := make(chan int, e.workerThreads)

		for j := 0; j < len(vecStacks[i]); j++ {
			jobs <- vecStacks[i][j]
		}

		node := NewNode(0, "end")
		for w := 1; w <= e.workerThreads; w++ {
			jobs <- node
		}

		for w := 1; w <= e.workerThreads; w++ {
			go e.evaluateExpressionPerLevel(w, jobs, done)
		}

		var completedThreadCount = 0
		for {
			for _ = range done {
				completedThreadCount += 1
				if completedThreadCount == e.workerThreads {
					break
				}
			}
			if completedThreadCount == e.workerThreads {
				break
			}
		}
		close(jobs)
		close(done)
	}

	glog.Info("Evaluated Binary Expression Tree value", e.root.data)

	return e.root.data, nil
}