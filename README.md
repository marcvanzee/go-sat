# go-sat: Simple Go SAT Solver

A simple SAT solver that can either use a recursive or an iterative algorithm. The code is based on [Knuth's SAT0W program](http://www-cs-faculty.stanford.edu/~uno/programs.html). Much of the code is based on the Python implementation by sahands, which can be downloaded [here](https://github.com/sahands/simple-sat). He also wrote a nice explanatory article that can be viewed [here](http://sahandsaba.com/understanding-sat-by-implementing-a-simple-sat-solver-in-python.html).

### Installing

No additional packages are required. Get this package as follows:

```go get github.com/marcvanzee/go-sat```

### Input syntax:

```
$ go run main.go -h
  -all
        Output all possible solutions (default true)                  
  -brief                                                              
        Only output variables assigned true                           
  -i string                                                           
        Read from given file instead of stdin                         
  -recursive                                                          
        Use recursive algorithm instead of iterative                
  -starting_with string                                               
        Only output variables with names starting with the given string
  -verbose                                          
        Verbose output (default true)              
```
