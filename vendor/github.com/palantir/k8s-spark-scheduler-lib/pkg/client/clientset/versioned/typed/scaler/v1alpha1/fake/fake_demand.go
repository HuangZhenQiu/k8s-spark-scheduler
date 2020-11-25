// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/palantir/k8s-spark-scheduler-lib/pkg/apis/scaler/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDemands implements DemandInterface
type FakeDemands struct {
	Fake *FakeScalerV1alpha1
	ns   string
}

var demandsResource = schema.GroupVersionResource{Group: "scaler.palantir.com", Version: "v1alpha1", Resource: "demands"}

var demandsKind = schema.GroupVersionKind{Group: "scaler.palantir.com", Version: "v1alpha1", Kind: "Demand"}

// Get takes name of the demand, and returns the corresponding demand object, and an error if there is any.
func (c *FakeDemands) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Demand, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(demandsResource, c.ns, name), &v1alpha1.Demand{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Demand), err
}

// List takes label and field selectors, and returns the list of Demands that match those selectors.
func (c *FakeDemands) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DemandList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(demandsResource, demandsKind, c.ns, opts), &v1alpha1.DemandList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DemandList{ListMeta: obj.(*v1alpha1.DemandList).ListMeta}
	for _, item := range obj.(*v1alpha1.DemandList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested demands.
func (c *FakeDemands) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(demandsResource, c.ns, opts))

}

// Create takes the representation of a demand and creates it.  Returns the server's representation of the demand, and an error, if there is any.
func (c *FakeDemands) Create(ctx context.Context, demand *v1alpha1.Demand, opts v1.CreateOptions) (result *v1alpha1.Demand, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(demandsResource, c.ns, demand), &v1alpha1.Demand{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Demand), err
}

// Update takes the representation of a demand and updates it. Returns the server's representation of the demand, and an error, if there is any.
func (c *FakeDemands) Update(ctx context.Context, demand *v1alpha1.Demand, opts v1.UpdateOptions) (result *v1alpha1.Demand, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(demandsResource, c.ns, demand), &v1alpha1.Demand{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Demand), err
}

// Delete takes name of the demand and deletes it. Returns an error if one occurs.
func (c *FakeDemands) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(demandsResource, c.ns, name), &v1alpha1.Demand{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDemands) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(demandsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DemandList{})
	return err
}

// Patch applies the patch and returns the patched demand.
func (c *FakeDemands) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Demand, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(demandsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Demand{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Demand), err
}
