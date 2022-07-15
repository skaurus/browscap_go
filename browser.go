// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package browscap_go

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

type Browser struct {
	parent  string //name of the parent
	built   bool   // has complete data from parents
	buildMu sync.Mutex

	Parent string

	Comment string

	Browser         string
	BrowserVersion  string
	BrowserMajorVer string
	BrowserMinorVer string
	// Browser, Application, Bot/Crawler, Useragent Anonymizer, Offline Browser,
	// Multimedia Player, Library, Feed Reader, Email Client or unknown
	BrowserType string
	BrowserMaker string
	BrowserModus string
	BrowserBits string

	Platform        string
	PlatformShort   string
	PlatformVersion string
	PlatformDescription string
	PlatformBits string
	PlatformMaker string

	DeviceName string
	DeviceMaker string
	// Mobile Phone, Mobile Device, Tablet, Desktop, TV Device, Console,
	// FonePad, Ebook Reader, Car Entertainment System or unknown
	DeviceType     string
	DevicePointingMethod string
	DeviceName     string
	DeviceCodeName string
	DeviceBrandName    string

	SyndicationReader string
	Crawler string
	Fake string
	Anonymized string
	Modified string
	Alpha string
	Beta string
	Win16 string
	Win32 string
	Win64 string
	Frames string
	IFrames string
	Tables string
	Cookies    string
	JavaScript string
	BackgroundSounds string
	VBScript string
	JavaApplets string
	ActiveXControls string
	CssVersion string
	AolVersion string

	RenderingEngineName    string
	RenderingEngineVersion string
	RenderingEngineDescription string
	RenderingEngineMaker string
}

func (browser *Browser) build(browsers map[string]*Browser) {
	if browser.built {
		return
	}

	browser.buildMu.Lock()
	defer browser.buildMu.Unlock()
	// Check again after lock if another gorutine built the object while we were waiting
	if browser.built {
		return
	}

	n := reflect.ValueOf(*browser).NumField()

	current := reflect.ValueOf(browser)
	parent := browser.parent
	for parent != "" {
		b, ok := browsers[parent]
		if !ok {
			break
		}

		parentObj := reflect.ValueOf(b)
		for i := 0; i < n; i++ {
			cField := current.Elem().Field(i)
			if cField.String() != "" {
				continue
			}

			pField := parentObj.Elem().Field(i)
			if pField.String() == "" {
				continue
			}

			cField.SetString(pField.String())
		}

		parent = b.parent
	}

	browser.built = true
}

