
[![Bonsai Asset Badge](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/ArcticXWolf/sensu-sftp-check)

# Sensu Go SFTP Check

- [Overview](#overview)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Resource Configuration](#resource-configuration)
- [Functionality](#functionality)
- [Installation from source and contributing](#installation-from-source-and-contributing)

## Overview

TODO

## Usage examples

TODO

## Configuration

### Asset Registration

Assets are the best way to make use of this plugin. If you're not using an asset, please consider doing so! If you're using sensuctl 5.13 or later, you can use the following command to add the asset: 

`sensuctl asset add ArcticXWolf/sensu-sftp-check`

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index](https://bonsai.sensu.io/assets/ArcticXWolf/sensu-sftp-check).

### Resource configuration
TODO

### Functionality

TODO

## Installation from source and contributing

The preferred way of installing and deploying this plugin is to use it as an [asset][2]. If you would like to compile and install the plugin from source or contribute to it, download the latest version of the sensu-sftp-check from [releases][1]
or create an executable script from this source.

From the local path of the sensu-sftp-check repository:

```
go build -o /usr/local/bin/sensu-sftp-check cmd/sensu-sftp-check/main.go
```

For more information about contributing to this plugin, see https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md

[1]: https://github.com/ArcticXWolf/sensu-sftp-check/releases
[2]: #asset-registration

## Credits
