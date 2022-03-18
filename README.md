# yacal 
![](https://img.shields.io/github/workflow/status/bragov4ik/yacal/PR%20to%20master%20branch?label=tests)
![](https://img.shields.io/github/go-mod/go-version/bragov4ik/yacal)
![](https://img.shields.io/github/v/release/bragov4ik/yacal)
![](https://img.shields.io/github/license/bragov4ik/yacal)


Yet Another Compiler Adjustment of Lisp - interpreter for toy functional language written within Compiler Construction 
(S22) course in IU.


## How to run

### Naive

On Linux/MacOs download appropriate interpreter and simply run it from console 
```shell
./yacal <files>
```
On Windows the same
```shell
./yacal.exe <files>
```
If no files specified the interpreter will run in interactive mode.

### Sources compilation

Alternatively, you can compile the sources for your machine.

```shell
go build https://github.com/gragov4ik/yacal/cmd/yacal
```
or
```shell
go build cmd/yacal
```

## Language description

Yacal is a simple toy functional language.

Currently, there are implemented a small subset of very crucial built-in functions and constructs, such as 
functions/lambdas, simple arithmetic operations, predicates, logic, list constructors/destructors and some other. 

It follows the syntax rules of the Lisp programming language.

Many examples you can find in [examples](https://github.com/bragov4ik/yacal/tree/master/examples) folder.

For instance, [dot_product.yacal](https://github.com/bragov4ik/yacal/tree/master/examples/dot_product.yacal)
```
(print "Enter n:")
(set n (toint (input)))

(func enter_n (n) (
    cond (> n 0)
    (cons (toint (input)) (enter_n (- n 1)))
    '()
))

(print "Enter v1:")
(set v1 (enter_n n))
(print)
(print "Enter v2:")
(set v2 (enter_n n))
(print)

(func dot (v1 v2) (
    cond (isnull v1)
    '()
    (cons (* (head v1) (head v2)) (dot (tail v1) (tail v2)))
))

(print (dot v1 v2))
```

## Authors

[Ivan Rybin](https://github.com/i1i1), [Kirill Ivanov](https://github.com/bragov4ik), [Lev Lymarenko](https://github.com/sevenzing), [Ilya Kolomin](https://github.com/Ilya-Kolomin) 

## License
GNU General Public License v3.0, see [LICENSE](https://github.com/bragov4ik/yacal/tree/master/LICENSE)