package in

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"slime.io/slime/slime-modules/discovery/api/v1alpha1/wrapper"
	"testing"
)

func TestFileWatcher_WatchDir(t *testing.T) {
	buffer,err := ioutil.ReadFile("test/eureka.yaml")
	if err == nil {
		a := &wrapper.MeshSource{}
		err = yaml.Unmarshal(buffer,a)
		if err != nil {
			fmt.Printf("%v",err)
		}
		fmt.Printf("%v",a)
	}
	f,_ := fsnotify.NewWatcher()
	fw := FileWatcher{
		Watch: f,
		stop: make(chan struct{}),
	}
	fw.WatchDir("test")
	m := <-fw.out
	fmt.Printf("%v",m)
}
