package play

import (
   "154.pages.dev/encoding"
   "154.pages.dev/protobuf"
   "fmt"
   "io"
   "net/http"
)

func (d Details) Field13_1_3() (uint64, bool) {
   d.m = <-d.m.Get(13)
   d.m = <-d.m.Get(1)
   if v, ok := <-d.m.GetVarint(3); ok {
      return uint64(v), true
   }
   return 0, false
}

func (d Details) Field13_1_70() (uint64, bool) {
   d.m = <-d.m.Get(13)
   d.m = <-d.m.Get(1)
   if v, ok := <-d.m.GetVarint(70); ok {
      return uint64(v), true
   }
   return 0, false
}

func (d Details) Field13_1_16() (string, bool) {
   d.m = <-d.m.Get(13)
   d.m = <-d.m.Get(1)
   if v, ok := <-d.m.GetBytes(16); ok {
      return string(v), true
   }
   return "", false
}

func (d Details) Field13_1_4() (string, bool) {
   d.m = <-d.m.Get(13)
   d.m = <-d.m.Get(1)
   if v, ok := <-d.m.GetBytes(4); ok {
      return string(v), true
   }
   return "", false
}

func (d Details) Field13_1_82_1_1() (string, bool) {
   d.m = <-d.m.Get(13)
   d.m = <-d.m.Get(1)
   d.m = <-d.m.Get(82)
   d.m = <-d.m.Get(1)
   if v, ok := <-d.m.GetBytes(1); ok {
      return string(v), true
   }
   return "", false
}

func (d Details) Field6() (string, bool) {
   if v, ok := <-d.m.GetBytes(6); ok {
      return string(v), true
   }
   return "", false
}

func (d Details) Field8_1() (float64, bool) {
   d.m = <-d.m.Get(8)
   if v, ok := <-d.m.GetVarint(1); ok {
      return float64(v) / 1_000_000, true
   }
   return 0, false
}

func (d Details) Field8_2() (string, bool) {
   d.m = <-d.m.Get(8)
   if v, ok := <-d.m.GetBytes(2); ok {
      return string(v), true
   }
   return "", false
}

func (d Details) Field5() (string, bool) {
   if v, ok := <-d.m.GetBytes(5); ok {
      return string(v), true
   }
   return "", false
}

func (d Details) Field13_1_9() (uint64, bool) {
   d.m = <-d.m.Get(13)
   d.m = <-d.m.Get(1)
   if v, ok := <-d.m.GetVarint(9); ok {
      return uint64(v), true
   }
   return 0, false
}

func (d Details) Field13_1_17() chan uint64 {
   vs := make(chan uint64)
   d.m = <-d.m.Get(13)
   d.m = <-d.m.Get(1)
   go func() {
      for v := range d.m.Get(17) {
         if v, ok := <-v.GetVarint(1); ok {
            vs <- uint64(v)
         }
      }
      close(vs)
   }()
   return vs
}

func (d Details) String() string {
   var b []byte
   b = append(b, "downloads ="...)
   if v, ok := d.Field13_1_70(); ok {
      b = fmt.Append(b, " ", encoding.Cardinal(v))
   }
   b = append(b, "\nfiles ="...)
   for file := range d.Field13_1_17() {
      if file >= 1 {
         b = append(b, " OBB"...)
      } else {
         b = append(b, " APK"...)
      }
   }
   b = append(b, "\nname ="...)
   if v, ok := d.Field5(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\noffered by ="...)
   if v, ok := d.Field6(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\nprice ="...)
   if v, ok := d.Field8_1(); ok {
      b = fmt.Append(b, " ", v)
   }
   if v, ok := d.Field8_2(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\nrequires ="...)
   if v, ok := d.Field13_1_82_1_1(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\nsize ="...)
   if v, ok := d.Field13_1_9(); ok {
      b = fmt.Appendf(b, " %v (%v)", encoding.Size(v), v)
   }
   b = append(b, "\nupdated on ="...)
   if v, ok := d.Field13_1_16(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\nversion code ="...)
   if v, ok := d.Field13_1_3(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\nversion name ="...)
   if v, ok := d.Field13_1_4(); ok {
      b = fmt.Append(b, " ", v)
   }
   return string(b)
}

type Details struct {
   m protobuf.Message
}

func (g GoogleAuth) Details(
   checkin GoogleCheckin, doc string, single bool,
) (*Details, error) {
   req, err := http.NewRequest("GET", "https://android.clients.google.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/fdfe/details"
   req.URL.RawQuery = "doc=" + doc
   g.authorization(req)
   user_agent(req, single)
   if err := checkin.x_dfe_device_id(req); err != nil {
      return nil, err
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   var d Details
   if err := d.m.Consume(data); err != nil {
      return nil, err
   }
   d.m = <-d.m.Get(1)
   d.m = <-d.m.Get(2)
   d.m = <-d.m.Get(4)
   return &d, nil
}
