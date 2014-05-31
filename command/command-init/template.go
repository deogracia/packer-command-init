package commandinit

const defaultTemplate = `
{
	"description": "This is a default template!",
	"builders": [{
		"type": "virtualbox-iso",
		"iso_checksum_type": "md5",
		"iso_checksum": "2DEGDH5YXH5D",
		"iso_url": "http://example.com/dvd.iso",
		"shutdown_command": "shutdown -hP now",
		"ssh_username": "packer"
	}]
}
`
