# Watcher
> Awesome iOS scrapper

[![Go Version][swift-image]][swift-url]
[![codebeat badge](https://codebeat.co/badges/ed248580-942c-4ffc-919f-d3681d28a799)](https://codebeat.co/projects/github-com-vsouza-go-gin-boilerplate)
[![Build Status][travis-image]][travis-url]
[![License][license-image]][license-url]

![](header.png)

## Installation

`make install`

## Configuration

Inside of config folder, `example.yaml` file contains a sample configuration to setup your watcher.

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

* 0.0.1
    * Work in progress

## Contribute

We would love for you to contribute to **Watcher**, check the ``LICENSE`` file for more info.

## Meta

Vinicius Souza – [@iamvsouza](https://twitter.com/iamvsouza) – hi@vsouza.com

Distributed under the MIT license. See ``LICENSE`` for more information.

[https://github.com/vsouza](https://github.com/vsouza/)

[swift-image]:https://img.shields.io/badge/Go--version-1.6-blue.svg
[swift-url]: https://swift.org/
[license-image]: https://img.shields.io/badge/License-MIT-blue.svg
[license-url]: LICENSE
[travis-image]: https://img.shields.io/travis/dbader/node-datadog-metrics/master.svg
[travis-url]: https://travis-ci.org/dbader/node-datadog-metrics
[codebeat-image]: https://codebeat.co/badges/c19b47ea-2f9d-45df-8458-b2d952fe9dad
[codebeat-url]: https://codebeat.co/projects/github-com-vsouza-awesomeios-com
