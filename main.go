package main

import (
	"flag"
	"fmt"
	"github.com/swaince/image-repo/conf"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"os"
)

var path string

func init() {
	flag.StringVar(&path, "p", "conf/version.yaml", "version conf path")
}

func main() {
	flag.Parse()
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var rs []conf.Repository
	err = yaml.Unmarshal(b, &rs)

	if err != nil {
		panic(err)
	}

	for _, r := range rs {
		for _, p := range r.Projects {
			for _, v := range p.Versions {
				path := fmt.Sprintf("repo/%s/%s/%s", r.Workspace, p.Name, v.Tag)
				err := os.MkdirAll(path, fs.ModeDir)
				if err != nil {
					fmt.Println(err)
					continue
				}
				from := ""
				if v.Tag == "" && v.Digest == "" {
					from = fmt.Sprintf("FROM %s:latest", p.Url)
				} else if v.Tag == "" {
					from = fmt.Sprintf("FROM %s@%s", p.Url, v.Digest)
				} else {
					from = fmt.Sprintf("FROM %s:%s@%s", p.Url, v.Tag, v.Digest)
				}
				fp := path + "/Dockerfile"
				os.Chmod(fp, 0666)
				err = ioutil.WriteFile(fp, []byte(from), fs.ModeAppend)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
		}
	}

}
