package pkg

/*
NMAP的正确食用方式
*/
import (
	"fmt"
	"strings"
	"time"

	nmap "github.com/lair-framework/go-nmap"
)

type Scanner struct {
	Cmd     string   `json:"cmd"`
	Args    []string `json:"args"`
	Result  string   `json:"result"`
	Command string   `json:"command"`
	Path    string   `json:"path"`
}

func NewScanner() *Scanner {
	return &Scanner{
		Cmd: "nmap",
	}
}

// 合并命令
func (this *Scanner) GetCommand() {
	if this.Cmd == "" {
		this.Cmd = "nmap"
	}

	this.SetXmlSave()
	this.Command = fmt.Sprintf("%s %s", this.Cmd, strings.Join(this.Args, " "))
}

func (this *Scanner) SetArgs(args ...string) *Scanner {
	this.Args = append(this.Args, args...)
	return this
}

// help命令解释
func (this *Scanner) Help() (string, error) {
	s, e := ExecCommand("nmap -h")
	return s, e
}

// 执行scan命令并获取结果
func (this *Scanner) Scan() ([]byte, error) {
	var (
		result []byte
		err    error
	)
	// 获取命令
	this.GetCommand()
	// 执行命令
	es, err := ExecCommand(this.Command)
	if err != nil {
		return []byte(es), err
	}
	result, err = ReadFile(this.Path)
	if err != nil {
		return result, err
	}
	return result, err
}

func (this *Scanner) SetXmlSave() {
	now := time.Now().Nanosecond
	this.SetArgs(fmt.Sprintf("-oX /tmp/%d", now))
	this.Path = fmt.Sprintf("/tmp/%d", now)
}

func (this *Scanner) Parse() (*nmap.NmapRun, error) {
	var result *nmap.NmapRun
	info, err := this.Scan()
	if err != nil {
		return result, err
	}
	result, err = nmap.Parse(info)
	return result, err
}
