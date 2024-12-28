package kubeconfig

import (
	"github.com/ahmetb/kubectx/internal/kubeconfig/loader"
	"github.com/ahmetb/kubectx/internal/kubeconfig/single"
)

type CurrentKubeconfig interface {
	WithLoader(l loader.Loader) *single.Kubeconfig
	Close() error
	Parse() error
	GetCurrentContext() string
}

type DeleteKubeconfig interface {
	WithLoader(l loader.Loader) *single.Kubeconfig
	Close() error
	Parse() error
	GetCurrentContext() string
	ContextExists(name string) bool
	DeleteContextEntry(deleteName string) error
	Save() error
}

type SwitchKubeconfig interface {
	WithLoader(l loader.Loader) *single.Kubeconfig
	Close() error
	Parse() error
	Save() error
	GetCurrentContext() string
	ContextExists(name string) bool
	ModifyCurrentContext(name string) error
}

type CurrentKubeconfigForNamespace interface {
	WithLoader(l loader.Loader) *single.Kubeconfig
	Close() error
	Parse() error
	GetCurrentContext() string
	NamespaceOfContext(contextName string) (string, error)
}

type ListKubeconfigForNamespace interface {
	ByteKubeconfig
	WithLoader(l loader.Loader) *single.Kubeconfig
	Close() error
	Parse() error
	GetCurrentContext() string
	NamespaceOfContext(contextName string) (string, error)
}

type SwitchKubeconfigForNamespace interface {
	ByteKubeconfig
	WithLoader(l loader.Loader) *single.Kubeconfig
	Close() error
	Parse() error
	GetCurrentContext() string
	NamespaceOfContext(contextName string) (string, error)
	SetNamespace(ctxName string, ns string) error
	Save() error
}

type ByteKubeconfig interface {
	Bytes() ([]byte, error)
}
