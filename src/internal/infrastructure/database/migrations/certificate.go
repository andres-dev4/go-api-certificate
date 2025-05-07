package migrations

import "gorm.io/gorm"

func MigrateCertificatesTable(db *gorm.DB) error {
    type Certificate struct {
        ID          string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
        UserID      string `gorm:"type:uuid;not null"`
        Title       string `gorm:"size:255;not null"`
        Description string `gorm:"type:text"`
        IssueDate   string `gorm:"type:date;not null"`
        ExpiryDate  string `gorm:"type:date"`
        Credential  string `gorm:"type:text;not null"`
        CreatedAt   int64  `gorm:"autoCreateTime"`
        UpdatedAt   int64  `gorm:"autoUpdateTime"`
    }
    
    return db.AutoMigrate(&Certificate{})
}