import os

with open("config_paths.txt") as f:
    lines = f.readlines()
#    print(lines)

for line in lines:
    path = line.strip()
    # verify that is valid path and file exits
    filename = ".bashrc"

    command = f"cp {path} ./files/"
    print(command)
    os.system(command)
