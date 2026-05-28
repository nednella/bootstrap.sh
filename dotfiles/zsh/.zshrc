# ==============================================================================
# .zshrc
# ==============================================================================

# ----- variables --------------------------------------------------------------

export LANG=en_GB.UTF-8
export LC_ALL=en_GB.UTF-8
export EDITOR='code --new-window --wait'
export VISUAL=$EDITOR
export XDG_CONFIG_HOME="$HOME/.config"

# Upscope
export UPSCOPE_DIR="$HOME/code/upscope"
export PATH="$PATH:$UPSCOPE_DIR/bin"

# ----- history ----------------------------------------------------------------

HISTFILE="$HOME/.zsh_history"
HISTSIZE=20000
SAVEHIST=20000
setopt INC_APPEND_HISTORY
setopt HIST_IGNORE_DUPS
setopt HIST_REDUCE_BLANKS
setopt SHARE_HISTORY

# ----- autocomplete -----------------------------------------------------------

autoload -Uz compinit && compinit

# ----- aliases ----------------------------------------------------------------

alias path='echo $PATH | tr ":" "\n"'
alias reload!='source ~/.zshrc'

# ----- functions --------------------------------------------------------------

take() {
  mkdir -p "$1" && cd "$1"
}
