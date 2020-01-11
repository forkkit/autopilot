package codegen

import (
	"context"

	"github.com/solo-io/anyvendor/pkg/manager"
	"github.com/solo-io/autopilot/codegen/ap_anyvendor"
	"github.com/solo-io/autopilot/codegen/model"
	"github.com/solo-io/autopilot/codegen/proto"
	"github.com/solo-io/autopilot/codegen/render"
	"github.com/solo-io/autopilot/codegen/util"
	"github.com/solo-io/autopilot/codegen/writer"
)

// runs the codegen compilation for the current Go module
type Command struct {
	// the name of the app or component
	// used to label k8s manifests
	AppName string

	// search protos recursively starting from this directory
	// if left empty will default to vendor_any
	ProtoDir string

	// settings to configure anyvendor for easier proto imports
	AnyVendorConfig *ap_anyvendor.Imports

	// the k8s api groups for which to compile
	Groups []render.Group

	// the root directory for generated API code
	ApiRoot string

	// the root directory for generated Kube manfiests
	ManifestRoot string

	// Should we compile protos?
	RenderProtos bool

	// the go module of the project
	// set by Execute()
	goModule string

	// the path to the root dir of the module on disk
	moduleRoot string
}

// function to execute Autopilot code gen from another repository
func (c Command) Execute() error {
	c.goModule = util.GetGoModule()
	c.moduleRoot = util.GetModuleRoot()
	if err := c.compileProtos(); err != nil {
		return err
	}
	for _, group := range c.Groups {
		// init connects children to their parents
		group.Init()

		if err := c.writeGeneratedFiles(group); err != nil {
			return err
		}
	}
	return nil
}

func (c Command) writeGeneratedFiles(grp model.Group) error {
	writer := &writer.DefaultFileWriter{Root: c.moduleRoot}

	apiTypes, err := render.RenderApiTypes(c.goModule, c.ApiRoot, grp)
	if err != nil {
		return err
	}

	if err := writer.WriteFiles(apiTypes); err != nil {
		return err
	}

	manifests, err := render.RenderManifests(c.AppName, c.ManifestRoot, grp)
	if err != nil {
		return err
	}

	if err := writer.WriteFiles(manifests); err != nil {
		return err
	}

	if err := render.KubeCodegen(c.ApiRoot, grp); err != nil {
		return err
	}

	return nil
}

func (c Command) compileProtos() error {
	if !c.RenderProtos {
		return nil
	}

	mgr, err := manager.NewManager(context.TODO(), c.moduleRoot)
	if err != nil {
		return err
	}

	if err := mgr.Ensure(context.TODO(), c.AnyVendorConfig.ToAnyvendorConfig()); err != nil {
		return err
	}

	if err := proto.CompileProtos(
		c.goModule,
		c.ApiRoot,
		c.ProtoDir,
	); err != nil {
		return err
	}

	return nil
}
