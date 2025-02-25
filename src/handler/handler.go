package handler

import (
	"io"
	"os"

	"github.com/anhtu2808/VPSQuickSetup/src/config"
	"golang.org/x/crypto/ssh"
)

func RunScriptRemote(scriptFunc func() (string, error), config *config.Config) error {
	ip := config.IP
	user := config.User
	password := config.Password
	// Get script from function
	script, err := scriptFunc()
	if err != nil {
		return err
	}

	configSSH := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr := ip + ":22"

	// Connect to the server
	client, err := ssh.Dial("tcp", addr, configSSH)
	if err != nil {
		return err
	}
	defer client.Close()

	// Create a new session
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// Tạo pipe cho stdout và stderr
	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		return err
	}

	// In log realtime từ stdout và stderr
	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	// Lấy stdin pipe để truyền nội dung script
	stdin, err := session.StdinPipe()
	if err != nil {
		return err
	}

	// Chạy lệnh "bash -s" trên remote
	if err := session.Start("bash -s"); err != nil {
		return err
	}

	// Ghi nội dung script vào stdin và đóng pipe
	go func() {
		defer stdin.Close()
		stdin.Write([]byte(script))
	}()

	// Chờ cho đến khi session kết thúc
	if err := session.Wait(); err != nil {
		return err
	}

	return nil
}
