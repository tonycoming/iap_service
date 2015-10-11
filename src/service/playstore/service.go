package playstore

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	. "proto/playstore"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

type PlayStoreService struct {
	PublicKeys map[int32]*rsa.PublicKey
}

func RegisterService(s *grpc.Server, srv GoogleIapServiceServer) {
	RegisterGoogleIapServiceServer(s, srv)
}

func New() *PlayStoreService {
	return &PlayStoreService{
		PublicKeys: make(map[int32]*rsa.PublicKey),
	}
}

// playstroe iap verify
func (s *PlayStoreService) GooglePayVerify(ctx context.Context, in *Request) (*Response, error) {
	if nil == s {
		return nil, errors.New("IAPService is nil !")
	}

	pk, err := GPPublicKey(s, in.ProjectId)
	if nil != err {
		return nil, err
	}

	purchaseData := &InappPurchaseData{}
	jbin := []byte(in.InappPurchaseData)
	err = json.Unmarshal(jbin, purchaseData)
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
	err = rsa.VerifyPKCS1v15(pk, crypto.SHA1, sha1Hash.Sum(nil), signature)
	if nil != err {
		return nil, err
	}

	resp := &Response{}
	resp.Status = int32(purchaseData.PurchaseState)
	resp.ProductId = purchaseData.ProductId
	resp.Options = purchaseData.DelvelopePalyload
	return resp, nil
}

func GPPublicKey(s *PlayStoreService, projectId int32) (*rsa.PublicKey, error) {
	pk, ok := s.PublicKeys[projectId]
	if ok {
		return pk, nil
	}

	key, err := GetGPPublicKeyByProjectId(projectId)
	if nil != err {
		return nil, err
	}
	tmp, err := LoadPublicKey(key)
	if nil != err {
		return nil, err
	}
	s.PublicKeys[projectId] = tmp
	return tmp, nil
}

// 获取项目公钥
func GetGPPublicKeyByProjectId(id int32) (string, error) {
	return "====================", nil
}
