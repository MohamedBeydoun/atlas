#!/bin/bash

MAJOR=0
MINOR=0
PATCH=0

# Check for jq
if ! jq_loc="$(type -p "jq")" || [[ -z $jq_loc ]]; then
	echo "Install jq before running"
	exit 0
fi

# Parse commits
for message in $(git log --pretty=format:'{%n "subject": "%s" %n},' | sed "$ s/,$//" | sed ':a;N;$!ba;s/\r\n\([^{]\)/\\n\1/g'| awk 'BEGIN { print("[") } { print($0) } END { print("]") }'| jq -c '.[].subject' | cut -d '"' -f 2 | cut -d '/' -f 1 | tac); do
	if [ $message == "Bugfixes" ] || [ $message == "Patches" ] || [ $message == "Tests" ]; then
		PATCH=$(($PATCH+1))
	elif [ $message == "Features" ] || [ $message == "Enhancements" ] || [ $message == "Improvements" ]; then
		PATCH=0
		MINOR=$(($MINOR+1))
	elif [ $message == "MAJOR" ]; then
		PATCH=0
		MINOR=0
		MAJOR=$(($MAJOR+1))
	fi
done

BRANCH=$(git branch | grep \* | cut -d ' ' -f 2 | sed "$ s/\//-/")

# Output VERSION file
if [ $BRANCH == "master" ]; then
	echo $MAJOR.$MINOR.$PATCH > VERSION
else
	echo $MAJOR.$MINOR.$PATCH-$BRANCH > VERSION
fi