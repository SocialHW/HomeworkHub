#!/bin/sh


is_installed() {
	search="$(apt-cache pkgnames | grep -i $1)"

	if [ -z "${search}" ]; then
		return 1
	fi

	return 0
}


# Install mysql-server
get_mysql() {
	p="mysql-server"
	
	if [ ! $(is_installed $p) ]; then
		# is not installed
		echo "$p is not installed"
		
		echo "Installing... "
		sudo apt update
		sudo apt install mysql-server -y
		
	fi
	
	echo "Installed: "
	apt-cache pkgnames | grep -i $p
}

get_mysql
