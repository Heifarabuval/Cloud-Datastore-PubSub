package utils

import "Calicut/config"

func GetEnvVar(varKey string, defaultVar string) string {
	var envConst, projectIdError = config.GetEnvConst(varKey)
	if !projectIdError {
		envConst = defaultVar
	}
	return envConst
}
