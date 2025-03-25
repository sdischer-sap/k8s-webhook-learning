/*
Copyright 2025.

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

package controller

import (
	"context"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	v1alpha1 "github.com/sdischer-sap/webhook-learning/api/v1beta1"
)

var requeueAfter = time.Duration(10 * time.Second)

// GreeterReconciler reconciles a Greeter object
type GreeterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=friendly.orchestrate.cloud.sap,resources=greeters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=friendly.orchestrate.cloud.sap,resources=greeters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=friendly.orchestrate.cloud.sap,resources=greeters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Greeter object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.2/pkg/reconcile
func (r *GreeterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	greeter := v1alpha1.Greeter{}
	if err := r.Client.Get(ctx, req.NamespacedName, &greeter); err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	now := metav1.Now()
	in10Secs := greeter.Status.LastGreeting.Add(time.Duration(10 * time.Second))

	if (greeter.Status.LastGreeting == metav1.Time{}) || now.After(in10Secs) {
		log.Info("Hello " + strings.Join(greeter.Spec.People, ", "))
		greeter.Status.AmountOfGreetings++
		greeter.Status.LastGreeting = metav1.Now()
		if err := r.Client.Status().Update(ctx, &greeter); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{RequeueAfter: requeueAfter}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GreeterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Greeter{}).
		Named("greeter").
		Complete(r)
}
