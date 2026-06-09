export LANG=en_GB.UTF-8
export LC_ALL=en_GB.UTF-8

export EDITOR='code --new-window --wait'
export VISUAL=$EDITOR

export XDG_CONFIG_HOME="$HOME/.config"
export CODE="$HOME/code"
export PATH="$HOME/.local/bin:$PATH"

# Upscope
export UPSCOPE_DIR="$CODE/upscope"
export PATH="$PATH:$UPSCOPE_DIR/bin"

# Go
export GOPATH="$HOME/.local/share/go"
export GOBIN="$HOME/.local/bin"

# starship
export STARSHIP_CONFIG="$XDG_CONFIG_HOME/starship/starship.toml"

HISTFILE="$HOME/.zsh_history"
HISTSIZE=20000
SAVEHIST=20000
setopt INC_APPEND_HISTORY
setopt HIST_IGNORE_DUPS
setopt HIST_REDUCE_BLANKS
setopt SHARE_HISTORY

alias path='echo $PATH | tr ":" "\n"'
alias reload!='source ~/.zshrc'

take() {
  mkdir -p "$1" && cd "$1"
}

eval "$(fnm env --use-on-cd --shell zsh)"
eval "$(starship init zsh)"

source /opt/homebrew/share/zsh-autosuggestions/zsh-autosuggestions.zsh
bindkey '^I' autosuggest-accept   # Tab accepts the grey suggestion
bindkey '^[[Z' undo               # Shift+Tab undoes the accept (zsh undo)

# must be sourced last — it hooks the line editor
source /opt/homebrew/share/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh
