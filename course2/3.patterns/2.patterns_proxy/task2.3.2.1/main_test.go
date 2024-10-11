package main

import (
	"log"
	"reflect"
	"testing"
	"time"
)

func TestLinesProxy(t *testing.T) {
	data := KLines{Pair: "USD_BTC", Candles: []Candle{
		{H: 300, L: 100, C: 250},
		{H: 310, L: 250, C: 290},
		{H: 210, L: 100, C: 130},
	}}
	DataB, err := data.MarshalKLines()
	if err != nil {
		log.Fatal(err)
	}
	proxy := LoadKLinesProxy(DataB)

	{
		start := time.Now()
		k, d := proxy.MACD()
		elapsed := time.Since(start)
		startCache := time.Now()
		kCache, dCache := proxy.MACD()
		elapsedCash := time.Since(startCache)
		if !reflect.DeepEqual(k, kCache) {
			t.Errorf("KLinesProxy: expected %v, got %v", k, kCache)
		}
		if !reflect.DeepEqual(dCache, d) {
			t.Errorf("KLinesProxy: expected %v, got %v", d, dCache)
		}
		if elapsed < elapsedCash {
			t.Errorf("KLinesProxy: time with cache: %v, time without cache: %v", elapsedCash, elapsed)
		}
	}

	{
		start := time.Now()
		k, d := proxy.StochPrice()
		elapsed := time.Since(start)
		startCache := time.Now()
		kCache, dCache := proxy.StochPrice()
		elapsedCash := time.Since(startCache)
		if !reflect.DeepEqual(k, kCache) {
			t.Errorf("KLinesProxy: expected %v, got %v", k, kCache)
		}
		if !reflect.DeepEqual(dCache, d) {
			t.Errorf("KLinesProxy: expected %v, got %v", d, dCache)
		}
		if elapsed < elapsedCash {
			t.Errorf("KLinesProxy: time with cache: %v, time without cache: %v", elapsedCash, elapsed)
		}
	}

	{
		start := time.Now()
		k, d := proxy.RSI(10)
		elapsed := time.Since(start)
		startCache := time.Now()
		kCache, dCache := proxy.RSI(10)
		elapsedCash := time.Since(startCache)
		if !reflect.DeepEqual(k, kCache) {
			t.Errorf("KLinesProxy: expected %v, got %v", k, kCache)
		}
		if !reflect.DeepEqual(dCache, d) {
			t.Errorf("KLinesProxy: expected %v, got %v", d, dCache)
		}
		if elapsed < elapsedCash {
			t.Errorf("KLinesProxy: time with cache: %v, time without cache: %v", elapsedCash, elapsed)
		}
	}
	{
		start := time.Now()
		k, d := proxy.StochRSI(10)
		elapsed := time.Since(start)
		startCache := time.Now()
		kCache, dCache := proxy.StochRSI(10)
		elapsedCash := time.Since(startCache)
		if !reflect.DeepEqual(k, kCache) {
			t.Errorf("KLinesProxy: expected %v, got %v", k, kCache)
		}
		if !reflect.DeepEqual(dCache, d) {
			t.Errorf("KLinesProxy: expected %v, got %v", d, dCache)
		}
		if elapsed < elapsedCash {
			t.Errorf("KLinesProxy: time with cache: %v, time without cache: %v", elapsedCash, elapsed)
		}
	}
	{
		start := time.Now()
		k := proxy.SMA(10)
		elapsed := time.Since(start)
		startCache := time.Now()
		kCache := proxy.SMA(10)
		elapsedCash := time.Since(startCache)
		if !reflect.DeepEqual(k, kCache) {
			t.Errorf("KLinesProxy: expected %v, got %v", k, kCache)
		}
		if elapsed < elapsedCash {
			t.Errorf("KLinesProxy: time with cache: %v, time without cache: %v", elapsedCash, elapsed)
		}
	}
	{
		start := time.Now()
		k := proxy.EMA()
		elapsed := time.Since(start)
		startCache := time.Now()
		kCache := proxy.EMA()
		elapsedCash := time.Since(startCache)
		if !reflect.DeepEqual(k, kCache) {
			t.Errorf("KLinesProxy: expected %v, got %v", k, kCache)
		}
		if elapsed < elapsedCash {
			t.Errorf("KLinesProxy: time with cache: %v, time without cache: %v", elapsedCash, elapsed)
		}
	}
}

func TestLoadKLines(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("LoadKLinesProxy: expected panic")
		}
	}()
	data := []byte(`wrong:json}`)
	LoadKLinesProxy(data)

}
