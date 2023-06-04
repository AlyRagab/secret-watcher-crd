package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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
	SchemeBuilder.Register(&SecretWatcher{}, &SecretWatcherList{})
}
