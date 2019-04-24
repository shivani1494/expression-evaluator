An arithmetic expression:
- contains numbers, unary operators (+/-), binary operators (+,-,*,/, ^ exp, log) and parenthesis.
- can be arbitrarily complex (thousands of operations)
- is given as a string

This is a service, that given a large set of expressions can:
-- Print each of them as an expression tree
-- Evaluate all expressions and provide the results
-- Is able to provide visualizations of arithmetic expressions
-- Scales to handle millions of complex expressions with low latency.


## Evaluating expressions at scale

Divide and conquering expression and running parallel computations for non-dependent parts of the expression. 

Traversing the tree in a reverse level-by-level manner since nodes at each level are independent of each other for performing computations(can be proved by expression tree structure). So, threads could be spawned and work on processing nodes level-by-level without idle wait time.

Every expression is a collection of subexpressions (recursive definition) which can be reused to store state and be computed in parallel. The same concept applies to across expressions.

Storing and reusing states using concurrent hash maps where keys are expression strings of fixed length(2 operands and operator) and values are the expression values.


## Evaluating Performance

+ Memory

There can be lots of duplicates within the expression and across. So as we perform computations the values can be stored to avoid re-computations reducing the latency. However, with large hashmaps we would have to consider the tradeoff between hashmap lookups/storing (hashing) and operational latency. 

Although, let’s say the data set doesn’t have lots of duplicates in that case, hashing/lookups/stores can add to the performance latency instead of reducing it.

For large scale, we can use a key-value store for persistent storage across millions of calls to expression evaluator.

+ CPU

For small expressions and relatively small computations, spawning multiple threads can in fact stall the performance because in such cases the operational latency tends to be lower than the latency from spawning threads, syncs, managing locks. (proven through benchmark tests)

In fact, the size of expression and the intensity of computations and size of numbers is directly proportional to what would be the ideal number of threads to spawn in a given case to balance out the trade off between operational latency and thread management latency.

Larger expressions and larger numbers are more computationally-intensive and in those cases multithreading is helpful as proven through benchmark tests as well.

Did benchmarking for various sizes of expressions, sizes of numbers and different operators.


## Initial Approaches

## Approach 1 

Abandoned the approach to couple tree building and running computations simultaneously since was unable to parallelize building an expression tree from an expression and in effect could not run parallel computations. This is because we can’t safely divide an expression with duplicate operands and operators and get arithmetically(following bodmas) correct value. 

Also, note that building a tree itself is not as computationally-intensive(string traversal, stack operations, memcopies, pointer updates). The computationally intensive part is actually performing operations.

Abandoned the approach of storing variable length expressions because as expressions grow larger hashing larger strings can become computationally-intensive. Note, this approach would only help in a top-down traversal.

## Approach 2

Abandoned the approach to traverse the tree from top down, dividing at every level into LST and RST and spawning threads (*Limited by ideal maximum number of threads for a given size of data/expression) Ancestors nodes depend on their child nodes to return values to proceed, in effect, many threads would be idly waiting. Ideally, we would not want threads to be idly waiting. 

This meant it would be useful to know which nodes are independent/dependent on each other - leading to dependency graph or doing tree traversal in a way that extracts such dependencies allowing for parallel computations without threads idly waiting in large expressions.

Initial approach was to use C++ but threads and synchronization are not well-supported, for example no built-in concurrent hashmaps, lock managers or thread pools. There are external libraries but are limited in their support and need research to find what best meets the needs. Golang fit the needs well with channels, goroutines, well-supported testing/benchmarking.




## Assumptions 

+ unary operators unsupported

+ Expression is a well-formatted postfix expression i.e preprocessing step of fully parenthesizing the given expression using BODMAS rule (given expressions are not parenthesized) then converting the infix to postfix expression is already done.

+ tokenizing the passed string using space as a delimiter - trade off is one pass through the entire string as a preprocessing step which can be computationally intensive if there are thousands of operators/operands in the string and millions of such expressions. However, validating all possible numerical formats of an operand would have taken a lot of logic/code so to retain all possible numerical formats doing string splits.

+ Used threads synonymously with goroutines, but aware that goroutines are not OS threads but their behaviours are similar. also, wait groups (instead of channels) would be better to let all threads arrive before further processing.

+ will add more tests division and power operations

+ Used this for converting postfix to infix expression http://scanftree.com/Data_Structure/prefix-postfix-infix-online-converter

+ Used this for evaluating infix expressions to values
http://www.convertit.com/Go/ConvertIt/Calculators/Math/Expression_Calc.ASP

+ can improve with flags for # of threads + table-driven benchmark testing
