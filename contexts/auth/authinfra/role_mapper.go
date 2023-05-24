package authinfra

import (
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbpublic"
)

func FromRoleModelToAliasRoleEntity(model dbpublic.Role) (authdomain.AliasRoleEntity, error) {
	alias, aliasErr := authdomain.CreateRoleAlias(model.Alias)
	if aliasErr != nil {
		return authdomain.AliasRoleEntity{}, aliasErr
	}

	name, nameErr := authdomain.CreateRoleName(model.Name)
	if nameErr != nil {
		return authdomain.AliasRoleEntity{}, nameErr
	}

	scope, scopeErr := authdomain.CreateRoleScope(model.Scope)
	if scopeErr != nil {
		return authdomain.AliasRoleEntity{}, scopeErr
	}

	state, stateErr := authdomain.CreateRoleState(model.State)
	if stateErr != nil {
		return authdomain.AliasRoleEntity{}, stateErr
	}

	return authdomain.AliasRoleEntity{
		Alias: alias,
		RoleEntity: authdomain.RoleEntity{
			ID:    model.ID,
			Name:  name,
			Scope: scope,
			State: state,
		},
	}, nil
}
