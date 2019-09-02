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
			modTime: time.Date(2019, 9, 2, 9, 48, 32, 933122594, time.UTC),
		},
		"/handlers.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "handlers.tmpl",
			modTime:          time.Date(2019, 9, 2, 9, 48, 32, 932716629, time.UTC),
			uncompressedSize: 3845,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x57\x51\x93\xd3\x36\x10\x7e\x4e\x7e\xc5\xe2\x09\xd4\xee\xf8\x04\x9d\x61\x78\x38\x26\x9d\x81\xbb\x00\x69\xe1\x92\x5e\x42\xfb\xc0\x30\x77\xc2\x5e\x27\x6a\x1d\xd9\x27\xc9\x01\xc6\xf8\xbf\x77\x56\x72\x1c\x3b\x71\x81\x74\xa6\x0f\xf5\x93\x25\x7f\xda\x5d\xed\xb7\xfa\xb4\x2e\x4b\x18\xe5\x0a\x13\xf1\x09\xce\xc7\xe0\xbf\xe1\x7f\xe1\x34\x46\x69\x44\x22\x50\x01\x9b\xca\x24\x63\x4b\x61\x52\x0c\xa0\xaa\x86\xc3\x61\x59\x9e\x41\x8c\x89\x90\x08\xde\x9a\xcb\x38\x45\x75\x89\x51\xea\xd1\xd7\xa4\x90\x11\xf8\x1a\xd5\x16\x15\x94\x25\x73\x76\xab\x6a\x61\x67\x02\xb8\x69\xcd\xd9\x77\x6e\xd6\x6c\x96\xa3\xe2\x46\x64\x72\x7a\x59\x55\x37\xaf\x9c\x49\x3f\x82\x1f\x57\x42\xb2\x8b\x4c\x1a\xfc\x64\x02\x28\x87\x00\x00\x42\x86\x80\x4a\x51\xa4\x14\x68\xaf\x89\x6b\xbc\xf3\xa3\xc0\xc1\x13\x48\x51\xfa\xa8\x54\x00\x3f\xc3\xa3\xda\xc8\xee\xd9\x72\x45\xc6\xde\xe8\x15\x68\xa3\x84\x5c\x69\xf6\xbc\x10\x69\x8c\xaa\x03\x4b\x32\x05\x82\x3c\x2a\x2e\x57\x68\xdd\x77\xed\x58\xd0\xc6\xb0\x17\xb9\x12\xd2\xa4\xd2\x7f\xe0\xac\xda\x50\xdf\x89\xf7\x6c\xa2\x54\xa6\xfc\x20\xe8\xac\xaa\xec\x68\x30\x88\xd8\x2f\x8b\xd9\x95\xbf\x36\x26\x67\x0b\xc3\x4d\xa1\x9f\xf3\xf8\x1a\xef\x0a\xd4\x26\xa4\xc0\x8a\xc8\xd4\x0e\x07\x83\x81\x35\x55\x87\x0b\xb7\x7f\xea\x4c\x9e\x7b\x48\x73\xde\x6d\x0d\xa9\x4a\xe7\x9d\x2d\x2c\xc8\x0f\xaa\xa0\xfe\xa2\xd0\x14\x4a\xba\x81\xf3\xee\x98\x62\x0b\xb5\x65\xbd\xa9\xf4\x29\xdf\x51\x30\xac\x2c\xeb\x28\x63\x5b\x01\x9d\x12\x50\x78\xb7\xb0\x31\x36\x45\x40\x86\x09\x31\x72\xb1\x5f\xf1\x0d\xda\xca\xb2\xd9\x81\xb6\x03\xf0\xae\xf1\xce\x0b\xe0\xac\xaa\x86\xe6\x73\x8e\x40\xb5\xd8\x5a\x55\x55\xdd\xfd\x93\x55\x47\xc2\x28\xe7\x8a\x6f\xd0\xa0\x2d\x05\x36\xdf\x8d\x74\xdb\x7f\x83\x69\x42\x38\x28\xee\x3d\x82\x11\x24\x80\xde\xd5\x4b\x8a\x8c\x56\x5f\x64\x72\x8b\xca\xd8\x71\x6b\xe9\x22\x5a\xe3\x86\xb7\x16\x1f\x7a\xae\xaa\xee\xa4\x35\xd0\x72\x55\xe7\x75\x37\xfc\x28\xcc\x1a\x58\x5d\x01\xcf\xb3\xf8\xf3\x0e\x7b\x90\x82\x88\x8e\x86\xb4\xe1\x84\x30\xfa\x40\x40\xca\xc5\x85\x9b\xde\x2d\xb2\x06\xca\x12\x3a\xc1\x13\xb8\x8e\xfb\xd0\xb8\x0b\xa6\x3b\x6a\xe8\xb7\x0a\xd0\x66\x3f\xca\xa4\x63\x28\x53\x1d\xfe\x4f\xad\x80\xdd\xfe\x5b\x19\x58\x66\x17\x7b\xe3\x36\x6c\x66\x0b\xc5\x89\x8c\x3b\xfd\x07\xd5\x72\x24\x1a\xbe\x42\x0d\x65\xd9\x42\x55\x95\x3d\x95\x99\x7a\x2d\xb4\x81\x77\xef\xed\xfb\x4e\x5d\x6a\x41\x70\x00\x3b\x73\x03\x63\x1a\xb9\xe0\x1e\x3e\x84\xb7\xd7\x53\xf8\x20\x64\x2c\xe4\x6a\xb8\xd3\x86\x9b\x10\x2c\xb5\x7b\x89\x88\x5c\x41\xea\x96\x4e\xe8\x8f\xc2\x44\x6b\x07\x64\xbf\xe2\xe7\xd6\xa7\x16\xa5\x32\x84\x91\x4d\x15\x9b\xca\x39\x37\xeb\x36\x39\x11\xd7\x08\x1e\xed\x59\x42\x55\x79\xe7\x1d\x31\x51\xa8\x59\x59\xc2\x61\x7d\x13\x12\xc6\xb5\xd7\xdf\x79\x5a\xe0\x70\x38\x18\x1c\xd4\x1c\x3d\x31\x26\xbc\x48\x4d\xd7\xe8\x3e\x4f\x63\xe0\x79\x8e\x32\xf6\x9b\xa9\xd0\x2a\x9e\x55\xa3\xc4\xf7\xde\x4a\xfc\x94\x63\x64\x30\x86\x42\x09\x68\x2a\xfd\x1c\x6e\xef\xeb\x5b\x2f\xdc\xef\xbb\x25\x82\x6e\x6b\xad\xc2\xef\x49\xc2\x6f\x05\xaa\xa3\xfa\x17\x09\xe0\x1d\x8c\x90\xd9\xaa\xf0\x84\x34\xb8\x42\xe5\xb5\x61\x47\xd0\x17\x99\xda\x70\x63\xc1\x4f\x1e\x37\xd0\xaf\xa4\xcd\xdd\x33\x63\x12\xa0\x28\x93\x5b\x62\x54\xe3\x54\x1a\x3f\x62\x36\x28\x7f\x4f\x45\x10\xc2\x4f\x8f\x42\x78\xf2\xb8\xb9\x74\x68\xe9\xbd\x31\x48\x91\xb6\x78\xfe\x7a\x3a\xed\xab\x66\x7f\x28\x9e\xd3\x74\x08\xde\xc4\xa6\x94\x74\xde\x06\x0d\x42\xc2\xde\x65\x30\xdc\xa7\xb0\x73\x7c\x53\x8d\xa7\x6f\xee\x99\xc9\x44\xdf\xc6\xfe\xc3\x0d\x11\x65\xdf\xb9\xa5\x46\x91\x4e\xdb\x24\x8c\xa1\x67\x4b\x3d\x62\x77\x78\x20\xfe\xa1\x16\x5f\x21\x8f\x51\x7d\x47\xe9\x58\xc7\xb5\x7e\x33\xb7\x8a\xbd\x44\xd3\x89\xa2\xed\x73\x50\x57\x2a\x23\xad\xa6\x75\x42\x61\xdc\x48\xa2\x48\xf6\xd6\x60\xec\x38\xf8\xf2\xa5\xe5\xc1\x2a\xfc\xf8\x34\x72\xda\x87\xd7\xae\x17\x1a\xf8\x07\x8d\xd2\x74\x79\xe8\xc9\x8c\x48\x40\x66\xc6\x05\x7b\x62\x90\xf7\xc6\x60\xdb\x9c\xab\xcc\xdd\x4a\xff\x22\xd8\x96\xd2\xd0\x25\xf6\xcd\x68\x7b\xef\x4a\xc7\x68\x67\x03\xb5\x3e\x47\x86\x3e\xf5\xb3\x57\x5f\xac\x67\x64\xc4\x0b\x9e\x42\xd3\x96\x38\x65\xe6\x79\x9e\x8a\xc8\xde\x6b\x0f\xa9\x2f\x6b\x29\x74\x8c\x11\x99\xa5\x59\x76\x85\x1f\x2f\x31\xca\x62\x6a\x71\x3b\xd9\xa9\xbb\xb4\xfa\xb0\x9d\x8f\x69\x15\x73\x50\xff\x01\x55\x9b\x05\x3d\x3d\x3e\x8a\x83\xc1\xa9\x07\xf1\x82\xcb\x1f\x0c\x58\x45\x03\x6a\x3e\x77\x59\xac\xbb\xc2\xa3\xcb\xe0\x14\x7a\x74\x91\xe7\x99\x22\x7e\xea\x9c\x9f\x51\x67\x77\x0e\xf7\xb7\x5e\x08\x91\xf9\x1a\x5f\x75\x7f\xda\xd3\x77\x34\x3f\x1e\x35\x9d\x85\x4a\x1b\x1a\xe9\x9e\xd4\x3d\x94\x6f\xd0\xac\xb3\x38\x84\x51\x6e\x7b\x8f\x58\x44\x06\xbc\xf9\x6c\xb1\xf4\xe8\x3e\x98\x67\xda\x80\xf7\x72\xe2\x46\x2f\xd1\x80\x77\x39\x79\x3d\x59\x4e\xec\xf8\x12\x53\x34\x08\xde\xfc\x6d\x8d\x2e\x68\xed\xb3\xe5\xc5\x2b\x37\xe4\x54\x2b\xde\x6c\xbe\x9c\xce\xae\x16\x76\x6a\x96\x13\xf1\x3a\x38\xbc\xab\x6c\x2b\x33\xca\x69\xba\x2c\x0d\x6e\xf2\x94\x9b\xc3\x1f\xa7\x5d\x70\xee\xcf\xc8\x6b\x7e\xc7\x3c\x6a\xc8\x69\x08\x47\x66\xfb\xba\xb4\x53\x92\xf5\x7f\xca\xd5\xc1\x1f\x86\x03\x74\x10\x47\x5d\x68\x8d\xf9\x56\xc6\xf6\x6f\x7f\x07\x00\x00\xff\xff\x62\x39\xd7\xfc\x05\x0f\x00\x00"),
		},
		"/interface.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "interface.tmpl",
			modTime:          time.Date(2019, 9, 2, 9, 48, 32, 932831877, time.UTC),
			uncompressedSize: 586,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xcd\x6e\xab\x30\x10\x85\xf7\x3c\xc5\x91\xc5\x22\xb9\x4a\xc8\xfe\x4a\x77\x71\x95\xa0\x14\xa9\x29\xa8\xd0\x07\xb0\x60\x08\x56\x89\xa1\x66\x92\xb6\xb2\xfc\xee\x15\x3f\x49\xab\x26\x52\xbb\xf3\x8c\xe7\xe7\x3b\x67\x3c\x6b\x97\x28\xa8\x54\x9a\x20\x94\x66\x32\xa5\xcc\x69\x47\x5c\x35\x85\xc0\xd2\x39\x0f\x00\xac\x0d\xe2\x96\x8c\x64\xd5\xe8\x68\xe3\xdc\x4c\xe9\xab\xdc\x23\xbd\x2c\x90\xe3\xcf\x5e\xe9\x60\xdd\x68\xa6\x37\x9e\x0f\xd3\x49\x17\xce\x79\x1e\xbf\xb7\x04\x6b\xb1\x93\xcf\x14\x15\xa4\x59\x95\x8a\x0c\x82\x48\x97\x4d\x90\x29\xae\x09\xce\xe1\x82\x00\x3b\x6d\x5e\xc2\x48\xbd\x27\xf8\x47\x53\x2f\xe0\x13\xfe\xfe\x43\x90\x48\xae\xba\x09\xee\x5b\xd9\x61\x60\x5f\xc0\x6f\xfb\xca\x59\xa1\x72\x86\x48\xe2\x34\x13\xf0\x29\x48\x9a\x8e\x21\xb6\xe1\x18\x6d\x89\x21\x36\xe1\x7d\x98\x85\x43\xbc\xa1\x9a\x98\x20\x92\xa7\xa9\xfa\xd8\xf7\xfe\xcf\xd6\x77\x63\x28\x39\xaf\x20\xe2\x24\x8b\xe2\x87\x74\x48\xc5\x6d\xaf\xbf\x9b\xe3\x0b\xcc\x19\xe8\x55\x71\xd5\x63\x4c\x5f\xab\x55\x2f\xbf\x35\x4a\xf3\x99\x12\x02\x62\x10\x86\x8b\xd1\x4c\x87\xb6\x96\x7c\xeb\x1a\xe3\xa8\xab\x35\xa3\xbf\xb7\x32\x9f\xef\xdf\xfa\x9f\x92\x39\x91\x41\xc7\xe6\x98\xf3\x74\x82\xd4\x9c\xf0\x63\xa3\xe7\x3e\x02\x00\x00\xff\xff\x86\x24\xe5\x8a\x4a\x02\x00\x00"),
		},
		"/main.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "main.tmpl",
			modTime:          time.Date(2019, 9, 2, 9, 48, 32, 932947751, time.UTC),
			uncompressedSize: 334,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xce\x31\x6e\x42\x31\x0c\xc6\xf1\x99\x9c\xc2\xca\xd4\x0e\x24\x07\xe8\x58\x3a\x74\x29\x0b\x17\x08\x79\xc6\x2f\xe5\xc5\x8e\x1c\x53\x09\x3d\x71\xf7\x2a\x12\x43\xc5\x50\x26\xff\xa5\xef\x37\x38\x46\x78\x97\x09\x81\x90\x51\x93\xe1\x04\xc7\x2b\x90\x28\x76\x7b\x83\xdd\x1e\xbe\xf6\x07\xf8\xd8\x7d\x1e\x82\x73\x2d\xe5\x73\x22\x84\x75\x0d\xf7\xbc\xdd\x9c\x2b\xb5\x89\x1a\xbc\xb8\x8d\x47\xce\x32\x15\xa6\xf8\xdd\x85\xbd\xdb\xf8\x53\xb5\x71\xa8\xd8\x7c\x39\x86\x2c\x35\x52\xe1\x2d\x09\x97\x3c\xea\x61\x6b\x67\x8a\xa8\x2a\xda\xc7\xc0\x68\x71\x36\x6b\xa3\xbb\x69\x16\xfe\xb9\x67\x61\xea\xde\xbd\x3a\xb7\xae\x86\xb5\x2d\xc9\x10\x7c\x61\x43\x3d\xa5\x8c\xc1\x6a\x5b\x3c\x84\xde\x30\xc3\x76\x7c\xf8\x97\xcd\x89\xa7\x05\xb5\xff\xaf\xec\xda\xf0\x09\x51\xb9\x18\xea\xa3\xf9\x0d\x00\x00\xff\xff\x1f\x77\x1f\xd6\x4e\x01\x00\x00"),
		},
		"/router.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "router.tmpl",
			modTime:          time.Date(2019, 9, 2, 9, 48, 32, 933059739, time.UTC),
			uncompressedSize: 479,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x50\xd1\x4a\xc3\x40\x10\x7c\xee\x7d\xc5\x70\x04\x69\x25\xde\x07\x08\x3e\x48\x1b\xda\x80\x9a\xd0\xa6\xcf\x72\x34\x9b\xe4\x30\x5e\xc2\xe5\x52\x85\x70\xff\x2e\x77\x29\xda\x82\xfb\x36\xb3\xbb\x33\xb3\x3b\x4d\x88\x7a\x43\x95\xfa\xc6\xe3\x13\x96\xaf\xf2\x83\xd2\x92\xb4\x55\x95\x22\x03\x91\xea\xaa\x13\x85\xb2\x2d\xad\xe0\x1c\x63\xd5\xa8\x4f\xd8\x53\xad\x06\x4b\x66\xdf\x8d\x96\x86\xa5\xc1\x7d\xad\xb4\x48\x74\xad\x34\xc5\x90\xbd\xc2\x95\xaa\x73\x2b\x4c\x0c\x20\xaf\x7f\x77\xd3\x38\x90\x39\x93\x99\x64\xaf\x1c\x63\x8b\x69\x7a\x80\x91\xba\x26\x44\xa3\x69\x63\x44\x61\x43\xe4\xd2\x36\x83\x73\x6c\x01\x00\x57\x33\x9f\x64\x9b\xae\x8c\x11\xf5\x21\x78\xa9\x4e\x16\x3c\xcf\x0e\x05\x47\x44\x22\xef\x06\x0b\xbe\x4d\x66\xb4\x25\x0b\xbe\x49\x5e\x92\x22\x09\x78\x43\x2d\x59\x02\xcf\x8f\x97\xe9\xd1\xef\x3e\x17\xeb\xdd\x0c\xa5\x3d\x35\xe0\x59\x5e\xa4\xd9\xdb\x21\x50\x59\x6f\x55\xa7\x87\xf0\x84\xdf\x24\x5f\xca\x36\xde\xdf\x39\x06\x18\xb1\x93\xba\x6c\x69\xc9\xfd\x8d\x73\x3a\x38\xc7\x63\x78\x62\xdd\xe9\x33\x19\x7b\x34\x6d\xb8\x6e\x6e\x90\x78\xbf\xf9\xc7\x8c\x44\xd6\x93\x91\xde\x2e\xdd\x78\x72\x96\x35\x2b\x86\x4b\x79\x6b\xd2\x65\x70\xfd\x8f\xf9\x43\x8e\xfd\x04\x00\x00\xff\xff\xbe\x70\x58\x84\xdf\x01\x00\x00"),
		},
		"/types.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "types.tmpl",
			modTime:          time.Date(2019, 9, 2, 9, 48, 32, 933168599, time.UTC),
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