func (browser *Browser) setValue(key, item string) {
	if key == "Parent" {
		browser.parent = item
		browser.Parent = item
	} else if key == "Comment" {
		browser.Comment = item
	} else if key == "Browser" {
		browser.Browser = item
	} else if key == "Version" {
		browser.BrowserVersion = item
	} else if key == "MajorVer" {
		browser.BrowserMajorVer = item
	} else if key == "MinorVer" {
		browser.BrowserMinorVer = item
	} else if key == "Browser_Type" {
		browser.BrowserType = item
	} else if key == "Browser_Maker" {
		browser.BrowserMaker = item
	} else if key == "Browser_Modus" {
		browser.BrowserModus = item
	} else if key == "Browser_Bits" {
		browser.BrowserBits = item
	} else if key == "JavaScript" {
		browser.JavaScript = item
	} else if key == "BackgroundSounds" {
		browser.BackgroundSounds = item
	} else if key == "VBScript" {
		browser.VBScript = item
	} else if key == "JavaApplets" {
		browser.JavaApplets = item
	} else if key == "ActiveXControls" {
		browser.ActiveXControls = item
	} else if key == "CssVersion" {
		browser.CssVersion = item
	} else if key == "AolVersion" {
		browser.AolVersion = item
	} else if key == "Frames" {
		browser.Frames = item
	} else if key == "IFrames" {
		browser.IFrames = item
	} else if key == "Tables" {
		browser.Tables = item
	} else if key == "Cookies" {
		browser.Cookies = item
	} else if key == "isSyndicationReader" {
		browser.SyndicationReader = item
	} else if key == "Crawler" {
		browser.Crawler = item
	} else if key == "isFake" {
		browser.Fake = item
	} else if key == "isAnonymized" {
		browser.Anonymized = item
	} else if key == "isModified" {
		browser.Modified = item
	} else if key == "Alpha" {
		browser.Alpha = item
	} else if key == "Beta" {
		browser.Beta = item
	} else if key == "Win16" {
		browser.Win16 = item
	} else if key == "Win32" {
		browser.Win32 = item
	} else if key == "Win64" {
		browser.Win64 = item
	} else if key == "Platform" {
		browser.Platform = item
		browser.PlatformShort = strings.ToLower(item)

		if strings.HasPrefix(browser.PlatformShort, "win") {
			browser.PlatformShort = "win"
		} else if strings.HasPrefix(browser.PlatformShort, "mac") {
			browser.PlatformShort = "mac"
		}
	} else if key == "Platform_Version" {
		browser.PlatformVersion = item
	} else if key == "Platform_Description" {
		browser.PlatformDescription = item
	} else if key == "Platform_Bits" {
		browser.PlatformBits = item
	} else if key == "Platform_Maker" {
		browser.PlatformMaker = item
	} else if key == "RenderingEngine_Name" {
		browser.RenderingEngineName = item
	} else if key == "RenderingEngine_Version" {
		browser.RenderingEngineVersion = item
	} else if key == "RenderingEngine_Description" {
		browser.RenderingEngineDescription = item
	} else if key == "RenderingEngine_Maker" {
		browser.RenderingEngineMaker = item
	} else if key == "Device_Name" {
		browser.DeviceName = item
	} else if key == "Device_Maker" {
		browser.DeviceMaker = item
	} else if key == "Device_Type" {
		browser.DeviceType = item
	} else if key == "Device_Pointing_Method" {
		browser.DevicePointingMethod = item
	} else if key == "Device_Name" {
		browser.DeviceName = item
	} else if key == "Device_Code_Name" {
		browser.DeviceCodeName = item
	} else if key == "Device_Brand_Name" {
		browser.DeviceBrandName = item
	}
}

func extractBrowser(data map[string]string) *Browser {
	browser := &Browser{}

	if debug {
		fmt.Println("= Browser ==================")
		for k, v := range data {
			fmt.Printf("%s = %s\n", k, v)
		}
		fmt.Println("============================")
	}

	// Browser
	if item, ok := data["Browser"]; ok {
		browser.Browser = item
	}
	if item, ok := data["Version"]; ok {
		browser.BrowserVersion = item
	}
	if item, ok := data["MajorVer"]; ok {
		browser.BrowserMajorVer = item
	}
	if item, ok := data["MinorVer"]; ok {
		browser.BrowserMinorVer = item
	}
	if item, ok := data["Browser_Type"]; ok {
		browser.BrowserType = item
	}
	if item, ok := data["Browser_Maker"]; ok {
		browser.BrowserMaker = item
	}
	if item, ok := data["Browser_Bits"]; ok {
		browser.BrowserBits = item
	}

	// Platform
	if item, ok := data["Platform"]; ok {
		browser.Platform = item
		browser.PlatformShort = strings.ToLower(item)

		if strings.HasPrefix(browser.PlatformShort, "win") {
			browser.PlatformShort = "win"
		} else if strings.HasPrefix(browser.PlatformShort, "mac") {
			browser.PlatformShort = "mac"
		}
	}
	if item, ok := data["Platform_Version"]; ok {
		browser.PlatformVersion = item
	}
	if item, ok := data["Platform_Description"]; ok {
		browser.PlatformDescription = item
	}
	if item, ok := data["Platform_Bits"]; ok {
		browser.PlatformBits = item
	}
	if item, ok := data["Platform_Maker"]; ok {
		browser.PlatformMaker = item
	}

	// Device
	if item, ok := data["Device_Name"]; ok {
		browser.DeviceName = item
	}
	if item, ok := data["Device_Maker"]; ok {
		browser.DeviceMaker = item
	}
	if item, ok := data["Device_Type"]; ok {
		browser.DeviceType = item
	}
	if item, ok := data["Device_Pointing_Method"]; ok {
                browser.DevicePointingMethod = item
        }
	if item, ok := data["Device_Name"]; ok {
		browser.DeviceName = item
	}
	if item, ok := data["Device_Code_Name"]; ok {
		browser.DeviceCodeName = item
	}
	if item, ok := data["Device_Brand_Name"]; ok {
		browser.DeviceBrandName = item
	}

	return browser
}

