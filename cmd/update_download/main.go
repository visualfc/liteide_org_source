// update_download project main.go
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var tmpl_page = `---
title: Download
description:
---


<div class="panel-group">
	
	<div class="panel panel-default">
`
var tmpl_dir_name = `<div class="panel-heading">
		<h4 class="panel-title">
		  <a class="accordion-toggle" data-toggle="collapse" data-parent="#accordion" href="#collapse0">
			$DIRNAME
		  </a>
		</h4>
	  </div>
`
var tmpl_files_table = `<div class="accordion-body collapse in">
		<div class="accordion-inner">
		  <table class="table table-striped table-hover table-condensed">
			<thead>
			  	<tr>
				  	<th>File</th>
					<th>Size</th>
					<th></th>
				</tr>
			</thead>
			<tbody>
`
var tmpl_file_name = `
			  <tr>
				<td>$FILENAME</td>
				<td>$FILESIZE</td>
				<td>
				  <a href="{{site.home}}/assets/media/download/$DIRNAME/$FILENAME" class="btn">download</a>
				</td>
			  </tr>
`

var tmpl_files_table_end = `
			</tbody>
		  </table>
`

var tmpl_page_end = `
		</div>
	  </div>
	</div>
	
  </div>
`

var (
	root     = "./../.."
	flagRoot string
)

func init() {
	flag.StringVar(&flagRoot, "root", "", "setup liteide root")
}

func main() {
	flag.Parse()
	if flagRoot != "" {
		root = flagRoot
	}
	root, _ := filepath.Abs(root)
	log.Println("process download root", root)
	update(root)
}

func IsDir(filename string) bool {
	s, err := os.Stat(filename)
	return err == nil && s.IsDir()
}

func SizeInfo(bytes int64) string {
	const kb float64 = 1024
	const mb float64 = 1024 * kb
	const gb float64 = 1024 * mb
	const tb float64 = 1024 * gb
	fbytes := float64(bytes)
	if fbytes >= tb {
		return fmt.Sprintf("%.3f TB", fbytes/tb)
	} else if fbytes >= gb {
		return fmt.Sprintf("%.2f GB", fbytes/gb)
	} else if fbytes >= mb {
		return fmt.Sprintf("%.1f MB", fbytes/mb)
	} else if fbytes >= kb {
		return fmt.Sprintf("%.1f KB", fbytes/kb)
	}
	return fmt.Sprintf("%v bytes", bytes)
}

func update(root string) {
	pagedir := filepath.Join(root, "pages")
	if !IsDir(pagedir) {
		log.Fatalln("error read pages path", pagedir)
	}
	downdir := filepath.Join(root, "media/download")
	dirs, err := ioutil.ReadDir(downdir)
	if err != nil {
		log.Fatalln("error read download path", downdir)
	}
	var datas []string
	datas = append(datas, tmpl_page)
	for _, dir := range dirs {
		if dir.IsDir() {
			files, _ := ioutil.ReadDir(filepath.Join(downdir, dir.Name()))
			if len(files) > 0 {
				datas = append(datas, strings.Replace(tmpl_dir_name, "$DIRNAME", dir.Name(), -1))
				datas = append(datas, tmpl_files_table)
				for _, file := range files {
					if !file.IsDir() {
						data := strings.NewReplacer("$DIRNAME", dir.Name(), "$FILENAME", file.Name(), "$FILESIZE", SizeInfo(file.Size())).Replace(tmpl_file_name)
						datas = append(datas, data)
					}
				}
				datas = append(datas, tmpl_files_table_end)
			}
		}
	}
	datas = append(datas, tmpl_page_end)
	//fmt.Println(strings.Join(datas, "\n"))
	err = ioutil.WriteFile(filepath.Join(pagedir, "download.html"), []byte(strings.Join(datas, "\n")), 0777)
	if err != nil {
		log.Fatalln("error write downlaod file", err)
	} else {
		log.Println("success update download page", filepath.Join(pagedir, "download.html"))
	}
}
