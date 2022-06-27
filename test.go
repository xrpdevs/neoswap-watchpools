package main

import (
	"context"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"

	"github.com/rodkranz/fetch"
	b "github.com/xrpdevs/golang-sgbnft-binding/SGBFTSO_ERC721"
	a "github.com/xrpdevs/golang-sgbnft-binding/auctionmanager"
	"log"
	"math/big"
	_ "math/big"
	"runtime"
	"strings"
)

var sqlH = "sgbnft:7Pft7LP46MJYKgX@tcp(10.64.45.1)/sgbnft"

var dbHandleR *sql.DB // global database handle accessible from all the threads

var cPaySent = make(chan *a.AMPaymentSent)
var cNoWinner = make(chan *a.AMAuctionNoWinner)
var cWinner = make(chan *a.AMAuctionWinner)
var c2981Fee = make(chan *a.AMErc2981Fee)
var cIncBid = make(chan *a.AMIncreaseBid)
var cNewBid = make(chan *a.AMNewBid)
var cOutBid = make(chan *a.AMOutBid)
var cNewList = make(chan *a.AMNewListing)
var cRemoved = make(chan *a.AMRemovedListing)

var AuctionManagerContract *a.AM

var sgbferc721 *b.SGBFERC721

type Cs struct {
	RpcNode  string
	AManager common.Address
}

type metaItem struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Dna         string `json:"dna"`
	Edition     int    `json:"edition"`
	Date        int64  `json:"date"`
	Attributes  []struct {
		TraitType string `json:"trait_type"`
		Value     string `json:"value"`
	} `json:"attributes"`
}

var Config = Cs{RpcNode: "wss://rpc.sgbftso.com/testws", AManager: common.HexToAddress("0x9191F9f9dB2d8AbCd5C787F9956852e65fBF83e5")}

var MainNet = Cs{"ws://10.64.45.2:9650/ext/bc/C/ws", common.Address{0x0}}

var logs = make(chan types.Log)
var headers = make(chan *types.Header)

var funcMap = make(map[string]string)

var wlClient2, _ = ethclient.Dial(MainNet.RpcNode)
var wlClient0, _ = ethclient.Dial(Config.RpcNode)

