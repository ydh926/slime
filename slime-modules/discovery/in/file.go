package in

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/types"
	"os"
	"path/filepath"
	"reflect"
	"slime.io/slime/slime-framework/util"
	"slime.io/slime/slime-modules/discovery/api/v1alpha1"
	"slime.io/slime/slime-modules/discovery/api/v1alpha1/wrapper"
	"slime.io/slime/slime-modules/discovery/model"
	meshsource "slime.io/slime/slime-modules/discovery/source"
	"slime.io/slime/slime-modules/discovery/source/eureka"
	"slime.io/slime/slime-modules/discovery/source/zookeeper"
	"strings"
)

type FileWatcher struct {
	Watch     *fsnotify.Watcher
	stop      chan struct{}
	out       chan *wrapper.MeshSource
	configDir string
	sources   map[types.NamespacedName]meshsource.Source
	outer    model.Actor
}

func NewFileWatcher(configDir string, outer model.Actor)*FileWatcher{
	return nil
}

func (w *FileWatcher) InitProcess() {
	files, _ := ioutil.ReadDir(w.configDir)
	for _,f := range files{
		if  !f.IsDir() && (strings.HasSuffix(f.Name(),".yaml") || strings.HasSuffix(f.Name(),"yml")){
			buffer,err := ioutil.ReadFile(strings.Join([]string{w.configDir,f.Name()},"/"))
			if err == nil {
				ms := &wrapper.MeshSource{}
				err = yaml.Unmarshal(buffer, ms)
				if err == nil {
					pb, err := util.FromJSONMap("slime.microservice.v1alpha1.MeshSource", ms.Spec)
					if err != nil {
						fmt.Println(err.Error())
						continue
					}
					if a, ok := pb.(*v1alpha1.MeshSourceSpec); ok {
						w.Process(a, types.NamespacedName{
							Name: ms.Name,
							Namespace: ms.Namespace,
						})
					}
				}
			}
		}
	}

}

func (w *FileWatcher) Process(message *v1alpha1.MeshSourceSpec, loc types.NamespacedName) {
	old := w.sources[loc]
	if !reflect.DeepEqual(message, old) {
		old.Stop()
		switch m := message.Source.(type) {
		case *v1alpha1.MeshSourceSpec_Zookeeper:
			if s, err := zookeeper.New(m.Zookeeper, w.outer, message.MappingNamespace); err == nil {
				s.Start()
				w.sources[loc] = s
			} else {
				// todo log
			}
		case *v1alpha1.MeshSourceSpec_Eureka:
			if s, err := eureka.New(m.Eureka, w.outer, message.MappingNamespace); err == nil {
				s.Start()
				w.sources[loc] = s
			}
		}
	}
}

func (w *FileWatcher) WatchDir(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			path, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			err = w.Watch.Add(path)
			if err != nil {
				return err
			}
			fmt.Println("monitor：", path)
		}
		return nil
	})

	go func() {
		for {
			select {
			case ev := <-w.Watch.Events:
				buffer, err := os.ReadFile(ev.Name)
				if err != nil {
					fmt.Printf(err.Error())
				}
				ms := &wrapper.MeshSource{}
				_ = yaml.Unmarshal(buffer, ms)
				if err != nil {
					fmt.Printf(err.Error())
				}
				fmt.Printf("%v", ms)
				w.out <- ms
			case <-w.stop:
				return
			}
		}
	}()
}
