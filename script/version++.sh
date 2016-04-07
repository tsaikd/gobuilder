#!/bin/bash

gitroot="$(git rev-parse --show-toplevel)"
versionfile="${gitroot}/builder/version.go"

currentversion="$(cat "${versionfile}" | awk '/version\.VERSION/{gsub("\"", "", $3);print $3}')"
nextversion="${currentversion%.*}.$(awk -F '.' '{print $3+1}' <<< "${currentversion}")"

sed -i '' 's/version\.VERSION = .*$/version.VERSION = "'"${nextversion}"'"/' "${versionfile}"

git commit -s -m "${nextversion}" "${versionfile}"
echo "bump version: ${currentversion} -> ${nextversion}"
