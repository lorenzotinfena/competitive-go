# competitive-go

# Usage
# Limitations
in all the code (main package + non standard libraries imported):
- all the exported and unexported names must be different
- in the same file: literal strings or characters cannot contain aliases of non standard libraries
- it should be formatted with gofmt
- You cannot import packages with a local path
- it shouldn't import standard libraries with the dot '.'