package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/big"
	_ "math/big"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

var sqlH = "dexcharts:Eibeu4gahquuch2EiRie@tcp(10.64.45.1)/dex"

var dbHandleR *sql.DB // global database handle accessible from all the threads

var swapEvents = make(chan *NeoLPSwap)

type Cs struct {
	RpcNode  string
	AManager common.Address
}

type PairInfo struct {
	Pair0        common.Address
	Pair0Name    string
	Pair0Balance *big.Int
	Pair1        common.Address
	Pair1Name    string
	Pair1Balance *big.Int
	LpTokens     *big.Int
}

var Config = Cs{RpcNode: "ws://10.64.45.1:9660/ext/bc/C/ws", AManager: common.HexToAddress("0xA08db6fe8D16067E3a85073cADFf5e4660ce8200")}
var MainNet = Cs{"ws://10.64.45.1:9660/ext/bc/C/ws", common.Address{0x0}}
var Flare = Cs{"ws://10.64.45.2:9650/ext/bc/C/ws", common.Address{0x0}}
var HBAR = Cs{"ws://10.64.45.2:9650/ext/bc/C/ws", common.Address{0x0}}
var NEAR = Cs{"ws://10.64.45.2:9650/ext/bc/C/ws", common.Address{0x0}}

type PriceResponse struct {
	poolAddr string
	c0Addr   string
	c1Addr   string
	ts       string
	price    string
}

var configs = make(map[int]Cs)

var callOptsFac bind.CallOpts

var wlClient2 *ethclient.Client

var info = make(map[common.Address]PairInfo)

func main() {

	configs[0] = Config
	configs[1] = MainNet
	//configs[2]=Flare
	//configs[3]=HBAR
	//configs[4]=NEAR

	http.HandleFunc("/", httpHandler)

	var err error
	check(err)

	dbHandleR, err = sql.Open("mysql", sqlH)

	check(err)

	wlClient2, _ = ethclient.Dial(MainNet.RpcNode)

	vactory, err2 := NewNeoF(Config.AManager, wlClient2)

	check(err2)

	callOptsFac = bind.CallOpts{Pending: false, From: common.HexToAddress("0xDA6FF90214CB6D0CEee462f2D788b2556E657422"), Context: context.Background()}

	apl, err := vactory.AllPairsLength(&callOptsFac)

	check(err)

	var i int64
	//i = 0

	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	var errors = make(map[common.Address]error)
	var clients = make(map[common.Address]*NeoLP)

	for i = 0; i < apl.Int64(); i++ {
		addr, err := vactory.AllPairs(&callOptsFac, big.NewInt(i))
		toAdd, _ := NewNeoLP(addr, wlClient2)
		clients[addr] = toAdd
		_, errors[addr] = clients[addr].WatchSwap(watchOpts, swapEvents, nil, nil)
		info[addr] = getPairInfo(clients[addr], addr)

		check(err)

	}

	//log.Println(clients)
	//log.Println(info)
	log.Println("watchPools loaded, watching", len(clients), "pools...")

	go serve()

	for {
		select {
		case msg := <-swapEvents:

			store, sender := isFactory(msg.Sender)

			log.Println("\nSwap Event!\n")
			log.Println("ClientAddr:", msg.To)
			log.Println("SenderAddr:", sender)
			log.Println("Pool Addr :", msg.Raw.Address)

			var c0 string
			var c1 string
			var val3 float64
			if msg.Amount0In.Cmp(big.NewInt(0)) == 0 && msg.Amount1Out.Cmp(big.NewInt(0)) == 0 {

				p1float := ParseBigFloat(msg.Amount1In)
				p2float := ParseBigFloat(msg.Amount0Out)
				val1, _ := p1float.Float64()
				val2, _ := p2float.Float64()
				val3 = val2 / val1
				c0 = info[msg.Raw.Address].Pair1.String()
				c1 = info[msg.Raw.Address].Pair0.String()

				log.Println("Swapping  :", fmt.Sprintf("%f", p1float), info[msg.Raw.Address].Pair1Name, "for", fmt.Sprintf("%f", p2float), info[msg.Raw.Address].Pair0Name)
				log.Printf("Coin Value: 1 %s = %f %s", info[msg.Raw.Address].Pair1Name, val3, info[msg.Raw.Address].Pair0Name)
			} else {

				p1float := ParseBigFloat(msg.Amount0In)
				p2float := ParseBigFloat(msg.Amount1Out)
				val1, _ := p1float.Float64()
				val2, _ := p2float.Float64()
				val3 = val2 / val1
				c0 = info[msg.Raw.Address].Pair0.String()
				c1 = info[msg.Raw.Address].Pair1.String()

				log.Println("Swapping  :", fmt.Sprintf("%f", p1float), info[msg.Raw.Address].Pair0Name, "for", fmt.Sprintf("%f", p2float), info[msg.Raw.Address].Pair1Name)
				//log.Println(big.NewInt(1).Div(msg.Amount0In, msg.Amount1Out))
				log.Printf("Coin Value: 1 %s = %f %s", info[msg.Raw.Address].Pair0Name, val3, info[msg.Raw.Address].Pair1Name)
			}
			log.Println("Amt0In    :", msg.Amount0In)
			log.Println("Amt1In    :", msg.Amount1In)
			log.Println("Amt0Out   :", msg.Amount0Out)
			log.Println("Amt1Out   :", msg.Amount1Out)
			log.Println("DB Commit :", store, "\n")

			if store {

				ts := fmt.Sprintf("%d", time.Now().Unix())

				sql := "INSERT INTO `prices` ( `ts`, `pooladdr`, `c0addr`, `c1addr`, `price`) " +
					"VALUES ( '" + ts + "', '" + msg.Raw.Address.String() + "', '" + c0 + "', '" + c1 + "', '" + fmt.Sprintf("%f", val3) + "') "

				_, err := dbHandleR.Exec(sql)
				check(err)

			}

		}
	}

}

