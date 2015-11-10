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
//Businessloan 每月月供额=(贷款本金÷还款月数)+(贷款本金-已归还本金累计额)×月利率
func blv(x,y,a,b,c float64) (float64, float64) {
	var blm,gjm float64
	blm = (x/ 240) + (x-0)*(c*a)/100/12
	gjm = (y / 240) + (y-0)*b/100/12
return blm,gjm
}
var (
    configFile = flag.String("configfile", "config.ini", "General configuration file")
)

var TOPIC = make(map[string]string)

func main() {
 runtime.GOMAXPROCS(runtime.NumCPU())
    flag.Parse()
 
    //set config file std
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
  
  fmt.Printf("\n贷款总额  按揭部分%s元；公积金部分%s元；按揭利率及折扣%s(%s)；公积金利率%s\n" ,
  TOPIC["bdz"],TOPIC["gjz"],TOPIC["sv"],TOPIC["zk"],TOPIC["gv"])

var sv,gv,zk,blw, gjw, blm, gjm,djb,djg float64
p := [] float64 {}
q := [] float64 {}


 blw,err =  strconv.ParseFloat(TOPIC["bdz"],64) 
 gjw,err =  strconv.ParseFloat(TOPIC["gjz"],64) 
 sv,err =  strconv.ParseFloat(TOPIC["sv"],64) 
 gv,err =  strconv.ParseFloat(TOPIC["gv"],64) 
 zk,err =  strconv.ParseFloat(TOPIC["zk"],64) 

blm, gjm = blv(blw,gjw,sv,gv,zk)

//每月月供递减额=每月应还本金×月利率=贷款本金÷还款月数×月利率
	djb=(blw/240)*(sv*zk)/100/12
	djg=(gjw/240)*(gv)/100/12

fmt.Printf("\n每月应还本金=贷款本金÷还款月数；按揭部分%.2f元；公积金部分%.2f元\n",blw/240,gjw/240)

//总利息=〔(总贷款额÷还款月数+总贷款额×月利率)+总贷款额÷还款月数×(1+月利率)〕÷2×还款月数-总贷款额

fmt.Printf("\n总利息按揭部分：%.2f元；总利息公积金部分：%.2f元\n", ((blw/240+blw*sv*zk/1200)+blw*(1+sv*zk/1200)/240)*240/2-blw,
		((gjw/240+gjw*gv/1200)+gjw*(1+gv/1200)/240)*240/2-gjw)

fmt.Printf("\n按揭部分每月递减值%.2f元；公积金部分每月递减值 %.2f元 。",djb,djg)

fmt.Printf("   按揭部分月利率：%.2f元；公积金部分月利率 ：%.2f 元 \n ",sv*zk/12,gv/12)

fmt.Printf("\n还款时间  2015年/11月 [ 第 1 期 ] = 按揭部分 %.2f元； 公积金部分 %.2f元； 当月还款总额  %.2f元\n",
				blm, gjm,blm+gjm)

	for i := 0; i < 240; i++ {
		blm  =blm- djb
		p = append(p, blm)
	}

	for i := 0; i < 239; i++ {
		gjm  =gjm- djg
		q = append(q, gjm)
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
		//每月应还利息=剩余本金×月利率=(贷款本金-已归还本金累计额)×月利率
fmt.Printf("本月应还利息 按揭部分: %.2f元；公积金部分 %.2f元\n"  , 
(blw-blw*float64(i+2)/240)*(sv*zk)/100/12,(gjw-gjw*float64(i+2)/240)*(gv)/100/12)
}
}