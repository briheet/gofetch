package fetch

type Version interface {
	GetOsVersion() string
}

type Name interface {
	GetName() string
}

type Time interface {
	GetTime() string
}

type Package interface {
	GetPackage() string
}

type Shell interface {
	GetShellName() string
}

type Resolution interface {
	GetResolution() string
}

type Environment interface {
	GetEnvironment() string
}

type Terminal interface {
	GetTerminalName() string
}

type CPU interface {
	GetCPU() string
}

type GPU interface {
	GetGPU() string
}

type Usage interface {
	GetUsage() string
}

type Kernel interface {
	GetKernel() string
}

type Fetch interface {
	Version
	Name
	Time
	Package
	Shell
	Resolution
	Environment
	Terminal
	CPU
	GPU
	Usage
	Kernel
}