func serve() {

	errhttp := http.ListenAndServe("10.65.0.1"+":"+strconv.Itoa(6221), nil)

	check(errhttp)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "content-type")
	w.Header().Add("Content-Type", "text/plain; version=0.0.4; charset=utf-8")

	log.Printf(r.RequestURI)

	var intP int64

	period := r.FormValue("period")
	paddr := r.FormValue("paddr")

	addrchk := common.HexToAddress(paddr)
	var pi PairInfo

	if val, ok := info[addrchk]; ok {
		pi = val
	}

	if period == "y" {
		intP = 60 * 60 * 24 * 365
	} else if period == "6m" {
		intP = 60 * 60 * 12 * 365
	} else if period == "3m" {
		intP = 60 * 60 * 6 * 365
	} else if period == "1m" {
		intP = 60 * 60 * 24 * 30
	} else if period == "2w" {
		intP = 60 * 60 * 24 * 15
	} else if period == "1w" {
		intP = 60 * 60 * 24 * 7
	} else if period == "3d" {
		intP = 60 * 60 * 24 * 3
	} else if period == "1d" {
		intP = 60 * 60 * 24
	} else if period == "12h" {
		intP = 60 * 60 * 24 * 3
	} else if period == "6h" {
		intP = 60 * 60 * 24 * 3
	} else if period == "3h" {
		intP = 60 * 60 * 24 * 3
	} else {
		intP = 60 * 60 * 24
	}

	type PairRow struct {
		ts     string
		c0addr string
		c1addr string
		price  string
	}

	record := PairRow{}

	rows := make(map[int64]PairRow)

	var count int64 = 0
	// SELECT * FROM (SELECT @row:=0) temp, prices WHERE (@row:=@row + 1) % 2 = 1 and id > 5
	query := "SELECT @bob := MAX(`ts`) from prices where `poolAddr` = '" + paddr + "'; select (ts*1000),c0addr,c1addr,price from prices where `poolAddr` = '" + paddr + "' and `ts` > (@bob - " + fmt.Sprintf("%d", intP) + ") order by `ts` desc limit 100"

	log.Println(query)
	res, err := dbHandleR.Query(query)
	check(err)
	for res.Next() {
		err := res.Scan(&record.ts, &record.c0addr, &record.c1addr, &record.price)
		check(err)
		rows[count] = record
		//log.Println(record)
		count++
	}
	err = res.Close()
	check(err)

	type xy struct {
		X int64  `json:"x"`
		Y string `json:"y"`
	}

	type dataset struct {
		Label string `json:"label"`
		Data  []xy   `json:"data"`
	}

	var ds [2]dataset

	ds[0].Label = pi.Pair0Name + "-" + pi.Pair1Name
	ds[1].Label = pi.Pair1Name + "-" + pi.Pair0Name

	for i := int64(0); i < count; i++ {
		if rows[i].c0addr == pi.Pair0.String() {
			tmp, _ := strconv.ParseInt(rows[i].ts, 10, 64)
			ds[0].Data = append(ds[0].Data, xy{tmp, rows[i].price})
			// add to dataset 1
		} else {
			tmp, _ := strconv.ParseInt(rows[i].ts, 10, 64)
			ds[1].Data = append(ds[1].Data, xy{tmp, rows[i].price})
			// add to dataset 2
		}
	}

	var output string

	jsonString1, err := json.Marshal(ds)

	//log.Println(jsonString1)

	check(err)

	output += string(jsonString1)

	log.Printf("Output: \n" + output)

	_, err = fmt.Fprintf(w, output)
	check(err)

}

func ParseBigFloat(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
}

func getPairInfo(LP *NeoLP, addr common.Address) PairInfo {
	var pi PairInfo

	pi.Pair0, _ = LP.Token0(&callOptsFac)
	pi.Pair1, _ = LP.Token1(&callOptsFac)
	pi.LpTokens, _ = LP.TotalSupply(&callOptsFac)
	pi.Pair0Balance = nil
	pi.Pair1Balance = nil

	coin0, _ := NewNeoLP(pi.Pair0, wlClient2)
	pi.Pair0Name, _ = coin0.Symbol(&callOptsFac)
	pi.Pair0Balance, _ = coin0.BalanceOf(&callOptsFac, addr)

	coin1, _ := NewNeoLP(pi.Pair1, wlClient2)
	pi.Pair1Name, _ = coin1.Symbol(&callOptsFac)
	pi.Pair1Balance, _ = coin1.BalanceOf(&callOptsFac, addr)

	return pi

}

func isFactory(addr common.Address) (bool, string) {
	if addr.String() == "0xF27fC7D2F8E96fe690Cd3C8bddD26A78F8805bdA" {
		return true, "NeoSwap Factory Contract"
	} else if addr.String() == "0xe41F6319d2a95646Ef1432591a09d61e448222c9" {
		return false, "Oracle Swap Factory"
	} else {
		return false, addr.String()
	}
}

func check(err error) {
	if err != nil {
		log.Printf("\033[41;30mERROR: %v", err)
		pc := make([]uintptr, 15)
		n := runtime.Callers(2, pc)
		frames := runtime.CallersFrames(pc[:n])
		frame, _ := frames.Next()
		log.Printf("%s:%d %s\033[0m\n", frame.File, frame.Line, frame.Function)
	}
} // fatal error check
