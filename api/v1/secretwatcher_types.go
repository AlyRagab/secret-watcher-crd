package v1

import (
	"github.com/rs/zerolog/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

// SecretWatcherSpec defines the desired state of SecretWatcher
type SecretWatcherSpec struct {
}

// SecretWatcherStatus defines the observed state of SecretWatcher
type SecretWatcherStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SecretWatcher is the Schema for the secretwatchers API
type SecretWatcher struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecretWatcherSpec   `json:"spec,omitempty"`
	Status SecretWatcherStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretWatcherList contains a list of SecretWatcher
type SecretWatcherList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretWatcher `json:"items"`
}

func init() {
	err := AddToScheme(scheme.Scheme)
	if err != nil {
		log.Printf("Failed to add custom types to scheme: %v", err)
	}
}

// Add custom types to the scheme
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&SecretWatcher{},
		&SecretWatcherList{},
	)
	// Register CRD types with the Kubernetes scheme
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
