package gplayapi

import (
   "154.pages.dev/google/gplayapi/gpproto"
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "fmt"
   "google.golang.org/protobuf/proto"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (g GooglePlayClient) Acquire(doc string, version uint64) error {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["X-Dfe-Device-Id"] = []string{g.AuthData.GsfID}
   req.Header["Authorization"] = []string{"Bearer " + g.AuthData.AuthToken}
   // org.videolan.vlc
   body := protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes(doc)},
            protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(1)},
            protobuf.Field{Number: 3, Type: 0,  Value: protobuf.Varint(3)},
         }},
         protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(1)},
         protobuf.Field{Number: 3, Type: 2,  Value: protobuf.Bytes("")},
      }},
      protobuf.Field{Number: 8, Type: 2,  Value: protobuf.Bytes("")},
      protobuf.Field{Number: 12, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(version)},
         protobuf.Field{Number: 3, Type: 0,  Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 13, Type: 0,  Value: protobuf.Varint(1)},
      protobuf.Field{Number: 15, Type: 0,  Value: protobuf.Varint(0)},
      protobuf.Field{Number: 22, Type: 2,  Value: protobuf.Bytes("nonce=qlIhuESfWlLW7GI7k6fWej7z403Mynf3o0dl5B9RYfWxmHxGGSGqBARF_TxpzL5RfVPW3oFX0zAHhSETtuUBv7TvrzWOx5hEgolPjFDs1lr_Po9lyH1HJ8UskVSkyMe_gImmY0-hA-I0SSaBfUXyInciRuuMtSNiXsMclJwWoW1bPgjYsQKCn5szQnDPlMvSqz4hBCbIGxGKiWe6L9f6IcmfIwlz8eSRQUA02YN633zvXDbptBIKfrpwE9_P_N0sWrOhc3k9LAQlrn4f4kXor4g98ZQ6BN6U3us7US-2ES-xiCaFvdlbIMiMvpp7_AsnLon1KyvxS_rujvoDaUyOOQ")},
      protobuf.Field{Number: 25, Type: 0,  Value: protobuf.Varint(2)},
      protobuf.Field{Number: 30, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(2)},
      }},
   }
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "play-fe.googleapis.com"
   req.URL.Path = "/fdfe/acquire"
   val := make(url.Values)
   val["theme"] = []string{"2"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(body.Append(nil)))
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return fmt.Errorf(res.Status)
   }
   {
      b, err := io.ReadAll(res.Body)
      if err != nil {
         return err
      }
      if bytes.Contains(b, []byte("Error")) {
         return fmt.Errorf("%q", b)
      }
   }
   return nil
}

func (client *GooglePlayClient) doAuthedReq(r *http.Request) (_ *gpproto.Payload, err error) {
   client.setDefaultHeaders(r)
   b, _, err := doReq(r)
   if err != nil {
      return
   }
   resp := &gpproto.ResponseWrapper{}
   err = proto.Unmarshal(b, resp)
   if err != nil {
      return
   }
   return resp.Payload, nil
}
func doReq(r *http.Request) ([]byte, int, error) {
   res, err := http.DefaultClient.Do(r)
   if err != nil {
      return nil, 0, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, 0, errors.New(res.Status)
   }
   b, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, 0, err
   }
   return b, res.StatusCode, err
}

const (
   ImageTypeAppScreenshot = iota + 1
   ImageTypePlayStorePageBackground
   ImageTypeYoutubeVideoLink
   ImageTypeAppIcon
   ImageTypeCategoryIcon
   ImageTypeYoutubeVideoThumbnail = 13

   UrlBase               = "https://android.clients.google.com"
   UrlFdfe               = UrlBase + "/fdfe"
   UrlAuth               = UrlBase + "/auth"
   UrlCheckIn            = UrlBase + "/checkin"
   UrlDetails            = UrlFdfe + "/details"
   UrlDelivery           = UrlFdfe + "/delivery"
   UrlPurchase           = UrlFdfe + "/purchase"
   UrlToc                = UrlFdfe + "/toc"
   UrlTosAccept          = UrlFdfe + "/acceptTos"
   UrlUploadDeviceConfig = UrlFdfe + "/uploadDeviceConfig"
)

var ErrNilPayload = errors.New("got nil payload from google play")

func ptrBool(b bool) *bool {
   return &b
}

func ptrStr(str string) *string {
   return &str
}

func ptrInt32(i int32) *int32 {
   return &i
}

func parseResponse(res string) map[string]string {
   ret := map[string]string{}
   for _, ln := range strings.Split(res, "\n") {
      keyVal := strings.SplitN(ln, "=", 2)
      if len(keyVal) >= 2 {
         ret[keyVal[0]] = keyVal[1]
      }
   }
   return ret
}
