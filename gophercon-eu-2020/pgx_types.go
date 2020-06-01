type Role int

func (r Role) Value() (driver.Value, error) {
	switch r {
	case RoleAdmin:
		return "admin", nil
	case RoleMember:
		return "member", nil
	default:
		return nil, fmt.Errorf("invalid Role: %T", r)
	}
}
func (r *Role) Scan(in interface{}) error {
	switch in.(string) {
	case "admin":
		*r = RoleAdmin
		return nil
	case "member":
		*r = RoleMember
		return nil
	default:
		return fmt.Errorf("invalid Role: %q", in.(string))
	}
}
