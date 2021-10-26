#!/bin/zsh

# vars
ZSH_RC=~/.zshrc
BASH_RC=~/.bashrc
SCRIPTS_URL=https://github.com/RWEngelbrecht/scripts/archive/refs/heads/usable.zip
SCRIPTS_ZIP=scripts.zip
SCRIPTS_DIR=$HOME/scripts

#compile yarvis
echo "Building Yarvis..."
sudo go build -o /bin/yarvis .

# check if unzip and wget installed
# note: `2>/dev/null` quiets a warning. 
#   Redirect (>) 2nd output channel (2 i.e. stderr) to /dev/null
if ! apt list 2>/dev/null wget | grep -Fq "[installed]"; then
  echo "wget not installed, installing now..."
  sudo apt-get install wget
fi
if ! apt list 2>/dev/null unzip | grep -Fq "[installed]"; then
  echo "unzip not installed, installing now..."
  sudo apt-get install unzip
fi

# download scripts repo
if [ ! -d $SCRIPTS_DIR ]; then
  echo "Downloading scripts from $SCRIPTS_URL..."
  wget -q $SCRIPTS_URL -O $SCRIPTS_ZIP
  echo "Downloaded..."

  echo "Unzipping $SCRIPTS_ZIP..."
  unzip -q $SCRIPTS_ZIP -d $SCRIPTS_DIR
  rm $SCRIPTS_ZIP
  echo "Unzipped and cleaned up..."
fi

# add aliases
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