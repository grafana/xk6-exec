#!/bin/bash

# Output to stdout
echo "This is the normal output."

# Output to stderr
echo "This is the error." >&2

# Exit with a non-zero status
exit 12