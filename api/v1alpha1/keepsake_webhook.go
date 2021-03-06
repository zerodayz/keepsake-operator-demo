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

package v1alpha1

import (
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var keepsakelog = logf.Log.WithName("keepsake-resource")

func (r *Keepsake) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-keepsake-example-com-v1alpha1-keepsake,mutating=true,failurePolicy=fail,groups=keepsake.example.com,resources=keepsakes,verbs=create;update,versions=v1alpha1,name=mkeepsake.kb.io

var _ webhook.Defaulter = &Keepsake{}

func validateOdd(n int32) error {
	if n%2 == 0 {
		return errors.New("Cluster size must be an odd number")
	}
	return nil
}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Keepsake) Default() {
	keepsakelog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-keepsake-example-com-v1alpha1-keepsake,mutating=false,failurePolicy=fail,groups=keepsake.example.com,resources=keepsakes,versions=v1alpha1,name=vkeepsake.kb.io

var _ webhook.Validator = &Keepsake{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Keepsake) ValidateCreate() error {
	log.Info("validate create", "name", r.Name)
	return validateOdd(r.Spec.Size)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Keepsake) ValidateUpdate(old runtime.Object) error {
	log.Info("validate update", "name", r.Name)
	return validateOdd(r.Spec.Size)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Keepsake) ValidateDelete() error {
	keepsakelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
