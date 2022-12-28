package file_system

import content_modifier "otp-filemanager/permission-controller/id-manager/content-modifier"

func CreateFileSystemModifier() content_modifier.Modifier {
	return content_modifier.Modifier{
		OtpModifier: content_modifier.OtpModifier{
			InitializeOTPModifier: InitializeOTPModifier,
			ReadAllIdentities:     ReadAllIdentities,
			ReadIdentity:          ReadIdentity,
			WriteIdentity:         WriteIdentity,
			DeleteIdentity:        DeleteIdentity,
			ReadFilesOfIdentity:   ReadFilesOfIdentity,
		},
		FileModifier: content_modifier.FileModifier{
			WriteFile:  WriteFile,
			ReadFile:   ReadFile,
			DeleteFile: DeleteFile,
		},
	}
}
