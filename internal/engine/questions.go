package engine

// Question represents a single step in the quiz.
type Question struct {
	ID          string   `json:"id"`
	Text        string   `json:"text"`
	Subtitle    string   `json:"subtitle,omitempty"`
	MultiSelect bool     `json:"multi_select"`
	Options     []Option `json:"options"`
}

// Option is a selectable choice within a question.
type Option struct {
	ID      string             `json:"id"`
	Label   string             `json:"label"`
	Desc    string             `json:"desc,omitempty"`
	Icon    string             `json:"icon,omitempty"`
	Weights map[string]float64 `json:"-"`
}

// OptionByID returns the option matching the given id, or nil.
func (q *Question) OptionByID(id string) *Option {
	for i := range q.Options {
		if q.Options[i].ID == id {
			return &q.Options[i]
		}
	}
	return nil
}

var questionIndex map[string]*Question

func init() {
	questionIndex = make(map[string]*Question, len(Questions))
	for i := range Questions {
		questionIndex[Questions[i].ID] = &Questions[i]
	}
}

// QuestionByID returns the question matching the given id, or nil.
func QuestionByID(id string) *Question {
	return questionIndex[id]
}

// Questions is the ordered quiz sequence.
var Questions = []Question{
	{
		ID: "current_os", Text: "What are you using right now?",
		Subtitle: "We'll suggest something that feels familiar",
		Options: []Option{
			{ID: "windows", Label: "Windows", Icon: "🪟", Weights: map[string]float64{"windows_like": 1.5, "beginner_friendly": 0.5}},
			{ID: "macos", Label: "macOS", Icon: "🍎", Weights: map[string]float64{"macos_like": 1.5, "beginner_friendly": 0.5}},
			{ID: "chromeos", Label: "ChromeOS", Icon: "💻", Weights: map[string]float64{"beginner_friendly": 1.5, "lightweight": 1.0}},
			{ID: "linux", Label: "Already on Linux", Icon: "🐧", Weights: map[string]float64{"customizable": 0.5}},
		},
	},
	{
		ID: "experience", Text: "How much Linux experience do you have?",
		Options: []Option{
			{ID: "none", Label: "Never used it", Icon: "🌱", Desc: "Completely new to Linux",
				Weights: map[string]float64{"beginner_friendly": 2.5, "community": 1.0, "stability": 0.5}},
			{ID: "brief", Label: "Tried it briefly", Icon: "🌿", Desc: "Booted a live USB or used it for a bit",
				Weights: map[string]float64{"beginner_friendly": 1.5, "community": 0.5, "stability": 0.3}},
			{ID: "daily", Label: "Daily driver", Icon: "🌳", Desc: "Comfortable using Linux day-to-day",
				Weights: map[string]float64{"dev_tools": 0.5, "customizable": 0.5}},
			{ID: "sysadmin", Label: "Power user", Icon: "🏔️", Desc: "Sysadmin-level, comfortable with anything",
				Weights: map[string]float64{"customizable": 1.5, "bleeding_edge": 0.5, "beginner_friendly": -1.0}},
		},
	},
	{
		ID: "use_case", Text: "What will you mainly use it for?",
		Subtitle: "Select all that apply", MultiSelect: true,
		Options: []Option{
			{ID: "browsing", Label: "Browsing & Office", Icon: "🌐",
				Weights: map[string]float64{"beginner_friendly": 0.5, "stability": 0.5}},
			{ID: "dev", Label: "Software Development", Icon: "💻",
				Weights: map[string]float64{"dev_tools": 2.0, "customizable": 0.5}},
			{ID: "gaming", Label: "Gaming", Icon: "🎮",
				Weights: map[string]float64{"gaming": 2.5, "nvidia_support": 1.0}},
			{ID: "creative", Label: "Creative & Multimedia", Icon: "🎨",
				Weights: map[string]float64{"multimedia": 2.0, "stability": 0.5}},
			{ID: "server", Label: "Server & Homelab", Icon: "🖥️",
				Weights: map[string]float64{"stability": 1.5, "lightweight": 1.0, "dev_tools": 0.5}},
		},
	},
	{
		ID: "form_factor", Text: "What hardware are you installing on?",
		Options: []Option{
			{ID: "laptop", Label: "Laptop", Icon: "💻",
				Weights: map[string]float64{"laptop_friendly": 2.0, "hardware_compat": 0.5}},
			{ID: "desktop", Label: "Desktop", Icon: "🖥️",
				Weights: map[string]float64{}},
			{ID: "both", Label: "Both", Icon: "🔄",
				Weights: map[string]float64{"laptop_friendly": 1.0, "hardware_compat": 0.5}},
		},
	},
	{
		ID: "hardware_age", Text: "How old is your hardware?",
		Options: []Option{
			{ID: "new", Label: "Less than 2 years", Icon: "✨",
				Weights: map[string]float64{"bleeding_edge": 0.5, "hardware_compat": 0.5}},
			{ID: "mid", Label: "2–5 years", Icon: "👍",
				Weights: map[string]float64{"hardware_compat": 0.5}},
			{ID: "old", Label: "5–10 years", Icon: "🔧",
				Weights: map[string]float64{"lightweight": 1.5, "stability": 0.5, "hardware_compat": 1.0}},
			{ID: "ancient", Label: "10+ years", Icon: "🏚️",
				Weights: map[string]float64{"lightweight": 2.5, "hardware_compat": 1.5}},
		},
	},
	{
		ID: "gpu", Text: "What GPU do you have?",
		Subtitle: "This affects driver compatibility",
		Options: []Option{
			{ID: "nvidia", Label: "NVIDIA", Icon: "💚",
				Weights: map[string]float64{"nvidia_support": 2.5, "gaming": 0.3}},
			{ID: "amd", Label: "AMD", Icon: "❤️",
				Weights: map[string]float64{"foss_purity": 0.3, "gaming": 0.3}},
			{ID: "intel", Label: "Intel", Icon: "💙",
				Weights: map[string]float64{"lightweight": 0.3}},
			{ID: "unsure", Label: "Not sure", Icon: "❓",
				Weights: map[string]float64{"beginner_friendly": 0.5, "nvidia_support": 0.5}},
		},
	},
	{
		ID: "stability", Text: "Stability or latest packages?",
		Options: []Option{
			{ID: "stable", Label: "Rock solid", Icon: "🪨", Desc: "I want zero surprises",
				Weights: map[string]float64{"stability": 2.5, "bleeding_edge": -1.0}},
			{ID: "balanced", Label: "Balanced", Icon: "⚖️", Desc: "Reasonably current, reasonably stable",
				Weights: map[string]float64{"stability": 0.5, "bleeding_edge": 0.5}},
			{ID: "bleeding", Label: "Bleeding edge", Icon: "🚀", Desc: "Always the newest versions",
				Weights: map[string]float64{"bleeding_edge": 2.5, "stability": -0.5}},
		},
	},
	{
		ID: "customization", Text: "How much do you want to customize?",
		Options: []Option{
			{ID: "just_works", Label: "Just works", Icon: "📦", Desc: "Set it up and forget it",
				Weights: map[string]float64{"beginner_friendly": 1.0, "stability": 0.5}},
			{ID: "some", Label: "Some tweaking", Icon: "🔩", Desc: "Adjust settings, install themes",
				Weights: map[string]float64{"customizable": 0.5}},
			{ID: "total", Label: "Total control", Icon: "🏗️", Desc: "Build my system from the ground up",
				Weights: map[string]float64{"customizable": 2.5, "beginner_friendly": -0.5}},
		},
	},
	{
		ID: "package_pref", Text: "Package manager preference?",
		Subtitle: "Skip if you're not sure — it won't affect results much",
		Options: []Option{
			{ID: "none", Label: "No preference", Icon: "🤷",
				Weights: map[string]float64{}},
			{ID: "apt", Label: "APT (deb)", Icon: "📦", Desc: "Ubuntu, Debian, Mint",
				Weights: map[string]float64{"pkg_deb": 3.0}},
			{ID: "dnf", Label: "DNF / Zypper (rpm)", Icon: "📦", Desc: "Fedora, openSUSE",
				Weights: map[string]float64{"pkg_rpm": 3.0}},
			{ID: "pacman", Label: "Pacman", Icon: "📦", Desc: "Arch, Manjaro, EndeavourOS",
				Weights: map[string]float64{"pkg_pacman": 3.0}},
			{ID: "nix", Label: "Nix", Icon: "❄️", Desc: "NixOS",
				Weights: map[string]float64{"pkg_nix": 3.0}},
		},
	},
	{
		ID: "community", Text: "How important is community support?",
		Options: []Option{
			{ID: "essential", Label: "Essential", Icon: "🤝", Desc: "I'll need help, I want active forums",
				Weights: map[string]float64{"community": 2.0, "beginner_friendly": 0.5}},
			{ID: "good_docs", Label: "Good docs are enough", Icon: "📖",
				Weights: map[string]float64{"community": 1.0}},
			{ID: "self_reliant", Label: "I'll figure it out", Icon: "🧭",
				Weights: map[string]float64{"customizable": 0.3}},
		},
	},
	{
		ID: "philosophy", Text: "How do you feel about proprietary software?",
		Options: []Option{
			{ID: "foss", Label: "FOSS only", Icon: "🔓", Desc: "Free and open-source everything",
				Weights: map[string]float64{"foss_purity": 2.5}},
			{ID: "practical", Label: "Practical mix", Icon: "🔀", Desc: "Open-source preferred, proprietary when needed",
				Weights: map[string]float64{"foss_purity": 0.3}},
			{ID: "dont_care", Label: "Don't care", Icon: "🤷",
				Weights: map[string]float64{}},
		},
	},
	{
		ID: "ram", Text: "How much RAM does your machine have?",
		Options: []Option{
			{ID: "lt4", Label: "Less than 4 GB", Icon: "🪶",
				Weights: map[string]float64{"lightweight": 3.0, "beginner_friendly": -0.3}},
			{ID: "4to8", Label: "4–8 GB", Icon: "💾",
				Weights: map[string]float64{"lightweight": 1.0}},
			{ID: "8to16", Label: "8–16 GB", Icon: "🧠",
				Weights: map[string]float64{}},
			{ID: "gt16", Label: "16 GB+", Icon: "🚀",
				Weights: map[string]float64{"gaming": 0.3, "dev_tools": 0.3}},
		},
	},
}
