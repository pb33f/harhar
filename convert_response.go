// Copyright 2023-2024 Princess Beef Heavy Industries, LLC / Dave Shanley
// https://pb33f.io

package harhar

import (
    "io"
    "net/http"
    "strings"
    "time"
)

func ConvertResponseIntoHttpResponse(r Response) *http.Response {

    resp := &http.Response{
        StatusCode: r.StatusCode,
        Status:     r.StatusText,
        Proto:      r.HTTPVersion,
    }

    if r.Cookies != nil {
        h := http.Header{}
        for _, cookie := range r.Cookies {
            exp, _ := time.Parse(time.RFC3339Nano, cookie.Expires)
            c := &http.Cookie{
                Name:     cookie.Name,
                Path:     cookie.Path,
                Value:    cookie.Value,
                Domain:   cookie.Domain,
                Expires:  exp,
                HttpOnly: cookie.HTTPOnly,
                Secure:   cookie.Secure,
            }
            h.Add("Set-Cookie", c.String())
        }
        for _, header := range r.Headers {
            h.Add(header.Name, header.Value)
        }
        if r.RedirectURL != "" {
            h.Add("Location", r.RedirectURL)
        }
        if r.Body.MIMEType != "" {
            h.Add("Content-Type", r.Body.MIMEType)
        }
        if r.Body.Encoding != "" {
            h.Add("Content-Encoding", r.Body.Encoding)
        }
        if r.Body.Compression > 0 {
            h.Add("Content-Length", string(rune(r.Body.Compression)))
        }
        if r.Body.Size > 0 {
            h.Add("Content-Length", string(rune(r.Body.Size)))
        }
        if r.Body.Content != "" {
            resp.Body = io.NopCloser(strings.NewReader(r.Body.Content))
        }
        resp.Header = h
    }

    return resp

    /*
       r := Response{
             StatusCode:  hr.StatusCode,
             StatusText:  http.StatusText(hr.StatusCode),
             HTTPVersion: hr.Proto,
             HeadersSize: -1,
             BodySize:    -1,
         }

         h2 := hr.Header.Clone()
         buf := &bytes.Buffer{}
         h2.Write(buf)
         r.HeadersSize = buf.Len() + 4 // incl. CRLF CRLF

         // parse out headers
         r.Headers = make([]NameValuePair, 0, len(hr.Header))
         for name, vals := range hr.Header {
             for _, val := range vals {
                 r.Headers = append(r.Headers, NameValuePair{Name: name, Value: val})
             }
         }
         rurl, err := hr.Location()
         if err == nil {
             r.RedirectURL = rurl.String()
         }

         // parse out cookies
         r.Cookies = make([]Cookie, 0, len(hr.Cookies()))
         for _, c := range hr.Cookies() {
             nc := Cookie{
                 Name:     c.Name,
                 Path:     c.Path,
                 Value:    c.Value,
                 Domain:   c.Domain,
                 Expires:  c.Expires.Format(time.RFC3339Nano),
                 HTTPOnly: c.HttpOnly,
                 Secure:   c.Secure,
             }
             r.Cookies = append(r.Cookies, nc)
         }

         // FIXME: net/http transparently decompresses content,
         // so r.Body.Size and r.Body.Compression are not true to the server's response
         // also, if the response is not utf-8, then r.Body.Content and r.Body.Encoding
         // are not properly handled (spec says to decode anything into UTF-8)
         //
         // see hr.Uncompressed for next steps

         // read in all the data and replace the ReadCloser
         bodyData, err := io.ReadAll(hr.Body)
         if err != nil {
             return r, err
         }
         hr.Body.Close()
         hr.Body = io.NopCloser(bytes.NewReader(bodyData))
         r.Body.Content = string(bodyData)
         r.Body.Compression = 0
         r.Body.Size = len(bodyData)
         r.BodySize = r.Body.Size

         r.Body.MIMEType = hr.Header.Get("Content-Type")
         if r.Body.MIMEType == "" {
             // default per RFC2616
             r.Body.MIMEType = "application/octet-stream"
         }

    */

}
