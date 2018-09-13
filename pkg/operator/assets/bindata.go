// Code generated by go-bindata.
// sources:
// manifests/bootstrap-pod.yaml
// manifests/controllerconfig.crd.yaml
// manifests/machineconfig.crd.yaml
// manifests/machineconfigcontroller/clusterrole.yaml
// manifests/machineconfigcontroller/clusterrolebinding.yaml
// manifests/machineconfigcontroller/controllerconfig.yaml
// manifests/machineconfigcontroller/deployment.yaml
// manifests/machineconfigcontroller/sa.yaml
// manifests/machineconfigdaemon/clusterrole.yaml
// manifests/machineconfigdaemon/clusterrolebinding.yaml
// manifests/machineconfigdaemon/daemonset.yaml
// manifests/machineconfigdaemon/sa.yaml
// manifests/machineconfigpool.crd.yaml
// manifests/machineconfigserver/clusterrole.yaml
// manifests/machineconfigserver/clusterrolebinding.yaml
// manifests/machineconfigserver/daemonset.yaml
// manifests/machineconfigserver/node-bootstrapper-sa.yaml
// manifests/machineconfigserver/node-bootstrapper-token.yaml
// manifests/machineconfigserver/sa.yaml
// manifests/master.machineconfigpool.yaml
// manifests/scc.yaml
// manifests/worker.machineconfigpool.yaml
// DO NOT EDIT!

package assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _manifestsBootstrapPodYaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  name: bootstrap-machine-config-operator
  namespace: {{.TargetNamespace}}
spec:
  initContainers:
  - name: machine-config-controller
    image: {{.Images.MachineConfigController}}
    args:
    - "bootstrap"
    - "--manifest-dir=/etc/mcc/bootstrap/manifests"
    - "--dest-dir=/etc/mcc/bootstrap/server"
    resources:
      limits:
        cpu: 20m
        memory: 50Mi
      requests:
        cpu: 20m
        memory: 50Mi
    securityContext:
      privileged: true
    volumeMounts:
    - name: bootstrap-manifests
      mountPath: /etc/mcc/bootstrap/manifests
    - name: server-basedir
      mountPath: /etc/mcc/bootstrap/server
  containers:
  - name: machine-config-server
    image: {{.Images.MachineConfigServer}}
    args:
      - "bootstrap"
    volumeMounts:
    - name: certs
      mountPath: /etc/ssl/mcs
    - name: etc-kubernetes
      mountPath: /etc/kubernetes/kubeconfig
    - name: server-basedir
      mountPath: /etc/mcs/bootstrap
    - name:  etcd-certs
      mountPath: /etc/ssl/etcd
    securityContext:
      privileged: true
  hostNetwork: true
  tolerations:
    - key: node-role.kubernetes.io/master
      operator: Exists
      effect: NoSchedule
  restartPolicy: Always
  volumes:
  - name: certs
    hostPath:
      path: /etc/ssl/mcs
  - name: etc-kubernetes
    hostPath:
      path: /etc/kubernetes/kubeconfig
  - name: server-basedir
    hostPath:
      path: /etc/mcs/bootstrap
  - name: etcd-certs
    hostPath:
      path: /etc/ssl/etcd
  - name: bootstrap-manifests
    hostPath:
      path: /etc/mcc/bootstrap/manifests
`)

func manifestsBootstrapPodYamlBytes() ([]byte, error) {
	return _manifestsBootstrapPodYaml, nil
}

func manifestsBootstrapPodYaml() (*asset, error) {
	bytes, err := manifestsBootstrapPodYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/bootstrap-pod.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsControllerconfigCrdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  # name must match the spec fields below, and be in the form: <plural>.<group>
  name: controllerconfigs.machineconfiguration.openshift.io
spec:
  # group name to use for REST API: /apis/<group>/<version>
  group: machineconfiguration.openshift.io
  # list of versions supported by this CustomResourceDefinition
  versions:
    - name: v1
      # Each version can be enabled/disabled by Served flag.
      served: true
      # One and only one version must be marked as the storage version.
      storage: true
  # either Namespaced or Cluster
  scope: Namespaced
  names:
    # plural name to be used in the URL: /apis/<group>/<version>/<plural>
    plural: controllerconfigs
    # singular name to be used as an alias on the CLI and for display
    singular: controllerconfig
    # kind is normally the CamelCased singular type. Your resource manifests use this.
    kind: ControllerConfig
`)

