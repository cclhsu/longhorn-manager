package client

const (
	SETTING_TYPE = "setting"
)

type Setting struct {
	Resource `yaml:"-"`

	Definition SettingDefinition `json:"definition,omitempty" yaml:"definition,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}

type SettingCollection struct {
	Collection
	Data   []Setting `json:"data,omitempty"`
	client *SettingClient
}

type SettingClient struct {
	rancherClient *RancherClient
}

type SettingOperations interface {
	List(opts *ListOpts) (*SettingCollection, error)
	Create(opts *Setting) (*Setting, error)
	Update(existing *Setting, updates interface{}) (*Setting, error)
	ById(id string) (*Setting, error)
	Delete(container *Setting) error
}

func newSettingClient(rancherClient *RancherClient) *SettingClient {
	return &SettingClient{
		rancherClient: rancherClient,
	}
}

func (c *SettingClient) Create(container *Setting) (*Setting, error) {
	resp := &Setting{}
	err := c.rancherClient.doCreate(SETTING_TYPE, container, resp)
	return resp, err
}

func (c *SettingClient) Update(existing *Setting, updates interface{}) (*Setting, error) {
	resp := &Setting{}
	err := c.rancherClient.doUpdate(SETTING_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SettingClient) List(opts *ListOpts) (*SettingCollection, error) {
	resp := &SettingCollection{}
	err := c.rancherClient.doList(SETTING_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *SettingCollection) Next() (*SettingCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &SettingCollection{}
		err := cc.client.rancherClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *SettingClient) ById(id string) (*Setting, error) {
	resp := &Setting{}
	err := c.rancherClient.doById(SETTING_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *SettingClient) Delete(container *Setting) error {
	return c.rancherClient.doResourceDelete(SETTING_TYPE, &container.Resource)
}
