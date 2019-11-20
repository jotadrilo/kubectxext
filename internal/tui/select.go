package tui

import (
	"strings"

	"github.com/manifoldco/promptui"
)

// Context contains basic information about a kubernetes context.
type Context struct {
	Name    string
	Cluster string
	Auth    string
}

// ContextSelect will show a series of contexts stored. To display the content
// of the struct, the usual dot notation is used inside the templates to
// select the fields and color them.
func ContextSelect(contexts []Context) *promptui.Select {
	// The Active and Selected templates set a small helm icon next to the
	// name colored and the cluster for the active template. The details
	// template is show at the bottom of the select's list and displays the
	// full info for that context in a multi-line template.
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U00002388 {{ .Name | cyan }} ({{ .Cluster | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .Cluster | red }})",
		Selected: "\U00002388 {{ .Name | red | cyan }}",
		Details: `
--------- Context ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Cluster:" | faint }}	{{ .Cluster }}
{{ "Auth:" | faint }}	{{ .Auth }}`,
	}

	searcher := func(input string, index int) bool {
		ctx := contexts[index]
		name := strings.Replace(strings.ToLower(ctx.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(name, input)
	}

	return &promptui.Select{
		Label:     "Context Name",
		Items:     contexts,
		Templates: templates,
		Size:      10,
		Searcher:  searcher,
	}
}
