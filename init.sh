#!/bin/sh
#compile yarvis
sudo go build -o /bin/yarvis .

# add aliases
ZSH_RC=~/.zshrc
BASH_RC=~/.bashrc
if [ ! -f "$ZSH_RC" ]; then
  touch .zshrc
elif [ ! -f "$BASH_RC" ]; then
  touch .bashrc
fi

if ! grep -Fxq "alias yarvis" $ZSH_RC && ! grep -Fxq "alias yarvis" $BASH_RC; then
  if [ -f "$ZSH_RC" ]; then
    echo 'alias yarvis="/bin/yarvis"' >> $ZSH_RC
    
  elif [ -f "$BASH_RC" ]; then
    echo 'alias yarvis="/bin/yarvis"' >> $BASH_RC
    source $BASH_RC
  fi
fi