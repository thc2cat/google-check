# google-check

`google-check` cli purpose was to automate google web search for sign of hacking on our domain.

```shell
$ go run main.go -h
Usage of main.exe:
  -i string
        ignore regexp
  -min int
        minimal web hosts links limit for reporting (default 19)
  -p string
        google search pattern. (default "site:uvsq.fr viagra")
  -v    verbose mode
```

If there is more than min links for your search, `google-check` will emit a warning line.
Verbose mode allow you to see exacts links.