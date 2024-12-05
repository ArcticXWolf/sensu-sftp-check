package main

import (
	"fmt"

	"github.com/pkg/sftp"
	"github.com/sensu/sensu-go/types"
	"github.com/sensu/sensu-plugin-sdk/sensu"
	"golang.org/x/crypto/ssh"
)

// Config represents the check plugin config.
type Config struct {
	sensu.PluginConfig
	sftpAddress  string
	sftpUsername string
	sftpPassword string
	sftpHostkey  string
	sftpFilePath string
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "sensu-sftp-check",
			Short:    "SFTP check for Sensu",
			Keyspace: "sensu.io/plugins/sensu-sftp-check/config",
		},
	}

	options = []sensu.ConfigOption{
		&sensu.PluginConfigOption[string]{
			Path:      "sftp-address",
			Env:       "SFTP_ADDRESS",
			Argument:  "sftp-address",
			Shorthand: "a",
			Secret:    false,
			Usage:     "Address of the sftp server to connect to",
			Value:     &plugin.sftpAddress,
		},
		&sensu.PluginConfigOption[string]{
			Path:      "sftp-username",
			Env:       "SFTP_USERNAME",
			Argument:  "sftp-username",
			Shorthand: "u",
			Secret:    true,
			Usage:     "Username of the sftp server to connect to",
			Value:     &plugin.sftpUsername,
		},
		&sensu.PluginConfigOption[string]{
			Path:      "sftp-password",
			Env:       "SFTP_PASSWORD",
			Argument:  "sftp-password",
			Shorthand: "c",
			Secret:    true,
			Usage:     "Password of the sftp server to connect to",
			Value:     &plugin.sftpPassword,
		},
		&sensu.PluginConfigOption[string]{
			Path:      "sftp-hostkey",
			Env:       "SFTP_HOSTKEY",
			Argument:  "sftp-hostkey",
			Shorthand: "k",
			Secret:    true,
			Usage:     "Hostkey of the sftp server to connect to",
			Value:     &plugin.sftpHostkey,
		},
		&sensu.PluginConfigOption[string]{
			Path:      "sftp-filepath",
			Env:       "SFTP_FILEPATH",
			Argument:  "sftp-filepath",
			Shorthand: "f",
			Secret:    false,
			Usage:     "Filepath to check on the sftp server, if nil then we just connect to the server",
			Value:     &plugin.sftpFilePath,
		},
	}
)

func main() {
	check := sensu.NewCheck(&plugin.PluginConfig, options, validateFunction, executeFunction, false)
	check.Execute()
}

func validateFunction(event *types.Event) (int, error) {
	return 0, nil
}

func executeFunction(event *types.Event) (int, error) {
	fmt.Println("Ssh connecting")
	sshclient, err := connectToHost()
	if err != nil {
		return 2, err
	}

	fmt.Println("Sftp connecting")
	client, err := sftp.NewClient(sshclient)
	if err != nil {
		return 2, err
	}
	defer client.Close()

	fmt.Println("Get WD")
	_, err = client.Getwd()
	if err != nil {
		return 2, err
	}
	fmt.Println("Done")

	return 0, nil
}

func connectToHost() (*ssh.Client, error) {
	fmt.Printf("Address: %s\nUser: %s\nHostkey: %s\n", plugin.sftpAddress, plugin.sftpUsername, plugin.sftpHostkey)

	sshConfig := &ssh.ClientConfig{
		User: plugin.sftpUsername,
		Auth: []ssh.AuthMethod{ssh.Password(plugin.sftpPassword)},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", plugin.sftpAddress, sshConfig)
	if err != nil {
		return nil, err
	}

	return client, nil
}
