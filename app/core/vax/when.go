package vax

import "context"

// When returns a validation rule that executes the given list of rules when the condition is true.
func When(condition bool, rules ...Rule) WhenRule {
	return WhenRule{
		condition: condition,
		rules:     rules,
		elseRules: []Rule{},
	}
}

// WhenRule is a validation rule that executes the given list of rules when the condition is true.
type WhenRule struct {
	condition bool
	rules     []Rule
	elseRules []Rule
}

func (r WhenRule) Code(code string) Rule {
	return r
}

func (r WhenRule) Msg(msg string) Rule {
	return r
}

// Validate checks if the condition is true and if so, it validates the value using the specified rules.
func (r WhenRule) Validate(lang string, value interface{}) error {
	return r.ValidateWithContext(context.Background(), lang, value)
}

// ValidateWithContext checks if the condition is true and if so, it validates the value using the specified rules.
func (r WhenRule) ValidateWithContext(ctx context.Context, lang string, value interface{}) error {
	if r.condition {
		if ctx == nil {
			return Validate(lang, value, r.rules...)
		}
		return ValidateWithContext(ctx, lang, value, r.rules...)
	}

	if ctx == nil {
		return Validate(lang, value, r.elseRules...)
	}
	return ValidateWithContext(ctx, lang, value, r.elseRules...)
}

// Else returns a validation rule that executes the given list of rules when the condition is false.
func (r WhenRule) Else(rules ...Rule) WhenRule {
	r.elseRules = rules
	return r
}
