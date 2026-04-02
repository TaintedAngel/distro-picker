package engine

// Distro holds metadata and scoring traits for a Linux distribution.
type Distro struct {
	ID             string             `json:"id"`
	Name           string             `json:"name"`
	Version        string             `json:"version"`
	Tagline        string             `json:"tagline"`
	Homepage       string             `json:"homepage"`
	DownloadURL    string             `json:"download_url"`
	ISOSize        string             `json:"iso_size"`
	PackageManager string             `json:"package_manager"`
	BasedOn        string             `json:"based_on,omitempty"`
	DefaultDE      string             `json:"default_de"`
	Traits         map[string]float64 `json:"-"`
	Pros           []string           `json:"pros"`
	Cons           []string           `json:"cons"`
}

// Desktop holds metadata and scoring traits for a desktop environment or window manager.
type Desktop struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	MinRAMGB    float64            `json:"-"`
	Traits      map[string]float64 `json:"-"`
}

// Distros is the full catalog of supported distributions.
var Distros = []Distro{
	{
		ID: "linux-mint", Name: "Linux Mint", Version: "22",
		Tagline:     "Elegant, easy to use, up to date and comfortable",
		Homepage:    "https://linuxmint.com",
		DownloadURL: "https://linuxmint.com/download.php", ISOSize: "2.8 GB",
		PackageManager: "apt", BasedOn: "Ubuntu", DefaultDE: "cinnamon",
		Traits: map[string]float64{
			"beginner_friendly": 0.95, "stability": 0.85, "bleeding_edge": 0.15,
			"gaming": 0.55, "dev_tools": 0.60, "multimedia": 0.65,
			"lightweight": 0.50, "customizable": 0.50, "community": 0.90,
			"foss_purity": 0.60, "nvidia_support": 0.75, "hardware_compat": 0.80,
			"laptop_friendly": 0.80, "pkg_deb": 1.0,
		},
		Pros: []string{"Extremely beginner-friendly", "Familiar taskbar-and-menu layout", "Large community with active forums", "Stable base with curated updates"},
		Cons: []string{"Packages lag behind upstream Ubuntu", "Not ideal for bleeding-edge software", "Cinnamon can feel heavy on older machines"},
	},
	{
		ID: "ubuntu", Name: "Ubuntu", Version: "24.04 LTS",
		Tagline:     "The leading platform for developers and enterprises",
		Homepage:    "https://ubuntu.com",
		DownloadURL: "https://ubuntu.com/download/desktop", ISOSize: "5.7 GB",
		PackageManager: "apt", DefaultDE: "gnome",
		Traits: map[string]float64{
			"beginner_friendly": 0.88, "stability": 0.80, "bleeding_edge": 0.25,
			"gaming": 0.60, "dev_tools": 0.78, "multimedia": 0.65,
			"lightweight": 0.30, "customizable": 0.45, "community": 0.95,
			"foss_purity": 0.50, "nvidia_support": 0.80, "hardware_compat": 0.85,
			"laptop_friendly": 0.85, "pkg_deb": 1.0,
		},
		Pros: []string{"Largest community and documentation", "Industry standard for Linux", "Excellent hardware vendor support", "Huge software repository"},
		Cons: []string{"Snap packages can be polarizing", "GNOME defaults feel restrictive to some", "Heavier resource usage than minimal distros"},
	},
	{
		ID: "fedora", Name: "Fedora Workstation", Version: "41",
		Tagline:     "Leading-edge features in a polished, community-driven OS",
		Homepage:    "https://fedoraproject.org",
		DownloadURL: "https://fedoraproject.org/workstation/download", ISOSize: "2.2 GB",
		PackageManager: "dnf", DefaultDE: "gnome",
		Traits: map[string]float64{
			"beginner_friendly": 0.65, "stability": 0.70, "bleeding_edge": 0.75,
			"gaming": 0.55, "dev_tools": 0.85, "multimedia": 0.50,
			"lightweight": 0.38, "customizable": 0.60, "community": 0.80,
			"foss_purity": 0.75, "nvidia_support": 0.58, "hardware_compat": 0.85,
			"laptop_friendly": 0.90, "pkg_rpm": 1.0,
		},
		Pros: []string{"Excellent developer platform", "Best Wayland and GNOME integration", "Strong laptop support", "Fresh packages without sacrificing stability"},
		Cons: []string{"Multimedia codecs not included by default", "Smaller gaming community than Ubuntu", "Six-month release cycle requires regular upgrades"},
	},
	{
		ID: "pop-os", Name: "Pop!_OS", Version: "22.04 LTS",
		Tagline:     "An operating system for STEM and creative professionals",
		Homepage:    "https://pop.system76.com",
		DownloadURL: "https://pop.system76.com", ISOSize: "2.6 GB",
		PackageManager: "apt", BasedOn: "Ubuntu", DefaultDE: "cosmic",
		Traits: map[string]float64{
			"beginner_friendly": 0.82, "stability": 0.75, "bleeding_edge": 0.35,
			"gaming": 0.80, "dev_tools": 0.82, "multimedia": 0.68,
			"lightweight": 0.35, "customizable": 0.58, "community": 0.70,
			"foss_purity": 0.55, "nvidia_support": 0.95, "hardware_compat": 0.85,
			"laptop_friendly": 0.90, "pkg_deb": 1.0,
		},
		Pros: []string{"Best out-of-the-box NVIDIA support", "Dedicated ISO for NVIDIA GPUs", "Tiling window manager built in", "Great for development and gaming"},
		Cons: []string{"Smaller community than Ubuntu", "Moving to Cosmic DE (still maturing)", "Limited to System76 packaging cadence"},
	},
	{
		ID: "nobara", Name: "Nobara Linux", Version: "40",
		Tagline:     "A modified Fedora designed for gaming and streaming",
		Homepage:    "https://nobaraproject.org",
		DownloadURL: "https://nobaraproject.org/download/", ISOSize: "3.5 GB",
		PackageManager: "dnf", BasedOn: "Fedora", DefaultDE: "gnome",
		Traits: map[string]float64{
			"beginner_friendly": 0.70, "stability": 0.55, "bleeding_edge": 0.65,
			"gaming": 0.95, "dev_tools": 0.60, "multimedia": 0.78,
			"lightweight": 0.30, "customizable": 0.55, "community": 0.48,
			"foss_purity": 0.38, "nvidia_support": 0.90, "hardware_compat": 0.75,
			"laptop_friendly": 0.65, "pkg_rpm": 1.0,
		},
		Pros: []string{"Gaming-optimized out of the box", "Pre-configured OBS, Steam, Lutris", "Custom kernel patches for gaming", "Codecs and drivers pre-installed"},
		Cons: []string{"Smaller community, primarily one maintainer", "Can lag behind Fedora updates", "Less tested for non-gaming workloads"},
	},
	{
		ID: "bazzite", Name: "Bazzite", Version: "3.0",
		Tagline:     "The next generation of Linux gaming — SteamOS for every PC",
		Homepage:    "https://bazzite.gg",
		DownloadURL: "https://bazzite.gg/#image-picker", ISOSize: "6.0 GB",
		PackageManager: "rpm-ostree", BasedOn: "Fedora", DefaultDE: "kde",
		Traits: map[string]float64{
			"beginner_friendly": 0.75, "stability": 0.72, "bleeding_edge": 0.60,
			"gaming": 0.95, "dev_tools": 0.30, "multimedia": 0.60,
			"lightweight": 0.22, "customizable": 0.30, "community": 0.55,
			"foss_purity": 0.45, "nvidia_support": 0.85, "hardware_compat": 0.75,
			"laptop_friendly": 0.60, "pkg_rpm": 1.0,
		},
		Pros: []string{"SteamOS-like experience on any hardware", "Immutable base prevents system breakage", "Automatic updates with rollback", "Excellent controller and handheld support"},
		Cons: []string{"Immutable design limits traditional package installs", "Relies heavily on Flatpak for apps", "Not ideal for development or server use"},
	},
	{
		ID: "cachyos", Name: "CachyOS", Version: "Rolling",
		Tagline:     "Performance-optimized Arch with custom kernels and toolchains",
		Homepage:    "https://cachyos.org",
		DownloadURL: "https://cachyos.org/download/", ISOSize: "2.5 GB",
		PackageManager: "pacman", BasedOn: "Arch", DefaultDE: "kde",
		Traits: map[string]float64{
			"beginner_friendly": 0.48, "stability": 0.50, "bleeding_edge": 0.92,
			"gaming": 0.90, "dev_tools": 0.82, "multimedia": 0.72,
			"lightweight": 0.55, "customizable": 0.88, "community": 0.50,
			"foss_purity": 0.55, "nvidia_support": 0.85, "hardware_compat": 0.85,
			"laptop_friendly": 0.68, "pkg_pacman": 1.0,
		},
		Pros: []string{"Custom performance-tuned kernels", "x86-64-v3/v4 optimized packages", "Excellent gaming performance", "Friendly installer for an Arch derivative"},
		Cons: []string{"Smaller community than EndeavourOS or Manjaro", "Rolling release requires comfort with updates", "Performance patches may introduce edge cases"},
	},
	{
		ID: "garuda", Name: "Garuda Linux", Version: "Rolling",
		Tagline:     "A performance-focused, visually striking Arch-based distro",
		Homepage:    "https://garudalinux.org",
		DownloadURL: "https://garudalinux.org/downloads", ISOSize: "3.5 GB",
		PackageManager: "pacman", BasedOn: "Arch", DefaultDE: "kde",
		Traits: map[string]float64{
			"beginner_friendly": 0.52, "stability": 0.42, "bleeding_edge": 0.85,
			"gaming": 0.85, "dev_tools": 0.70, "multimedia": 0.68,
			"lightweight": 0.22, "customizable": 0.80, "community": 0.45,
			"foss_purity": 0.50, "nvidia_support": 0.78, "hardware_compat": 0.75,
			"laptop_friendly": 0.58, "pkg_pacman": 1.0,
		},
		Pros: []string{"Eye-catching default theming", "Btrfs with automatic snapshots", "Gaming edition with pre-installed tools", "Calamares installer makes setup easy"},
		Cons: []string{"Heavy default installs with lots of bloat", "Visual polish over substance for some tastes", "Smaller community support"},
	},
	{
		ID: "opensuse-tw", Name: "openSUSE Tumbleweed", Version: "Rolling",
		Tagline:     "The original rolling release — tested, stable, and always current",
		Homepage:    "https://get.opensuse.org/tumbleweed",
		DownloadURL: "https://get.opensuse.org/tumbleweed/", ISOSize: "4.5 GB",
		PackageManager: "zypper", DefaultDE: "kde",
		Traits: map[string]float64{
			"beginner_friendly": 0.50, "stability": 0.68, "bleeding_edge": 0.85,
			"gaming": 0.50, "dev_tools": 0.80, "multimedia": 0.58,
			"lightweight": 0.32, "customizable": 0.75, "community": 0.70,
			"foss_purity": 0.72, "nvidia_support": 0.62, "hardware_compat": 0.80,
			"laptop_friendly": 0.75, "pkg_rpm": 1.0,
		},
		Pros: []string{"Automated openQA testing for every snapshot", "YaST system management tool", "Rolling release with enterprise-grade QA", "Btrfs with snapper rollback by default"},
		Cons: []string{"Zypper is slower than apt or pacman", "Codec and NVIDIA setup takes extra steps", "Smaller gaming community"},
	},
	{
		ID: "opensuse-leap", Name: "openSUSE Leap", Version: "15.6",
		Tagline:     "Enterprise-grade stability for workstations and servers",
		Homepage:    "https://get.opensuse.org/leap",
		DownloadURL: "https://get.opensuse.org/leap/", ISOSize: "4.0 GB",
		PackageManager: "zypper", DefaultDE: "kde",
		Traits: map[string]float64{
			"beginner_friendly": 0.52, "stability": 0.92, "bleeding_edge": 0.10,
			"gaming": 0.32, "dev_tools": 0.75, "multimedia": 0.52,
			"lightweight": 0.38, "customizable": 0.65, "community": 0.70,
			"foss_purity": 0.72, "nvidia_support": 0.58, "hardware_compat": 0.78,
			"laptop_friendly": 0.72, "pkg_rpm": 1.0,
		},
		Pros: []string{"Shares codebase with SUSE Enterprise", "Extremely stable for production use", "YaST makes system admin accessible", "Good bridge between desktop and server"},
		Cons: []string{"Packages are often quite dated", "Not suitable for gaming", "Smaller community than Debian or Ubuntu"},
	},
	{
		ID: "debian", Name: "Debian", Version: "12",
		Tagline:     "The universal operating system — rock-solid and free",
		Homepage:    "https://www.debian.org",
		DownloadURL: "https://www.debian.org/distrib/", ISOSize: "3.7 GB",
		PackageManager: "apt", DefaultDE: "gnome",
		Traits: map[string]float64{
			"beginner_friendly": 0.40, "stability": 0.95, "bleeding_edge": 0.05,
			"gaming": 0.28, "dev_tools": 0.78, "multimedia": 0.42,
			"lightweight": 0.60, "customizable": 0.72, "community": 0.85,
			"foss_purity": 0.88, "nvidia_support": 0.42, "hardware_compat": 0.72,
			"laptop_friendly": 0.62, "pkg_deb": 1.0,
		},
		Pros: []string{"Unmatched stability and reliability", "The foundation most distros build on", "Massive package repository", "Strong commitment to software freedom"},
		Cons: []string{"Very old packages in stable release", "Non-free firmware requires extra steps", "Not beginner-friendly out of the box"},
	},
	{
		ID: "endeavouros", Name: "EndeavourOS", Version: "Rolling",
		Tagline:     "A terminal-centric, friendly gateway to Arch Linux",
		Homepage:    "https://endeavouros.com",
		DownloadURL: "https://endeavouros.com/download/", ISOSize: "2.1 GB",
		PackageManager: "pacman", BasedOn: "Arch", DefaultDE: "kde",
		Traits: map[string]float64{
			"beginner_friendly": 0.58, "stability": 0.52, "bleeding_edge": 0.90,
			"gaming": 0.72, "dev_tools": 0.82, "multimedia": 0.65,
			"lightweight": 0.68, "customizable": 0.90, "community": 0.68,
			"foss_purity": 0.60, "nvidia_support": 0.75, "hardware_compat": 0.80,
			"laptop_friendly": 0.70, "pkg_pacman": 1.0,
		},
		Pros: []string{"Close to vanilla Arch with an installer", "Welcoming community for Arch newcomers", "Minimal default install you customize", "Full access to AUR"},
		Cons: []string{"Still requires comfort with the terminal", "Less hand-holding than Manjaro", "Rolling release can break occasionally"},
	},
	{
		ID: "arch", Name: "Arch Linux", Version: "Rolling",
		Tagline:     "A lightweight, flexible distro that tries to Keep It Simple",
		Homepage:    "https://archlinux.org",
		DownloadURL: "https://archlinux.org/download/", ISOSize: "0.9 GB",
		PackageManager: "pacman", DefaultDE: "none",
		Traits: map[string]float64{
			"beginner_friendly": 0.08, "stability": 0.50, "bleeding_edge": 0.95,
			"gaming": 0.72, "dev_tools": 0.90, "multimedia": 0.65,
			"lightweight": 0.85, "customizable": 0.98, "community": 0.88,
			"foss_purity": 0.62, "nvidia_support": 0.70, "hardware_compat": 0.82,
			"laptop_friendly": 0.58, "pkg_pacman": 1.0,
		},
		Pros: []string{"Build exactly the system you want", "Best documentation (Arch Wiki)", "Bleeding-edge packages", "Massive AUR ecosystem"},
		Cons: []string{"Manual installation process", "Requires significant Linux knowledge", "Can break if updated carelessly"},
	},
	{
		ID: "manjaro", Name: "Manjaro", Version: "Rolling",
		Tagline:     "Enjoy the simplicity — Arch-based and beginner-friendly",
		Homepage:    "https://manjaro.org",
		DownloadURL: "https://manjaro.org/download/", ISOSize: "3.5 GB",
		PackageManager: "pacman", BasedOn: "Arch", DefaultDE: "kde",
		Traits: map[string]float64{
			"beginner_friendly": 0.72, "stability": 0.58, "bleeding_edge": 0.68,
			"gaming": 0.70, "dev_tools": 0.65, "multimedia": 0.68,
			"lightweight": 0.38, "customizable": 0.68, "community": 0.65,
			"foss_purity": 0.55, "nvidia_support": 0.75, "hardware_compat": 0.75,
			"laptop_friendly": 0.72, "pkg_pacman": 1.0,
		},
		Pros: []string{"Arch-based with a graphical installer", "Delayed package testing for stability", "Multiple DE editions available", "Hardware detection tool (mhwd)"},
		Cons: []string{"Delayed updates can cause AUR conflicts", "Controversial project management history", "Not as close to vanilla Arch as EndeavourOS"},
	},
	{
		ID: "zorin", Name: "Zorin OS", Version: "17",
		Tagline:     "The alternative to Windows and macOS designed to be easy",
		Homepage:    "https://zorin.com/os",
		DownloadURL: "https://zorin.com/os/download/", ISOSize: "3.3 GB",
		PackageManager: "apt", BasedOn: "Ubuntu", DefaultDE: "gnome",
		Traits: map[string]float64{
			"beginner_friendly": 0.95, "stability": 0.82, "bleeding_edge": 0.15,
			"gaming": 0.52, "dev_tools": 0.55, "multimedia": 0.60,
			"lightweight": 0.38, "customizable": 0.42, "community": 0.55,
			"foss_purity": 0.48, "nvidia_support": 0.70, "hardware_compat": 0.75,
			"laptop_friendly": 0.80, "pkg_deb": 1.0,
		},
		Pros: []string{"Gorgeous out-of-the-box experience", "Windows and macOS layout modes", "Very polished for non-technical users", "Stable Ubuntu LTS base"},
		Cons: []string{"Pro edition costs money for extra layouts", "Smaller community than Mint or Ubuntu", "Limited advanced customization"},
	},
	{
		ID: "elementary", Name: "elementary OS", Version: "8",
		Tagline:     "The thoughtful, capable, and ethical replacement for macOS",
		Homepage:    "https://elementary.io",
		DownloadURL: "https://elementary.io/", ISOSize: "2.8 GB",
		PackageManager: "apt", BasedOn: "Ubuntu", DefaultDE: "pantheon",
		Traits: map[string]float64{
			"beginner_friendly": 0.88, "stability": 0.80, "bleeding_edge": 0.15,
			"gaming": 0.32, "dev_tools": 0.55, "multimedia": 0.60,
			"lightweight": 0.40, "customizable": 0.22, "community": 0.48,
			"foss_purity": 0.68, "nvidia_support": 0.55, "hardware_compat": 0.70,
			"laptop_friendly": 0.85, "pkg_deb": 1.0,
		},
		Pros: []string{"Beautiful, cohesive macOS-like design", "Curated AppCenter", "Strong privacy focus", "Pay-what-you-want model supports devs"},
		Cons: []string{"Very limited customization by design", "Small app ecosystem", "Not for power users or gamers"},
	},
	{
		ID: "mx-linux", Name: "MX Linux", Version: "23",
		Tagline:     "Fast, friendly, and stable — designed for older hardware too",
		Homepage:    "https://mxlinux.org",
		DownloadURL: "https://mxlinux.org/download-links/", ISOSize: "2.2 GB",
		PackageManager: "apt", BasedOn: "Debian", DefaultDE: "xfce",
		Traits: map[string]float64{
			"beginner_friendly": 0.75, "stability": 0.88, "bleeding_edge": 0.15,
			"gaming": 0.30, "dev_tools": 0.58, "multimedia": 0.55,
			"lightweight": 0.85, "customizable": 0.60, "community": 0.65,
			"foss_purity": 0.65, "nvidia_support": 0.48, "hardware_compat": 0.88,
			"laptop_friendly": 0.80, "pkg_deb": 1.0,
		},
		Pros: []string{"Runs great on old and low-spec hardware", "Stable Debian base with backports", "MX Tools simplify system management", "Active community and forums"},
		Cons: []string{"XFCE default looks dated to some", "Not suited for gaming", "Smaller ecosystem than Ubuntu-based distros"},
	},
	{
		ID: "nixos", Name: "NixOS", Version: "24.05",
		Tagline:     "The purely functional Linux distribution — declarative and reproducible",
		Homepage:    "https://nixos.org",
		DownloadURL: "https://nixos.org/download/", ISOSize: "2.5 GB",
		PackageManager: "nix", DefaultDE: "gnome",
		Traits: map[string]float64{
			"beginner_friendly": 0.05, "stability": 0.72, "bleeding_edge": 0.75,
			"gaming": 0.42, "dev_tools": 0.95, "multimedia": 0.48,
			"lightweight": 0.48, "customizable": 0.92, "community": 0.60,
			"foss_purity": 0.80, "nvidia_support": 0.52, "hardware_compat": 0.68,
			"laptop_friendly": 0.58, "pkg_nix": 1.0,
		},
		Pros: []string{"Entire system defined in one config file", "Atomic upgrades with instant rollback", "Reproducible builds and environments", "Largest package repository (Nixpkgs)"},
		Cons: []string{"Steep learning curve — unique paradigm", "Nix language takes effort to learn", "Some software requires workarounds"},
	},
	{
		ID: "void", Name: "Void Linux", Version: "Rolling",
		Tagline:     "An independent distro with runit and musl — fast and minimal",
		Homepage:    "https://voidlinux.org",
		DownloadURL: "https://voidlinux.org/download/", ISOSize: "0.6 GB",
		PackageManager: "xbps", DefaultDE: "none",
		Traits: map[string]float64{
			"beginner_friendly": 0.08, "stability": 0.65, "bleeding_edge": 0.80,
			"gaming": 0.42, "dev_tools": 0.75, "multimedia": 0.48,
			"lightweight": 0.92, "customizable": 0.88, "community": 0.38,
			"foss_purity": 0.78, "nvidia_support": 0.48, "hardware_compat": 0.70,
			"laptop_friendly": 0.52,
		},
		Pros: []string{"Extremely lightweight and fast", "Runit init is simple and fast-booting", "Independent — not based on another distro", "Musl libc option for minimal setups"},
		Cons: []string{"Small community and limited docs", "Requires significant Linux knowledge", "Some software may not support musl"},
	},
}

