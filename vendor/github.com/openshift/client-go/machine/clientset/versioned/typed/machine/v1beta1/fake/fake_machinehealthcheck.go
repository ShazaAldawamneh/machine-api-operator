// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1beta1 "github.com/openshift/api/machine/v1beta1"
	machinev1beta1 "github.com/openshift/client-go/machine/applyconfigurations/machine/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMachineHealthChecks implements MachineHealthCheckInterface
type FakeMachineHealthChecks struct {
	Fake *FakeMachineV1beta1
	ns   string
}

var machinehealthchecksResource = schema.GroupVersionResource{Group: "machine.openshift.io", Version: "v1beta1", Resource: "machinehealthchecks"}

var machinehealthchecksKind = schema.GroupVersionKind{Group: "machine.openshift.io", Version: "v1beta1", Kind: "MachineHealthCheck"}

// Get takes name of the machineHealthCheck, and returns the corresponding machineHealthCheck object, and an error if there is any.
func (c *FakeMachineHealthChecks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.MachineHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(machinehealthchecksResource, c.ns, name), &v1beta1.MachineHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MachineHealthCheck), err
}

// List takes label and field selectors, and returns the list of MachineHealthChecks that match those selectors.
func (c *FakeMachineHealthChecks) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.MachineHealthCheckList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(machinehealthchecksResource, machinehealthchecksKind, c.ns, opts), &v1beta1.MachineHealthCheckList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.MachineHealthCheckList{ListMeta: obj.(*v1beta1.MachineHealthCheckList).ListMeta}
	for _, item := range obj.(*v1beta1.MachineHealthCheckList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machineHealthChecks.
func (c *FakeMachineHealthChecks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(machinehealthchecksResource, c.ns, opts))

}

// Create takes the representation of a machineHealthCheck and creates it.  Returns the server's representation of the machineHealthCheck, and an error, if there is any.
func (c *FakeMachineHealthChecks) Create(ctx context.Context, machineHealthCheck *v1beta1.MachineHealthCheck, opts v1.CreateOptions) (result *v1beta1.MachineHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(machinehealthchecksResource, c.ns, machineHealthCheck), &v1beta1.MachineHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MachineHealthCheck), err
}

// Update takes the representation of a machineHealthCheck and updates it. Returns the server's representation of the machineHealthCheck, and an error, if there is any.
func (c *FakeMachineHealthChecks) Update(ctx context.Context, machineHealthCheck *v1beta1.MachineHealthCheck, opts v1.UpdateOptions) (result *v1beta1.MachineHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(machinehealthchecksResource, c.ns, machineHealthCheck), &v1beta1.MachineHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MachineHealthCheck), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMachineHealthChecks) UpdateStatus(ctx context.Context, machineHealthCheck *v1beta1.MachineHealthCheck, opts v1.UpdateOptions) (*v1beta1.MachineHealthCheck, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(machinehealthchecksResource, "status", c.ns, machineHealthCheck), &v1beta1.MachineHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MachineHealthCheck), err
}

// Delete takes name of the machineHealthCheck and deletes it. Returns an error if one occurs.
func (c *FakeMachineHealthChecks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(machinehealthchecksResource, c.ns, name, opts), &v1beta1.MachineHealthCheck{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineHealthChecks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(machinehealthchecksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.MachineHealthCheckList{})
	return err
}

// Patch applies the patch and returns the patched machineHealthCheck.
func (c *FakeMachineHealthChecks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MachineHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(machinehealthchecksResource, c.ns, name, pt, data, subresources...), &v1beta1.MachineHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MachineHealthCheck), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied machineHealthCheck.
func (c *FakeMachineHealthChecks) Apply(ctx context.Context, machineHealthCheck *machinev1beta1.MachineHealthCheckApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.MachineHealthCheck, err error) {
	if machineHealthCheck == nil {
		return nil, fmt.Errorf("machineHealthCheck provided to Apply must not be nil")
	}
	data, err := json.Marshal(machineHealthCheck)
	if err != nil {
		return nil, err
	}
	name := machineHealthCheck.Name
	if name == nil {
		return nil, fmt.Errorf("machineHealthCheck.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(machinehealthchecksResource, c.ns, *name, types.ApplyPatchType, data), &v1beta1.MachineHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MachineHealthCheck), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeMachineHealthChecks) ApplyStatus(ctx context.Context, machineHealthCheck *machinev1beta1.MachineHealthCheckApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.MachineHealthCheck, err error) {
	if machineHealthCheck == nil {
		return nil, fmt.Errorf("machineHealthCheck provided to Apply must not be nil")
	}
	data, err := json.Marshal(machineHealthCheck)
	if err != nil {
		return nil, err
	}
	name := machineHealthCheck.Name
	if name == nil {
		return nil, fmt.Errorf("machineHealthCheck.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(machinehealthchecksResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1beta1.MachineHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MachineHealthCheck), err
}