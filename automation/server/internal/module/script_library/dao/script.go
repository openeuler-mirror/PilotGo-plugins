package dao

import (
	"gorm.io/gorm"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library/model"
)

func AddScript(script *model.Script, scriptVersion *model.ScriptVersion) error {
	return global.App.MySQL.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(script).Error; err != nil {
			return err
		}

		if err := tx.Save(scriptVersion).Error; err != nil {
			return err
		}

		return nil
	})
}
