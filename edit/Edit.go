/* package edit is a prototype of a StreamWork graph editor.*/
package edit

/*  Generated StreamWork component, Edit,
    by sw.sh on Wed Jun 26 19:54:41 EDT 2019
	1  input  ports
	1 output ports
*/

import (
	"fmt"
	config "github.com/zpatrick/go-config"
	"sync"
)

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

/* Edit starts the sw editor running.*/
func Edit(wg *sync.WaitGroup, arg []string, cs []chan interface{}) {
	var version string = "v0.0.0"

	defer wg.Done()
	cfg := pkgConfig("/home/tyoung3/.sw/edit.toml")
	bs, _ := cfg.IntOr("edit.buffersize", 1)
	title, _ := cfg.StringOr("edit.title", "n/a")

	fmt.Println(title,
		"Running",
		arg[0],
		version,
		"bs =", bs)

}
