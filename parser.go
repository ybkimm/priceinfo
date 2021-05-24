package main

import (
	"context"
	"log"
	"net/url"

	"github.com/chromedp/chromedp"
)

type ParseFunc func(out *[]PriceInfo) chromedp.Action

var parserMap = map[string]ParseFunc{}

func CrawPage(href string, opts ...chromedp.ExecAllocatorOption) ([]PriceInfo, error) {
	if opts == nil {
		opts = chromedp.DefaultExecAllocatorOptions[:]
	}

	u, err := url.Parse(href)
	if err != nil {
		return nil, err
	}

	parse, ok := parserMap[u.Host]
	if !ok {
		return nil, ErrUnsupportedHost
	}

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	allocCtx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	var out []PriceInfo

	err = chromedp.Run(
		taskCtx,
		chromedp.Navigate(href),
		parse(&out),
	)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func actions(actions ...chromedp.Action) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		var err error
		for _, action := range actions {
			err = action.Do(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
