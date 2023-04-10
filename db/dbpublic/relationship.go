package dbpublic

// User-Role data
type UserRole struct {
	UserID string `gorm:"column:user_id;type:uuid NOT NULL;primaryKey;"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleID string `gorm:"column:role_id;type:uuid NOT NULL;primaryKey;"`
	Role   Role   `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (UserRole) TableName() string {
	return "user_role"
}

// User-Role-User data
type UserRoleUser struct {
	UserID          string `gorm:"column:user_id;type:uuid NOT NULL;primaryKey;"`
	User            User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleID          string `gorm:"column:role_id;type:uuid NOT NULL;primaryKey;"`
	Role            Role   `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DependentUserID string `gorm:"column:dependent_user_id;type:uuid NOT NULL;primaryKey;"`
	DependentUser   User   `gorm:"foreignKey:DependentUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (UserRoleUser) TableName() string {
	return "user_role_user"
}

// User-Role-User data
type UserRoleOrganization struct {
	UserID         string `gorm:"column:user_id;type:uuid NOT NULL;primaryKey;"`
	User           User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleID         string `gorm:"column:role_id;type:uuid NOT NULL;primaryKey;"`
	Role           Role   `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrganizationID string `gorm:"column:organization_id;type:uuid NOT NULL;primaryKey;"`
	Organization   User   `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (UserRoleOrganization) TableName() string {
	return "user_role_organization"
}
