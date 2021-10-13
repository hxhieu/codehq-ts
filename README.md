# `.env` file

```
CODEHQ_TS_TIMESHEET_DSN='sqlserver://USER:PASS@HOST?database=DB&parseTime=true'
CODEHQ_TS_PIMP_DSN='sqlserver://USER:PASS@HOST?database=DB&parseTime=true'
```

Or it will fallback to the machine env vars if no `.env` file found.

Ultimately it will crash if nothing found :)

# To scaffold new command

`cobra add COMMAND --config ./cobra.yaml`

For more details on `cobra` generator, read [https://github.com/spf13/cobra/blob/master/cobra/README.md](https://github.com/spf13/cobra/blob/master/cobra/README.md)

For sub commands with the same name, just use `cobra add` with specific file name such as `timesheetCreate` and `timesheetGet` then modify the generated code with shorten command name

NOTES:

> All time is parsed in LOCAL time
