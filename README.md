# TOGO, the To Go Compiler

Using the "To" language, because it's an even worse name name "Go" for a language.

Additional Features:
1. The error operator, '!', which looks like: `x := doThing() !(result, ok) { fmt.Println(err) return err }` where the return is actually (result, ok, error)
2. The ternary operator, `?` which is just macro expansion for `if else` assignments
3. Comptime keyword?
4. Actual enums, with safety?