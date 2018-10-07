#!/bin/sh

systemctl stop mysql

sudo apt remove mysql-server

sudo apt autoremove
