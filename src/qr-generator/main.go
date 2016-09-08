package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	port := flag.Int("port", 80, "listening on port")
	flag.Parse()

	fmt.Println("starting qr-generator on port ", *port)
	http.HandleFunc("/", QrGenerator)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func QrGenerator(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	if data == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	s, err := url.QueryUnescape(data)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	code, err := qr.Encode(s, qr.L, qr.Auto)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		size = 250
	}

	// Scale the barcode to the appropriate size
	code, err = barcode.Scale(code, size, size)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// convert to 8bit image
	b := code.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, code, b.Min, draw.Src)

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, m); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

	if _, err := w.Write(buffer.Bytes()); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
