pkg update && pkg upgrade -y
pkg install git 
# pkg install 

# cat path/to/id_rsa.pub >> ~/.ssh/authorized_keys
mkdir ~/.ssh 
touch ~/.ssh/authorized_keys
chmod 600 ~/.ssh/authorized_keys
chmod 700 ~/.ssh



# ifconfig 
# ip you can get from either router/local network scan or just type ifconfig
# host api server to accept one new ssh key provied they type in the correct password. also implement rate limiting with 3 tries and then exponential backoff with max 1 hour wait time.  
# whoami

