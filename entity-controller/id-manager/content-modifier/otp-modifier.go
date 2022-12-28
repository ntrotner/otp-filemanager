package content_modifier

type InitializeOTPModifier func()
type ReadAllIdentities func() *map[string]*UserOtp
type ReadIdentity func(*string) (*UserOtp, error)
type WriteIdentity func(*string, *UserOtp) error
type DeleteIdentity func(*string) error
type ReadFilesOfIdentity func(*string) []string

type OtpModifier struct {
	InitializeOTPModifier InitializeOTPModifier
	ReadAllIdentities     ReadAllIdentities
	ReadIdentity          ReadIdentity
	WriteIdentity         WriteIdentity
	DeleteIdentity        DeleteIdentity
	ReadFilesOfIdentity   ReadFilesOfIdentity
}
