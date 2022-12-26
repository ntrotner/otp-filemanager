#!/bin/sh

# get up-to-date
git fetch
git checkout main
git pull

# create release branch
git checkout -b "release/$1"
go build -o "otp-filemanager-$1"

cp .env-example .env
tar -cvf "otp-filemanager-$1.tar" .env "otp-filemanager-$1"

git add "otp-filemanager-$1.tar"
git commit -m "Release $1"

git push --set-upstream origin release/$1
