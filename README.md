#ivylog (utility)

Another interpretation of working with logging (golang package to improve it).

**Current version: 0.1.0 (Apr 09, 2016)**

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
  ivylog.WriteWarn("1 is less than 2")
}
```

**Result:**
```

```