func main() {

	var err error

	dbHandleR, err = sql.Open("mysql", sqlH)

	check(err)

	enc1 := crypto.Keccak256([]byte("submitPriceHashes(uint256,uint256[],bytes32[])"))
	enc2 := crypto.Keccak256([]byte("revealPrices(uint256,uint256[],uint256[],uint256[])"))
	enc3 := crypto.Keccak256([]byte("safeTransferFrom(address,address,uint256)"))
	enc4 := crypto.Keccak256([]byte("safeTransferFrom(address,address,uint256,bytes)"))
	enc5 := crypto.Keccak256([]byte("transferFrom(address,address,uint256)"))
	enc6 := crypto.Keccak256([]byte("transferOwnership(address)"))

	hStr1 := hexutil.Encode(enc1)[2:10]
	hStr2 := hexutil.Encode(enc2)[2:10]
	hStr3 := hexutil.Encode(enc3)[2:10]
	hStr4 := hexutil.Encode(enc4)[2:10]
	hStr5 := hexutil.Encode(enc5)[2:10]
	hStr6 := hexutil.Encode(enc6)[2:10]

	log.Println("AAA:", hStr1)

	funcMap[hStr1] = "submitPriceHashes(uint256,uint256[],bytes32[])"
	funcMap[hStr2] = "revealPrices(uint256,uint256[],uint256[],uint256[])"
	funcMap[hStr3] = "safeTransferFrom(address,address,uint256)"
	funcMap[hStr4] = "safeTransferFrom(address,address,uint256,bytes)"
	funcMap[hStr5] = "transferFrom(address,address,uint256)"
	funcMap[hStr6] = "transferOwnership(address)"

	//	importToken(common.HexToAddress("0x2972ea6e6CC45c5837CE909DeF032DD325B48415"))
	//importToken(common.HexToAddress("0xd83Ae2C70916a2360e23683A0d3a3556b2c09935"))

	go watchAuctionManagerEvents()

	for {
		select {
		case msg := <-logs:
			fmt.Println(msg) // pointer to event log
		case header := <-headers:
			ctx := context.Background()
			//fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, errz := wlClient2.BlockByNumber(context.Background(), header.Number)

			if errz != nil {
				//		log.Fatal(err)
			} else {

				//fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
				//fmt.Println(block.Number().Uint64())   // 3477413
				//fmt.Println(block.Time())              // 1529525947
				//fmt.Println(block.Nonce())             // 130524141876765836
				//fmt.Println(block.Transactions()) // 7
				if len(block.Transactions()) > 0 {
					for _, s := range block.Transactions() {
						//fmt.Println(s.To())

						if s.To() != nil {
							cid := s.To().Hex()
							if s.To().Hex() == "0x1000000000000000000000000000000000000003" {
								cid = "PriceSubmitter"
							}
							if s.To().Hex() == "0x1000000000000000000000000000000000000002" {
								cid = "FlareDaemon"
							}
							if s.To().Hex() == "0x1000000000000000000000000000000000000001" {
								cid = "StateConnector"
							}
							if s.To().Hex() == "0xA1a9B8aB5BB798EeE536A23669AD744DCF8537a3" {
								cid = "FTSO: XRP"
							}
							if s.To().Hex() == "0xdA0621876081479666F5A0b127D3ea9D426658d3" {
								cid = "Flare Finance"
							}
							if s.To().Hex() == "0xc5478a1d5914cF9D0Ee20Da21459502eCb7E1646" {
								cid = "Flare Finance"
							}
							if s.To().Hex() == "0x3B46090e608cBC963356f30857F4DAcC09f5DDC4" {
								cid = "BestFTSO NFTP"
							}
							if s.To().Hex() == "0xc5738334b972745067fFa666040fdeADc66Cb925" {
								cid = "FTSORewardManager"
							}
							if s.To().Hex() == "0x02f0826ef6aD107Cfc861152B32B52fD11BaB9ED" {
								cid = "Wnat"
							}
							if s.To().Hex() == "0x554742076743b366504972F86609d64fd18BDC34" {
								cid = "SFIN"
							}
							//codeAt, _ := wlClient2.CodeAt(context.Background(), *s.To(), header.Number)
							//log.Println(common.Bytes2Hex(codeAt))
							sgbferc721_, _ := b.NewSGBFERC721(*s.To(), wlClient2)
							var inte [4]byte
							blah, _ := hex.DecodeString("80ac58cd") // IERC721
							inte[0] = blah[0]
							inte[1] = blah[1]
							inte[2] = blah[2]
							inte[3] = blah[3]

							//s.Hash()

							callOpts := bind.CallOpts{Pending: false, From: common.HexToAddress("0xDA6FF90214CB6D0CEee462f2D788b2556E657422"), Context: ctx}

							isERC721, _ := sgbferc721_.SupportsInterface(&callOpts, inte)

							txData := s.Data()

							txString := hex.EncodeToString(txData)
							if len(txString) > 5 {
								functionCall := txString[0:8]
								log.Println("FN CALL: ", functionCall)
								log.Println("ID: ", funcMap[functionCall])
							}

							if isERC721 {
								go importToken(*s.To())
							}
							symbol, _ := sgbferc721_.Symbol(&callOpts)

							name, _ := sgbferc721_.Name(&callOpts)

							//tracers.Tracer()

							//log.Println("Symbol: ", symbol)

							if strings.Contains(cid, "0x") {
								cid = "at " + cid
							} else {
								cid = "is " + cid
							}
							cid = fmt.Sprintf("%-45s", cid)

							log.Println("Block #", block.Number().String(), "Contract", cid, "Is ERC721:", isERC721, "Symbol:", symbol, "-", name)

						} else {
							//meh, _ := s.MarshalJSON()

							go func() {
								ctx := context.Background()
								log.Println("Waiting for receipt")
								txHash, _ := bind.WaitMined(context.Background(), wlClient2, s)

								//txHash, _ := wlClient2.TransactionReceipt(context.Background(), block.Hash())
								log.Println("Waiting for receipt")

								log.Println("Contract created: ", txHash)
								log.Println("Contract created: ", txHash.ContractAddress)
								sgbferc721_, _ := b.NewSGBFERC721(txHash.ContractAddress, wlClient2)
								var inte [4]byte
								blah, _ := hex.DecodeString("80ac58cd") // IERC721
								inte[0] = blah[0]
								inte[1] = blah[1]
								inte[2] = blah[2]
								inte[3] = blah[3]

								callOpts := bind.CallOpts{Pending: false, From: common.HexToAddress("0xDA6FF90214CB6D0CEee462f2D788b2556E657422"), Context: ctx}

								log.Println("Contract created: ", txHash.TxHash)
								log.Println("Contract created: ", txHash.ContractAddress)
								isERC721, _ := sgbferc721_.SupportsInterface(&callOpts, inte)

								symbol, _ := sgbferc721_.Symbol(&callOpts)
								log.Println("Is ERC721: ", isERC721)
								if isERC721 {
									go importToken(txHash.ContractAddress)
								}
								//tracers.Tracer()

								log.Println("Symbol: ", symbol)
							}()

							//fmt.Println(s.Hash())
						}
					}
				}
			}
			//	case msg := <-cPaySent:
			//		log.Printf("**Payment Sent**\nAuction ID: %s, Payee: %s, Reason: %s, Amount %s\n",
			//			msg.AuctionId.String(), msg.Payee.String(), msg.Reason, msg.Amount.String() )
			//
			//	case msg := <-cRemoved:
			//		log.Printf("**Auction REMOVED**\nAuction ID: %s, Owner: %s, Fee: %s, Contract: %s, TokenID: %s\n",
			//			msg.AuctionId.String(), msg.TokenOwner.String(), msg.RemovalFee.String(), msg.NftContract.String(), msg.TokenId.String())
			//
			//	case msg := <-cOutBid:
			//		log.Printf("**OutBid**\nAuction ID: %s, LosingBidder: %s, HighBidder: %s, Bid: %s, Refunded: %s, Contract: %s, TokenID: %s\n",
			//			msg.AuctionId.String(), msg.LosingBidder.String(), msg.NewHighBidder.String(), msg.NewHighBid.String(), msg.RefundedAmount.String(), msg.NftContract.String(), msg.TokenId.String())
			//
			//	case msg := <-cNewList:
			//		log.Printf("**New Listing**\nAuction ID: %s, Listed By: %s, Reserve: %s, Contract: %s, TokenID: %s\n",
			//			msg.AuctionId.String(),   msg.ListedBy.String(), msg.ReserveAmount.String(), msg.NftContract.String(), msg.TokenId.String() )
			//
			//		// need to check if token is owned by member of the site
			//
			//		query := "INSERT INTO auctions (`_id`, `auctionID`, `listerAddress`, `reserveAmount`, `contractAddress`, `tokenID`) VALUES ( null, '" +
			//			msg.AuctionId.String() + "', '" +
			//			msg.ListedBy.String() + "', '" +
			//			msg.ReserveAmount.String() + "', '" +
			//			msg.NftContract.String() + "', '" +
			//			msg.TokenId.String() + "');"
			//
			//		fmt.Printf("SQL Query: %s \n", query)
			//
			//		fmt.Print(dbHandleR);
			//		_, err := dbHandleR.Exec(query)
			//		check(err)
			//		fmt.Printf("SQL Query: %s \n", query)
			//
			//	case msg := <-cNewBid:
			//		log.Printf("**New Bid**\nAuction ID: %s, Bidder: %s, Amount: %s, Contract: %s, TokenID: %s\n",
			//			msg.AuctionId.String(),   msg.Bidder.String(), msg.Amount.String(), msg.NftContract.String(), msg.TokenId.String() )
			//
			//	case msg := <-cNoWinner:
			//		log.Printf("**Ended: No Winner**\nAuction ID: %s, Reason: %s, Contract: %s, TokenID: %s\n",
			//			msg.AuctionId.String(), msg.Reason, msg.NftContract.String(), msg.TokenId.String())
			//
			//	case msg := <-cIncBid:
			//		log.Printf("**Increase Bid**\nAuction ID: %s, Bidder: %s, TotalAmount: %s, Contract: %s, TokenID: %s\n",
			//			msg.AuctionId.String(),   msg.Bidder.String(), msg.Totalamount.String(), msg.NftContract.String(), msg.TokenId.String() )
			//
			//	case msg := <-c2981Fee:
			//		log.Printf("**EIP-2981 Fee**\nAuction ID: %s, Recipient: %s, Amount: %s, Contract: %s\n",
			//			msg.AuctionId.String(), msg.RoyRecipient.String(), msg.RoyPayout.String(), msg.NftContract.String() )
			//
			//	case msg := <-cWinner:
			//		log.Printf("**Auction Won**\nAuction ID: %s, Bidder: %s, Winning Amount: %s, Contract: %s, TokenID: %s\n",
			//			msg.AuctionId.String(), msg.Bidder.String(), msg.WinningAmount, msg.NftContract, msg.TokenId.String())
		}

	}

}

