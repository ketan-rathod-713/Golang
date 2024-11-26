#!/bin/bash
# Fetch all branches
git fetch --all

# Get a list of all local branches
branches=$(git branch | sed 's/^\*//')

# Loop through each branch and push it if not already on the remote
for branch in $branches; do
    branch=$(echo $branch | xargs) # Trim leading/trailing whitespace
    if ! git show-ref --verify --quiet refs/remotes/origin/$branch; then
        echo "Publishing branch $branch to remote"
        git push origin $branch
    else
        echo "Branch $branch is already on the remote"
    fi
done
