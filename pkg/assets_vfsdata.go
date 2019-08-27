// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package pkg

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 8, 27, 7, 8, 1, 674218529, time.UTC),
		},
		"/handlers.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "handlers.tmpl",
			modTime:          time.Date(2019, 8, 26, 16, 50, 36, 366147175, time.UTC),
			uncompressedSize: 3308,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x56\x5f\x6f\xdb\xb6\x17\x7d\xb6\x3f\xc5\xfd\x09\x6e\x21\xfd\xa0\x30\x1d\x50\xf4\x21\x45\x06\xb4\x76\xda\x78\xeb\x62\x2f\x76\xf6\x9a\x2a\xd4\x95\xcd\x4d\xa6\x14\x92\x72\x12\xa8\xfa\xee\xc3\x25\x65\x59\xb2\x95\xb6\x79\x9c\x01\x03\x22\x79\x78\xee\xbf\xc3\x4b\x96\x25\x8c\x72\x85\x89\x78\x84\xb3\x73\xf0\xff\x88\xfe\xc1\x69\x8c\xd2\x88\x44\xa0\x02\x36\x95\x49\xc6\x96\xc2\xa4\x18\x40\x55\x0d\x87\xc3\xb2\x3c\x81\x18\x13\x21\x11\xbc\x75\x24\xe3\x14\xd5\x04\x79\xea\xd1\x6a\x52\x48\x0e\xbe\x46\xb5\x45\x05\x65\xc9\x1c\x6f\x55\x2d\xec\x4c\x00\xb7\xad\x39\xfb\x1d\x99\x35\x9b\xe5\xa8\x22\x23\x32\x39\x9d\x54\xd5\xed\xa5\xa3\xf4\x39\xfc\x7f\x25\x24\x1b\x67\xd2\xe0\xa3\x09\xa0\x1c\x02\x00\x08\x19\xc2\x2d\xf9\x49\x6e\xf6\x12\x5c\xe3\xbd\xcf\x03\x0b\x76\x7e\xb0\x85\xda\xb2\x5e\xa8\x4f\x6c\x3c\x18\x56\x36\x26\x94\xb1\x8d\xaf\x13\xa0\xc2\xfb\x85\x51\x05\x37\x4d\x88\x44\x4c\x88\x91\xb6\xf3\x57\xd1\x06\x6d\xde\x72\x25\xa4\x81\xb6\x01\xf0\xae\xf1\xde\x0b\xe0\xa4\xaa\x86\xe6\x29\x47\xa0\x4c\xb7\x76\x55\x15\xb8\x51\x1d\x1b\xb1\xaa\x48\xae\x10\x46\x79\xa4\xa2\x0d\x1a\x54\x44\xcd\xe6\xbb\x91\x6e\xdb\x6f\x30\x8d\x0b\x07\xa5\xdb\x23\x18\x41\x02\xe8\xdd\xbd\x24\xcf\x68\xf7\x38\x93\x5b\x54\xc6\x8e\x5b\x5b\x17\x7c\x8d\x9b\xa8\xb5\xf9\xd0\x72\x55\x75\x27\x2d\x41\xcb\x54\x9d\xd7\xdd\xf0\x41\x98\x35\xb0\x6b\xbc\x2f\x50\x9b\x8f\x59\xfc\xb4\xc3\x1e\xa4\x80\x53\xe1\xa5\x75\x27\x84\xd1\x1d\x01\x29\x17\x63\x37\xbd\xdb\x64\x09\xca\x12\x3a\xce\x13\xb8\xf6\xfb\x90\xdc\x39\xd3\x1d\x35\xe5\xb7\xfa\x6e\x57\x9f\x67\xd2\x55\x28\x53\x9d\xfa\xbf\x54\x01\xbb\xf8\x5b\x19\x58\x66\xe3\x3d\xb9\x75\x9b\x59\xa1\xb8\x23\xe4\xd4\x7d\xa0\x96\xa3\x23\xe1\x2b\xd4\x50\x96\x2d\x54\x55\x85\x80\x4a\xd1\x3f\x53\xbb\x33\xb3\x8d\xea\x09\x4d\x7a\x13\x72\xa5\xd9\xc7\x42\xa4\x31\x2a\xe7\xd5\xe9\x29\xdc\x5c\x4f\xe1\x4e\xc8\x58\xc8\x95\x9d\x4a\x32\x05\xb7\x21\xd8\x9a\x52\x6c\xae\x26\xdc\x29\x51\xd7\xbc\xf6\x88\x3d\x08\xc3\xd7\x0e\xc8\x7e\xc7\xa7\xd6\x52\xab\x96\x32\x84\x91\xcd\x11\x9b\xca\x79\x64\xd6\xed\xaa\xf0\x48\x23\x78\x14\xac\x84\xaa\xf2\xce\x9a\x05\xfa\x29\xd4\xac\x2c\xe1\x50\xd8\x84\x84\xf3\xda\xea\x5f\x51\x5a\xe0\x70\x38\x18\x1c\x88\x8d\x7e\x31\x26\x51\x91\x9a\x2e\x69\xb2\x31\xec\x93\x2d\x56\xe2\xbf\x76\x89\x09\xc1\xbb\x91\xf8\x98\x23\x37\x18\x43\xa1\x04\x34\x72\x3e\x83\xaf\xaf\xf4\x57\x2f\xdc\xc7\x18\x34\x6c\x2e\x8a\x96\xb8\x7b\xe2\xfd\xb3\x40\x75\xa4\x71\x91\x00\xde\xc3\x08\x99\xad\xbc\x27\xa4\xc1\x15\x2a\xaf\x0d\x3b\x82\x7e\xca\xd4\x26\x32\x16\xfc\xee\x6d\x03\xfd\x7e\x86\xb4\x51\x3c\x93\x5b\xaa\x9b\xc6\xa9\x34\x3e\x67\xd6\x1f\x7f\x9f\xf0\x20\x84\x5f\xde\x84\xf0\xee\x6d\x70\x64\x1a\x53\x8d\x2f\x32\xf3\xc1\x64\xa2\xcf\x44\x0f\x75\x73\x0e\x5f\x6c\xac\x87\xbf\xe7\x88\x1f\xaa\xe1\x99\xea\x5c\x62\x14\xa3\xfa\xb9\x64\xf2\x5d\xd7\x62\x6e\x17\xfb\x8c\xa6\xe3\x45\xdb\xe6\xa0\xae\x1d\xa3\x0e\x45\xfb\x84\xc2\xb8\x69\x04\x22\xd9\xb3\xc1\xf9\x39\x48\x91\xc2\xb7\x6f\x2d\x0b\xb6\xaf\xd5\x0b\xfb\x23\xd5\xaf\x5c\x8b\x15\x1a\xa2\x3b\x8d\xd2\xd4\xd9\x78\xb6\xff\x8a\x04\x64\x66\x9c\x5f\x2f\xf4\xe7\x7f\x3f\xe7\x4f\xeb\x24\x51\x27\xfe\x91\x43\xbd\xfd\xde\xd5\xa7\xe3\x63\xdd\x6a\xb8\xa1\xa5\xfe\x5a\xd4\x97\xc3\x09\x91\x78\xc1\x7b\x68\xae\x56\xd7\x64\xa2\x3c\x4f\x05\xb7\xbd\xf9\xf4\x6f\x9d\xc9\x56\xb3\x89\x91\x13\x2d\xcd\xb2\x2b\x7c\x98\x20\xcf\x62\x7a\x84\x74\x12\xe0\x02\x19\xd0\x89\x54\xf6\x66\x8e\x91\x33\x07\xf5\x5f\x93\x76\x2c\xe8\xbd\x5d\xed\xe4\x6a\x30\xe8\xcf\xd4\x38\x92\x86\xfa\x8a\x46\xf8\x6d\x31\xbb\xf2\x02\x38\x3d\x5d\xce\x26\xb3\x33\x10\x72\x8b\xda\x88\x55\x64\xd0\x36\xe3\x4d\xa6\x10\x62\x34\x91\x48\x31\x76\xdd\xdc\x31\xbb\xb4\x1e\xf5\xb9\xe7\x2a\xa3\x8b\x3c\xcf\x14\x95\xa6\xce\xf7\x09\xbd\x4c\xce\xe0\xd5\xd6\x0b\x81\x9b\xef\x94\xca\x45\x9d\x29\xcd\xbe\xa0\xf4\x03\xf8\x15\xde\xb4\x94\x40\x21\x9f\x5b\xab\x17\x04\x4a\xfc\x1a\xbb\xb0\x17\x8e\x1f\xb4\x89\x15\x9a\x42\xc9\x61\xcf\x05\xdc\xbc\x2f\x6b\x4d\x14\x2a\x6d\xb4\x40\xf7\x86\xee\xd1\xcd\x06\xcd\x3a\x8b\x43\x18\xe5\xf6\x12\x8e\x05\x37\xe0\xcd\x67\x8b\xa5\x47\x4d\x73\x9e\x69\x03\xde\xe7\x0b\x37\xfa\x8c\x06\xbc\xc9\xc5\x97\x8b\xe5\x85\x1d\x4f\x30\x45\x83\xe0\xcd\x6f\x6a\x74\x41\x7b\x3f\x2c\xc7\x97\x6e\x18\x91\xe0\xbc\xd9\x7c\x39\x9d\x5d\x2d\xec\xd4\x2c\x27\xf5\xe8\xe0\xb0\xa1\xdb\x3b\x7d\x94\xd3\x74\x59\x1a\xdc\xe4\x29\x55\xae\xfb\x3e\xde\x39\xe7\x1e\xc0\x5e\xf3\xea\xf6\xe8\x65\x4a\x43\x38\xa2\xed\x7b\xae\xbc\x24\x59\xff\xa5\x5c\x1d\x3c\xb5\x1d\xa0\x83\x38\x7a\x8e\xd5\x98\x1f\x65\x6c\xff\xf5\x6f\x00\x00\x00\xff\xff\xb1\x2d\xdb\xd1\xec\x0c\x00\x00"),
		},
		"/interface.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "interface.tmpl",
			modTime:          time.Date(2019, 8, 26, 16, 44, 26, 113606321, time.UTC),
			uncompressedSize: 586,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xcd\x6e\xab\x30\x10\x85\xf7\x3c\xc5\x91\xc5\x22\xb9\x4a\xc8\xfe\x4a\x77\x71\x95\xa0\x14\xa9\x29\xa8\xd0\x07\xb0\x60\x08\x56\x89\xa1\x66\x92\xb6\xb2\xfc\xee\x15\x3f\x49\xab\x26\x52\xbb\xf3\x8c\xe7\xe7\x3b\x67\x3c\x6b\x97\x28\xa8\x54\x9a\x20\x94\x66\x32\xa5\xcc\x69\x47\x5c\x35\x85\xc0\xd2\x39\x0f\x00\xac\x0d\xe2\x96\x8c\x64\xd5\xe8\x68\xe3\xdc\x4c\xe9\xab\xdc\x23\xbd\x2c\x90\xe3\xcf\x5e\xe9\x60\xdd\x68\xa6\x37\x9e\x0f\xd3\x49\x17\xce\x79\x1e\xbf\xb7\x04\x6b\xb1\x93\xcf\x14\x15\xa4\x59\x95\x8a\x0c\x82\x48\x97\x4d\x90\x29\xae\x09\xce\xe1\x82\x00\x3b\x6d\x5e\xc2\x48\xbd\x27\xf8\x47\x53\x2f\xe0\x13\xfe\xfe\x43\x90\x48\xae\xba\x09\xee\x5b\xd9\x61\x60\x5f\xc0\x6f\xfb\xca\x59\xa1\x72\x86\x48\xe2\x34\x13\xf0\x29\x48\x9a\x8e\x21\xb6\xe1\x18\x6d\x89\x21\x36\xe1\x7d\x98\x85\x43\xbc\xa1\x9a\x98\x20\x92\xa7\xa9\xfa\xd8\xf7\xfe\xcf\xd6\x77\x63\x28\x39\xaf\x20\xe2\x24\x8b\xe2\x87\x74\x48\xc5\x6d\xaf\xbf\x9b\xe3\x0b\xcc\x19\xe8\x55\x71\xd5\x63\x4c\x5f\xab\x55\x2f\xbf\x35\x4a\xf3\x99\x12\x02\x62\x10\x86\x8b\xd1\x4c\x87\xb6\x96\x7c\xeb\x1a\xe3\xa8\xab\x35\xa3\xbf\xb7\x32\x9f\xef\xdf\xfa\x9f\x92\x39\x91\x41\xc7\xe6\x98\xf3\x74\x82\xd4\x9c\xf0\x63\xa3\xe7\x3e\x02\x00\x00\xff\xff\x86\x24\xe5\x8a\x4a\x02\x00\x00"),
		},
		"/main.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "main.tmpl",
			modTime:          time.Date(2019, 8, 27, 6, 15, 43, 456090745, time.UTC),
			uncompressedSize: 297,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xcd\x31\x4e\x43\x31\x0c\xc6\xf1\xb9\x3e\x85\x95\x09\x86\xe6\x1d\x80\x91\x32\xb0\xd0\xa5\x17\x48\xf3\x5c\x37\xf0\x62\x47\x8e\x8b\x54\x3d\xf5\xee\x28\x52\x07\xc4\x40\x27\xff\x25\xff\xa4\x6f\x9a\xf0\x55\x67\x42\x26\x21\x4b\x4e\x33\x1e\xaf\xc8\x6a\xd4\xfd\x05\x77\x7b\xfc\xd8\x1f\xf0\x6d\xf7\x7e\x88\x00\x2d\xe5\xaf\xc4\x84\xeb\x1a\xef\x79\xbb\x01\x94\xda\xd4\x1c\x9f\x60\x13\x48\xb2\xce\x45\x78\xfa\xec\x2a\x01\x36\xe1\x54\x7d\x1c\x2e\x7e\xbe\x1c\x63\xd6\x3a\x71\x91\x2d\xab\x94\x3c\x6a\xfc\xba\x5b\x56\xf9\xbe\x67\x11\xee\x01\x9e\x01\xd6\xd5\xa9\xb6\x25\x39\x61\x28\xe2\x64\xa7\x94\x29\x7a\x6d\x4b\xc0\xd8\x1b\x65\xdc\x8e\xf1\xdf\xec\x9c\x64\x5e\xc8\xfa\xff\xca\xaf\x8d\x1e\x10\xd3\x8b\x93\xfd\x35\x3f\x01\x00\x00\xff\xff\xe8\x2a\xfe\x84\x29\x01\x00\x00"),
		},
		"/router.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "router.tmpl",
			modTime:          time.Date(2019, 8, 26, 16, 45, 27, 349693703, time.UTC),
			uncompressedSize: 468,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xd1\x6a\xb3\x40\x10\x85\xaf\xb3\x4f\x71\x58\xe4\x27\xf9\xb1\xfb\x00\x85\x5e\x94\x46\x12\xa1\xad\x92\xd8\xeb\xb2\xc4\x51\x97\xda\x55\xd6\xb1\x2d\xc8\xbe\x7b\x71\x0d\x6d\x02\xf5\xee\x3b\xce\x9c\x73\x76\xa6\x09\x51\xef\xa8\x32\x5f\xb8\xbd\xc3\xfa\x49\xbf\x51\x5a\x92\x65\x53\x19\x72\x50\xa9\xad\x3a\x55\x18\x6e\x69\x03\xef\x85\xa8\x46\x7b\xc2\x81\x6a\x33\x30\xb9\x43\x37\x32\x0d\x6b\x87\xff\xb5\xb1\x2a\xb1\xb5\xb1\x14\x43\xf7\x06\x17\xae\xde\x6f\x30\x09\x80\x66\xff\x7f\x57\x3f\x8e\xe4\x3e\xc8\x4d\xba\x37\x5e\x88\xd5\x34\xdd\xc0\x69\x5b\x13\xa2\xd1\xb5\x31\xa2\xb0\xa1\x72\xcd\xcd\xe0\xbd\x58\x01\xc0\xc5\xcc\x3b\x71\xd3\x95\x31\xa2\x3e\x14\x2f\xcd\x89\x21\xf3\xec\x58\x48\x44\xa4\xf2\x6e\x60\xc8\x5d\xb2\xd0\x8e\x18\x72\x9b\x3c\x26\x45\x12\x78\x4b\x2d\x31\x41\xe6\x2f\xe7\xe9\x71\xde\xbd\x2f\x1e\xf6\x0b\x6a\x3e\x35\x90\x59\x5e\xa4\xd9\xf3\x31\x48\x59\xcf\xa6\xb3\x43\x38\xc2\x4f\x93\x4f\xc3\xcd\x9c\xef\xbd\x00\x9c\x6a\xb4\x2d\x5b\x5a\xcb\xf9\x8d\x4b\x3b\x78\x2f\x63\x04\x61\x74\xed\x42\xa4\x5e\xaf\x8e\xb0\x90\xca\x7a\x72\x7a\xce\x48\xb7\xb3\xb8\x0f\x5e\x6e\x23\x70\xfe\xe6\x3c\xb2\x65\x88\xfa\x4b\xf9\x25\x2f\xbe\x03\x00\x00\xff\xff\x6b\xc1\x56\x47\xd4\x01\x00\x00"),
		},
		"/types.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "types.tmpl",
			modTime:          time.Date(2019, 8, 27, 7, 8, 1, 673651083, time.UTC),
			uncompressedSize: 727,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\x41\x4e\xf3\x30\x10\x85\xf7\x3e\xc5\xc8\xea\xa2\x95\xfe\xe6\x00\x95\xba\xf8\xd5\x15\x8b\x22\x04\x1c\xa0\x56\xfc\x1a\x0c\x8d\x6d\x6c\x03\xaa\xac\xb9\x3b\x72\x9c\x96\x50\x65\xc5\x2e\x8a\xfd\x7d\xe3\xf7\x34\x39\xaf\x49\xe3\x68\x2c\x48\xb6\xae\xf7\xce\xc2\xa6\xa7\xf6\x05\xbd\x92\xc4\x2c\xca\xb9\x39\x12\xde\xa9\xd1\x88\x6d\xf3\x7c\xf6\x20\x19\x53\x30\xb6\x93\xb4\x66\x16\xa9\xfc\xc9\xb9\xb1\xaa\x07\x33\xd5\x23\x91\x33\xe1\x14\x31\xc3\x1a\x9b\xd0\x21\xfc\x0d\x56\x21\xa8\xf3\x2c\x9a\x33\x2d\x77\xce\x7e\x22\xa4\xe1\xea\x62\xc0\x56\x35\xc3\x55\x57\x5d\xff\xb5\x36\xc9\x38\xab\x4e\x0f\xc1\x79\x84\x64\x10\xe7\x9c\x5b\x7a\x8d\xce\x36\x8f\xea\x6b\x8f\x18\x55\x87\xab\x6a\xfe\xed\x1f\x6d\xa2\x2c\x88\x4a\x69\x41\xd9\x0e\xb4\xf0\xc1\xf9\x72\xe1\x5f\xfd\xa4\xcd\x76\x7c\xc3\x64\x32\xb3\x20\xaa\xd4\xc2\xab\xa0\x7a\x24\x84\x7b\xd5\xa3\xdc\x5e\xee\xd5\x1b\xee\x34\x6c\x32\x47\x83\xf0\x63\x5c\xcd\x73\x43\xf8\xc2\xfd\x2a\xa3\x40\x13\xe0\x76\x4e\xed\xef\x46\xc2\x4c\x87\x92\x7f\x23\x73\xbe\x4e\x65\x96\x87\x31\x21\xac\x66\x16\xb5\x5e\xab\x2f\xcb\x52\xff\x0a\x21\x26\x25\x8c\x05\x94\xdc\x43\x01\xbb\xcb\xa2\xc5\xa6\xae\x5a\x1c\xe1\x84\xde\x9f\x54\x9a\xdb\xc5\xa5\x36\x6d\x22\x59\x54\xb2\x1a\x49\x16\xa1\x1c\xbd\x43\x3a\x91\x73\x9d\xff\x1d\x00\x00\xff\xff\xc4\x91\x7f\x4d\xd7\x02\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/handlers.tmpl"].(os.FileInfo),
		fs["/interface.tmpl"].(os.FileInfo),
		fs["/main.tmpl"].(os.FileInfo),
		fs["/router.tmpl"].(os.FileInfo),
		fs["/types.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