func tokenURItoClean(tURI string) string {
	if strings.Contains(tURI, "ipfs://") {
		return strings.Replace(tURI, "ipfs://", "/ipfs/", -1)
	} else if strings.Contains(tURI, "/ipfs") {
		return "/ipfs" + strings.Split(tURI, "/ipfs")[1]
	} else if strings.Contains(tURI, "ipns://") {
		return strings.Replace(tURI, "ipns://", "/ipns/", -1)
	} else if strings.Contains(tURI, "/ipns") {
		return "/ipns" + strings.Split(tURI, "/ipns")[1]
	}
	return tURI
}

func metaFetch(url string) {
	opts := fetch.Options{
		Header: http.Header{
			"User-Agent": []string{"geterc721/1.0"},
		},
	}

	f := fetch.New(&opts)

	res, err := f.GetWithContext(context.Background(), url, nil)

	var iJson metaItem

	json.Unmarshal([]byte(res.String()), &iJson)

	if strings.Contains(iJson.Image, "http://") || strings.Contains(iJson.Image, "https://") {
		go imageFetch(iJson.Image)
	} else {
		//		imgurl := strings.Replace(iJson.Image, "http:/", "", -1)
		//		imgurl = strings.Replace(imgurl, "https:/", "", -1)
		var imgurl string
		if strings.Contains(iJson.Image, "ipfs://") {
			imgurl = strings.Split(iJson.Image, "ipfs://")[1]
			go imageFetch("https://cloudflare-ipfs.com/ipfs/" + imgurl)
		}
		if strings.Contains(iJson.Image, "ipns://") {
			imgurl = strings.Split(iJson.Image, "ipns://")[1]
			go imageFetch("https://cloudflare-ipfs.com/ipns/" + imgurl)
		}
		if strings.Contains(iJson.Image, "ipfs/") {
			imgurl = strings.Split(iJson.Image, "ipfs/")[1]
			go imageFetch("https://cloudflare-ipfs.com/ipfs/" + imgurl)
		}
		if strings.Contains(iJson.Image, "ipns/") {
			imgurl = strings.Split(iJson.Image, "ipns/")[1]
			go imageFetch("https://cloudflare-ipfs.com/ipns/" + imgurl)
		}

	}

	log.Println("\033[36m", iJson, err, "\033[0m")

}

