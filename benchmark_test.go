package Expression_Evaluator

import (
	"testing"
)

//variable length expressions
var (

	//best perf - 1 thread
	expr1 = "2111 3444 + 4555 588888 + * 100.0 + 20 8 * ^"

	//best perf - 2 threads
	expr2 = "2 3 + 4 5 + * 100.0 + 20 8298988899 log + 298988899 3 + 4321112111 5321112111 + ^ 100.0 + 20 211121112111 * " +
		"+ + 2 321112111 + 4 5 + * 100.0 + 20 8 * + 321112111321112111 3 ^ 4321112111 5 + * 100.0 + 20 8 * + + +"

	//best perf - 3 threads
	expr3 = "2111 3 + 4 5 + * 100.0 + 20321112111 8 * + 2111321112111 3 + 4 5 + * 100.0 + 20321112111 8 * + + 21112111 3 " +
		"+ 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 " +
		"+ 20 8 * + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + + +"

	//best perf - 3 threads
	expr4 = "2111 3 + 4 5 + * 100.0 + 20321112111 8 * + 2111321112111 3 + 4 5 + * 100.0 + 20321112111 8 * + + 21112111 3" +
		" + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 " +
		"+ 20 8 * + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + + + 2111 3 + 4 5 + * 100.0 + " +
		"20321112111 8 * + 2111321112111 3 + 4 5 + * 100.0 + 20321112111 8 * + + 21112111 3 + 4 5 + * 100.0 + " +
		"20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + " +
		"20 8 * + + 2 3 + 4 5 + * 100.0 + 20 8 * + 2 3 + 4 5 + * 100.0 + 20 8 * + + + + +"
)

func SetupBenchmarks(str string) ExpressionEvaluator {
	var expr ExpressionEvaluator
	expr.NewExpressionEvaluator()
	expr.BuildExpressionTree(str)
	return expr
}

//can change the values of expr
var expr = SetupBenchmarks(expr4)

func BenchmarkEvaluateExpressionThread1(b *testing.B) {

	expr.workerThreads = 1
	for n := 1; n < b.N; n++ {
		_, err := expr.EvaluateExpressionTree()
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkEvaluateExpressionThread2(b *testing.B) {

	expr.workerThreads = 2
	for n := 1; n < b.N; n++ {
		_, err := expr.EvaluateExpressionTree()
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkEvaluateExpressionThread3(b *testing.B) {

	expr.workerThreads = 3
	for n := 1; n < b.N; n++ {
		_, err := expr.EvaluateExpressionTree()
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkEvaluateExpressionThread4(b *testing.B) {

	expr.workerThreads = 4
	for n := 1; n < b.N; n++ {
		_, err := expr.EvaluateExpressionTree()
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkEvaluateExpressionThread5(b *testing.B) {

	expr.workerThreads = 5
	for n := 1; n < b.N; n++ {
		_, err := expr.EvaluateExpressionTree()
		if err != nil {
			panic(err)
		}
	}
}

/*
table driven benchamarking
func benchmarkEvaluateExpression(numThreads int, b *testing.B) {

	expr.workerThreads = numThreads

	for n := 0; n < b.N; n++ {
		expr.EvaluateExpressionTree()
	}
}

func BenchmarkEvaluateExpression2(b *testing.B) { benchmarkEvaluateExpression(2, b) }

func BenchmarkEvaluateExpression3(b *testing.B) { benchmarkEvaluateExpression(3, b) }

func BenchmarkEvaluateExpression4(b *testing.B) { benchmarkEvaluateExpression(4, b) }

func BenchmarkEvaluateExpression5(b *testing.B) { benchmarkEvaluateExpression(5, b) }
*/
