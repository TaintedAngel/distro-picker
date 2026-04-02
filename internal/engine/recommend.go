package engine

import (
	"fmt"
	"sort"
	"strings"
)

// Answer holds the user's response to a single question.
type Answer struct {
	QuestionID string   `json:"question_id"`
	Choices    []string `json:"choices"`
}

// RecommendRequest is the payload sent by the client.
type RecommendRequest struct {
	Answers []Answer `json:"answers,omitempty"`
	Profile string   `json:"profile,omitempty"`
}

// DistroResult is a scored distro with a personalized explanation.
type DistroResult struct {
	Distro
	Score     float64  `json:"score"`
	Reason    string   `json:"reason"`
	TopTraits []string `json:"top_traits"`
}

// DesktopResult is the recommended DE with an explanation.
type DesktopResult struct {
	Desktop
	Reason string `json:"reason"`
}

// Recommendation is the complete response returned to the client.
type Recommendation struct {
	Distros []DistroResult `json:"distros"`
	Desktop DesktopResult  `json:"desktop"`
}

// Profile is a preset answer set for common archetypes.
type Profile struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Desc    string   `json:"description"`
	Icon    string   `json:"icon"`
	Answers []Answer `json:"answers"`
}

// Profiles defines the quick-pick presets shown on the landing page.
var Profiles = []Profile{
	{
		ID: "gamer_windows", Name: "Gamer from Windows", Icon: "🎮",
		Desc: "I play games on Windows and want to switch to Linux",
		Answers: []Answer{
			{QuestionID: "current_os", Choices: []string{"windows"}},
			{QuestionID: "experience", Choices: []string{"brief"}},
			{QuestionID: "use_case", Choices: []string{"gaming"}},
			{QuestionID: "form_factor", Choices: []string{"desktop"}},
			{QuestionID: "hardware_age", Choices: []string{"new"}},
			{QuestionID: "gpu", Choices: []string{"nvidia"}},
			{QuestionID: "stability", Choices: []string{"balanced"}},
			{QuestionID: "customization", Choices: []string{"just_works"}},
			{QuestionID: "package_pref", Choices: []string{"none"}},
			{QuestionID: "community", Choices: []string{"essential"}},
			{QuestionID: "philosophy", Choices: []string{"dont_care"}},
			{QuestionID: "ram", Choices: []string{"gt16"}},
		},
	},
	{
		ID: "dev_stable", Name: "Developer, wants stability", Icon: "💻",
		Desc: "I write code all day and need a rock-solid platform",
		Answers: []Answer{
			{QuestionID: "current_os", Choices: []string{"linux"}},
			{QuestionID: "experience", Choices: []string{"daily"}},
			{QuestionID: "use_case", Choices: []string{"dev"}},
			{QuestionID: "form_factor", Choices: []string{"laptop"}},
			{QuestionID: "hardware_age", Choices: []string{"mid"}},
			{QuestionID: "gpu", Choices: []string{"amd"}},
			{QuestionID: "stability", Choices: []string{"stable"}},
			{QuestionID: "customization", Choices: []string{"some"}},
			{QuestionID: "package_pref", Choices: []string{"none"}},
			{QuestionID: "community", Choices: []string{"good_docs"}},
			{QuestionID: "philosophy", Choices: []string{"practical"}},
			{QuestionID: "ram", Choices: []string{"8to16"}},
		},
	},
	{
		ID: "old_laptop", Name: "Old laptop rescue", Icon: "🔧",
		Desc: "I have an old machine and want to breathe new life into it",
		Answers: []Answer{
			{QuestionID: "current_os", Choices: []string{"windows"}},
			{QuestionID: "experience", Choices: []string{"none"}},
			{QuestionID: "use_case", Choices: []string{"browsing"}},
			{QuestionID: "form_factor", Choices: []string{"laptop"}},
			{QuestionID: "hardware_age", Choices: []string{"ancient"}},
			{QuestionID: "gpu", Choices: []string{"intel"}},
			{QuestionID: "stability", Choices: []string{"stable"}},
			{QuestionID: "customization", Choices: []string{"just_works"}},
			{QuestionID: "package_pref", Choices: []string{"none"}},
			{QuestionID: "community", Choices: []string{"essential"}},
			{QuestionID: "philosophy", Choices: []string{"dont_care"}},
			{QuestionID: "ram", Choices: []string{"lt4"}},
		},
	},
	{
		ID: "power_user", Name: "Power user / ricer", Icon: "🏗️",
		Desc: "I want full control and love customizing every pixel",
		Answers: []Answer{
			{QuestionID: "current_os", Choices: []string{"linux"}},
			{QuestionID: "experience", Choices: []string{"sysadmin"}},
			{QuestionID: "use_case", Choices: []string{"dev", "gaming"}},
			{QuestionID: "form_factor", Choices: []string{"desktop"}},
			{QuestionID: "hardware_age", Choices: []string{"new"}},
			{QuestionID: "gpu", Choices: []string{"amd"}},
			{QuestionID: "stability", Choices: []string{"bleeding"}},
			{QuestionID: "customization", Choices: []string{"total"}},
			{QuestionID: "package_pref", Choices: []string{"pacman"}},
			{QuestionID: "community", Choices: []string{"self_reliant"}},
			{QuestionID: "philosophy", Choices: []string{"practical"}},
			{QuestionID: "ram", Choices: []string{"gt16"}},
		},
	},
}

