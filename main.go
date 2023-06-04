package main

import (
	"flag"
	"fmt"
	"os"

	v1 "github.com/AlyRagab/secret-watcher-crd/api/v1"
	"github.com/AlyRagab/secret-watcher-crd/controllers"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func main() {
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

	scheme := runtime.NewScheme()
	corev1.AddToScheme(scheme)
	v1.AddToScheme(scheme)

	kubeConfig := config.GetConfigOrDie()
	mgr, err := ctrl.NewManager(kubeConfig, ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: "0",
		Port:               9443,
		LeaderElection:     true,
		LeaderElectionID:   "secretwatcher.aly.com",
	})
	if err != nil {
		fmt.Printf("Failed to create controller manager: %v", err)
		os.Exit(1)
	}

	err = controllers.SetupWithManager(mgr)
	if err != nil {
		fmt.Printf("Failed to setup controller: %v", err)
		os.Exit(1)
	}

	err = mgr.Start(ctrl.SetupSignalHandler())
	if err != nil {
		fmt.Printf("Failed to start manager: %v", err)
		os.Exit(1)
	}
}
