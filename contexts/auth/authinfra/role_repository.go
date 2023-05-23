package authinfra

import (
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func (repo RoleRepository) GetSystemRolesByAlias(aliases []authdomain.RoleAlias) ([]authdomain.SystemRoleEntity, error) {
	return []authdomain.SystemRoleEntity{}, nil
}