func (browser *Browser) IsCrawler() bool {
	return browser.BrowserType == "Bot/Crawler" || browser.Crawler == "true"
}

func (browser *Browser) IsMobile() bool {
	return browser.DeviceType == "Mobile Phone" || browser.DeviceType == "Mobile Device"
}

func (browser *Browser) IsTablet() bool {
	return browser.DeviceType == "Tablet" || browser.DeviceType == "FonePad" || browser.DeviceType == "Ebook Reader"
}

func (browser *Browser) IsDesktop() bool {
	return browser.DeviceType == "Desktop"
}

func (browser *Browser) IsConsole() bool {
	return browser.DeviceType == "Console"
}

func (browser *Browser) IsTv() bool {
	return browser.DeviceType == "TV Device"
}

func (browser *Browser) IsAndroid() bool {
	return browser.Platform == "Android"
}

func (browser *Browser) IsIPhone() bool {
	return browser.Platform == "iOS" && browser.DeviceCodeName == "iPhone"
}

func (browser *Browser) IsIPad() bool {
	return browser.Platform == "iOS" && browser.DeviceCodeName == "iPad"
}

func (browser *Browser) IsWinPhone() bool {
	return strings.Index(browser.Platform, "WinPhone") != -1 || browser.Platform == "WinMobile"
}

func (browser *Browser) IsJavaScriptSupports() bool {
	return browser.JavaScript == "true"
}

func (browser *Browser) IsCookiesSupports() bool {
	return browser.Cookies == "true"
}

func (browser *Browser) IsFramesSupports() bool {
	return browser.Frames == "true"
}

func (browser *Browser) IsIFramesSupports() bool {
	return browser.IFrames == "true"
}

func (browser *Browser) IsTablesSupports() bool {
	return browser.Tables == "true"
}

func (browser *Browser) IsFakeF() bool {
	return browser.Fake == "true"
}

func (browser *Browser) IsAnonymizedF() bool {
	return browser.Anonymized == "true"
}

func (browser *Browser) IsModifiedF() bool {
	return browser.Modified == "true"
}

func (browser *Browser) IsSyndicationReaderF() bool {
	return browser.SyndicationReader == "true"
}

func (browser *Browser) IsAlpha() bool {
	return browser.Alpha == "true"
}

func (browser *Browser) IsBeta() bool {
	return browser.Beta == "true"
}

func (browser *Browser) IsWin16() bool {
	return browser.Win16 == "true"
}

func (browser *Browser) IsWin32() bool {
	return browser.Win32 == "true"
}

func (browser *Browser) IsWin64() bool {
	return browser.Win64 == "true"
}

func (browser *Browser) IsBackgroundSoundsSupports() bool {
	return browser.BackgroundSounds == "true"
}

func (browser *Browser) IsVBScriptSupports() bool {
	return browser.VBScript == "true"
}

func (browser *Browser) IsJavaAppletsSupports() bool {
	return browser.JavaApplets == "true"
}

func (browser *Browser) IsActiveXControlsSupports() bool {
	return browser.ActiveXControls == "true"
}
