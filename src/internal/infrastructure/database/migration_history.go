package database

import "gorm.io/gorm"

type MigrationHistory struct {
	gorm.Model
	Version   string `gorm:"uniqueIndex"`
	Name      string
	AppliedAt int64 `gorm:"autoCreateTime"`
}

func (m *Migrator) checkIfMigrated(version string) (bool, error) {
	var count int64
	err := m.db.Model(&MigrationHistory{}).
		Where("version = ?", version).
		Count(&count).Error

	return count > 0, err
}

func (m *Migrator) recordMigration(version, name string) error {
	return m.db.Create(&MigrationHistory{
		Version: version,
		Name:    name,
	}).Error
}
