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
	"log"
	"strings"

	"sigs.k8s.io/controller-runtime/pkg/conversion"

	friendlyv1beta1 "github.com/sdischer-sap/webhook-learning/api/v1beta1"
)

// ConvertTo converts this Greeter (v1alpha1) to the Hub version (v1beta1).
func (src *Greeter) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*friendlyv1beta1.Greeter)
	log.Printf("ConvertTo: Converting Greeter from Spoke version v1alpha1 to Hub version v1beta1;"+
		"source: %s/%s, target: %s/%s", src.Namespace, src.Name, dst.Namespace, dst.Name)

	// simply mapping for most stuff
	dst.ObjectMeta = src.ObjectMeta
	dst.Status.LastGreeting = src.Status.LastGreeting
	dst.Status.AmountOfGreetings = src.Status.AmountOfGreetings

	// only format conversion
	splits := strings.Split(src.Spec.People, ",")
	dst.Spec.People = splits

	return nil
}

// ConvertFrom converts the Hub version (v1beta1) to this Greeter (v1alpha1).
func (dst *Greeter) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*friendlyv1beta1.Greeter)

	log.Printf("ConvertFrom: Converting Greeter from Hub version v1beta1 to Spoke version v1alpha1;"+
		"source: %s/%s, target: %s/%s", src.Namespace, src.Name, dst.Namespace, dst.Name)

	// simply mapping for most stuff
	dst.ObjectMeta = src.ObjectMeta
	dst.Status.AmountOfGreetings = src.Status.AmountOfGreetings
	dst.Status.LastGreeting = src.Status.LastGreeting

	// concat people list to comma separated string
	dst.Spec.People = strings.Join(src.Spec.People, ",")

	return nil
}
