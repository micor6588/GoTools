package config

import (
	"errors"
	"os"
)

/**
* 截入配置文件并读取其中内容
 * @Authr micor
 * @Date  2019/12/08
 *
 * 与INI配置文风格一样 根据顺序读取文件和每一行 如果在行首出现了;号，则认为是配置文件的注释
  当INI不规范时 如[mysql] 的注释被写为[mysql 则会返回错误
* 本包的配置文件是严格区分大小写的 需要禁止区分大小写 将在后期加入或自行加入
*本包是对文件的相关操作：进行文件检查和打开文件
*/

// FileCheckExists 检查文件或文件夹是否存在
func FileCheckExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

// NewFileReader 打开一个文件句柄
func NewFileReader(filePath string) (*os.File, error) {
	if !FileCheckExists(filePath) {
		return nil, errors.New("Error:File not exists:" + filePath)
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return f, nil
}
