# To scaffold new command

`cobra add COMMAND --config ./cobra.yaml`

For more details on `cobra` generator, read [https://github.com/spf13/cobra/blob/master/cobra/README.md](https://github.com/spf13/cobra/blob/master/cobra/README.md)

For sub commands with the same name, just use `cobra add` with specific file name such as `timesheetCreate` and `timesheetGet` then modify the generated code with shorten command name

NOTES:

> All time is parsed in LOCAL time
