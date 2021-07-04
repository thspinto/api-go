# api-go
Simple template for Golang APIs

## Supporting new cli commands

Cobra helps create a nice cli to manage the server and supporting operations. You can use the [code generator cli](https://github.com/spf13/cobra/blob/master/cobra/README.md) to add commands.

## Environment variable prefix
You might want to change the environment variable prefix in `cmd/roo.go`:

```
	viper.SetEnvPrefix("GO_API")
```

## Adding environment variables

Use the cli do add commands `cobra add <command name>`.

Add this line at the end of the `init()` function:
```
serverCmd.Flags().VisitAll(configureViper("server"))
```
Notice that the `server` string must be replace with the new command name. This line is for the config file work with hierarchy and also keep the environment variables working. See how the `port` configuration works as example. In the .config file it is under `server.port` and in the environment variable it will use `GO_API_SERVER_PORT`.
