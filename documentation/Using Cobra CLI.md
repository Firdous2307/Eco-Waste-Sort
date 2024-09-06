`cmd/ Package`
This package handles all the command-line commands and flags. It uses the Cobra library, a popular package for creating CLI applications in Go.

Since this is a CLI Project for now, I searched through Go frameworks for cmmand-line applications in Go and found out [Cobra](https://github.com/spf13/cobra)

It can be installed by running:

```go
go install github.com/spf13/cobra-cli@latest
```


The root.go file is typically used in projects built with the Cobra library to define the root command, which is the base for all other commands.
[Cobra](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md)