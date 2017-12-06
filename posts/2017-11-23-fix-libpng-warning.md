---
title: fix liteide libpng warning
date: '2017-11-23'
description:
categories:
- blog
- dev
- liteide
tags:
- blog
- liteide
---

<!-- ## fix liteide libpng warning -->

### liteide issues #867

Start LiteIDE build for Qt5 version suggests libpng warning. [liteide issues #867](https://github.com/visualfc/liteide/issues/867)

```
libpng warning: iCCP: profile 'ICC profile': 'GRAY': Gray color space not permitted on RGB PNG
libpng warning: iCCP: known incorrect sRGB profile
libpng warning: iCCP: cHRM chunk does not match sRGB
libpng warning: iCCP: profile 'ICC profile': 'GRAY': Gray color space not permitted on RGB PNG
libpng warning: iCCP: profile 'ICC profile': 'GRAY': Gray color space not permitted on RGB PNG
libpng warning: iCCP: profile 'ICC profile': 'GRAY': Gray color space not permitted on RGB PNG
libpng warning: iCCP: profile 'ICC profile': 'GRAY': Gray color space not permitted on RGB PNG
...
```

> CraigAD:	
> I found this description of the problem that suggests a possible solution:	
> https://stackoverflow.com/questions/43374187/qt-open-a-jpg-file-with-warning-image-format-not-supported
> 
> PNG has many Formats some old, some not standard for Qt. see more History_and_development(https://en.wikipedia.org/wiki/Portable_Network_Graphics#History_and_development) You need to use Standard format in Qt.


I wrote a new go tool **check_png** for convert liteide all png to standard png and fix this issues.


### check_png tool

https://github.com/visualfc/liteide/blob/master/liteidex/src/tools/check_png/main.go

code: convert png file to NRGBA format.

```
func (p *Process) ProcessPng(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	stat, err := f.Stat()
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	srcImage, err := png.Decode(f)
	if err != nil {
		return err
	}

	dstImage := image.NewNRGBA(srcImage.Bounds())
	draw.Draw(dstImage, dstImage.Bounds(), srcImage, srcImage.Bounds().Min, draw.Src)

	err = png.Encode(&buf, dstImage)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if bytes.Compare(data, buf.Bytes()) != 0 {
		log.Println("update png", path, len(data), buf.Len())
		ioutil.WriteFile(path, buf.Bytes(), stat.Mode())
	}
	return nil
}
```





