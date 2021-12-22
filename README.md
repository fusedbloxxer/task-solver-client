# task-solver-client

## Install dependencies

1. You have to first install the [Go Language](https://go.dev/dl/) on your system. I used the
**1.17.3** version for windows to develop the server application.
2. I used the [Fyne](https://developer.fyne.io/) toolkit to build a cross-platform user interface. According
to the [documentation](https://developer.fyne.io/started/) you have to install a C Compiler because the
toolkit uses OpenGL (GLFW) under the hood. I recommend you to install and set up the
[TDM-GCC](https://jmeubank.github.io/tdm-gcc/download/) if you are developing on Windows.

## Configure the Client

The file at ``./config/client.json`` contains the configuration options
for the client.

You can change the settings for your needs. In order for the client to communicate with
the server you should provide the correct ``api.baseUrl``.

## Run the Client

To run the client

```bash
go run ./src
```