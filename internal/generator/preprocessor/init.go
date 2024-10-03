package preprocessor

func NewPreprocessor(configDir string) *Preprocessor {
	return &Preprocessor{
		ConfigDir:             configDir,
		ForbiddenDeclarators:  forbiddenDeclarators,
	}
}
