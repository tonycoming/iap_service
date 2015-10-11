package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	. "proto"
	"strings"

	context "golang.org/x/net/context"
)

var _EVN = map[string]int8{
	"production": PRODUCTION,
	"sandbox":    SANDBOX,
}

var RESP_NIL = errors.New("Response is nil .")
var GP_PUBLIC_KEY = "xxxxxxxxxxxx"

type IAPService struct {
	applePayURL string // 不允许动态修改
	pk          *rsa.PublicKey
	Client      *http.Client
}

func New(evn int8, dialTimeout, rwTimeout int64) *IAPService {
	s := &IAPService{
		applePayURL: APPLE_VERIY_SANDBOX_URL, // default sand box
		Client: &http.Client{
			Transport: &http.Transport{
				Dial: TimeoutDialer(dialTimeout, rwTimeout),
			},
		},
	}

	// PRODUCTION
	if evn == PRODUCTION {
		s.applePayURL = APPLE_VERIY_PRODUCTION_URL
	}
	return s
}

/*
坑爹的苹果会更具不同的receipt_data 返回不同的数据 这里使用了推荐方式
if ([UIDevice iOSVersion] > 6.9f) {
	NSURLRequest *urlRequest = [NSURLRequest requestWithURL:[[NSBundle mainBundle] appStoreReceiptURL]];//苹果推荐
	NSError *error = nil;
	receiptData = [NSURLConnection sendSynchronousRequest:urlRequest returningResponse:nil error:&error];
}
else {
	receiptData = transaction.transactionReceipt;
}

*/
// Validating Receipts With the App Store
// NOTICE : gen order id by self
func (s *IAPService) ApplePayVerify(ctx context.Context, in *IosRequest) (*IosResponse, error) {
	if nil == s {
		return nil, errors.New("IAPService is nil !")
	}

	js := `{"receipt-data":"` + in.ReceiptData + `"}`
	// server verify
	resp, err := s.Client.Post(s.applePayURL, "application/json", strings.NewReader(js))
	if nil != err {
		return nil, err
	}
	if nil == resp {
		return nil, RESP_NIL
	}
	defer resp.Body.Close()

	r := &IosResponse{}
	response := &IAPResponse{}
	err = json.NewDecoder(resp.Body).Decode(r)
	if nil != err {
		return nil, err
	}

	var products []*Product
	app := response.Receipt.InApp
	for i := range app {
		p := &Product{ProductId: app[i].ProductID, ItemId: app[i].AppItemID}
		products = append(products, p)
	}
	r.Status = int32(response.Status)
	r.Products = products
	return r, nil
}

// playstroe iap verify
func (s *IAPService) GooglePayVerify(ctx context.Context, in *GPRequest) (*GPResponse, error) {
	if nil == s {
		return nil, errors.New("IAPService is nil !")
	}
	purchaseData := &InappPurchaseData{}
	jbin := []byte(in.InappPurchaseData)
	err := json.Unmarshal(jbin, purchaseData)
	if nil != err {
		return nil, err
	}

	signature, err := base64.StdEncoding.DecodeString(in.Signature)
	if nil != err {
		return nil, err
	}

	// defalut hash is sha1
	sha1Hash := sha1.New()
	io.WriteString(sha1Hash, in.InappPurchaseData)
	err = rsa.VerifyPKCS1v15(s.pk, crypto.SHA1, sha1Hash.Sum(nil), signature)
	if nil != err {
		return nil, err
	}

	resp := &GPResponse{}
	resp.Status = int32(purchaseData.PurchaseState)
	resp.ProductId = purchaseData.ProductId
	resp.Options = purchaseData.DelvelopePalyload
	return resp, nil
}
