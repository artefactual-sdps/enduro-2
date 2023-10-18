package sftp

import (
	"os"
	"path/filepath"
)

type Config struct {
	// Host address, e.g. 127.0.0.1 (default), sftp.example.org.
	Host string

	// User name.
	User string

	// Host port (default: 22).
	Port string

	// Path to known_hosts file as per https://linux.die.net/man/8/sshd
	// "SSH_KNOWN_HOSTS FILE FORMAT" (default: "$HOME/.ssh/known_hosts"). The
	// known_hosts file must include the public key of the SFTP server for
	// authentication to succeed.
	KnownHostsFile string

	// Private key used for authentication.
	PrivateKey PrivateKey

	// Default directory on SFTP server for file transfers.
	RemoteDir string
}

type PrivateKey struct {
	// Path to private key file used for authentication (default:
	// "$HOME/.ssh/id_rsa")
	Path string

	// Passphrase (if any) used to decrypt private key.
	Passphrase string
}

// SetDefaults sets default values for some configs.
func (c *Config) SetDefaults() {
	if c.Host == "" {
		c.Host = "localhost"
	}

	if c.Port == "" {
		c.Port = "22"
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return // Don't set default paths if homedir is unknown.
	}

	if c.KnownHostsFile == "" {
		c.KnownHostsFile = filepath.Join(home, ".ssh", "known_hosts")
	}

	if c.PrivateKey.Path == "" {
		c.PrivateKey.Path = filepath.Join(home, ".ssh", "id_rsa")
	}
}