func manifestsControllerconfigCrdYamlBytes() ([]byte, error) {
	return _manifestsControllerconfigCrdYaml, nil
}

func manifestsControllerconfigCrdYaml() (*asset, error) {
	bytes, err := manifestsControllerconfigCrdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/controllerconfig.crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigCrdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  # name must match the spec fields below, and be in the form: <plural>.<group>
  name: machineconfigs.machineconfiguration.openshift.io
spec:
  # group name to use for REST API: /apis/<group>/<version>
  group: machineconfiguration.openshift.io
  # list of versions supported by this CustomResourceDefinition
  versions:
    - name: v1
      # Each version can be enabled/disabled by Served flag.
      served: true
      # One and only one version must be marked as the storage version.
      storage: true
  # either Namespaced or Cluster
  scope: Cluster
  names:
    # plural name to be used in the URL: /apis/<group>/<version>/<plural>
    plural: machineconfigs
    # singular name to be used as an alias on the CLI and for display
    singular: machineconfig
    # kind is normally the CamelCased singular type. Your resource manifests use this.
    kind: MachineConfig
`)

func manifestsMachineconfigCrdYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigCrdYaml, nil
}

func manifestsMachineconfigCrdYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigCrdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfig.crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: machine-config-controller
  namespace: {{.TargetNamespace}}
rules:
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["get", "list", "watch", "patch"]
- apiGroups: ["machineconfiguration.openshift.io"]
  resources: ["*"]
  verbs: ["*"]
- apiGroups: [""]
  resources: ["configmaps"]
  verbs: ["*"]
`)

func manifestsMachineconfigcontrollerClusterroleYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigcontrollerClusterroleYaml, nil
}

func manifestsMachineconfigcontrollerClusterroleYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigcontrollerClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigcontroller/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: machine-config-controller
  namespace: {{.TargetNamespace}}
roleRef:
  kind: ClusterRole
  name: machine-config-controller
subjects:
- kind: ServiceAccount
  namespace: {{.TargetNamespace}}
  name: machine-config-controller
`)

func manifestsMachineconfigcontrollerClusterrolebindingYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigcontrollerClusterrolebindingYaml, nil
}

func manifestsMachineconfigcontrollerClusterrolebindingYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigcontrollerClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigcontroller/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerControllerconfigYaml = []byte(`apiVersion: machineconfiguration.openshift.io/v1
kind: ControllerConfig
metadata:
  name: machine-config-controller
  namespace: {{.TargetNamespace}}
spec:
  clusterDNSIP: {{.ControllerConfig.ClusterDNSIP}}
  cloudProviderConfig: {{.ControllerConfig.CloudProviderConfig}}
  clusterName: {{.ControllerConfig.ClusterName}}
  platform: {{.ControllerConfig.Platform}}
  baseDomain: {{.ControllerConfig.BaseDomain}}
  etcdInitialCount: {{.ControllerConfig.EtcdInitialCount}}
  etcdCAData: {{.ControllerConfig.EtcdCAData | toString | b64enc}}
  rootCAData: {{.ControllerConfig.RootCAData | toString | b64enc}}
`)

func manifestsMachineconfigcontrollerControllerconfigYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigcontrollerControllerconfigYaml, nil
}

func manifestsMachineconfigcontrollerControllerconfigYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigcontrollerControllerconfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigcontroller/controllerconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: machine-config-controller
  namespace: {{.TargetNamespace}}
spec:
  selector:
    matchLabels:
      k8s-app: machine-config-controller
  template:
    metadata:
      labels:
        k8s-app: machine-config-controller
    spec:
      containers:
      - name: machine-config-controller
        image: {{.Images.MachineConfigController}}
        args:
        - "start"
        - "--resourcelock-namespace={{.TargetNamespace}}"
        resources:
          limits:
            cpu: 20m
            memory: 50Mi
          requests:
            cpu: 20m
            memory: 50Mi
      serviceAccountName: machine-config-controller
      nodeSelector:
        node-role.kubernetes.io/master: ""
      restartPolicy: Always
      tolerations:
      - key: "node-role.kubernetes.io/master"
        operator: "Exists"
        effect: "NoSchedule"`)

