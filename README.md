# TPM.tools

Source code for <tpm.tools>. Feel free to deploy and use this in accordance with the [license](LICENSE).

## Build

This project builds with `docker build`.

```sh
docker build -t tpm-tools:latest .
```

## Deploy

While this project is designed for deployment to <fly.io>, you can also just `docker run` it.

```sh
docker run -p 8080:8080 tpm-tools:latest
```
