package rule

import (
	"fmt"

	"github.com/conventionalcommit/commitlint/lint"
)

// HeadMaxLenRule to validate max length of header
type HeadMaxLenRule struct {
	CheckLen int
}

// Name return name of the rule
func (r *HeadMaxLenRule) Name() string { return "header-max-length" }

// Validate validates HeadMaxLenRule
func (r *HeadMaxLenRule) Validate(msg *lint.Commit) ([]string, bool) {
	return checkMaxLen(r.CheckLen, msg.Header())
}

// Apply sets the needed argument for the rule
func (r *HeadMaxLenRule) Apply(setting lint.RuleSetting) error {
	err := setIntArg(&r.CheckLen, setting.Argument)
	if err != nil {
		return errInvalidArg(r.Name(), err)
	}
	return nil
}

// BodyMaxLenRule to validate max length of body
type BodyMaxLenRule struct {
	CheckLen int
}

// Name return name of the rule
func (r *BodyMaxLenRule) Name() string { return "body-max-length" }

// Validate validates BodyMaxLenRule
func (r *BodyMaxLenRule) Validate(msg *lint.Commit) ([]string, bool) {
	return checkMaxLen(r.CheckLen, msg.Body())
}

// Apply sets the needed argument for the rule
func (r *BodyMaxLenRule) Apply(setting lint.RuleSetting) error {
	err := setIntArg(&r.CheckLen, setting.Argument)
	if err != nil {
		return errInvalidArg(r.Name(), err)
	}
	return nil
}

// FooterMaxLenRule to validate max length of footer
type FooterMaxLenRule struct {
	CheckLen int
}

// Name return name of the rule
func (r *FooterMaxLenRule) Name() string { return "footer-max-length" }

// Validate validates FooterMaxLenRule
func (r *FooterMaxLenRule) Validate(msg *lint.Commit) ([]string, bool) {
	return checkMaxLen(r.CheckLen, msg.Footer())
}

// Apply sets the needed argument for the rule
func (r *FooterMaxLenRule) Apply(setting lint.RuleSetting) error {
	err := setIntArg(&r.CheckLen, setting.Argument)
	if err != nil {
		return errInvalidArg(r.Name(), err)
	}
	return nil
}

// TypeMaxLenRule to validate max length of type
type TypeMaxLenRule struct {
	CheckLen int
}

// Name return name of the rule
func (r *TypeMaxLenRule) Name() string { return "type-max-length" }

// Validate validates TypeMaxLenRule
func (r *TypeMaxLenRule) Validate(msg *lint.Commit) ([]string, bool) {
	return checkMaxLen(r.CheckLen, msg.Type())
}

// Apply sets the needed argument for the rule
func (r *TypeMaxLenRule) Apply(setting lint.RuleSetting) error {
	err := setIntArg(&r.CheckLen, setting.Argument)
	if err != nil {
		return errInvalidArg(r.Name(), err)
	}
	return nil
}

// ScopeMaxLenRule to validate max length of type
type ScopeMaxLenRule struct {
	CheckLen int
}

// Name return name of the rule
func (r *ScopeMaxLenRule) Name() string { return "scope-max-length" }

// Validate validates ScopeMaxLenRule
func (r *ScopeMaxLenRule) Validate(msg *lint.Commit) ([]string, bool) {
	return checkMaxLen(r.CheckLen, msg.Scope())
}

// Apply sets the needed argument for the rule
func (r *ScopeMaxLenRule) Apply(setting lint.RuleSetting) error {
	err := setIntArg(&r.CheckLen, setting.Argument)
	if err != nil {
		return errInvalidArg(r.Name(), err)
	}
	return nil
}

// DescriptionMaxLenRule to validate max length of type
type DescriptionMaxLenRule struct {
	CheckLen int
}

// Name return name of the rule
func (r *DescriptionMaxLenRule) Name() string { return "description-max-length" }

// Validate validates DescriptionMaxLenRule
func (r *DescriptionMaxLenRule) Validate(msg *lint.Commit) ([]string, bool) {
	return checkMaxLen(r.CheckLen, msg.Description())
}

// Apply sets the needed argument for the rule
func (r *DescriptionMaxLenRule) Apply(setting lint.RuleSetting) error {
	err := setIntArg(&r.CheckLen, setting.Argument)
	if err != nil {
		return errInvalidArg(r.Name(), err)
	}
	return nil
}

func checkMaxLen(checkLen int, toCheck string) ([]string, bool) {
	if checkLen < 0 {
		return nil, true
	}
	actualLen := len(toCheck)
	if actualLen > checkLen {
		errMsg := fmt.Sprintf("length is %d, should have less than %d chars", actualLen, checkLen)
		return []string{errMsg}, false
	}
	return nil, true
}
