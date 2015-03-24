package hotomata

import (
	"fmt"
	"strconv"

	"golang.org/x/crypto/ssh"
)

type Machine struct {
	Hostname  string
	Port      int
	SSHConfig *ssh.Config
}

func Machines(hosts []map[string]string) []*Machine {
	var machines []*Machines
	for _, host := range hosts {
		hostname := host["name"]
		if h, ok := host["ssh_hostname"]; ok {
			hostname = h
		}

		port := 22
		if p, ok = host["ssh_port"]; ok {
			port, err = strconv.Atoi(p)
			if err != nil {
				fmt.Printf("Error parsing port for host [%s]\n", hostname)
				panic(err)
			}
		}

		username := "root"
		if u, ok = host["ssh_username"]; ok {
			username = u
		}

		sshAuthMethods := []ssh.AuthMethod{}
		// Password
		if password, ok := host["ssh_password"]; ok {
			sshAuthMethods = append(sshAuthMethods, ssh.Password(password))
		}
		// Key provided
		keyLocation, ok := host["ssh_key"]
		// Try default key
		if !ok {
			defaultKeyLocation := path.Join(os.Getenv("HOME"), ".ssh", "id_rsa")
			if _, err := os.Stat(defaultKeyLocation); err != nil {
				keyLocation = defaultKeyLocation
			}
		}
		if keyLocation != "" {
			authMethod, err := clientKeyAuth(keyLocation)
			if err != nil {
				fmt.Printf("Error loading key for host [%s]\n", hostname)
				panic(err)
			}
			sshAuthMethods = append(sshAuthMethods, authMethod)
		}

		config := &ssh.ClientConfig{
			User: username,
			Auth: sshAuthMethods,
		}
		machines = append(machines, &Machine{
			Hostname: hostname,
			Port:     port,
			Config:   config,
		})
	}
	return machines
}

func sshAuth(keyLocation string) (ssh.AuthMethod, error) {
	buf, err := ioutil.ReadFile(keyLocation)
	if err != nil {
		return nil, err
	}
	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(key), nil
}
