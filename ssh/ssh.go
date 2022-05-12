package ssh

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// SshClient ssh客户端对象
type SshClient struct {
	host     string
	username string
	password string
	Client   *ssh.Client
}

// NewSshClient 创建一个SSH客户端对象
func (sshClient *SshClient) NewSshClient(host, username, password string) *ssh.Client {
	// 免费的测试服务器
	flag.StringVar(&sshClient.host, "host", host, "通过ssh2登录linux的ip地址")
	flag.StringVar(&sshClient.username, "username", username, "通过ssh2登录linux的用户名")
	flag.StringVar(&sshClient.password, "password", password, "通过ssh2登录linux的密码")
	flag.Parse()

	auth := make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(sshClient.password))

	clientConfig := &ssh.ClientConfig{
		User:    sshClient.username,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	addr := sshClient.host + ":22"
	client, err := ssh.Dial("tcp", addr, clientConfig) //连接ssh
	if err != nil {
		log.Fatal("连接ssh失败", err)
	}

	sshClient.Client = client
	return client
}

func (sshClient *SshClient) RunCmd(cmd string) string {
	session, err := sshClient.Client.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	runResult, err := session.CombinedOutput(cmd)
	if err != nil {
		panic(err)
	}
	return string(runResult)
}

func (sshClient *SshClient) NewFtpClient() *sftp.Client {
	ftpClient, err := sftp.NewClient(sshClient.Client)
	if err != nil {
		fmt.Println("创建ftp客户端失败", err)
		panic(err)
	}
	return ftpClient
}

func (sshClient *SshClient) MkdirAll(path string) {
	ftpClient := sshClient.NewFtpClient()
	defer ftpClient.Close()
	err := ftpClient.MkdirAll(path)
	if err != nil {
		fmt.Println("创建文件夹失败", err)
		panic(err)
	}
}
func (sshClient *SshClient) UploadFile(localPath string, remoteDir string, remoteFileName string) {
	ftpClient, err := sftp.NewClient(sshClient.Client)
	if err != nil {
		fmt.Println("创建ftp客户端失败", err)
		panic(err)
	}

	defer ftpClient.Close()

	fmt.Println(localPath, remoteFileName)
	srcFile, err := os.Open(localPath)
	if err != nil {
		fmt.Println("打开文件失败", err)
		panic(err)
	}
	defer srcFile.Close()

	fn := path.Join(remoteDir, remoteFileName)

	err = ftpClient.MkdirAll(remoteDir)
	if err != nil {
		fmt.Println("创建文件夹失败", err)
		panic(err)
	}

	dstFile, e := ftpClient.Create(fn)
	if e != nil {
		fmt.Println("创建文件失败", e)
		panic(e)
	}
	defer dstFile.Close()

	buffer := make([]byte, 1024000)
	for {
		n, err := srcFile.Read(buffer)
		dstFile.Write(buffer[:n])
		//注意，由于文件大小不定，不可直接使用buffer，否则会在文件末尾重复写入，以填充1024的整数倍
		if err != nil {
			if err == io.EOF {
				fmt.Println("已读取到文件末尾")
				break
			} else {
				fmt.Println("读取文件出错", err)
				panic(err)
			}
		}
	}
}

func (sshClient *SshClient) DownloadFile(remotePath string, localDir string, localFilename string) {
	ftpClient, err := sftp.NewClient(sshClient.Client)
	if err != nil {
		fmt.Println("创建ftp客户端失败", err)
		panic(err)
	}

	defer ftpClient.Close()

	srcFile, err := ftpClient.Open(remotePath)
	if err != nil {
		fmt.Println("文件读取失败", err)
		panic(err)
	}
	defer srcFile.Close()

	dstFile, e := os.Create(path.Join(localDir, localFilename))
	if e != nil {
		fmt.Println("文件创建失败", e)
		panic(e)
	}
	defer dstFile.Close()
	if _, err = srcFile.WriteTo(dstFile); err != nil {
		fmt.Println("文件写入失败", err)
		panic(err)
	}
	fmt.Println("文件下载成功")
}
