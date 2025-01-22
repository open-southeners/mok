Mok ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/open-southeners/mok) [![codecov](https://codecov.io/gh/open-southeners/mok/branch/main/graph/badge.svg?token=zx68fmehzI)](https://codecov.io/gh/open-southeners/mok) [![Edit on VSCode online](https://img.shields.io/badge/vscode-edit%20online-blue?logo=visualstudiocode)](https://vscode.dev/github/open-southeners/mok)
===

Mock APIs as files with a NextJS routing directory style

## Getting started

Download the executable from the [latest release](https://github.com/opensoutheners/mok/releases/latest).

Run the following under a folder with some `.json` files:

```sh
mok
```

Then you can go to [http://localhost:8080](http://localhost:8080) and check for yourself

You can specify the address and the port to listen:

```sh
mok --listen 0.0.0.0 --port 80
```

Or use a different folder than the current:

```sh
mok path/to/my/json/files
```

### Random responses

To be able to randomise the content of our JSON files we can use all functions from [gofakeit](https://github.com/brianvoe/gofakeit?tab=readme-ov-file#functions) and passed request data.

Check [example](./example) folder for more info.

## Partners

[![skore logo](https://github.com/open-southeners/partners/raw/main/logos/skore_logo.png)](https://getskore.com)

## License

This package is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).
