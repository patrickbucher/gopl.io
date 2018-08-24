#!/bin/sh

# cleans up the exercise directory/file/package structure by performing the following steps:
# 1) Rename folders like ch02/ex02_03 to ch02/ex03
# 2) Rename files like ch02/ex02_03/ex02_03.go to ch02/ex03/ex03.go
# 3) Rename package declarations like ex02_03 to ex03

for e in ch*/ex*; do

    # 1) folders
    old_folder="$e"
    old_package="`echo "$old_folder" | xargs basename | sed -E 's/^ex([[:digit:]]{2})_([[:digit:]]{2})$/ex\1_\2/'`"
    new_folder="`echo "$old_folder" | sed -E 's/^ch([[:digit:]]{2})\/ex[[:digit:]]{2}_([[:digit:]]{2})$/ch\1\/ex\2/'`"
    new_package="`basename "$new_folder"`"
    git mv "$old_folder" "$new_folder"

    # 2) files
    for f in `find "$new_folder" -type f | grep -E 'ex[[:digit:]]{2}_([[:digit:]]{2}).go$'`; do
        old_file="$f"
        new_file="`echo "$f" | sed -E 's/ex[[:digit:]]{2}_([[:digit:]]{2}).go$/ex\1.go/'`"
        git mv "$old_file" "$new_file"
    done

    # 3) packages
    for f in `find "$new_folder" -type f | grep '\.go$' | xargs grep -l "^package ${old_package}$"`; do
        sed -i '' "s/^package ${old_package}$/package ${new_package}/" "$f"
        git add "$f"
    done

done
