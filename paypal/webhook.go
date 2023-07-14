package paypal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-pay/gopay"
	"net/http"
)

// VerifyWebhookSignature
// 文档：https://developer.paypal.com/docs/api/webhooks/v1/#verify-webhook-signature_post
func (c *Client) VerifyWebhookSignature(ctx context.Context, bm interface{}) (verifyRes *VerifyWebhookResponse, err error) {

	res, bs, err := c.doPayPalPost(ctx, bm, VerifyWebhookSignature)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return verifyRes, errors.New("request paypal url[verify-webhook-signature_post] error")
	}
	verifyRes = &VerifyWebhookResponse{}
	if err = json.Unmarshal(bs, verifyRes); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return verifyRes, nil
}
