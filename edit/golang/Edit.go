/* package edit is a prototype of a StreamWork graph editor.
This module contains constants and functions used elsewhere.

Overview: 
	Use Dotty to render graphs and return events. 
*/
package edit

/*  Generated StreamWork component, Edit,
    by sw.sh on Wed Jun 26 19:54:41 EDT 2019
*/

import (
	config "github.com/zpatrick/go-config"
)

const version string  = "v0.0.2"

/* pkgConfig initializes the go-config package.
See: https://github.com/zpatrick/go-config for details*/
func pkgConfig(configFile string) *config.Config {
	mappings := map[string]string{
		"EDIT_BUFFERSIZE": "edit.buffersize",
	}

	return config.NewConfig([]config.Provider{
		config.NewTOMLFile(configFile),
		config.NewEnvironment(mappings)})
}

