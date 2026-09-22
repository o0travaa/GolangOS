# GolangOS
## Hi everyone!

This is my first big project. I'm 14 yo and my dream is to become a good programmer. I hope you enjoy it ^^

This is a beta test of my upcoming large-scale project of the same name.

## First step: [Install Go](https://go.dev/doc/install)

## Second step: Install
### 1) Install to the system


```bash
git clone https://github.com/o0travaa/GolangOS.git
cd GolangOS
go install .
```

or

```bash
go install github.com/o0travaa/GolangOS@latest
```

### 2) Add to PATH (choose your shell)

**For Bash** (Default in Ubuntu/Debian):
```bash
echo 'export PATH=PATH:HOME/go/bin' >> ~/.bashrc && source ~/.bashrc
```

**For Zsh** (Default in macOS/Arch with OhMyZsh):
```bash
echo 'export PATH=PATH:HOME/go/bin' >> ~/.zshrc && source ~/.zshrc
```

**For Fish**:
```fish
fish_add_path $HOME/go/bin
```

## *Optional: Run (if you just want to run)

```bash
git clone https://github.com/o0travaa/GolangOS.git
cd GolangOS
go run .