// Desktops is the full catalog of desktop environments and window managers.
var Desktops = []Desktop{
	{
		ID: "gnome", Name: "GNOME", MinRAMGB: 4,
		Description: "Modern, clean, and focused — a polished workflow-driven desktop",
		Traits: map[string]float64{
			"beginner_friendly": 0.80, "lightweight": 0.25,
			"customizable": 0.40, "windows_like": 0.20, "macos_like": 0.70,
		},
	},
	{
		ID: "kde", Name: "KDE Plasma", MinRAMGB: 4,
		Description: "Feature-rich and deeply customizable — a familiar layout with modern power",
		Traits: map[string]float64{
			"beginner_friendly": 0.75, "lightweight": 0.48,
			"customizable": 0.95, "windows_like": 0.88, "macos_like": 0.30,
		},
	},
	{
		ID: "cinnamon", Name: "Cinnamon", MinRAMGB: 4,
		Description: "Traditional desktop with a modern touch — Linux Mint's flagship DE",
		Traits: map[string]float64{
			"beginner_friendly": 0.92, "lightweight": 0.40,
			"customizable": 0.65, "windows_like": 0.92, "macos_like": 0.15,
		},
	},
	{
		ID: "xfce", Name: "XFCE", MinRAMGB: 2,
		Description: "Lightweight and reliable — fast on any hardware without sacrificing usability",
		Traits: map[string]float64{
			"beginner_friendly": 0.70, "lightweight": 0.88,
			"customizable": 0.60, "windows_like": 0.65, "macos_like": 0.15,
		},
	},
	{
		ID: "cosmic", Name: "COSMIC", MinRAMGB: 4,
		Description: "System76's new Rust-based DE — modern, tiling-capable, and fast",
		Traits: map[string]float64{
			"beginner_friendly": 0.70, "lightweight": 0.52,
			"customizable": 0.82, "windows_like": 0.50, "macos_like": 0.45,
		},
	},
	{
		ID: "mate", Name: "MATE", MinRAMGB: 2,
		Description: "A continuation of GNOME 2 — traditional, stable, and resource-friendly",
		Traits: map[string]float64{
			"beginner_friendly": 0.75, "lightweight": 0.78,
			"customizable": 0.55, "windows_like": 0.72, "macos_like": 0.15,
		},
	},
	{
		ID: "budgie", Name: "Budgie", MinRAMGB: 4,
		Description: "Simple, elegant, and tightly integrated — modern without complexity",
		Traits: map[string]float64{
			"beginner_friendly": 0.85, "lightweight": 0.55,
			"customizable": 0.45, "windows_like": 0.45, "macos_like": 0.55,
		},
	},
	{
		ID: "lxqt", Name: "LXQt", MinRAMGB: 1,
		Description: "Ultra-lightweight Qt-based desktop — maximizes performance on minimal hardware",
		Traits: map[string]float64{
			"beginner_friendly": 0.55, "lightweight": 0.95,
			"customizable": 0.50, "windows_like": 0.70, "macos_like": 0.10,
		},
	},
	{
		ID: "hyprland", Name: "Hyprland", MinRAMGB: 2,
		Description: "A dynamic tiling Wayland compositor — fast, flashy, and endlessly configurable",
		Traits: map[string]float64{
			"beginner_friendly": 0.05, "lightweight": 0.85,
			"customizable": 0.98, "windows_like": 0.05, "macos_like": 0.05,
		},
	},
}
