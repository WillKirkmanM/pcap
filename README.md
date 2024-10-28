# Pascal Compiler for the Alpha Processor

> [!NOTE]
> This is a Compiler for the Pascal Programming Language which is read by the executed by the interpreter.

![Pascal Code](https://raw.githubusercontent.com/ALANVF/vscode-pascal-magic/master/assets/example.png)

Here is how it works:

`ast.go`, The [Abstract Syntax Tree](https://en.wikipedia.org/wiki/Abstract_syntax_tree), is the data structure that the parser generates from the source code. It is a tree of nodes that represent the structure of the program.

It includes the interfaces **Node**, **Statement**, and **Expression**, which are required methods for returning token literals and string representations. Think of these as different types of LEGO bricks. The Program struct holds a list of statements and implements methods to return the token literal of the first statement and a string representation of all statements. The Identifier, IntegerLiteral, and StringLiteral structs represent expressions with corresponding token and value fields, implementing the required methods. These are special LEGO bricks that represent names, numbers, and words. Various statement types each with methods to return their token literal and string representation. These statements are responsible for the different operations like variable assignment, printing, reading input, and grouping multiple statements, similar to specific instructions in a LEGO booklet, completing the structure of the program.

![AST](https://upload.wikimedia.org/wikipedia/commons/c/c7/Abstract_syntax_tree_for_Euclidean_algorithm.svg)

`compiler.go`, The Compiler that compiles the Pascal source code into an interpretable format.

The compiler is responsible for converting the abstract syntax tree into a string of instructions, called the generated code, which will then be interpreted and executed.

The main function in this file is `generateCode`, which takes a `Program` (a collection of statements) from the `parser` and `ast` and generates the corresponding code for each statement type. It iterates over all the statements in the program and calls specific functions to generate code for each type of statement.

For example, the `generateLetStatement` function generates code for `let` statements, which assign values to variables. If the statement is `let x = 'hello'`, this function will generate the code `MOV x, 'hello'`.

Similarly, the `generateWritelnStatement` function generates code for `writeln` statements, which print values to the console. For the statement `writeln 'hello'`, this function will generate the code `WRITELN 'hello'`.


`interpreter.go`, The Interpreter that reads the compiled code and executes it.

The interpreter starts by setting up a place to store variables. It then reads the code, splits it into individual lines, and processes each line by splitting it into parts to understand the command. For example, the `MOV` command stores a value in a variable, the `WRITELN` command prints the value of a variable, and the `READLN` command reads input from the user and stores it in a variable. The interpreter handles these commands and performs their respective actions.


`lexer.go`, The [Lexer](https://en.wikipedia.org/wiki/Lexical_analysis) that reads the source code and generates tokens.

The lexer reads the source code character by character and groups them into tokens, which are the smallest units of meaning in the source code. Tokens include keywords, identifiers, operators, and literals like numbers and strings.

`parser.go`, The [Parser](https://en.wikipedia.org/wiki/Parsing) Implements the parser that processes the sequence of tokens produced by the lexer and constructs the AST. The parser ensures the syntactic correctness of the source code.

The parser reads the tokens generated by the lexer and constructs the abstract syntax tree (AST) that represents the structure of the program. It does this by defining grammar rules for the language and using them to parse the tokens.

`token.go`, Defines the token types and structures used by the lexer and parser. Tokens represent the smallest units of meaning in the source code, such as keywords, identifiers, and operators.

`example.pas` - The Pascal Source Code that is compiled and then interpreted.

```pascal
program Hello;
begin
    let output = "Hello, world.";
    writeln(output);
    readln(output);
    writeln(output);
end
```