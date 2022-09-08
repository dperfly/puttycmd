package putty

import (
	"errors"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"time"
)

func SSH(host string, port string, username string, password string) (*ssh.Session, error) {
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", host, port), &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	session, err := client.NewSession()
	if err != nil {

		return nil, err
	}

	return session, nil
}

func CheckSSH(session *ssh.Session) bool {
	defer session.Close()
	if _, err := session.Output("ls"); err != nil {
		return false
	} else {
		return true
	}
}

func SftpConnect() (*sftp.Client, error) { //参数: 远程服务器用户名, 密码, ip, 端口
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("SftpConnect err: ", err)
		}
	}()
	auth := make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig := &ssh.ClientConfig{
		User:    username,
		Auth:    auth,
		Timeout: 1 * time.Minute,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	addr := host + ":" + port
	sshClient, err := ssh.Dial("tcp", addr, clientConfig) //连接ssh
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ssh.Dial :%s", err.Error()))
	}
	if sftpClient, err := sftp.NewClient(sshClient); err != nil { //创建客户端
		return nil, errors.New(fmt.Sprintf("sftp.NewClient :%s", err.Error()))
	} else {
		return sftpClient, nil
	}
}

func GetSession() (*ssh.Session, error) {
	session, err := SSH(host, port, username, password)
	if err != nil {
		return nil, err
	}
	return session, nil
}
