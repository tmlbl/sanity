#!/bin/bash
#
# The Sanity setup script

# Initial package installation
apt-get update
echo "Installing packages..."
apt-get install -q -y \
	tmux \
	sudo \
	rsync \
	git \
	vim \
	chromium \
	redshift \
	nitrogen \
	htop \
	build-essential \
	terminator \
	wicd \
	pulseaudio \
	xorg \
	openbox \
	lightdm \
	lightdm-gtk-greeter \
	tint2

# Fetch and place config files
fetch_config() {
	echo -e "Fetching $1..."
	wget -q -O /$1 https://raw.githubusercontent.com/tmlbl/sanity/master/$1
}

fetch_config etc/lightdm/lightdm-gtk-greeter.conf
fetch_config etc/xdg/openbox/autostart
fetch_config etc/xdg/openbox/rc.xml
fetch_config etc/xdg/tint2/tint2rc

