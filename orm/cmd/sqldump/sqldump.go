package sqldump

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/fzf-labs/fctl/utils"
)

type SQLDump struct {
	dsn           string // 数据库连接
	outPutPath    string // 输出路径
	targetTables  string // 指定表
	fileOverwrite bool   // 是否覆盖
}

func NewSQLDump(dsn, outPutPath, targetTables string, fileCover bool) *SQLDump {
	return &SQLDump{dsn: dsn, outPutPath: outPutPath, targetTables: targetTables, fileOverwrite: fileCover}
}

func (s *SQLDump) Run() {
	// 连接数据库
	db := utils.NewDB(s.dsn)
	// 查找命令的可执行文件
	_, err := exec.LookPath("pg_dump")
	if err != nil {
		log.Print("command pg_dump not found,please install")
		return
	}
	var tableList []string
	if s.targetTables != "" {
		tableList = strings.Split(s.targetTables, ",")
	} else {
		tableList, err = db.Migrator().GetTables()
		if err != nil {
			return
		}
	}
	dsnParse := PostgresDsnParse(s.dsn)
	outPutPath := filepath.Join(strings.Trim(s.outPutPath, "/"), dsnParse.Dbname)
	err = os.MkdirAll(outPutPath, os.ModePerm)
	if err != nil {
		log.Print("DumpPostgres create path err:", err)
		return
	}
	for _, v := range tableList {
		outFile := filepath.Join(outPutPath, fmt.Sprintf("%s.sql", v))
		cmdArgs := []string{
			"-h", dsnParse.Host,
			"-p", strconv.Itoa(dsnParse.Port),
			"-U", dsnParse.User,
			"-s", dsnParse.Dbname,
			"-t", v,
		}
		// 创建一个 Cmd 对象来表示将要执行的命令
		cmd := exec.Command("pg_dump", cmdArgs...)
		// 添加一个环境变量到命令中
		cmd.Env = append(cmd.Env, fmt.Sprintf("PGPASSWORD=%s", dsnParse.Password))
		// 执行命令，并捕获输出和错误信息
		output, err := cmd.Output()
		if err != nil {
			log.Print("cmd exec err2:", err)
			return
		}
		if _, err2 := os.Stat(outFile); !os.IsNotExist(err2) {
			if !s.fileOverwrite {
				log.Printf("postgres sql file %s already exists,if you want new,please manually delete!\n", outFile)
				return
			}
		}
		err2 := utils.WriteContentCover(outFile, remove(string(output)))
		if err2 != nil {
			log.Print("postgres sql file write err2:", err2)
			return
		}
	}
}

type PostgresDsn struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

// PostgresDsnParse  数据库解析
func PostgresDsnParse(dsn string) *PostgresDsn {
	result := new(PostgresDsn)
	// 分割连接字符串
	params := strings.Split(dsn, " ")

	// 解析参数
	for _, param := range params {
		keyValue := strings.Split(param, "=")
		if len(keyValue) != 2 {
			continue
		}
		key := keyValue[0]
		value := keyValue[1]
		switch key {
		case "host":
			result.Host = value
		case "port":
			if p, err := strconv.Atoi(value); err == nil {
				result.Port = p
			}
		case "user":
			result.User = value
		case "password":
			result.Password = value
		case "dbname":
			result.Dbname = value
		}
	}
	return result
}

func remove(str string) string {
	var result string
	reader := strings.NewReader(str)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "--") || strings.HasPrefix(line, "SELECT") || strings.HasPrefix(line, "SET") || regexp.MustCompile(`(ALTER TABLE .*? OWNER TO postgres)`).MatchString(line) {
			continue
		}
		result += fmt.Sprintln(line)
	}
	return result
}
