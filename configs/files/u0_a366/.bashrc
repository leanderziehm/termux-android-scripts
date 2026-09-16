echo "loaded .bashrc termux"
h(){ cd ~ ; }
ob(){ vim ~/.bashrc ; }
sb(){ source ~/.bashrc && cat ~/.bashrc ; }
vimrc(){ vim ~/.vimrc ; } 
doc(){ cd ~/storage/shared/Documents/ ; }
t(){ cd ~/dev/termux-android-scripts ; }
ll(){ ls -a ; }
sync(){
	cd ~/termux-android-scripts/configs
	python3 sync_configs.py #make
	git add . && git commit -m "sync" && git push
}
blog(){ cd ~/storage/shared/Documents/git-repos/blogs/ ; }
alias b=blog
alias lg=lazygit
gigi(){
printf "git commit message: " && read -r commitmessage && git add . && git commit -m "$commitmessage" && git push
}
vps(){ cd ~/devops/vps ; } 
todo(){ vim ~/todo.md ; }
copyparty(){ uv tool run copyparty -a user:123 ;}
