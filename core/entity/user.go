package entity

// User Config
const (
	UserSerialPrefix = "USER"
	UserSerialLength = 10
)

type User struct {
	ID                int
	Serial			  string
	Email             string
	Role			  string
	Password          string
	FullName          string
	PhotoURL          string
	PhoneNumber       string
	Address           string
	AccessStatus      UserStatus
	Status            string
	RegistrationOTP   string
	ForgotPasswordToken string
}

type UserRole string

const (
	UserRoleSuperAdmin = "super-admin"
	UserRoleAdmin	   = "admin"
	UserRoleUser	   = "user"
	UserRoleGuest	   = "guest"
)

type UserStatus string

const (
	UserStatusDisabled       = "disabled"
	UserStatusActive         = "enabled"
	UserStatusForgotPassword = "forgot_password"
)