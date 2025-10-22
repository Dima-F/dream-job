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

## Air

### Install

```sh
go install github.com/air-verse/air@latest
```

### Using

```sh
air
air -c .air.debug.toml
```

## Процес дебага

> air -c .air.debug.toml

основний рядок тут:
...
full_bin = "dlv exec ./tmp/main --listen=127.0.0.1:2344 --headless=true --api-version=2 --accept-multiclient --continue --log -- "
...

#в .vscode/launch.json знаходиться конфіг для запуску дебагера

## Ієрархія UI 

В компонентах не можна використовувати віджети і інші компоненти

Віджети повині використовувати компоненти (не інші віджети)

Шаблони можуть використовувати як віджети, так і компоненти

# PG

```sh
go get github.com/jackc/pgx/v5
go get github.com/jackc/pgx/v5/pgxpool
```