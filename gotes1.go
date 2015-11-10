package main
//使用时间做种子生产随机数。
import (
    "flag"
    "math/rand"
    "time"
    "fmt"
    "github.com/larspensjo/config"
    "log"
    "runtime"
    "strconv"
)
 
var (
    configFile = flag.String("configfile", "config.ini", "General configuration file")
)
 
//topic list
var TOPIC = make(map[string]string)
 
func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    flag.Parse()
 
    //set config file std
    cfg, err := config.ReadDefault(*configFile)
    if err != nil {
        log.Fatalf("Fail to find", *configFile, err)
    }
    //set config file std End
 
    //Initialized topic from the configuration
    if cfg.HasSection("topicArr") {
        section, err := cfg.SectionOptions("topicArr")
        if err == nil {
            for _, v := range section {
                options, err := cfg.String("topicArr", v)
                if err == nil {
                    TOPIC[v] = options
                }
            }
        }
    }
    //Initialized topic from the configuration END
 
    fmt.Println(TOPIC["gjz"])
    fmt.Printf("test %T",TOPIC["gjz"])
    var t float64
    t,err =  strconv.ParseFloat(TOPIC["bdz"],64) 
fmt.Printf(" %T",t)
fmt.Printf("   %.2f\n",t)

    	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<3;i++{
	fmt.Println(r.Intn(100))
	}
}

