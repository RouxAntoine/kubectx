package main

import (
	"fmt"
	"io"

	"github.com/pkg/errors"

	"github.com/ahmetb/kubectx/internal/kubeconfig"
)

// UnsetOp indicates intention to remove current-context preference.
type UnsetOp struct{}

func (_ UnsetOp) Run(_, stderr io.Writer) error {
	kc := new(kubeconfig.Kubeconfig).WithLoader(defaultLoader)
	defer kc.Close()
	if err := kc.Parse(); err != nil {
		return errors.Wrap(err, "kubeconfig error")
	}

	if err := kc.UnsetCurrentContext(); err != nil {
		return errors.Wrap(err, "error while modifying current-context")
	}
	if err := kc.Save(); err != nil {
		return errors.Wrap(err, "failed to save kubeconfig file after modification")
	}

	_, err := fmt.Fprintln(stderr, "Successfully unset the active context for kubectl.")
	return errors.Wrap(err, "write error")
}