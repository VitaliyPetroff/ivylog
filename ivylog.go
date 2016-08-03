// Another interpretation of working with logging in golang.
package ivylog

import (
	"errors"
	"os"
	"time"
)

// Settings for log file.
// If File_path is empty file will be cteated in current dir.
// If Success_file_name same as Error_file_name or one of them not set,
// all messages will be written in one file.
// If Write_date set to false, time also will not be written.
type LogSettings struct {
	File_path         string
	Success_file_name string
	Error_file_name   string
	Write_date        bool
	Write_time        bool
}

var log_settings LogSettings

//	Initialization log settings.
//	Check or create file dir, check or create files.
func InitLog(lsettings LogSettings) error {
	var err error

	log_settings = lsettings

	if log_settings.File_path == "" {
		curdir, err := os.Getwd()
		if err != nil {
			return err
		}
		log_settings.File_path = curdir + "/"
	} else {
		if _, err := os.Stat(log_settings.File_path); os.IsNotExist(err) {
			err = os.Mkdir(log_settings.File_path, 0755)
			if err != nil {
				return err
			}
		}
	}

	if (log_settings.Success_file_name == log_settings.Error_file_name) && log_settings.Success_file_name == "" {
		return errors.New("Log file name not set")
	}

	if log_settings.Success_file_name == "" {
		log_settings.Success_file_name = log_settings.Error_file_name
	} else if log_settings.Error_file_name == "" {
		log_settings.Error_file_name = log_settings.Success_file_name
	}

	flog, err := os.OpenFile(log_settings.File_path+log_settings.Success_file_name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer flog.Close()

	if log_settings.Success_file_name != log_settings.Error_file_name {
		elog, err := os.OpenFile(log_settings.File_path+log_settings.Error_file_name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		defer elog.Close()
	}

	return nil
}

// Write info message to success log file.
func (l LogSettings) WriteInfo(info_message string) {
	flog, err := os.OpenFile(log_settings.File_path+log_settings.Success_file_name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	defer flog.Close()

	log_message := setTime() + " " + "INFO: " + info_message + "\n"

	flog.WriteString(log_message)
}

// Write err message to error log file.
func (l LogSettings) WriteErr(err_message string) {
	elog, err := os.OpenFile(log_settings.File_path+log_settings.Error_file_name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	defer elog.Close()

	log_message := setTime() + " " + "ERROR: " + err_message + "\n"

	elog.WriteString(log_message)
}

// Write warning message to error log file.
func (l LogSettings) WriteWarn(warn_message string) {
	wlog, err := os.OpenFile(log_settings.File_path+log_settings.Error_file_name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	defer wlog.Close()

	log_message := setTime() + " " + "WARNING: " + warn_message + "\n"

	wlog.WriteString(log_message)
}

func setTime() string {
	if log_settings.Write_date {
		if log_settings.Write_time {
			return time.Now().Format("2006-01-02 15:04:05")
		} else {
			return time.Now().Format("2006-01-02")
		}
	}
	return ""
}
