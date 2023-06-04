# secret-watcher-crd

A SecretWatcher CRD which watches for any change made against a Secret Resource

## Getting Started

```bash
operator-sdk init --domain secretwatcher.aly.com --repo github.com/AlyRagab/secret-watcher-crd # Scafollds the project
operator-sdk create api --group secretwatcher --version v1 --kind SecretWatcher --resource --controller # Creates the API
make manifests # Creates the Yaml definition
```

If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```bash
make manifests
```

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/secret-watcher-crd:tag
```

3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/secret-watcher-crd:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller from the cluster:

```sh
make undeploy
```

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`
