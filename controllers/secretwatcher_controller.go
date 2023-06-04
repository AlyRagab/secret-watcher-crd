package controllers

import (
	"context"
	"fmt"

	v1 "github.com/AlyRagab/secret-watcher-crd/api/v1"
	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// SecretWatcherReconciler reconciles a SecretWatcher object
type SecretWatcherReconciler struct {
	client.Client
}

// +kubebuilder:rbac:groups=aly.com,resources=secretwatchers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch

func (r *SecretWatcherReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// Fetch the SecretWatcher instance
	secretWatcher := &v1.SecretWatcher{}
	err := r.Get(ctx, req.NamespacedName, secretWatcher)
	if err != nil {
		// Error reading the object - requeue the request
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// List all Secrets in all namespaces
	secrets := &corev1.SecretList{}
	err = r.List(ctx, secrets, client.InNamespace(corev1.NamespaceAll))
	if err != nil {
		return ctrl.Result{}, err
	}

	// Handle each Secret
	for _, secret := range secrets.Items {
		// Process the Secret
		fmt.Printf("Processing Secret: %s/%s\n", secret.Namespace, secret.Name)
	}

	return ctrl.Result{}, nil
}

func SetupWithManager(mgr manager.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.SecretWatcher{}).
		Watches(&source.Kind{Type: &corev1.Secret{}}, &handler.EnqueueRequestForObject{}).
		Complete(&SecretWatcherReconciler{
			Client: mgr.GetClient(),
		})
}