func manifestsMachineconfigcontrollerDeploymentYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigcontrollerDeploymentYaml, nil
}

func manifestsMachineconfigcontrollerDeploymentYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigcontrollerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigcontroller/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: {{.TargetNamespace}}
  name: machine-config-controller
`)

func manifestsMachineconfigcontrollerSaYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigcontrollerSaYaml, nil
}

func manifestsMachineconfigcontrollerSaYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigcontrollerSaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigcontroller/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: machine-config-daemon
  namespace: {{.TargetNamespace}}
rules:
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["get", "list", "watch", "patch", "update"]
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["*"]
- apiGroups: ["extensions"]
  resources: ["daemonsets"]
  verbs: ["get"]
- apiGroups: [""]
  resources: ["pods/eviction"]
  verbs: ["create"]
- apiGroups: ["machineconfiguration.openshift.io"]
  resources: ["machineconfigs"]
  verbs: ["*"]
`)

func manifestsMachineconfigdaemonClusterroleYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigdaemonClusterroleYaml, nil
}

func manifestsMachineconfigdaemonClusterroleYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigdaemonClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigdaemon/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: machine-config-daemon
  namespace: {{.TargetNamespace}}
roleRef:
  kind: ClusterRole
  name: machine-config-daemon
subjects:
- kind: ServiceAccount
  namespace: {{.TargetNamespace}}
  name: machine-config-daemon
`)

func manifestsMachineconfigdaemonClusterrolebindingYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigdaemonClusterrolebindingYaml, nil
}

func manifestsMachineconfigdaemonClusterrolebindingYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigdaemonClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigdaemon/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonDaemonsetYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: machine-config-daemon
  namespace: {{.TargetNamespace}}
spec:
  selector:
    matchLabels:
      k8s-app: machine-config-daemon
  template:
    metadata:
      name: machine-config-daemon
      labels:
        k8s-app: machine-config-daemon
    spec:
      containers:
      - name: machine-config-daemon
        image: {{.Images.MachineConfigDaemon}}
        args:
          - "start"
        securityContext:
          privileged: true
        volumeMounts:
          - mountPath: /rootfs
            name: rootfs
          - mountPath: /var/run/dbus
            name: var-run-dbus
          - mountPath: /run/systemd
            name: run-systemd
          - mountPath: /etc/ssl/certs
            name: etc-ssl-certs
            readOnly: true
          - mountPath: /etc/machine-config-daemon
            name: etc-mcd
            readOnly: true
        env:
          - name: NODE_NAME
            valueFrom:
              fieldRef:
                fieldPath: spec.nodeName
      hostNetwork: true
      hostPID: true
      serviceAccountName: machine-config-daemon
      tolerations:
        - key: node-role.kubernetes.io/master
          operator: Exists
          effect: NoSchedule
        - key: node-role.kubernetes.io/etcd
          operator: Exists
          effect: NoSchedule
      volumes:
        - name: rootfs
          hostPath:
            path: /
        - name: var-run-dbus
          hostPath:
            path: /var/run/dbus
        - name: run-systemd
          hostPath:
            path: /run/systemd
        - name: etc-ssl-certs
          hostPath:
            path: /etc/ssl/certs
        - name: etc-mcd
          hostPath:
            path: /etc/machine-config-daemon
`)

func manifestsMachineconfigdaemonDaemonsetYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigdaemonDaemonsetYaml, nil
}

func manifestsMachineconfigdaemonDaemonsetYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigdaemonDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigdaemon/daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: {{.TargetNamespace}}
  name: machine-config-daemon
