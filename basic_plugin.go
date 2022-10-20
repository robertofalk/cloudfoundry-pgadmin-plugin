package main

import (
	"fmt"
	"os"

	"code.cloudfoundry.org/cli/plugin"
)

type RouteMapping struct {
	Resources []struct {
		Metadata struct {
			GUID string `json:"guid"`
			URL  string `json:"url"`
		} `json:"metadata"`
		Entity struct {
			RouteGUID string `json:"route_guid"`
		} `json:"entity"`
	} `json:"resources"`
}

// BasicPlugin is the struct implementing the interface defined by the core CLI. It can
// be found at  "code.cloudfoundry.org/cli/plugin/plugin.go"
type BasicPlugin struct{}

func handleError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

// Run must be implemented by any plugin because it is part of the
// plugin interface defined by the core CLI.
//
// Run(....) is the entry point when the core CLI is invoking a command defined
// by a plugin. The first parameter, plugin.CliConnection, is a struct that can
// be used to invoke cli commands. The second paramter, args, is a slice of
// strings. args[0] will be the name of the command, and will be followed by
// any additional arguments a cli user typed in.
//
// Any error handling should be handled with the plugin itself (this means printing
// user facing errors). The CLI will exit 0 if the plugin exits 0 and will exit
// 1 should the plugin exits nonzero.
func (c *BasicPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	// Ensure that we called the command basic-plugin-command
	if args[0] == "pgadmin" {

		fmt.Println("Running the pgadmin plugin")
		var appname string
		fmt.Printf("Enter the app name: ")

		fmt.Scanf("%s", &appname)
		cliConnection.CliCommand("push", appname, "--docker-image", "dpage/pgadmin4", "--no-start", "--random-route")

		// cliConnection.CliCommand("set-env", appname, "PGADMIN_DEFAULT_EMAIL", "admin@admin.com")
		// cliConnection.CliCommand("set-env", appname, "PGADMIN_DEFAULT_PASSWORD", "admin")
		// cliConnection.CliCommand("set-env", appname, "PGADMIN_LISTEN_ADDRESS", "0.0.0.0")
		// cliConnection.CliCommand("set-env", appname, "PGADMIN_LISTEN_PORT", "8080")

		// app, err := cliConnection.GetApp(appname)
		// handleError(err)

		// //CliCommandWithoutTerminalOutput
		// url := fmt.Sprintf("/v2/apps/%s", app.Guid)
		// _, err = cliConnection.CliCommand("curl", url, "-X", "PUT", "-d", "'{\"ports\": [8080]}'")
		// handleError(err)

		// url = fmt.Sprintf("/v2/apps/%s/route_mappings", app.Guid)
		// env, err := cliConnection.CliCommand("curl", url)
		// handleError(err)

		// var route RouteMapping
		// json.Unmarshal([]byte(strings.Join(env, "")), &route)

		// body := map[string]interface{}{"app_guid": app.Guid, "route_guid": route.Resources[0].Entity.RouteGUID, "app_port": 8080}
		// param, _ := json.Marshal(body)
		// _, err = cliConnection.CliCommand("curl", "/v2/route_mappings", "-X", "POST", "-d", fmt.Sprintf("%s", param))
		// handleError(err)

		// _, err = cliConnection.CliCommand("curl", route.Resources[0].Metadata.URL, "-X", "DELETE")
		// handleError(err)
		// cliConnection.CliCommand("restage", appname)

	}
}

// GetMetadata must be implemented as part of the plugin interface
// defined by the core CLI.
//
// GetMetadata() returns a PluginMetadata struct. The first field, Name,
// determines the name of the plugin which should generally be without spaces.
// If there are spaces in the name a user will need to properly quote the name
// during uninstall otherwise the name will be treated as seperate arguments.
// The second value is a slice of Command structs. Our slice only contains one
// Command Struct, but could contain any number of them. The first field Name
// defines the command `cf basic-plugin-command` once installed into the CLI. The
// second field, HelpText, is used by the core CLI to display help information
// to the user in the core commands `cf help`, `cf`, or `cf -h`.
func (c *BasicPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "cf-pgadmin",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "pgadmin",
				HelpText: "Deploy pgadmin and configure it to run on CF environment",

				// UsageDetails is optional
				// It is used to show help of usage of each command
				UsageDetails: plugin.Usage{
					Usage: "pgadmin\n   cf pgadmin",
				},
			},
		},
	}
}

// Unlike most Go programs, the `Main()` function will not be used to run all of the
// commands provided in your plugin. Main will be used to initialize the plugin
// process, as well as any dependencies you might require for your
// plugin.
func main() {
	// Any initialization for your plugin can be handled here
	//
	// Note: to run the plugin.Start method, we pass in a pointer to the struct
	// implementing the interface defined at "code.cloudfoundry.org/cli/plugin/plugin.go"
	//
	// Note: The plugin's main() method is invoked at install time to collect
	// metadata. The plugin will exit 0 and the Run([]string) method will not be
	// invoked.
	plugin.Start(new(BasicPlugin))
	// Plugin code should be written in the Run([]string) method,
	// ensuring the plugin environment is bootstrapped.
}
