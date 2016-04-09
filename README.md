#ivylog (utility)

Another interpretation of working with logging in golang.

**Current version: 0.1.0 (Apr 10, 2016)**

##Quick Start

Install package:
```
go get "github.com/VitaliyPetroff/ivylog"
```

Set logging setttings:
```
var mylog ivylog.LogSettings

func init() {
  mylog.File_path = ""                  // application dir
  mylog.Success_file_name = "mylog.log" // success log file name
  mylog.Error_file_name = ""            // same as success file name
  mylog.Write_date = true
  mylog.Write_time = true

  err := ivylog.InitLog(mylog) // initialization
  if err != nil {
    fmt.Println(err.Error())
  }
}
```

Use it:
```
func main() {
  mylog.WriteWarn("1 is less than 2")
}
```

**Result:**
```
2016-04-10 01:16:23 WARNING: 1 is less than 2
```