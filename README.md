# cobra-ui
Corbra-ui is a web system for [Cobra](https://github.com/spf13/cobra), which is a library for creating powerful modern CLI applications.

You can simply add several code lines to add web support to your exist cobra based applications. With cobra-ui, you can run commands via web browser. This is useful, because you maybe want someone else run theses commands, who has no chance to run as CLI.

Give a ⭐️ if this project helped you!

## Install

Juat import cobra-ui package into the place where the root command is...

```go
import (
    cobraui "github.com/ax2/cobra-ui"

    ...
)

and then add cobra-ui command to the root command

```go
rootCmd.AddCommand(cobraui.UICmd)
```

and more, you should change your cobra application, add following line to the command to enable it shown on web:

```go
	urlEncodeCmd = &cobra.Command{
		Use:         "urlencode text",
		Short:       "url encode",
		Args:        cobra.MinimumNArgs(1),
		Annotations: cobraui.Options(),  // Add this line
		Run:         urlencode,
	}
```

Now, you can start the web application by running the `ui` command,

```sh
./your-cobra-application ui
```

Code in examples/devx is full cobra application example, here is the screenshot

![cobra-ui](screenshots.jpg)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!

Feel free to check [issues page](https://github.com/ax2/cobra-ui/issues). 

## Roadmap

- Authorization support, use basic auth for seperated commands.
- File upload.
- Flags and args type support, default values.
- Realtime command output display.
- Configurations.
- Download results.
- Scheduled jobs.
- Run history.
- Favorite commands.
- Personalization.

## 📝 License

Copyright © 2022 [Alex Xiang](https://github.com/ax2).

This project is [MIT](https://github.com/ryanande/battlegrip/blob/master/LICENSE) licensed.
