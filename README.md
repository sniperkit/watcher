# Watcher
> Awesome iOS scrapper

Get all data from [awesome-ios repo](http://github.com/vsouza/awesome-ios), extracting all links and github repos, collecting all relevant data on Github API (owner , package managers, issues, pull request and more) , indexing on Elasticsearch.

This is part of project of [awesome-ios website](http://awesomeios.com) (new version).

[![Go Version][go-image]][go-url]
[![License][license-image]][license-url]

![](header.png)

## Installation

`make install`

## Configuration

At the first, you need to generate a Github Token, see this [instructions](https://github.com/blog/1509-personal-api-tokens).

For the development you can use a local Elasticsearch, to install:

__OSX__

`brew install elasticsearch`

__Linux__

`apt-get install elasticsearch`



After install you have to change configuration file, with your credential and Elasticsearch host.
Inside of config folder, `example.yaml` file contains a sample configuration to setup your watcher.

*Watcher runs with this commando `./watcher -e {{ENVIROMENT}}` it's better you create a file called `development` with you configuration data, then run `./watcher -e development`.

```yaml

app:
  name: app_example
  version: 0.0.1
db:
  host: http://127.0.0.1:9200
github:
  repo:
    url: https://api.github.com/repos/vsouza/awesome-ios/git/trees/HEAD
    readme: README.md
  auth:
    app: app_example
    token: ----------

```


## Release History

* 1.0.0
  * Repo, Links, Package Managers and Categories data.


## Contribute

We would love for you to contribute to **Watcher**, check the ``LICENSE`` file for more info.

## Meta

Vinicius Souza – [@iamvsouza](https://twitter.com/iamvsouza) – hi@vsouza.com

Distributed under the MIT license. See ``LICENSE`` for more information.

[https://github.com/vsouza](https://github.com/vsouza/)

[go-image]:https://img.shields.io/badge/Go--version-1.6-blue.svg
[go-url]: https://golang.org/
[license-image]: https://img.shields.io/badge/License-MIT-blue.svg
[license-url]: LICENSE
