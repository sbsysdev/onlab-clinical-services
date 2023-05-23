package authdomain

import (
	"errors"

	"github.com/google/uuid"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
)

// Role Alias Value Object
type RoleAlias string

const (
	ALIAS_PATIENT      RoleAlias = "patient"
	ALIAS_OWNER        RoleAlias = "owner"
	ALIAS_COLLABORATOR RoleAlias = "collaborator"
)

const (
	ERRORS_ROLE_ALIAS_EMPTY     shareddomain.DomainError = "ERRORS_ROLE_ALIAS_EMPTY"
	ERRORS_ROLE_ALIAS_NOT_VALID shareddomain.DomainError = "ERRORS_ROLE_ALIAS_NOT_VALID"
)

func CreateRoleAlias(alias string) (RoleAlias, error) {
	if len(alias) == 0 {
		return RoleAlias(""), errors.New(string(ERRORS_ROLE_ALIAS_EMPTY))
	}

	if alias != string(ALIAS_PATIENT) && alias != string(ALIAS_OWNER) && alias != string(ALIAS_COLLABORATOR) {
		return RoleAlias(""), errors.New(string(ERRORS_ROLE_ALIAS_NOT_VALID))
	}

	return RoleAlias(alias), nil
}

// Role Name Value Object
type RoleName map[string]string

const (
	ERRORS_ROLE_NAME_EMPTY         shareddomain.DomainError = "ERRORS_ROLE_NAME_EMPTY"
	ERRORS_ROLE_KEY_NAME_NOT_VALID shareddomain.DomainError = "ERRORS_ROLE_KEY_NAME_NOT_VALID"
	ERRORS_ROLE_VALUE_NAME_EMPTY   shareddomain.DomainError = "ERRORS_ROLE_VALUE_NAME_EMPTY"
)

func CreateRoleName(names map[string]string) (RoleName, error) {
	if len(names) == 0 {
		return RoleName{}, errors.New(string(ERRORS_ROLE_NAME_EMPTY))
	}

	for key, value := range names {
		if len(key) != 2 {
			return RoleName{}, errors.New(string(ERRORS_ROLE_KEY_NAME_NOT_VALID))
		}

		if len(value) == 0 {
			return RoleName{}, errors.New(string(ERRORS_ROLE_VALUE_NAME_EMPTY))
		}
	}

	return names, nil
}

// Role Scope Value Object
type RoleScope string

const (
	SCOPE_SYSTEM RoleScope = "system"
	SCOPE_USER   RoleScope = "user"
	SCOPE_ORG    RoleScope = "org"
	SCOPE_BRANCH RoleScope = "branch"
)

const (
	ERRORS_ROLE_SCOPE_EMPTY     shareddomain.DomainError = "ERRORS_ROLE_SCOPE_EMPTY"
	ERRORS_ROLE_SCOPE_NOT_VALID shareddomain.DomainError = "ERRORS_ROLE_SCOPE_NOT_VALID"
)

func CreateRoleScope(scope string) (RoleScope, error) {
	if len(scope) == 0 {
		return RoleScope(""), errors.New(string(ERRORS_ROLE_SCOPE_EMPTY))
	}

	if scope != string(SCOPE_SYSTEM) && scope != string(SCOPE_USER) && scope != string(SCOPE_ORG) && scope != string(SCOPE_BRANCH) {
		return RoleScope(""), errors.New(string(ERRORS_ROLE_SCOPE_NOT_VALID))
	}

	return RoleScope(scope), nil
}

// Role State Value Object
type RoleState string

const (
	ROLE_STATE_ACTIVE   RoleState = "active"
	ROLE_STATE_INACTIVE RoleState = "inactive"
)

const (
	ERRORS_ROLE_STATE_NOT_VALID shareddomain.DomainError = "ERRORS_ROLE_STATE_NOT_VALID"
)

func CreateRoleState(state string) (RoleState, error) {
	if state != string(ROLE_STATE_ACTIVE) && state != string(ROLE_STATE_INACTIVE) {
		return RoleState(""), errors.New(string(ERRORS_ROLE_STATE_NOT_VALID))
	}

	return RoleState(state), nil
}

// Role entity
type RoleEntity struct {
	ID    string    `json:"id"`
	Name  RoleName  `json:"name"`
	Scope RoleScope `json:"scope"`
	State RoleState `json:"state"`
}

type SystemRoleEntity struct {
	Alias      RoleAlias `json:"alias"`
	RoleEntity `json:"role"`
}

// Role Entity Factory
func CreateRoleEntity(name RoleName, scope RoleScope, state RoleState) RoleEntity {
	return RoleEntity{
		ID:    uuid.New().String(),
		Name:  name,
		Scope: scope,
		State: state,
	}
}

const (
	ERRORS_SYSTEM_ROLE_SCOPE_NOT_VALID shareddomain.DomainError = "ERRORS_SYSTEM_ROLE_SCOPE_NOT_VALID"
)

func CreateSystemRole(id string, alias RoleAlias, name RoleName, scope RoleScope, state RoleState) (SystemRoleEntity, error) {
	if scope != SCOPE_SYSTEM && scope != SCOPE_USER {
		return SystemRoleEntity{}, errors.New(string(ERRORS_SYSTEM_ROLE_SCOPE_NOT_VALID))
	}

	return SystemRoleEntity{
		Alias: alias,
		RoleEntity: RoleEntity{
			ID:    id,
			Name:  name,
			Scope: scope,
			State: state,
		},
	}, nil
}
