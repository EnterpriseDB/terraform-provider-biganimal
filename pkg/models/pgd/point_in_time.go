package pgd

import (
	"encoding/json"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
)

type PointInTime string

// UnmarshalJSON to implement json.Unmarshaler for custom unmarshalling
func (p *PointInTime) UnmarshalJSON(d []byte) error {
	var pointIntTimeAPI models.PointInTime
	if err := json.Unmarshal(d, &pointIntTimeAPI); err != nil {
		return err
	}
	stringTime := time.Unix(int64(pointIntTimeAPI.Seconds), int64(pointIntTimeAPI.Nanos)).String()
	pointIntTimeTF := PointInTime(stringTime)
	*p = pointIntTimeTF
	return nil
}
