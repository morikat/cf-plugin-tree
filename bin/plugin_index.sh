#!/bin/bash

#export COMPANY=""
export AUTHOR=${AUTHOR:-"Takeshi Morikawa"}
export EMAIL=${EMAIL:-"morika-t@outlook.com"}
export GH_AUTHOR=${GH_AUTHOR:-morikat}
export HOMEPAGE=${HOMEPAGE:-"http://github.com/$GH_AUTHOR"}
#export GH_ORG=${GH_ORG:-cloudfoundry-community}
export GH_ORG=${GH_ORG:-morikat}
export GH_REPO=${GH_REPO:-cf-plugin-tree}
export NAME=${NAME:-"tree"}
export DESCRIPTION=${DESCRIPTION:-"output application instance tree for an application."}
export PKG_DIR=${PKG_DIR:=out}
export PROJECT_CREATED="2015-03-28"

VERSION=$(<VERSION)

if [[ "$(which shasum)X" == "X" ]]; then
  echo "shasum not installed"
  exit 1
fi

cat << EOS
- name: $NAME
  description: $DESCRIPTION
  version: $VERSION
  created: $PROJECT_CREATED
  updated: $(date +%F)
  company: "$COMPANY"
  authors:
  - name: "$AUTHOR"
    homepage: $HOMEPAGE
    contact: $EMAIL
  homepage: http://github.com/$GH_ORG/$GH_REPO
  binaries:
  - platform: win64
    url: "https://github.com/$GH_ORG/$GH_REPO/releases/download/v$VERSION/${GH_REPO}_windows_amd64.exe"
    checksum: "$(shasum out/${GH_REPO}_windows_amd64.exe | awk '{print $1}')"
  - platform: win32
    url: "https://github.com/$GH_ORG/$GH_REPO/releases/download/v$VERSION/${GH_REPO}_windows_386.exe"
    checksum: "$(shasum out/${GH_REPO}_windows_386.exe | awk '{print $1}')"
  - platform: linux64
    url: "https://github.com/$GH_ORG/$GH_REPO/releases/download/v$VERSION/${GH_REPO}_linux_amd64"
    checksum: "$(shasum out/${GH_REPO}_linux_amd64 | awk '{print $1}')"
  - platform: linux32
    url: "https://github.com/$GH_ORG/$GH_REPO/releases/download/v$VERSION/${GH_REPO}_linux_386"
    checksum: "$(shasum out/${GH_REPO}_linux_386 | awk '{print $1}')"
  - platform: osx
    url: "https://github.com/$GH_ORG/$GH_REPO/releases/download/v$VERSION/${GH_REPO}_darwin_amd64"
    checksum: "$(shasum out/${GH_REPO}_darwin_amd64 | awk '{print $1}')"
EOS
