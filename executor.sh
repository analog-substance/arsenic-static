#!/bin/bash

# This is so we can execute the script with arguments
# without modifying the script to use SCRIPT_ARGS
execute_script() {
	eval "$SCRIPT_CONTENT"
}

execute_script $SCRIPT_ARGS
