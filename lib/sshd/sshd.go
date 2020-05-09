package sshd

import (
	"bufio"
	"io"
	"strings"

	"golang.org/x/crypto/ssh"
)

type Client struct {
	client       *ssh.Client
	config       *ssh.ClientConfig
	host         string
	port         string
	rootPassword string
}

func (c *Client) EnrollWithKey(username string, keyData string) error {
	k, err := c.publicKey(keyData)
	if err != nil {
		return err
	}
	c.config = &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			k,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return nil
}

func (c *Client) publicKey(keyData string) (ssh.AuthMethod, error) {
	signer, err := ssh.ParsePrivateKey([]byte(keyData))
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(signer), nil
}

func (c *Client) EnrollWithPassword(username string, password string) {
	c.config = &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

func (c *Client) Connect(host string, port string) error {
	c.host = host
	c.port = port
	conn, err := ssh.Dial("tcp", c.host+":"+c.port, c.config)
	c.client = conn
	return err
}

//e.g. output, err := remoteRun("root", "MY_IP", "PRIVATE_KEY", "ls")
func (c Client) SyncCommand(cmds []string) ([]byte, error) {
	session, err := c.client.NewSession()
	if err != nil {
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	err = session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		session.Close()
		return []byte{}, err
	}

	in, err := session.StdinPipe()
	if err != nil {
		session.Close()
		return nil, err
	}

	out, err := session.StdoutPipe()
	if err != nil {
		session.Close()
		return nil, err
	}

	var output []byte

	go func(in io.WriteCloser, out io.Reader, output *[]byte) {
		var (
			line string
			r    = bufio.NewReader(out)
		)
		for {
			b, err := r.ReadByte()
			if err != nil {
				break
			}

			*output = append(*output, b)

			if b == byte('\n') {
				line = ""
				continue
			}

			line += string(b)

			if strings.HasPrefix(line, "[sudo] password for ") && strings.HasSuffix(line, ": ") {
				_, err = in.Write([]byte(c.rootPassword + "\n"))
				if err != nil {
					break
				}
			}
		}
	}(in, out, &output)

	cmd := strings.Join(cmds, "; ")
	_, err = session.Output(cmd)
	if err != nil {
		session.Close()
		return []byte{}, err
	}
	session.Close()
	return output, nil
}
