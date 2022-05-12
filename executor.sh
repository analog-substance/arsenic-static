#!/bin/bash

# This is so we can execute the script with arguments
# without modifying the script to use SCRIPT_ARGS
execute_script() {
	eval "$SCRIPT_CONTENT"
}

BASH_ARGV0="$SCRIPT_NAME"
if [[ -z "$SCRIPT_STDIN" ]]; then
	execute_script $SCRIPT_ARGS
else
	echo "$SCRIPT_STDIN" | execute_script $SCRIPT_ARGS
fi
