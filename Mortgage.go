//等额本金还款法:
//总利息=〔(总贷款额÷还款月数+总贷款额×月利率)+总贷款额÷还款月数×(1+月利率)〕÷2×还款月数-总贷款额
package main
import (
	"fmt"
        "flag"
         "github.com/larspensjo/config"
         "log"
         "runtime"
	 "strconv"
)
var (
    configFile = flag.String("configfile", "config.ini", "General configuration file")
)
var TOPIC = make(map[string]string)
func strcon(x string) (y float64) {
        var err error
	 y,err=  strconv.ParseFloat(x,64)
	 if err != nil {
        fmt.Println("Fail to find")
	}
return y
}

func main() {

 runtime.GOMAXPROCS(runtime.NumCPU())
    flag.Parse()

    cfg, err := config.ReadDefault(*configFile)
    if err != nil {
        log.Fatalf("Fail to find", *configFile, err)
    }
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
 p,q := [] float64 {}, [] float64 {}
var sv,gv,zk,blw, gjw float64
blw,gjw =strcon(TOPIC["bdz"]) ,strcon(TOPIC["gjz"])
sv,gv,zk =  strcon(TOPIC["sv"]),strcon(TOPIC["gv"]) ,strcon(TOPIC["zk"])
zlb,zlg:=((blw/240+blw*sv*zk/1200)+blw*(1+sv*zk/1200)/240)*240/2-blw,((gjw/240+gjw*gv/1200)+gjw*(1+gv/1200)/240)*240/2-gjw
b,g:=(blw/ 240) + blw*(sv*zk)/1200,(gjw/ 240) + gjw*gv/1200

fmt.Printf("\n贷款总额  按揭部分%.2f元；公积金部分%.2f元；按揭利率及折扣%.2f(%.2f)；公积金利率%.2f\n" ,blw,gjw,sv,zk,gv)
fmt.Printf("\n每月应还本金=贷款本金÷还款月数；按揭部分%.2f元；公积金部分%.2f元\n",blw/240,gjw/240)
fmt.Printf("\n总利息按揭部分：%.2f元；总利息公积金部分：%.2f元\n",zlb,zlg)
fmt.Printf("\n按揭部分每月递减值%.2f元；公积金部分每月递减值 %.2f元 。",(blw/240)*(sv*zk)/100/12,(gjw/240)*(gv)/100/12)
fmt.Printf("按揭部分月利率：%.2f元；公积金部分月利率: %.2f 元 \n ",sv*zk/12,gv/12)
fmt.Printf("\n还款时间  2015年/11月 [ 第 1 期 ] = 按揭部分 %.2f元；公积金部分 %.2f元；当月还款总额  %.2f元\n",b,g,b+g)

for i := 0; i < 240; i++ {
		b  =b- (blw/240)*(sv*zk)/100/12
		g  =g- (gjw/240)*(gv)/100/12
		p = append(p, b)
		q = append(q, g)
	}

	for i,t := 0,0; i < len(p)-1; i++ {
		
		if i%12 == 0 {
			t = 12
		} else {
			t = i % 12
		}
		fmt.Printf(" \n还款时间 %d年/%d月",  (i+11)/12+2015, t)
		fmt.Printf(" [ 第 %d 期 ] ： 还款按揭 %.2f；公积金 %.2f元； 总额  %.2f元\n",
				i+2, p[i],q[i],q[i]+p[i])
		fmt.Printf("本月应还利息 按揭部分: %.2f元；公积金部分 %.2f元\n"  , 
		(blw-blw*float64(i+2)/240)*(sv*zk)/100/12,(gjw-gjw*float64(i+2)/240)*(gv)/100/12)
}
}