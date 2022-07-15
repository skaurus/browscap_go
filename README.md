## Disclaimer

This is a fork of a wonderful [digitalcrab/browscap_go](https://github.com/digitalcrab/browscap_go) library, which I hope to support in a more robust way. But time will show.

So far I merged PRs from [Khan](https://github.com/digitalcrab/browscap_go/pull/16), [DDosT](https://github.com/skaurus/browscap_go/pull/2), [LiveRamp](https://github.com/digitalcrab/browscap_go/pull/9).

If anything was broken during the merges - the fault is entirely mine. Tests pass though and basic usage works.

I have not changed the copyright in the LICENSE as I think it would be appropriate only if maintaince is formally transferred to me.

# Browser Capabilities GoLang Project

PHP has `get_browser()` function which tells what the user's browser is capable of.
You can check original documentation [here](http://php.net/get_browser). 
This is GoLang analog of `get_browser()` function.

[![Build Status](https://app.travis-ci.com/skaurus/browscap_go.png?branch=master)](https://app.travis-ci.com/github/skaurus/browscap_go)

## Introduction

The [browscap.ini](http://browscap.org/) file is a database which provides a lot of details about 
browsers and their capabilities, such as name, versions, Javascript support and so on.

## Quick start

First of all you need initialize library with [browscap.ini](http://browscap.org/) file. 
And then you can get Browser information as `Browser` structure.
Some fields will be available only with a full browscap.ini

```go
package main

import (
	"fmt"

	bgo "github.com/skaurus/browscap_go"
)

func main() {
	if err := bgo.InitBrowsCap("browscap.ini", false); err != nil {
		panic(err)
	}

	browser, ok := bgo.GetBrowser("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.120 Safari/537.36")
	if !ok || browser == nil {
    	panic("Browser not found")
	} else {
    	fmt.Printf("Browser = %s [%s] v%s\n", browser.Browser, browser.BrowserType, browser.BrowserVersion)
    	fmt.Printf("Platform = %s v%s\n", browser.Platform, browser.PlatformVersion)
    	fmt.Printf("Device = %s [%s] %s\n", browser.DeviceName, browser.DeviceType, browser.DeviceBrandName)
    	fmt.Printf("IsCrawler = %t\n", browser.IsCrawler())
    	fmt.Printf("IsMobile = %t\n", browser.IsMobile())
	}
}
```

## License

```
The MIT License (MIT)

Copyright (c) 2015 Maksim Naumov

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
