# Get the latest tag
latestTag=$(git describe --tags --abbrev=0)

# Print the tag
echo "Latest tag: $latestTag"
echo -e "\033[33m\nCurreny tag: $latestTag\033[0m"


newTag=$(cat version.txt)
echo -e "\033[32m\nNew tag: $newTag\033[0m"

# Check if the new tag is the same as the existing tag
if [ "$newTag" = "$latestTag" ]; then
    echo -e "\033[31mError: New tag is the same as the existing tag\033[0m"
    exit 1
fi

git tag $newTag

git push origin $newTag
