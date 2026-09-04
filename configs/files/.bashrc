echo "loaded .bashrc"
h(){ cd ~ ; }
home(){ cd ~ ; }
doc(){ cd ~/storage/shared/Documents/ ; }
dev(){ cd ~/dev/termux-android-scripts ; }
vib(){ vim ~/.bashrc ; }
ob(){ vim ~/.bashrc ; }
sb(){
source ~/.bashrc
# cat ~/.bashrc
bat ~/.bashrc
}
ll(){ ls -a ; }
sync(){
python3 ~/dev/termux-android-scripts/configs/sync_configs.py
}
