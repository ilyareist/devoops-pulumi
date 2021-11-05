#!/bin/bash

#Installing nginx
sudo amazon-linux-extras enable nginx1
sudo yum clean metadata
sudo yum -y install nginx

# Changing a message
echo "Hello, DevOps ! " | sudo tee /usr/share/nginx/html/index.html
sudo systemctl start nginx
