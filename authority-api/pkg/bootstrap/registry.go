package bootstrap

import (
	"authority-api/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"path/filepath"

	// etcd
	etcdKratos "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	etcdClient "go.etcd.io/etcd/client/v3"

	// consul
	consulKratos "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	consulClient "github.com/hashicorp/consul/api"

	// kubernetes
	k8sRegistry "github.com/go-kratos/kratos/contrib/registry/kubernetes/v2"
	k8s "k8s.io/client-go/kubernetes"
	k8sRest "k8s.io/client-go/rest"
	k8sTools "k8s.io/client-go/tools/clientcmd"
	k8sUtil "k8s.io/client-go/util/homedir"
)

type RegistryType string

const (
	RegistryTypeConsul     RegistryType = "consul"
	RegistryTypeEtcd       RegistryType = "etcd"
	RegistryTypeKubernetes RegistryType = "kubernetes"
)

// NewRegistry 创建一个注册客户端
func NewRegistry(cfg *conf.Registry) registry.Registrar {
	if cfg == nil {
		return nil
	}

	switch RegistryType(cfg.Type) {
	case RegistryTypeConsul:
		return NewConsulRegistry(cfg)
	case RegistryTypeEtcd:
		return NewEtcdRegistry(cfg)
	case RegistryTypeKubernetes:
		return NewKubernetesRegistry(cfg)
	}

	return nil
}

// NewConsulRegistry 创建一个注册发现客户端 - Consul
func NewConsulRegistry(c *conf.Registry) *consulKratos.Registry {
	cfg := consulClient.DefaultConfig()
	cfg.Address = c.Consul.Address
	cfg.Scheme = c.Consul.Scheme

	var cli *consulClient.Client
	var err error
	if cli, err = consulClient.NewClient(cfg); err != nil {
		log.Fatal(err)
	}

	reg := consulKratos.New(cli, consulKratos.WithHealthCheck(c.Consul.HealthCheck))

	return reg
}

// NewEtcdRegistry 创建一个注册发现客户端 - Etcd
func NewEtcdRegistry(c *conf.Registry) *etcdKratos.Registry {
	cfg := etcdClient.Config{
		Endpoints: c.Etcd.Endpoints,
	}

	var err error
	var cli *etcdClient.Client
	if cli, err = etcdClient.New(cfg); err != nil {
		panic(err)
	}

	reg := etcdKratos.New(cli)

	return reg
}

// NewKubernetesRegistry 创建一个注册发现客户端 - Kubernetes
func NewKubernetesRegistry(_ *conf.Registry) *k8sRegistry.Registry {
	restConfig, err := k8sRest.InClusterConfig()
	if err != nil {
		home := k8sUtil.HomeDir()
		kubeConfig := filepath.Join(home, ".kube", "config")
		restConfig, err = k8sTools.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			log.Fatal(err)
			return nil
		}
	}

	clientSet, err := k8s.NewForConfig(restConfig)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	reg := k8sRegistry.NewRegistry(clientSet)

	return reg
}

// NewDiscovery 创建一个发现客户端
func NewDiscovery(cfg *conf.Registry) registry.Discovery {
	if cfg == nil {
		return nil
	}

	switch RegistryType(cfg.Type) {
	case RegistryTypeConsul:
		return NewConsulRegistry(cfg)
	case RegistryTypeEtcd:
		return NewEtcdRegistry(cfg)
	case RegistryTypeKubernetes:
		return NewKubernetesRegistry(cfg)
	}

	return nil
}
