package api

// TyclContract describes the config written to ~/run/config.tycl.
var TyclContract = `
flexible {
	scripts: objects = strict {
		name: string,
		script: string,
		description: string,
		tags: strings,
		ext: string,
	},
	templates: objects = strict {
		ext: string,
		file: string,
	},
}
`
