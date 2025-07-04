package templates

import _ "embed"

// RulesMD contains default RULES.md file.
//
//go:embed rules.md
var RulesMD string

// GeneratePRPMD contains default GENERATE_PRP.md file.
//
//go:embed generate_prp.md
var GeneratePRPMD string

// ExecutePRPMD contains default EXECUTE_PRP.md file.
//
//go:embed execute_prp.md
var ExecutePRPMD string

// FeatureTemplateMD contains default feature template.
//
//go:embed feature_template.md
var FeatureTemplateMD string

// PrpTemplateMD contains default prp_template.md file.
//
//go:embed prp_template.md
var PrpTemplateMD string
