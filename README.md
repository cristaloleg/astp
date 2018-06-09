## ASTP

### Go AST predicates

```go
func Walk(var node ast.Node) {
    switch {
    case astp.IsValueSpec(node):
        value := astp.AsValueSpec()

        for _, name := range value.Names {
            if name != nil {
                println("value name: " + name.Name)
            }
        }

    case astp.IsFuncDecl(node):
        fn := astp.AsFuncDecl()
        if fn.Name != nil {
            println("function: " + fn.Name.Name)
        }
    }
}
```