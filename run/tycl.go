package runlib

var tyclContract = `
strict {
	scripts: objects = strict {
		name: string,
		script: string,
		description: string,
		tags: strings,
		ext: string,
	},
	templates: objects = strict {
		ext: string,
		template: string,
	},
}
`
