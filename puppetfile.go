package main

import "os"
import "log"
import "url"
import "errors"
import "path/filepath"
import "io/ioutil"
import "encoding/json"
import "gopkg.in/libgit2/git2go.v22"

type Puppetfile struct {
	Forge   string   `json:"forge"`
	Modules []Module `json:"modules"`
}

type Module struct {
	Name      string `json:"name"`
	Source    string `json:"source"`
	Ref       string `json:"ref"`
	Directory string `json:"directory"`
}

func (p *Puppetfile) Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err == nil {
		err = json.Unmarshal([]byte(data), p)
	}
	return err
}

func CachePath(m *Module, cacheroot string) (string, error) {
	source, err := url.Parse(m.Source)
	if err == nil {
		return filepath.Join(cacheroot, source.Host, source.Path), nil
	}
	return "", err
}

func MkCacheDir(cachedir string) error {
	if f, err := file.Stat(modulecache); err != nil {
		return file.Mkdirall(modulecache)
	} else {
		if f.IsDir != true {
			return errors.New("The cache directory %s already exists as a regular file!", modulecache)
		}
	}
	return nil
}

func FetchCache(m *Module, cacheroot string) {
	modulecache, err := CachePath(m, cacheroot)
	if err == nil {
		if err := MkCacheDir(modulecache); err != nil {
			log.Fatal(err)
		}
	}
}
