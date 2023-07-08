# gcli
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/MrBanja/gcli/blob/main/LICENSE)

**gcli** is a command-line interface (CLI) tool for using ChatGPT. It allows you to interact with the ChatGPT language model directly from your terminal. The current version supports GPT-4, with plans to add support for new versions in the future. The project is written in pure Go.

## Table of Contents
- [gcli](#gcli)
    - [Building from Source](#building-from-source)
    - [Dependencies](#dependencies)
    - [Usage](#usage)
        - [gcli [query]](#gcli-query)
        - [gcli conv [command]](#gcli-conv-command)
        - [gcli config [command]](#gcli-config-command)
        - [gcli conv history [command]](#gcli-conv-history-command)
    - [License](#license)

You can use this table of contents as a quick reference to navigate through the document.

## Building from Source

If you prefer to build **gcli** from source, follow these steps:

1. Clone the repository:

   Open a terminal window and run the following command:
    ```bash
    git clone https://github.com/MrBanja/gcli.git
    ```
   This will create a new directory named "gcli" in your current path.

2. Navigate into the directory:

   Use the `cd` command to enter into the directory:
    ```bash
    cd gcli
    ```

3. Build the application:

   Run the following command to compile the project.
    ```bash
    go build
    ```
   This will create an executable file named "gcli".

4. Install application:

   To install the gcli on your system run the following command
    ```bash
    go install
    ```

Congratulations, you have installed `gcli` from source on your system. To verify the installation; try executing `gcli`. If the installation is successful you should be able to see usage options for gcli command.
## Dependencies

The project utilizes the following open-source libraries:

- [glamour](https://github.com/charmbracelet/glamour): A library for rendering Markdown.
- [screen](https://github.com/inancgumus/screen): A library for clearing the console.
- [cobra](https://github.com/spf13/cobra): A CLI framework for Go.
- [viper](https://github.com/spf13/viper): A library for configuration management.
- [openaiAPI](https://github.com/MrBanja/openaiAPI): A library for querying the OpenAI API.

## Usage

Before using **gcli**, you need to obtain an OpenAI API token from [https://platform.openai.com/account/api-keys](https://platform.openai.com/account/api-keys). Once you have the token, set it using the following command:

```shell
gcli config set openai.token sk-xxx
```

Here are the available commands and their usage:

### gcli [query]

Ask ChatGPT a question or provide an input query. Example usage:

```shell
gcli "What is the capital of France?"
```

### gcli conv [command]

Command for working with conversations. Available sub-commands:

- `gcli conv del`: Delete the current conversation.
- `gcli conv history`: Command for working with the current conversation history.
- `gcli conv ls`: Display all conversations.
- `gcli conv use`: Use a conversation with the given ID. If the conversation does not exist, it will be created.

### gcli config [command]

Set configuration options. Available sub-commands:

- `gcli config set`: Set configuration values.

### gcli conv history [command]

Command for working with the current conversation history. Available sub-commands:

- `gcli conv history all`: Show all messages in the conversation.
- `gcli conv history clear`: Clear the current conversation history.
- `gcli conv history last`: Show the last message of the current conversation.

For more information about each command and their flags, you can use the `-h` or `--help` flag with the respective command.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/MrBanja/gcli/blob/main/LICENSE) file for more information.