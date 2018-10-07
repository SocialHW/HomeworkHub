#!/bin/bash

is_installed() {
	
	if [ -z "$1" ]; then

		search="$(apt list --installed | grep -i $1)"

		if [ -z "${search}" ]; then
			echo "$1 is not installed"
			return 0
		else
			echo "$1 is installed"
			return 1
		fi
	else 
		echo "Please enter a parameter!"
	fi
}

my_sql_installed=0

# Install mysql-server
get_mysql() {
	if $(is_installed vim); then
		echo "bleh"
	else
		echo "bloh"
	fi
}

get_mysql
#echo "$(is_installed)"
