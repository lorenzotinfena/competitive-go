# competitive-go

It currently supports only `github.com/lorenzotinfena/goji` library

# Supported platforms
- Codeforces

# Usage
For each problem:

1. You can create and open a codespace on this repo or if you want to run it locally: open the folder in vscode and using remote container extension, select reopen in container
3. Install all tools suggested by vscode
1. Copy paste an example testcase in `input.txt`
2. Code and debug in solution.go
    P.s. Use io for I/O operations and use "log" to print while debugging
3. Run `Aggregate and copy` command and paste the self-contained code on your platform.

If you want to solve multiple problems simultanously (like in contests) you can just backup your multiple solution.go somewhere.

# Limitations
All the code (main package + non standard libraries imported (testing files are excluded)):
- Should have all the exported and unexported names different
- In the same file: literal strings or characters cannot contain aliases of non standard libraries
- Should be formatted with gofmt
- Shouldn't import packages with a local path
- Shouldn't import standard libraries with the dot '.'