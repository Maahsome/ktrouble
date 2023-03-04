# Local Build

Use this set of commands to perform a local build for tesing.

```bash
SEMVER=v0.0.999; echo ${SEMVER}
BUILD_DATE=$(gdate --utc +%FT%T.%3NZ); echo ${BUILD_DATE}
GIT_COMMIT=$(git rev-parse HEAD); echo ${GIT_COMMIT}

MODNAME=ktrouble
go build -ldflags "-X ${MODNAME}/cmd.semVer=${SEMVER} -X ${MODNAME}/cmd.buildDate=${BUILD_DATE} -X ${MODNAME}/cmd.gitCommit=${GIT_COMMIT} -X ${MODNAME}/cmd.gitRef=/refs/tags/${SEMVER}" && \
./ktrouble version | jq .
if [[ -d ~/tbin ]]; then
  cp ./ktrouble ~/tbin
fi
```
