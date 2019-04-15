package Expression_Evaluator

import (
	"bytes"
	"errors"
	"github.com/deckarep/golang-set"
	"github.com/golang-collections/collections/queue"
	"github.com/golang/glog"
	"math"
	"strconv"
)

func binaryTreeLevelOrderTraversal(root *node) [][]*node {

	if root == nil {
		return nil
	}

	queueNodes := queue.New()
	queueNodes.Enqueue(root)
	var nodeCount int

	vecStacks := make([][]*node, 1)

	currLevel := 0

	for queueNodes.Len() > 0 {

		nodeCount = queueNodes.Len()

		for nodeCount > 0 {

			currNode := queueNodes.Peek().(*node)
			vecStacks[currLevel] = append(vecStacks[currLevel], currNode)
			queueNodes.Dequeue()
			if currNode.left != nil {
				queueNodes.Enqueue(currNode.left)
			}

			if currNode.right != nil {
				queueNodes.Enqueue(currNode.right)
			}
			nodeCount--
		}

		if nodeCount == 0 {
			currLevel++
			vecStacks = append(vecStacks, []*node{})
		}
	}

	return vecStacks
}

func isOperator(currElem string, operandSet mapset.Set) bool {

	if operandSet.Contains(currElem) {
		return true
	}
	return false
}

//gets tested when printing the binary expression tree
func postOrder(curr *node, finalTree *bytes.Buffer) {

	if curr == nil {
		return
	}

	postOrder(curr.left, finalTree)
	postOrder(curr.right, finalTree)

	if curr.op == "" {
		str := strconv.FormatFloat(curr.data, 'f', 1, 64)
		finalTree.WriteString(str)
		finalTree.WriteString(" ")

	} else {
		finalTree.WriteString(curr.op)
		finalTree.WriteString(" ")
	}
}

func compute(op string, num1 float64, num2 float64) (float64, error) {

	glog.Info("Performing Computation %+v, %+v, %+v", num1, op, num2)
	switch op {

	case "+":
		return num1 + num2, nil

	case "-":
		return num1 - num2, nil
		break

	case "*":
		return num1 * num2, nil
		break

	case "/":
		return num1 / num2, nil
		break

	case "^":
		return math.Pow(num1, num2), nil
		break

	case "log":
		return math.Log(num1) / math.Log(num2), nil
		break

	default:
		return -1, errors.New("Operator not found " + op)
	}

	return -1, nil
}
