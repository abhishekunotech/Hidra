package main

import (
        "fmt"
        "log"
	"path/filepath"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	jww "github.com/spf13/jwalterweatherman"
	//"encoding/json"
	//"bytes"
)

type config struct{
	configPaths []string
	configName string
	configFile string
	configType string
	onConfigChange func(fsnotify.Event)
}

type jsonobject struct {
	Version string `json:"version"`
	Routes []RoutesType 
	Components []ComponentsType 
}

type RoutesType struct
{
	Name string `json:"name"`
	Method string `json:"method"`
	URI string `json:"URI"`
	Handler string `json:"handler"`
}

type ComponentsType struct
{
        ComponentName string `json:"componentName"`
        URL string `json:"url"`
        API []APIs 
}

type APIs struct
{
        Name string `json:"name"`
        URI string `json:"URI"` 
        Parameters []Params 
}

type Params struct
{
	TicketId string `json:"TicketId"`
	UserLogin string `json:"UserLogin"`
	Password string	`json:"Password"`

}


func main()
{

	file, e := ioutil.ReadFile("/etc/Hydra/conf.d/Hydra.js")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
	WatchConfig()

}

func WatchConfig() {
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()

		// we have to watch the entire directory to pick up renames/atomic saves in a cross-platform way
		filename, err := v.getConfigFile()
		if err != nil {
			log.Println("error:", err)
			return
		}

		configFile := filepath.Clean(filename)
		configDir, _ := filepath.Split(configFile)

		done := make(chan bool)
		go func() {
			for {
				select {
				case event := <-watcher.Events:
					// we only care about the config file
					if filepath.Clean(event.Name) == configFile {
						if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
							err := v.ReadInConfig()
							if err != nil {
								log.Println("error:", err)
							}
							v.onConfigChange(event)
						}
					}
				case err := <-watcher.Errors:
					log.Println("error:", err)
				}
			}
		}()

		watcher.Add(configDir)
		<-done
	}()
}

func (v *Viper) ReadInConfig() error {
	jww.INFO.Println("Attempting to read in config file")
	filename, err := v.getConfigFile()
	if err != nil {
		return err
	}

	if !stringInSlice(v.getConfigType(), SupportedExts) {
		return UnsupportedConfigError(v.getConfigType())
	}

	file, err := afero.ReadFile(v.fs, filename)
	if err != nil {
		return err
	}

	config := make(map[string]interface{})

	err = v.unmarshalReader(bytes.NewReader(file), config)
	if err != nil {
		return err
	}

	v.config = config
	return nil
}

func (v *Viper) OnConfigChange(run func(in fsnotify.Event)) {
	v.onConfigChange = run
}
func (v *Viper) getConfigFile() (string, error) {
	// if explicitly set, then use it
	if v.configFile != "" {
		return v.configFile, nil
	}

	cf, err := v.findConfigFile()
	if err != nil {
		return "", err
	}

	v.configFile = cf
	return v.getConfigFile()
}
func (v *Viper) getConfigType() string {
	if v.configType != "" {
		return v.configType
	}

	cf, err := v.getConfigFile()
	if err != nil {
		return ""
	}

	ext := filepath.Ext(cf)

	if len(ext) > 1 {
		return ext[1:]
	}

	return ""
}

func (v *Viper) findConfigFile() (string, error) {
	jww.INFO.Println("Searching for config in ", v.configPaths)

	for _, cp := range v.configPaths {
		file := v.searchInPath(cp)
		if file != "" {
			return file, nil
		}
	}
	return "", ConfigFileNotFoundError{v.configName, fmt.Sprintf("%s", v.configPaths)}
}

func (v *Viper) searchInPath(in string) (filename string) {
	jww.DEBUG.Println("Searching for config in ", in)
	for _, ext := range SupportedExts {
		jww.DEBUG.Println("Checking for", filepath.Join(in, v.configName+"."+ext))
		if b, _ := exists(filepath.Join(in, v.configName+"."+ext)); b {
			jww.DEBUG.Println("Found: ", filepath.Join(in, v.configName+"."+ext))
			return filepath.Join(in, v.configName+"."+ext)
		}
	}

	return ""
}

