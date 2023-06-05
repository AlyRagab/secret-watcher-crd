package controllers

import (
	"context"
	"fmt"
	"os"

	v1 "github.com/AlyRagab/secret-watcher-crd/api/v1"
	"github.com/nlopes/slack"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// Create a global logger instance
var Logger = initLogger()

// SecretWatcherReconciler reconciles a SecretWatcher object
type SecretWatcherReconciler struct {
	client.Client
}

// +kubebuilder:rbac:groups=secretwatcher.aly.com,resources=secretwatchers,verbs=get;list;watch;create;update;patch;delete
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
		// fmt.Printf("Processing Secret: %s/%s\n", secret.Namespace, secret.Name)
		// Send a Slack alert for each Secret
		message := fmt.Sprintf("Secret in Namespace=%s, Name=%s", secret.Namespace, secret.Name)
		err := sendAlertToSlack(message)
		if err != nil {
			Logger.WithFields(logrus.Fields{
				"request": req,
			}).Error("Failed to reconcile.")
		}
	}

	return ctrl.Result{}, nil
}

func sendAlertToSlack(message string) error {
	fmt.Println("Building Slack Object")
	api := slack.New("xoxp-5079192664693-5105825226784-5356530773671-3d2e572cc24bf6828c695a15f4b2292c")
	_, _, err := api.PostMessage("C05ATJV2F4K", slack.MsgOptionText(message, false))
	if err != nil {
		fmt.Printf("Failed to send Slack alert: %v\n", err)
	}
	return nil
}

func SetupWithManager(mgr manager.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.SecretWatcher{}).
		Watches(&corev1.Secret{}, &handler.EnqueueRequestForObject{}).
		Complete(&SecretWatcherReconciler{
			Client: mgr.GetClient(),
		})
}

func initLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Set the desired log level
	logger.SetLevel(logrus.DebugLevel)

	return logger
}
