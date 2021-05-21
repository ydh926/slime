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
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"slime.io/slime/slime-modules/discovery/api/v1alpha1/wrapper"
	"slime.io/slime/slime-modules/discovery/model"
	meshsource "slime.io/slime/slime-modules/discovery/source"

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
	if
	if reflect.DeepEqual(instance.Spec, r.lastUpdatePolicy) {
		r.lastUpdatePolicyLock.RUnlock()
		return reconcile.Result{}, nil
	} else {
		r.lastUpdatePolicyLock.RUnlock()
		r.lastUpdatePolicyLock.Lock()
		r.lastUpdatePolicy = instance.Spec
		r.lastUpdatePolicyLock.Unlock()
		r.source.WatchAdd(req.NamespacedName)
	}
	r.source.WatchAdd(req.NamespacedName)

	return ctrl.Result{}, nil
}

func (r *MeshSourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&wrapper.MeshSource{}).
		Complete(r)
}