var traitLabels = map[string]string{
	"beginner_friendly": "beginner-friendly with a gentle learning curve",
	"stability":         "rock-solid and reliable for everyday use",
	"bleeding_edge":     "always running the latest software",
	"gaming":            "great out-of-the-box gaming support",
	"dev_tools":         "an excellent platform for software development",
	"multimedia":        "strong support for creative and multimedia work",
	"lightweight":       "lightweight and efficient on resources",
	"customizable":      "highly customizable to fit your workflow",
	"community":         "backed by a large and helpful community",
	"foss_purity":       "committed to free and open-source principles",
	"nvidia_support":    "reliable NVIDIA GPU driver support",
	"hardware_compat":   "broad hardware compatibility out of the box",
	"laptop_friendly":   "excellent laptop support and power management",
}

// Recommend scores all distros and desktops against the user's answers.
func Recommend(req RecommendRequest) *Recommendation {
	answers := req.Answers
	if req.Profile != "" {
		for _, p := range Profiles {
			if p.ID == req.Profile {
				answers = p.Answers
				break
			}
		}
	}

	weights := buildWeights(answers)
	distros := scoreDistros(weights)
	ramGB := ramFromAnswers(answers)
	desktop := scoreDesktops(weights, ramGB)

	return &Recommendation{
		Distros: distros,
		Desktop: desktop,
	}
}

func buildWeights(answers []Answer) map[string]float64 {
	w := make(map[string]float64)
	for _, a := range answers {
		q := QuestionByID(a.QuestionID)
		if q == nil {
			continue
		}
		for _, choice := range a.Choices {
			opt := q.OptionByID(choice)
			if opt == nil {
				continue
			}
			for trait, val := range opt.Weights {
				w[trait] += val
			}
		}
	}
	return w
}

type scored struct {
	index int
	score float64
}

func scoreDistros(weights map[string]float64) []DistroResult {
	entries := make([]scored, len(Distros))
	for i, d := range Distros {
		var s float64
		for trait, w := range weights {
			s += d.Traits[trait] * w
		}
		entries[i] = scored{index: i, score: s}
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].score > entries[j].score
	})

	maxScore := entries[0].score
	if maxScore <= 0 {
		maxScore = 1
	}

	cap := 5
	if len(entries) < cap {
		cap = len(entries)
	}

	results := make([]DistroResult, cap)
	for i := 0; i < cap; i++ {
		d := Distros[entries[i].index]
		norm := entries[i].score / maxScore
		results[i] = DistroResult{
			Distro:    d,
			Score:     norm,
			Reason:    generateReason(d, weights),
			TopTraits: topTraits(d, weights, 3),
		}
	}
	return results
}

func scoreDesktops(weights map[string]float64, ramGB float64) DesktopResult {
	var best int
	var bestScore float64 = -1

	for i, de := range Desktops {
		if de.MinRAMGB > ramGB {
			continue
		}
		var s float64
		for trait, w := range weights {
			s += de.Traits[trait] * w
		}
		if s > bestScore {
			bestScore = s
			best = i
		}
	}

	de := Desktops[best]
	return DesktopResult{
		Desktop: de,
		Reason:  generateDEReason(de, weights),
	}
}

func generateReason(d Distro, weights map[string]float64) string {
	traits := topTraits(d, weights, 3)
	phrases := make([]string, 0, len(traits))
	for _, t := range traits {
		if label, ok := traitLabels[t]; ok {
			phrases = append(phrases, label)
		}
	}
	switch len(phrases) {
	case 0:
		return fmt.Sprintf("%s is a solid choice for your needs.", d.Name)
	case 1:
		return fmt.Sprintf("%s is %s.", d.Name, phrases[0])
	case 2:
		return fmt.Sprintf("%s is %s and %s.", d.Name, phrases[0], phrases[1])
	default:
		return fmt.Sprintf("%s is %s, %s, and %s.",
			d.Name, phrases[0], phrases[1], phrases[2])
	}
}

func generateDEReason(de Desktop, weights map[string]float64) string {
	type kv struct {
		trait string
		val   float64
	}
	var contribs []kv
	for trait, w := range weights {
		v := de.Traits[trait] * w
		if v > 0 {
			contribs = append(contribs, kv{trait, v})
		}
	}
	sort.Slice(contribs, func(i, j int) bool {
		return contribs[i].val > contribs[j].val
	})

	var parts []string
	for i := 0; i < 2 && i < len(contribs); i++ {
		switch contribs[i].trait {
		case "windows_like":
			parts = append(parts, "has a familiar Windows-like layout")
		case "macos_like":
			parts = append(parts, "has a clean, macOS-inspired workflow")
		case "lightweight":
			parts = append(parts, "is lightweight on resources")
		case "customizable":
			parts = append(parts, "is deeply customizable")
		case "beginner_friendly":
			parts = append(parts, "is easy to learn")
		}
	}

	if len(parts) == 0 {
		return fmt.Sprintf("%s is a great fit for your setup.", de.Name)
	}
	return fmt.Sprintf("%s %s.", de.Name, strings.Join(parts, " and "))
}

func topTraits(d Distro, weights map[string]float64, n int) []string {
	type kv struct {
		trait string
		val   float64
	}
	var contribs []kv
	for trait, w := range weights {
		v := d.Traits[trait] * w
		if v > 0 {
			contribs = append(contribs, kv{trait, v})
		}
	}
	sort.Slice(contribs, func(i, j int) bool {
		return contribs[i].val > contribs[j].val
	})

	cap := n
	if len(contribs) < cap {
		cap = len(contribs)
	}
	out := make([]string, cap)
	for i := range out {
		out[i] = contribs[i].trait
	}
	return out
}

func ramFromAnswers(answers []Answer) float64 {
	for _, a := range answers {
		if a.QuestionID != "ram" || len(a.Choices) == 0 {
			continue
		}
		switch a.Choices[0] {
		case "lt4":
			return 3
		case "4to8":
			return 6
		case "8to16":
			return 12
		case "gt16":
			return 32
		}
	}
	return 8 // safe default
}
