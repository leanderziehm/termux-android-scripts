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
	cd ~/dev/termux-android-scripts/configs
	python3 sync_configs.py #make
	git add . && git commit -m "sync" && git push
}

