# WithENV

[![](https://licensebuttons.net/p/zero/1.0/88x31.png)](http://creativecommons.org/publicdomain/zero/1.0/)

WithENV allows one to set the environment variables from `.env` file and
execute the command, optionally with args.

Usage:

```shell
withenv [-f file] [-must] <command> [args...]
```

Default behaviour: if `.env` file not found, it will print a log message and
continue with execution of the command.

If `-must` specified, it will fail if there's no `.env` file.

If the `.env` file is called differently, one can specify `-f` flag with the
desired name, i.e.:

```shell
withenv -f .i_dont_like_default_names env
```
