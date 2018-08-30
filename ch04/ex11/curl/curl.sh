#!/bin/sh

#curl https://api.github.com/?access_token=`cat .github_token`
#curl -H "Authorization: token `cat ../.github_token`" https://api.github.com

# create a new issue
#curl -H "Authorization: token `cat ../.github_token`" -X POST https://api.github.com/repos/patrickbucher/gopl.io/issues -d @new_issue.json

# modify an existing issue
#curl -H "Authorization: token `cat ../.github_token`" -X POST https://api.github.com/repos/patrickbucher/gopl.io/issues/2 -d @edit_issue.json

# lock an issue (deleting is not possible)
#curl -H "Authorization: token `cat ../.github_token`" -X PUT https://api.github.com/repos/patrickbucher/gopl.io/issues/2/lock -d @lock_issue.json

# unlock an issue (deleting a lock)
#curl -H "Authorization: token `cat ../.github_token`" -X DELETE https://api.github.com/repos/patrickbucher/gopl.io/issues/2/lock

# list issues
#curl -H "Authorization: token `cat ../.github_token`" https://api.github.com/repos/patrickbucher/gopl.io/issues

# get an issue
curl -H "Authorization: token `cat ../.github_token`" https://api.github.com/repos/patrickbucher/gopl.io/issues/2
