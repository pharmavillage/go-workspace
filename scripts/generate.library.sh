#!/bin/bash

echo "Enter the name for the library: "
read libraryName

gen_lib_command="nx g @nx-go/nx-go:library data-access --name=$libraryName --directory=modules --projectNameAndRootFormat=derived --tags=database,data --skipFormat=false"

echo "Executing command: $gen_lib_command"

# Execute the command
$gen_lib_command
