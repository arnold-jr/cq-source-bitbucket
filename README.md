# CloudQuery bitbucket Source Plugin

[![test](https://github.com/arnold-jr/cq-source-bitbucket/actions/workflows/test.yaml/badge.svg)](https://github.com/arnold-jr/cq-source-bitbucket/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/arnold-jr/cq-source-bitbucket)](https://goreportcard.com/report/github.com/arnold-jr/cq-source-bitbucket)

A bitbucket source plugin for CloudQuery that loads data from bitbucket to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Links

 - [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
 - [Supported Tables](docs/tables/README.md)


## Configuration

The following source configuration file will sync to a PostgreSQL database. See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

```yaml
kind: source
spec:
  name: "bitbucket"
  path: "arnold-jr/bitbucket"
  version: "${VERSION}"
  destinations:
    - "postgresql"
  spec:
    # plugin spec section
```

## Development

### Run tests

```bash
make test
```

### Run linter

```bash
make lint
```

### Generate docs

```bash
make gen-docs
```

### Release a new version

1. Run `git tag v1.0.0` to create a new tag for the release (replace `v1.0.0` with the new version number)
2. Run `git push origin v1.0.0` to push the tag to GitHub  

Once the tag is pushed, a new GitHub Actions workflow will be triggered to build the release binaries and create the new release on GitHub.
To customize the release notes, see the Go releaser [changelog configuration docs](https://goreleaser.com/customization/changelog/#changelog).
