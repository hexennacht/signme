// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/hexennacht/signme/services/sm-user-api/ent/credential"
	"github.com/hexennacht/signme/services/sm-user-api/ent/schema"
	"github.com/hexennacht/signme/services/sm-user-api/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	credentialFields := schema.Credential{}.Fields()
	_ = credentialFields
	// credentialDescCreatedAt is the schema descriptor for created_at field.
	credentialDescCreatedAt := credentialFields[3].Descriptor()
	// credential.DefaultCreatedAt holds the default value on creation for the created_at field.
	credential.DefaultCreatedAt = credentialDescCreatedAt.Default.(func() time.Time)
	// credentialDescUpdatedAt is the schema descriptor for updated_at field.
	credentialDescUpdatedAt := credentialFields[4].Descriptor()
	// credential.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	credential.DefaultUpdatedAt = credentialDescUpdatedAt.Default.(func() time.Time)
	// credential.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	credential.UpdateDefaultUpdatedAt = credentialDescUpdatedAt.UpdateDefault.(func() time.Time)
	// credentialDescID is the schema descriptor for id field.
	credentialDescID := credentialFields[0].Descriptor()
	// credential.IDValidator is a validator for the "id" field. It is called by the builders before save.
	credential.IDValidator = credentialDescID.Validators[0].(func(int64) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[2].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescProfilePicture is the schema descriptor for profile_picture field.
	userDescProfilePicture := userFields[4].Descriptor()
	// user.DefaultProfilePicture holds the default value on creation for the profile_picture field.
	user.DefaultProfilePicture = userDescProfilePicture.Default.(string)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[6].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[7].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(int64) error)
}
