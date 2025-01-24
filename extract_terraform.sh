#!/bin/bash

# Ensure GNU awk and GNU sed are available
if ! command -v gawk &> /dev/null || ! command -v gsed &> /dev/null; then
    echo "This script requires gawk and gsed. Please install them using 'brew install gawk gnu-sed'."
    exit 1
fi

# Check if at least one argument is provided
if [ "$#" -lt 1 ]; then
    echo "Usage: $0 file1.md [file2.md ...]"
    exit 1
fi

# Loop through each provided Markdown file
for file in "$@"; do
    if [ -f "$file" ]; then
        # Extract text between ```terraform and the next ```
        gawk '/```terraform/{flag=1; next} /```/{flag=0} flag' "$file"
    else
        echo "File not found: $file"
    fi
done
