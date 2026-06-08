export LANG=en_GB.UTF-8
export LC_ALL=en_GB.UTF-8

export EDITOR='code --new-window --wait'
export VISUAL=$EDITOR

export XDG_CONFIG_HOME="$HOME/.config"
export PATH="$HOME/.local/bin:$PATH"

export STARSHIP_CONFIG="$XDG_CONFIG_HOME/starship/starship.toml"

export UPSCOPE_DIR="$HOME/dev/upscope"
export PATH="$PATH:$UPSCOPE_DIR/bin"

HISTFILE="$HOME/.zsh_history"
HISTSIZE=20000
SAVEHIST=20000
setopt INC_APPEND_HISTORY
setopt HIST_IGNORE_DUPS
setopt HIST_REDUCE_BLANKS
setopt SHARE_HISTORY

autoload -Uz compinit && compinit

alias path='echo $PATH | tr ":" "\n"'
alias reload!='source ~/.zshrc'

take() {
  mkdir -p "$1" && cd "$1"
}

eval "$(fnm env --use-on-cd --shell zsh)"
eval "$(starship init zsh)"
