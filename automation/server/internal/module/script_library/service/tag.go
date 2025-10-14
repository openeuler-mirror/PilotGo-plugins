package service

import (
	"strings"
	"time"

	"github.com/mozillazg/go-pinyin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/dao"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
)

func GetTags() ([]model.Tag, error) {
	return dao.GetTags()
}

func CreateTag(tag *model.Tag) error {
	return dao.CreateTag(&model.Tag{
		Name:        tag.Name,
		EnName:      tagChineseToEnglish(tag.Name),
		Description: tag.Description,
		ModifyUser:  tag.ModifyUser,
		ModifyTime:  time.Now().Format("2006-01-02 15:04:05"),
	})
}

func UpdateTag(tag *model.Tag) error {
	return dao.UpdateTag(&model.Tag{
		ID:          tag.ID,
		Name:        tag.Name,
		EnName:      tagChineseToEnglish(tag.Name),
		Description: tag.Description,
		ModifyUser:  tag.ModifyUser,
		ModifyTime:  time.Now().Format("2006-01-02 15:04:05"),
	})
}

func DeleteTag(id int) error {
	return dao.DeleteTag(id)
}

func tagChineseToEnglish(s string) string {
	args := pinyin.NewArgs()
	pys := pinyin.Pinyin(s, args)
	var result string
	for _, py := range pys {
		if len(py) > 0 {
			result += py[0]
		}
	}
	return strings.ToLower(result)
}
