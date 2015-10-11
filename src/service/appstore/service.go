package appstore

import (
	"encoding/json"
	"errors"
	"net/http"
	. "proto/appstore"
	"strings"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

type AppStoreService struct {
	applePayURL string // 不允许动态修改
	Client      *http.Client
}

func RegisterService(s *grpc.Server, srv AppleIapServiceServer) {
	RegisterAppleIapServiceServer(s, srv)
}

func New(evn int8, dialTimeout, rwTimeout int64) *AppStoreService {
	s := &AppStoreService{
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
func (s *AppStoreService) ApplePayVerify(ctx context.Context, in *Request) (*Response, error) {
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

	response := &IAPResponse{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if nil != err {
		return nil, err
	}
	r := &Response{}
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
