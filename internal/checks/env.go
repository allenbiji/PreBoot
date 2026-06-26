package checks

import (
	"fmt"

	"github.com/allenbiji/clone-sage/internal/detect"
	"github.com/allenbiji/clone-sage/internal/model"
	"github.com/allenbiji/clone-sage/internal/registry"
)

type EnvCheck struct {
	Key string
}

//execute method for the check
func (e *EnvCheck) Execute() error {
	envMap, err := detect.ExtractEnvKeys(".env")
	if err != nil{
		return fmt.Errorf("There was an error in scanning the .env file")
	}

	val, exists := envMap[e.Key]
	if !exists {
		return fmt.Errorf("Key '%s' does not exist", e.Key)
	} 
	if val == "" {
		return fmt.Errorf("Key '%s' exists but the value is empty", e.Key)
	}

	return nil
}

//build the EnvCheck factory
func buildEnvExistsCheck(cfg model.CheckConfig) (registry.Check, error){
	key, ok := cfg.Options["key"]
	if !ok || key == "" {
		return nil, fmt.Errorf("env_exists check requires a 'key' option")
	}

	return &EnvCheck{
		Key: key,
	}, nil
}

//register the check in the registry
func init(){
	registry.Register(model.TypeEnvExists, buildEnvExistsCheck)
}