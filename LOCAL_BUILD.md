# Local Build

Use this set of commands to perform a local build for tesing.

```bash
CURRENT_VER=$(git tag --list | sort --version-sort | tail -n 1)
VER_SELECTIONS="${CURRENT_VER}\nv$(semver bump patch ${CURRENT_VER})\nv$(semver bump minor ${CURRENT_VER})\nv$(semver bump major ${CURRENT_VER})"
SEMVER=$(echo "${VER_SELECTIONS}" | fzf)
echo ${SEMVER}
BUILD_DATE=$(gdate --utc +%FT%T.%3NZ); echo ${BUILD_DATE}
GIT_COMMIT=$(git rev-parse HEAD); echo ${GIT_COMMIT}

if [[ ! -d cmd/changelog ]]; then
  mkdir cmd/changelog
fi
cp changelog/*.md cmd/changelog/

MODNAME=ktrouble
go build -ldflags "-X ${MODNAME}/cmd.semVer=${SEMVER} -X ${MODNAME}/cmd.buildDate=${BUILD_DATE} -X ${MODNAME}/cmd.gitCommit=${GIT_COMMIT} -X ${MODNAME}/cmd.gitRef=/refs/tags/${SEMVER}" && \
./${MODNAME} version | jq .
if [[ -d ~/tbin ]]; then
  cp ./${MODNAME} ~/tbin
fi

./${MODNAME} genhelp --format markdown > HELP.md && \
# remove the double empty line at the end of the file
sed -i '$ d' HELP.md
```
