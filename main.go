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
	// Create a new Zap logger
	logger := zap.New(zap.UseDevMode(true))

	// Set the logger
	ctrl.SetLogger(logger)
	kubeconfig := flag.String("config", "", "Path to the kubeconfig file")
	flag.Parse()
	// Load the kubeconfig file
	cfg, err := config.GetConfigWithContext(*kubeconfig)
	if err != nil {
		fmt.Printf("Failed to load kubeconfig: %v", err)
		os.Exit(1)
	}
	scheme := runtime.NewScheme()
	corev1.AddToScheme(scheme)
	v1.AddToScheme(scheme)

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		MetricsBindAddress: "0",
	})
	if err != nil {
		panic(err)
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
