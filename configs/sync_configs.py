import os

with open("config_paths.txt") as f:
    lines = f.readlines()
#    print(lines)

for line in lines:
    path = line.strip()
    # verify that is valid path and file exits
    filename = ".bashrc"
    
    os.system(f"mkdir -p ./files/$(whoami)")
    command = f"cp {path} ./files/$(whoami)"
    print(command)
    os.system(command)
