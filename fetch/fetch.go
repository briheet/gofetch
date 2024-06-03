package fetch

type Versioner interface {
	GetOsVersion() string
}

type Namer interface {
	GetName() string
}

type Timer interface {
	GetTime() string
}

type Packager interface {
	GetPackage() string
}

type Sheller interface {
	GetShellName() string
}

type Resolutioner interface {
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

type Memory interface {
	GetMemory() string
}

type Kernel interface {
	GetKernel() string
}

type Fetcher interface {
	Versioner
	Namer
	Timer
	Packager
	Sheller
	Resolutioner
	Environment
	Terminal
	CPU
	GPU
	Memory
	Kernel
}
