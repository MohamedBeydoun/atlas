MAJOR=0
MINOR=0
PATCH=0

# Check for jq
if ! jq_loc="$(type -p "jq")" || [[ -z $jq_loc ]]; then
	echo "Install jq before running"
	exit 0
fi

# Parse commits
for message in $(curl -s https://api.github.com/repos/MohamedBeydoun/atlas/commits | jq -c '.[].commit.message' | cut -d '"' -f 2 | cut -d '/' -f 1 | tac); do
	if [ $message == "Bugfixes" ]; then
		PATCH=$(($PATCH+1))
	elif [ $message == "Features" ]; then
		PATCH=0
		MINOR=$(($MINOR+1))
	fi
done

# Output VERSION file
echo $MAJOR.$MINOR.$PATCH > VERSION