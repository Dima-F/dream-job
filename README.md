golang version: 1.24.5


install templ globaly:

```sh
go install github.com/a-h/templ/cmd/templ@latest
```

golang and .bashrc:

```sh
export PATH="$PYENV_ROOT/bin:$PATH"
if command -v pyenv 1>/dev/null 2>&1; then
 eval "$(pyenv init -)"
fi

# goenv
export GOENV_ROOT="$HOME/.goenv"
export PATH="$GOENV_ROOT/bin:$PATH"
eval "$(goenv init -)"

# Go
export GOROOT="$GOENV_ROOT/versions/$(goenv version-name)"
export GOPATH="$HOME/workspace/go"

# Add Go tools and binaries to PATH
export PATH="$GOROOT/bin:$GOPATH/bin:$PATH"

```
