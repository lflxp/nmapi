package models

/*
NMAP的正确食用方式
*/
import (
	"fmt"
	"strings"

	nmap "github.com/lair-framework/go-nmap"
	"github.com/lflxp/nmapi/pkg"
)

type Scanner struct {
	Cmd     string   `json:"cmd"`
	Args    []string `json:"args"`
	Result  string   `json:"result"`
	Command string   `json:"command"`
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
func (this *Scanner) Help() ([]byte, error) {
	s, e := pkg.ExecCommand("nmap -h")
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
	// fmt.Println(this.Command)
	result, err = pkg.ExecCommand(this.Command)

	return result, err
}

func (this *Scanner) SetXmlSave() {
	this.SetArgs("-oX -")
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