`)

func manifestsMachineconfigdaemonSaYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigdaemonSaYaml, nil
}

func manifestsMachineconfigdaemonSaYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigdaemonSaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigdaemon/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigpoolCrdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  # name must match the spec fields below, and be in the form: <plural>.<group>
  name: machineconfigpools.machineconfiguration.openshift.io
spec:
  # group name to use for REST API: /apis/<group>/<version>
  group: machineconfiguration.openshift.io
  # list of versions supported by this CustomResourceDefinition
  versions:
    - name: v1
      # Each version can be enabled/disabled by Served flag.
      served: true
      # One and only one version must be marked as the storage version.
      storage: true
  # either Namespaced or Cluster
  scope: Cluster
  subresources:
    status: {}
  names:
    # plural name to be used in the URL: /apis/<group>/<version>/<plural>
    plural: machineconfigpools
    # singular name to be used as an alias on the CLI and for display
    singular: machineconfigpool
    # kind is normally the CamelCased singular type. Your resource manifests use this.
    kind: MachineConfigPool
`)

func manifestsMachineconfigpoolCrdYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigpoolCrdYaml, nil
}

func manifestsMachineconfigpoolCrdYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigpoolCrdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigpool.crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: machine-config-server
  namespace: {{.TargetNamespace}}
rules:
- apiGroups: ["machineconfiguration.openshift.io"]
  resources: ["machineconfigs", "machineconfigpools"]
  verbs: ["*"]
`)

func manifestsMachineconfigserverClusterroleYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigserverClusterroleYaml, nil
}

func manifestsMachineconfigserverClusterroleYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigserverClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigserver/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: machine-config-server
  namespace: {{.TargetNamespace}}
roleRef:
  kind: ClusterRole
  name: machine-config-server
subjects:
- kind: ServiceAccount
  namespace: {{.TargetNamespace}}
  name: machine-config-server
`)

func manifestsMachineconfigserverClusterrolebindingYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigserverClusterrolebindingYaml, nil
}

func manifestsMachineconfigserverClusterrolebindingYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigserverClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigserver/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverDaemonsetYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: machine-config-server
  namespace: {{.TargetNamespace}}
spec:
  selector:
    matchLabels:
      k8s-app: machine-config-server
  template:
    metadata:
      name: machine-config-server
      labels:
        k8s-app: machine-config-server
    spec:
      containers:
      - name: machine-config-server
        image: {{.Images.MachineConfigServer}}
        args:
          - "start"
          - "--apiserver-url=https://{{.ControllerConfig.ClusterName}}-api.{{.ControllerConfig.BaseDomain}}:6443"
        volumeMounts:
        - name: certs
          mountPath: /etc/ssl/mcs
        - name: node-bootstrap-token
          mountPath: /etc/mcs/bootstrap-token
      hostNetwork: true
      nodeSelector:
        node-role.kubernetes.io/master: ""
      serviceAccountName: machine-config-server
      tolerations:
        - key: node-role.kubernetes.io/master
          operator: Exists
          effect: NoSchedule
      volumes:
      - name: node-bootstrap-token
        secret:
          secretName: node-bootstrapper-token
      - name: certs
        secret:
          secretName: machine-config-server-tls
`)

func manifestsMachineconfigserverDaemonsetYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigserverDaemonsetYaml, nil
}

func manifestsMachineconfigserverDaemonsetYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigserverDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigserver/daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverNodeBootstrapperSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: {{.TargetNamespace}}
  name: node-bootstrapper
`)

func manifestsMachineconfigserverNodeBootstrapperSaYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigserverNodeBootstrapperSaYaml, nil
}

func manifestsMachineconfigserverNodeBootstrapperSaYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigserverNodeBootstrapperSaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigserver/node-bootstrapper-sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverNodeBootstrapperTokenYaml = []byte(`apiVersion: v1
kind: Secret
metadata:
  annotations:
    kubernetes.io/service-account.name: node-bootstrapper
  name: node-bootstrapper-token
  namespace: {{.TargetNamespace}}
type: kubernetes.io/service-account-token
`)

func manifestsMachineconfigserverNodeBootstrapperTokenYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigserverNodeBootstrapperTokenYaml, nil
}

func manifestsMachineconfigserverNodeBootstrapperTokenYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigserverNodeBootstrapperTokenYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigserver/node-bootstrapper-token.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: {{.TargetNamespace}}
  name: machine-config-server
