package authinfra

import (
	"gorm.io/gorm"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbpublic"
)

type RoleRepository struct {
	DB *gorm.DB
}

func (repo RoleRepository) GetAliasRoleModelsByAlias(aliases []authdomain.RoleAlias) ([]dbpublic.Role, error) {
	var founds []dbpublic.Role

	if err := repo.DB.Exec("SET search_path=public;").Error; err != nil {
		return []dbpublic.Role{}, err
	}

	if err := repo.DB.Table("roles").Where("alias IN ?", aliases).Find(&founds).Error; err != nil {
		return []dbpublic.Role{}, err
	}

	return founds, nil
}

func (repo RoleRepository) GetAliasRolesByAlias(aliases []authdomain.RoleAlias) ([]authdomain.AliasRoleEntity, error) {
	founds, foundErr := repo.GetAliasRoleModelsByAlias(aliases)

	if foundErr != nil {
		return []authdomain.AliasRoleEntity{}, foundErr
	}

	roles := make([]authdomain.AliasRoleEntity, len(founds))

	for i, v := range founds {
		role, roleErr := FromRoleModelToAliasRoleEntity(v)

		if roleErr != nil {
			return []authdomain.AliasRoleEntity{}, roleErr
		}

		roles[i] = role
	}

	return roles, nil
}
