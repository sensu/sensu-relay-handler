# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic
Versioning](http://semver.org/spec/v2.0.0.html).

## Unreleased

### Changed
- Use go 1.13.x for builds
- Slim down Bonsai and Goreleaser Builds to supported platforms
- Switch to GitHub Actions
- Use Community Plugins library and enforce Enterprise licensing
- Cleanup options processing
- Replace HTTPWrapper with traditional Go http package use

## [0.0.9] - 2019-10-23

### Changed
- Docs image update
- Added Bonsai Badge

## [0.0.8] - 2019-08-13

### Changed
- README & diagram changes
- Use go modules instead of dep for dependency management
- Use go 1.12.x for builds
- Switch to open-source plugins library (sensu-plugins-go-library)

## [0.0.7] - 2019-03-22

### Fixed
- Fixed build tars, added goreleaser config

## [0.0.6] - 2019-03-21

### Changed
- Readme YAML example Handler definition

## [0.0.5] - 2019-03-21

### Fixed
- Removed timeout flag from usage example (not yet implemented)

## [0.0.4] - 2019-03-07

### Added
- Options to configure Event Check and Metrics handling for relayed Events

## [0.0.3] - 2019-03-06

### Fixed
- Updated Travis GitHub token for private repos

## [0.0.2] - 2019-03-06

### Fixed
- Agent Events API response code is either 201 or 202

## [0.0.1] - 2019-03-06

### Added
- Initial release
