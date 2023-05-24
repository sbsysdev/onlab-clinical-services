package authdomain

type RoleRepository interface {
	GetAliasRolesByAlias([]RoleAlias) ([]AliasRoleEntity, error)
}
