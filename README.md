# go-mvn
> Auto Maven Dependency Updater written in go

# Requirements
- maven project
- mvn cli available
- docker cli available

# FeatureList (planned)
- [X] Locate pom files in given directory
- [X] Locate sub / child poms for located parent pom in given directory
- [X] Parse Metadata
- [ ] Parse Properties
  - maybe done by mvnparser
- [X] Parse Dependencies
- [ ] Parse Plugins
  - maybe done by mvnparser
- [ ] Build Map of (Dependencies / Plugins) and Versions
- [ ] make use of mvn versions plugin to receive newest versions (ideal update versions) for testing
(write into second map to keep track of current versions and ideal update versions)
- [ ] Spawn Docker container with given image
- [ ] Evaluate maven build in docker container with current versions
- [ ] modify versions and start new docker container
- [ ] create a cli program out of it for easy use
- tbc

## Execute

```shell
go run .
```