package workflow

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/common/enum/common"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/pkg/response"
)

type PublishStatus int

const (
	Develop   PublishStatus = 1
	Published PublishStatus = 2
)

var PublishStatusMap = common.EnumMap{
	int(Develop):   "开发中",
	int(Published): "已发布",
}

func (p PublishStatus) String() string {
	return PublishStatusMap.String(int(p))
}
func ParsePublishStatus(s string) PublishStatus {
	for k, v := range PublishStatusMap {
		if v == s {
			return PublishStatus(k)
		}
	}
	return 0
}

func (p PublishStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(PublishStatusMap[int(p)])
}

func (p *PublishStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	for k, v := range PublishStatusMap {
		if v == s {
			*p = PublishStatus(k)
			return nil
		}
	}
	return fmt.Errorf("invalid PublishStatus: '%s', allowed: %v", s, PublishStatusMap)
}

func (p PublishStatus) IsValid() bool {
	_, exists := PublishStatusMap[int(p)]
	return exists
}

func (p PublishStatus) Value() (driver.Value, error) {
	if !p.IsValid() {
		return nil, fmt.Errorf("invalid PublishStatus: %d, allowed: %v", int(p), PublishStatusMap)
	}
	return int64(p), nil
}

func (p *PublishStatus) Scan(value interface{}) error {
	if value == nil {
		*p = 0
		return nil
	}
	if v, ok := value.(int64); ok {
		status := PublishStatus(int(v))
		if !status.IsValid() {
			return fmt.Errorf("invalid PublishStatus: %d, allowed: %v", int(v), PublishStatusMap)
		}
		*p = status
		return nil
	}
	return nil
}

func getPublishStatus() []common.Item {
	return PublishStatusMap.ToItems()
}

func WorkflowPublishStatusListHandler(c *gin.Context) {
	status := getPublishStatus()
	response.Success(c, status, "success")
}