`)

func manifestsMachineconfigserverSaYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigserverSaYaml, nil
}

func manifestsMachineconfigserverSaYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigserverSaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigserver/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMasterMachineconfigpoolYaml = []byte(`apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfigPool
metadata:
  name: master
spec:
  machineConfigSelector:
    matchLabels:
      "machineconfiguration.openshift.io/role": "master"
  machineSelector:
    matchLabels:
      node-role.kubernetes.io/master: ""`)

func manifestsMasterMachineconfigpoolYamlBytes() ([]byte, error) {
	return _manifestsMasterMachineconfigpoolYaml, nil
}

func manifestsMasterMachineconfigpoolYaml() (*asset, error) {
	bytes, err := manifestsMasterMachineconfigpoolYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/master.machineconfigpool.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsSccYaml = []byte(`apiVersion: security.openshift.io/v1
kind: SecurityContextConstraints
metadata:
  annotations:
    kubernetes.io/description: "privileged-openshift-machine-config-operator for running priviledged in openshift-machine-config-operator namespace."
  name: privileged-openshift-machine-config-operator
allowHostDirVolumePlugin: true
allowHostIPC: true
allowHostNetwork: true
allowHostPID: true
allowHostPorts: true
allowPrivilegedContainer: true
allowedCapabilities:
- "*"
fsGroup:
  type: RunAsAny
groups:
- system:serviceaccounts:{{.TargetNamespace}}
readOnlyRootFilesystem: false
runAsUser:
  type: RunAsAny
seLinuxContext:
  type: RunAsAny
seccompProfiles:
- "*"
supplementalGroups:
  type: RunAsAny
users: []
volumes:
- "*"`)

func manifestsSccYamlBytes() ([]byte, error) {
	return _manifestsSccYaml, nil
}

func manifestsSccYaml() (*asset, error) {
	bytes, err := manifestsSccYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/scc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsWorkerMachineconfigpoolYaml = []byte(`apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfigPool
metadata:
  name: worker
