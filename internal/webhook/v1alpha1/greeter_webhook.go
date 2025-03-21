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

package v1alpha1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	friendlyv1alpha1 "github.com/sdischer-sap/webhook-learning/api/v1alpha1"
)

// nolint:unused
// log is for logging in this package.
var greeterlog = logf.Log.WithName("greeter-resource")

// SetupGreeterWebhookWithManager registers the webhook for Greeter in the manager.
func SetupGreeterWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&friendlyv1alpha1.Greeter{}).
		WithValidator(&GreeterCustomValidator{}).
		WithDefaulter(&GreeterCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-friendly-orchestrate-cloud-sap-v1alpha1-greeter,mutating=true,failurePolicy=fail,sideEffects=None,groups=friendly.orchestrate.cloud.sap,resources=greeters,verbs=create;update,versions=v1alpha1,name=mgreeter-v1alpha1.kb.io,admissionReviewVersions=v1

// GreeterCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind Greeter when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type GreeterCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &GreeterCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind Greeter.
func (d *GreeterCustomDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	greeter, ok := obj.(*friendlyv1alpha1.Greeter)

	if !ok {
		return fmt.Errorf("expected an Greeter object but got %T", obj)
	}
	greeterlog.Info("Defaulting for Greeter", "name", greeter.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-friendly-orchestrate-cloud-sap-v1alpha1-greeter,mutating=false,failurePolicy=fail,sideEffects=None,groups=friendly.orchestrate.cloud.sap,resources=greeters,verbs=create;update,versions=v1alpha1,name=vgreeter-v1alpha1.kb.io,admissionReviewVersions=v1

// GreeterCustomValidator struct is responsible for validating the Greeter resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type GreeterCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &GreeterCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type Greeter.
func (v *GreeterCustomValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	greeter, ok := obj.(*friendlyv1alpha1.Greeter)
	if !ok {
		return nil, fmt.Errorf("expected a Greeter object but got %T", obj)
	}
	greeterlog.Info("Validation for Greeter upon creation", "name", greeter.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type Greeter.
func (v *GreeterCustomValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	greeter, ok := newObj.(*friendlyv1alpha1.Greeter)
	if !ok {
		return nil, fmt.Errorf("expected a Greeter object for the newObj but got %T", newObj)
	}
	greeterlog.Info("Validation for Greeter upon update", "name", greeter.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type Greeter.
func (v *GreeterCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	greeter, ok := obj.(*friendlyv1alpha1.Greeter)
	if !ok {
		return nil, fmt.Errorf("expected a Greeter object but got %T", obj)
	}
	greeterlog.Info("Validation for Greeter upon deletion", "name", greeter.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
