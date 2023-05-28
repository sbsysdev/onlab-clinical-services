package dbshared

// Operation data
type Operation struct {
	ID      string `gorm:"column:operation_id;type:uuid;not null;unique;primaryKey;default:gen_random_uuid()"`
	UseCase string `gorm:"column:use_case;type:VARCHAR(64);not null;unique"`
	State   string `gorm:"type:public.COMMON_STATE_ENUM;not null;default:'inactive'"`

	Permissions []*Permission `gorm:"many2many:permission_operation"`
}

// Permission data
type Permission struct {
	ID      string        `gorm:"column:permission_id;type:uuid;not null;unique;primaryKey;default:gen_random_uuid()"`
	UseCase string        `gorm:"column:use_case;type:VARCHAR(64);not null;unique"`
	Name    MultiLanguage `gorm:"not null;unique"`
	Scope   string        `gorm:"type:public.COMMON_SCOPE_ENUM;not null;default:'system'"`
	State   string        `gorm:"type:public.COMMON_STATE_ENUM;not null;default:'inactive'"`

	Operations []*Operation `gorm:"many2many:permission_operation"`
	Roles      []*Role      `gorm:"many2many:role_permission"`
}

// Permission-Operation data
type PermissionOperation struct {
	PermissionID string     `gorm:"column:permission_id;type:uuid;not null;primaryKey"`
	Permission   Permission `gorm:"foreignKey:PermissionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	OperationID  string     `gorm:"column:operation_id;type:uuid;not null;primaryKey"`
	Operation    Operation  `gorm:"foreignKey:OperationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (PermissionOperation) TableName() string {
	return "permission_operation"
}

// Role data
type Role struct {
	ID    string        `gorm:"column:role_id;type:uuid;not null;unique;primaryKey;default:gen_random_uuid()"`
	Name  MultiLanguage `gorm:"not null;unique"`
	Scope string        `gorm:"type:public.COMMON_SCOPE_ENUM;not null;default:'system'"`
	State string        `gorm:"type:public.COMMON_STATE_ENUM;not null;default:'inactive'"`
	Time  TimeAt        `gorm:"embedded"`

	Permissions []*Permission `gorm:"many2many:role_permission"`
}

// Role-Permission data
type RolePermission struct {
	RoleID       string     `gorm:"column:role_id;type:uuid;not null;primaryKey"`
	Role         Role       `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PermissionID string     `gorm:"column:permission_id;type:uuid;not null;primaryKey"`
	Permission   Permission `gorm:"foreignKey:PermissionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (RolePermission) TableName() string {
	return "role_permission"
}
