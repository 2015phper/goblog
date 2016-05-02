package cache

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/deepzz0/go-common/log"
	"github.com/deepzz0/goblog/models/manage"
)

type cache struct {
	BackgroundLeftBar  map[string]string
	BackgroundLeftBars []*manage.Leftbar
}

var Cache = NewCache()

func NewCache() *cache {
	return &cache{BackgroundLeftBar: make(map[string]string)}
}

func init() {
	doReadBackLeftBarConfig()
}

func doReadBackLeftBarConfig() {
	path, _ := os.Getwd()
	f, err := os.Open(path + "/conf/backleft.conf")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &Cache.BackgroundLeftBars)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range Cache.BackgroundLeftBars {
		if v.ID != "" {
			Cache.BackgroundLeftBar[v.ID] = v.ID
		}
	}
}
