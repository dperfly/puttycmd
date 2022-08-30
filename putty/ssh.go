package putty

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
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
	if res, err := session.Output("ls"); err != nil {
		return false
	} else {
		fmt.Println("res", string(res))
		return true
	}
}

func SftpConnect() (sftpClient *sftp.Client, err error) { //参数: 远程服务器用户名, 密码, ip, 端口
	auth := make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig := &ssh.ClientConfig{
		User:    username,
		Auth:    auth,
		Timeout: 10 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	addr := host + ":" + port
	sshClient, err := ssh.Dial("tcp", addr, clientConfig) //连接ssh
	if err != nil {
		fmt.Println(err)
		return
	}

	if sftpClient, err = sftp.NewClient(sshClient); err != nil { //创建客户端
		fmt.Println(err)
		return
	}

	return
}

func GetSession() (*ssh.Session, error) {
	session, err := SSH(host, port, username, password)
	fmt.Println(host, port, username, password)
	if err != nil {

		return nil, err
	}
	return session, nil
}
