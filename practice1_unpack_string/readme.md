## Practical Task #1 "String Unpacking"

You need to write a Go function that performs a primitive unpacking of a string containing repeated characters/runes, 
for example:
* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "3abc" => "" (некорректная строка)
* "45" => "" (некорректная строка)
* "aaa10b" => "" (некорректная строка)
* "aaa0b" => "aab"
* "" => ""
* "d\n5abc" => "d\n\n\n\n\nabc"

As seen in the examples, digits are allowed, but numbers are not.

If an invalid string is provided, the function should return an error.
You may create additional functions or errors if needed.

**(*) Bonus Task: Support escaping via `\`:**

**(note the backticks)**
* \`qwe\4\5\` => "qwe45"
* \`qwe\45\` => "qwe44444"
* \`qwe\\\5\` => \`qwe\\\\\\\\\\`
* \`qw\ne\`  => "" (некорректная строка)

As seen in the examples, only digits or a backslash can be escaped.

### Hints
- https://golang.org/ref/spec#String_literals
- `strings.Builder`
- `strings.Repeat`
- `strconv.Atoi`
