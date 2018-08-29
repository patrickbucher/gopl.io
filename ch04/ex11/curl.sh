#!/bin/sh

#curl https://api.github.com/?access_token=`cat .github_token`
curl -H "Authorization: token `cat .github_token`" https://api.github.com
