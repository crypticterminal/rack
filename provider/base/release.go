package aws

import (
	"fmt"

	"github.com/convox/rack/structs"
)

func (p *Provider) ReleaseCreate(app string, opts structs.ReleaseCreateOptions) (*structs.Release, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (p *Provider) ReleaseGet(app, id string) (*structs.Release, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (p *Provider) ReleaseList(app string, opts structs.ReleaseListOptions) (structs.Releases, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (p *Provider) ReleasePromote(app, id string) error {
	return fmt.Errorf("unimplemented")
}