func imageFetch(url string) {
	log.Println("\033[31mFetching: ", url, "\033[0m")
	var metaBaseUrl_ = strings.Split(url, "/")

	od := metaBaseUrl_[(len(metaBaseUrl_) - 2)]
	of := metaBaseUrl_[(len(metaBaseUrl_) - 1)]

	if _, err := os.Stat("./images/" + od); os.IsNotExist(err) {
		_ = os.Mkdir("./images/"+od, 0777)
		// TODO: handle error
	}

	outfile, _ := os.OpenFile("./images/"+od+"/"+of, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	opts := fetch.Options{
		Header: http.Header{
			"User-Agent": []string{"geterc721/1.0"},
		},
	}

	f := fetch.New(&opts)

	res, _ := f.GetWithContext(context.Background(), url, nil)

	outfile.Write([]byte(res.String()))
	outfile.Close()

}

//func updateTokenInfo(string addr, )

func importToken(contractAddress common.Address) {
	log.Println("\033[33mAdding Token\033[0m")
	ctx := context.Background()
	// set up a websocket & subscription to listen to events on WNat contract
	//var err1 error
	wlClient, err1 := ethclient.Dial(MainNet.RpcNode)
	check(err1)
	//contractAddress := Config.AManager

	var err2 error

	sgbferc721, err2 = b.NewSGBFERC721(contractAddress, wlClient)
	check(err2)
	//watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}
	callOpts := bind.CallOpts{Pending: false, From: common.HexToAddress("0xDA6FF90214CB6D0CEee462f2D788b2556E657422"), Context: ctx}
	tot, _ := sgbferc721.TotalSupply(&callOpts)
	name, _ := sgbferc721.Symbol(&callOpts)
	symbol, _ := sgbferc721.Name(&callOpts)
	ownerAddr, _ := sgbferc721.Owner(&callOpts)
	metaBaseUrl, _ := sgbferc721.TokenURI(&callOpts, big.NewInt(1))
	if strings.Contains(metaBaseUrl, "ipfs/") {
		tmp := strings.Split(metaBaseUrl, "ipfs/")[1]
		metaBaseUrl = "ipfs/" + tmp
	} else if strings.Contains(metaBaseUrl, "ipns/") {
		tmp := strings.Split(metaBaseUrl, "ipns/")[1]
		metaBaseUrl = "ipns/" + tmp
	}

	var metaBaseUrl_ = strings.Split(metaBaseUrl, "/")

	baseurl := ""
	for i := 1; i < len(metaBaseUrl_)-1; i++ {
		baseurl += metaBaseUrl_[i] + "/"
	}
	log.Println("BASEURL", baseurl)
	go metaFetch("https://cloudflare-ipfs.com/ipfs" + baseurl + "1.json")
	firstu, _ := sgbferc721.TokenURI(&callOpts, big.NewInt(1))

	query := `SELECT id FROM collections where contractAddress LIKE '` + contractAddress.String() + `' LIMIT 1`
	res := dbHandleR.QueryRow(query)
	var out string

	err := res.Scan(&out)
	if err == nil {
		log.Printf("Found: %s, %s, Total Supply: %s, Owner: %s\n", name, symbol, tot, ownerAddr.String())
		query = "UPDATE collections set totalSupply = '" + tot.String() + "' WHERE contractAddress = '" + contractAddress.String() + "' LIMIT 1;"
		_, err = dbHandleR.Exec(query)
		check(err)
		fmt.Printf("SQL Query: %s \n", query)
		//		log.Println(out)
	} else {
		log.Println("Found")
		tots := tot.String()
		if tot == nil {
			tots = "-1"
		}

		query = "INSERT INTO collections (`id`, `owner`, `name`, `symbol`, `contractAddress`, `coll_id`, `site_userid`, `totalSupply`, `imageUrl`, `metaBaseUrl`) VALUES ( null, '" +
			ownerAddr.String() + "', '" +
			name + "', '" +
			symbol + "', '" +
			contractAddress.String() + "', '" +
			"0" + "', '" +
			"0" + "', '" +
			tots + "', '" +
			firstu + "', '" +
			metaBaseUrl + "');"

		//fmt.Printf("SQL Query: %s \n", query)
		fmt.Print(dbHandleR)
		_, err = dbHandleR.Exec(query)
		check(err)
		fmt.Printf("SQL Query: %s \n", query)
	}

	log.Println(name, symbol)
	//for i := 0; int64(i)<tot.Int64(); i++ {
	//	_, err3 := sgbferc721.OwnerOf(&callOpts, big.NewInt(int64(i)))
	//	_, err4 := sgbferc721.TokenURI(&callOpts, big.NewInt(int64(i)))
	//	query = "INSERT INTO table (id, name, age) VALUES(1, \"A\", 19) ON DUPLICATE KEY UPDATE"+
	//	"name=\"A\", age=19"
	//
	//	check(err3)
	//	check(err4)
	//	//log.Println(token, info)
	//
	//
	//}

}

func watchAuctionManagerEvents() {
	//ctx := context.Background()
	// set up a websocket & subscription to listen to events on WNat contract
	//var err1 error
	//wlClient, err1 := ethclient.Dial(Config.RpcNode)
	//check(err1)
	//contractAddress := Config.AManager

	_, err := wlClient2.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	//
	//var err2 error
	//
	//AuctionManagerContract, err2 = a.NewAM (contractAddress, wlClient)
	//check(err2)
	//watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}
	//
	//go func() {
	//	log.Printf("Watchers init")
	//	var err error
	//	_, err = AuctionManagerContract.WatchAuctionWinner  (watchOpts, cWinner,  nil)
	//	_, err = AuctionManagerContract.WatchAuctionNoWinner(watchOpts, cNoWinner,nil)
	//	_, err = AuctionManagerContract.WatchErc2981Fee     (watchOpts, c2981Fee, nil)
	//	_, err = AuctionManagerContract.WatchIncreaseBid    (watchOpts, cIncBid,  nil)
	//	_, err = AuctionManagerContract.WatchNewBid         (watchOpts, cNewBid,  nil)
	//	_, err = AuctionManagerContract.WatchNewListing     (watchOpts, cNewList, nil)
	//	_, err = AuctionManagerContract.WatchOutBid         (watchOpts, cOutBid,  nil)
	//	_, err = AuctionManagerContract.WatchRemovedListing (watchOpts, cRemoved, nil)
	//	_, err = AuctionManagerContract.WatchPaymentSent    (watchOpts, cPaySent, nil)
	//
	//	check(err)
	//	check(err2)
	//}()
} // subscribe to PriceSubmitter's events

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