spec:
  machineConfigSelector:
    matchLabels:
      "machineconfiguration.openshift.io/role": "worker"
  machineSelector:
    matchLabels:
      node-role.kubernetes.io/worker: ""`)

func manifestsWorkerMachineconfigpoolYamlBytes() ([]byte, error) {
	return _manifestsWorkerMachineconfigpoolYaml, nil
}

func manifestsWorkerMachineconfigpoolYaml() (*asset, error) {
	bytes, err := manifestsWorkerMachineconfigpoolYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/worker.machineconfigpool.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"manifests/bootstrap-pod.yaml": manifestsBootstrapPodYaml,
	"manifests/controllerconfig.crd.yaml": manifestsControllerconfigCrdYaml,
	"manifests/machineconfig.crd.yaml": manifestsMachineconfigCrdYaml,
	"manifests/machineconfigcontroller/clusterrole.yaml": manifestsMachineconfigcontrollerClusterroleYaml,
	"manifests/machineconfigcontroller/clusterrolebinding.yaml": manifestsMachineconfigcontrollerClusterrolebindingYaml,
	"manifests/machineconfigcontroller/controllerconfig.yaml": manifestsMachineconfigcontrollerControllerconfigYaml,
	"manifests/machineconfigcontroller/deployment.yaml": manifestsMachineconfigcontrollerDeploymentYaml,
	"manifests/machineconfigcontroller/sa.yaml": manifestsMachineconfigcontrollerSaYaml,
	"manifests/machineconfigdaemon/clusterrole.yaml": manifestsMachineconfigdaemonClusterroleYaml,
	"manifests/machineconfigdaemon/clusterrolebinding.yaml": manifestsMachineconfigdaemonClusterrolebindingYaml,
	"manifests/machineconfigdaemon/daemonset.yaml": manifestsMachineconfigdaemonDaemonsetYaml,
	"manifests/machineconfigdaemon/sa.yaml": manifestsMachineconfigdaemonSaYaml,
	"manifests/machineconfigpool.crd.yaml": manifestsMachineconfigpoolCrdYaml,
	"manifests/machineconfigserver/clusterrole.yaml": manifestsMachineconfigserverClusterroleYaml,
	"manifests/machineconfigserver/clusterrolebinding.yaml": manifestsMachineconfigserverClusterrolebindingYaml,
	"manifests/machineconfigserver/daemonset.yaml": manifestsMachineconfigserverDaemonsetYaml,
	"manifests/machineconfigserver/node-bootstrapper-sa.yaml": manifestsMachineconfigserverNodeBootstrapperSaYaml,
	"manifests/machineconfigserver/node-bootstrapper-token.yaml": manifestsMachineconfigserverNodeBootstrapperTokenYaml,
	"manifests/machineconfigserver/sa.yaml": manifestsMachineconfigserverSaYaml,
	"manifests/master.machineconfigpool.yaml": manifestsMasterMachineconfigpoolYaml,
	"manifests/scc.yaml": manifestsSccYaml,
	"manifests/worker.machineconfigpool.yaml": manifestsWorkerMachineconfigpoolYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"manifests": &bintree{nil, map[string]*bintree{
		"bootstrap-pod.yaml": &bintree{manifestsBootstrapPodYaml, map[string]*bintree{}},
		"controllerconfig.crd.yaml": &bintree{manifestsControllerconfigCrdYaml, map[string]*bintree{}},
		"machineconfig.crd.yaml": &bintree{manifestsMachineconfigCrdYaml, map[string]*bintree{}},
		"machineconfigcontroller": &bintree{nil, map[string]*bintree{
			"clusterrole.yaml": &bintree{manifestsMachineconfigcontrollerClusterroleYaml, map[string]*bintree{}},
			"clusterrolebinding.yaml": &bintree{manifestsMachineconfigcontrollerClusterrolebindingYaml, map[string]*bintree{}},
			"controllerconfig.yaml": &bintree{manifestsMachineconfigcontrollerControllerconfigYaml, map[string]*bintree{}},
			"deployment.yaml": &bintree{manifestsMachineconfigcontrollerDeploymentYaml, map[string]*bintree{}},
			"sa.yaml": &bintree{manifestsMachineconfigcontrollerSaYaml, map[string]*bintree{}},
		}},
		"machineconfigdaemon": &bintree{nil, map[string]*bintree{
			"clusterrole.yaml": &bintree{manifestsMachineconfigdaemonClusterroleYaml, map[string]*bintree{}},
			"clusterrolebinding.yaml": &bintree{manifestsMachineconfigdaemonClusterrolebindingYaml, map[string]*bintree{}},
			"daemonset.yaml": &bintree{manifestsMachineconfigdaemonDaemonsetYaml, map[string]*bintree{}},
			"sa.yaml": &bintree{manifestsMachineconfigdaemonSaYaml, map[string]*bintree{}},
		}},
		"machineconfigpool.crd.yaml": &bintree{manifestsMachineconfigpoolCrdYaml, map[string]*bintree{}},
		"machineconfigserver": &bintree{nil, map[string]*bintree{
			"clusterrole.yaml": &bintree{manifestsMachineconfigserverClusterroleYaml, map[string]*bintree{}},
			"clusterrolebinding.yaml": &bintree{manifestsMachineconfigserverClusterrolebindingYaml, map[string]*bintree{}},
			"daemonset.yaml": &bintree{manifestsMachineconfigserverDaemonsetYaml, map[string]*bintree{}},
			"node-bootstrapper-sa.yaml": &bintree{manifestsMachineconfigserverNodeBootstrapperSaYaml, map[string]*bintree{}},
			"node-bootstrapper-token.yaml": &bintree{manifestsMachineconfigserverNodeBootstrapperTokenYaml, map[string]*bintree{}},
			"sa.yaml": &bintree{manifestsMachineconfigserverSaYaml, map[string]*bintree{}},
		}},
		"master.machineconfigpool.yaml": &bintree{manifestsMasterMachineconfigpoolYaml, map[string]*bintree{}},
		"scc.yaml": &bintree{manifestsSccYaml, map[string]*bintree{}},
		"worker.machineconfigpool.yaml": &bintree{manifestsWorkerMachineconfigpoolYaml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

