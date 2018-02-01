# Asks the user a yes or no question, returning the boolean result
ask_yes_no() {
  local CONF="Y/n"
  local reply="Y"
  read -p "$1 ${CONF} " -n 1 -r reply
  if [[ $reply =~ ^[Yy]$ ]]; then
    echo 'yes'
  elif [[ $reply =~ ^[Nn]$ ]]; then
    return
  else
    echo 'yes'
  fi
}

