package authdomain

type RoleRepository interface {
	GetSystemRolesByAlias([]RoleAlias) ([]SystemRoleEntity, error)
}
