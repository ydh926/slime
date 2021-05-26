/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"slime.io/slime/slime-framework/util"
	"slime.io/slime/slime-modules/discovery/api/v1alpha1"
	"slime.io/slime/slime-modules/discovery/api/v1alpha1/wrapper"
	"slime.io/slime/slime-modules/discovery/model"
	meshsource "slime.io/slime/slime-modules/discovery/source"
	"slime.io/slime/slime-modules/discovery/source/eureka"
	"slime.io/slime/slime-modules/discovery/source/zookeeper"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MeshSourceReconciler reconciles a MeshSource object
type MeshSourceReconciler struct {
	client.Client
	Log     logr.Logger
	Scheme  *runtime.Scheme
	outer   model.Actor
	sources map[types.NamespacedName]meshsource.Source
}

// +kubebuilder:rbac:groups=microservice.slime.io.my.domain,resources=meshsources,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=microservice.slime.io.my.domain,resources=meshsources/status,verbs=get;update;patch

func (r *MeshSourceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("meshsource", req.NamespacedName)
	// your logic here
	instance := &wrapper.MeshSource{}
	err := r.Client.Get(context.TODO(), req.NamespacedName, instance)
	// 异常分支
	if err != nil && !errors.IsNotFound(err) {
		return reconcile.Result{}, err
	}

	// 资源删除
	if err != nil && errors.IsNotFound(err) {
		if s, ok := r.sources[req.NamespacedName]; ok {
			s.Stop()
		}
		delete(r.sources, req.NamespacedName)
		return reconcile.Result{}, nil
	}

	// 资源更新
	pb, err := util.FromJSONMap("slime.microservice.v1alpha1.MeshSource", instance.Spec)
	if err != nil {
		old := r.sources[req.NamespacedName]
		if !reflect.DeepEqual(pb, old) {
			old.Stop()
			if a, ok := pb.(*v1alpha1.MeshSourceSpec); ok {
				switch m := a.Source.(type) {
				case *v1alpha1.MeshSourceSpec_Zookeeper:
					if s, err := zookeeper.New(m.Zookeeper, r.outer, a.MappingNamespace); err == nil {
						s.Start()
						r.sources[req.NamespacedName] = s
					} else {
						// todo log

					}
				case *v1alpha1.MeshSourceSpec_Eureka:
					if s, err := eureka.New(m.Eureka,r.outer,a.MappingNamespace);err == nil {
						s.Start()
						r.sources[req.NamespacedName] = s
					}
				}
			}
		}
	}
	return ctrl.Result{}, nil
}

func (r *MeshSourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&wrapper.MeshSource{}).
		Complete(r)
}
