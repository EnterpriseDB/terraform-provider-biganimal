package api

type ClientType int

// type APIClient interface {
// 	Create(ctx context.Context, model any) (string, error)
// 	Read(ctx context.Context, id string) (any, error)
// 	Update(ctx context.Context, model any, id string) error
// 	Delete(ctx context.Context, id string) error
// }

const (
	ClusterClientType ClientType = 1
	RegionClientType             = iota
)
