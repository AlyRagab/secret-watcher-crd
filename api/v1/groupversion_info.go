// Package v1 contains API Schema definitions for the secretwatcher v1 API group
// +kubebuilder:object:generate=true
// +groupName=secretwatcher.aly.com
package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// GroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "secretwatcher.aly.com", Version: "v1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	//SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}
	// AddToScheme adds the types in this group-version to the given scheme.
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)
