#!/bin/zsh
#compile yarvis
sudo go build -o /bin/yarvis .

# add aliases
ZSH_RC=~/.zshrc
BASH_RC=~/.bashrc
if [ ! -f "$ZSH_RC" ]; then
  touch .zshrc
fi
if [ ! -f "$BASH_RC" ]; then
  touch .bashrc
fi

# if no new line at end of rc file, this will break
if ! grep -Fq "alias yarvis" $ZSH_RC; then
  echo "Alias not found, adding to $ZSH_RC...\n"
  echo 'alias yarvis="/bin/yarvis"' >> $ZSH_RC
fi

if ! grep -Fq "alias yarvis" $BASH_RC; then
  echo "Alias not found, adding to $BASH_RC...\n"
  echo 'alias yarvis="/bin/yarvis"' >> $BASH_RC
fi

echo "ACTION REQUIRED: \
\nFor these changes to take effect, run: \
\n\t'source <.*rc_file>'"